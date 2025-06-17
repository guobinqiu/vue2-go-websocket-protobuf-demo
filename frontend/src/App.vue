<template>
  <div id="app">
    <input v-model="text" placeholder="Say something..." @keyup.enter="sendMsg" />
    <div v-for="(msg, index) in messages" :key="index">
      {{ msg.user }}: {{ msg.text }}
    </div>
  </div>
</template>

<script>
import protobuf from 'protobufjs';

export default {
  data() {
    return {
      socket: null,
      ChatMessage: null,
      text: '',
      messages: [],
    };
  },
  mounted() {
    //使用 protobuf.js 从 public 目录加载 chat.proto 文件
    protobuf.load('/chat.proto').then(root => {
      this.ChatMessage = root.lookupType('chat.ChatMessage'); //查找包名是 chat，类型是 ChatMessage 的消息类型
      this.connect();
    }).catch(error => {
      console.error("Failed to load proto file:", error);
    });
  },
  beforeDestroy() {
    this.stopSocket();
  },
  methods: {
    connect() {
      this.socket = new WebSocket('ws://localhost:8080/ws');
      this.socket.binaryType = 'arraybuffer'; // 选项有 arraybuffer | blob

      this.socket.onmessage = (event) => {
        const msg = this.ChatMessage.decode(new Uint8Array(event.data)); // 将服务端的二进制数据解码成对应的消息对象
        this.messages.push({ user: msg.user, text: msg.text });
      };

      this.socket.onopen = () => {
        console.log("WebSocket connection established.");
      };

      // onclose 被触发时(如断网)，socket连接尚未关闭
      this.socket.onerror = (error) => {
        console.error("WebSocket error:", error);
        this.stopSocket();
      };

      // onclose 被触发时(如服务端主动关闭连接或者断网)，socket连接已经关闭了
      this.socket.onclose = () => {
        console.log("WebSocket connection closed.");
        setTimeout(() => {
          console.log("尝试重连...");
          this.connect();
        }, 3000);
      };
    },
    sendMsg() {
      if (!this.text.trim()) return;

      const msg = this.ChatMessage.create({ user: 'guobin', text: this.text }); // 创建一个新的消息对象
      const buffer = this.ChatMessage.encode(msg).finish(); // 将消息对象转换为二进制格式
      this.socket.send(buffer);
      this.text = '';
    },
    stopSocket() {
      if (this.socket && this.socket.readyState === WebSocket.OPEN) {
        this.socket.close();
        this.socket = null;
      }
    },
  }
};
</script>

<style scoped>
input {
  width: 300px;
  padding: 10px;
  margin-bottom: 10px;
}
</style>
