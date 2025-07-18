<template>
  <div class="container">
    <div class="app">
      <div class="hitokoto">
        <p>{{ hitokoto }}</p>
        <p class="author" v-if="hitokotoAuthor">—— {{ hitokotoAuthor }}</p>
        <!-- <p class="author">—— Hello</p> -->
      </div>

      <div class="url-form">
        <input v-model="longUrl" type="text" placeholder="请输入长链接"
          class="url-input" @keyup.enter="generateShortUrl">
        <button @click="generateShortUrl" class="generate-btn">生成</button>
      </div>

      <div v-if="shortUrl" class="result">
        <a :href="shortUrl" target="_blank" class="short-url">{{ shortUrl
        }}</a>
        <button @click="copyToClipboard" class="copy-btn">复制</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import axios from 'axios'
import { onMounted, ref } from 'vue'

const longUrl = ref('')
const shortUrl = ref('')
const hitokoto = ref('正在加载一言...')
const hitokotoAuthor = ref('')
const apiUrl = 'https://v1.hitokoto.cn/'
const generateShortUrl = () => {
  if (!longUrl) {
    alert('请输入有效的URL');
    return;
  }
  axios.post('http://127.0.0.1:8080/api/v1/add', {
    "url": longUrl.value
  })
    .then((response) => {
      const data = response.data
      shortUrl.value = data.shortUrl
    })
}

const copyToClipboard = () => {
  navigator.clipboard.writeText(shortUrl.value);
  alert('已复制到剪贴板');
}

const fetchHitokoto = () => {
  axios.get(apiUrl)
    .then((response) => {

      hitokoto.value = response.data.hitokoto || '一言API请求失败';
      hitokotoAuthor.value = response.data.from
    })
    .catch((response) => {

      hitokoto.value = '网络错误，无法获取一言';
    })

}

onMounted(() => {
  fetchHitokoto()
})

</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Helvetica Neue', Arial, sans-serif;
}

.container {
  /* display: flex; */
  min-height: 100vh;
  background-color: #f5f7fa;
  padding: 20px;
}

.app {
  /* width: 100%; */
  width: fit-content;
  margin: 0 auto;
  margin-top: 20vh;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 40px;
  text-align: center;
}

.title {
  color: #2c3e50;
  margin-bottom: 30px;
  font-size: 28px;
}

.hitokoto {
  display: flex;
  flex-direction: column;
  margin: 20px 0;
  padding: 15px;
  border-radius: 8px;
  font-style: bold;
  min-height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 40px;
}

.author {
  margin-top: 10px;
  font-size: 16px;
  color: #333;
  font-style: italic;
  align-self: flex-end;
  margin-right: -20px;
}

.url-form {
  display: flex;
  margin: 30px 0;
}

.url-input {
  flex: 1;
  padding: 12px 15px;
  border: 1px solid #ddd;
  border-radius: 6px 0 0 6px;
  font-size: 16px;
  outline: none;
  transition: border 0.3s;
}

.url-input:focus {
  border-color: #3498db;
}

.generate-btn {
  padding: 12px 20px;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 0 6px 6px 0;
  cursor: pointer;
  font-size: 16px;
  transition: background 0.3s;
}

.generate-btn:hover {
  background-color: #2980b9;
}

.result {
  margin-top: 20px;
  padding: 15px;
  background: #f8f9fa;
  border-radius: 8px;
}

.short-url {
  color: #3498db;
  text-decoration: none;
  word-break: break-all;
  display: inline-block;
  margin: 10px 0;
  margin-right: 10px;
}

.short-url:hover {
  text-decoration: underline;
}

.copy-btn {
  margin-top: 10px;
  padding: 8px 15px;
  background-color: #2ecc71;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: background 0.3s;
}

.copy-btn:hover {
  background-color: #27ae60;
}
</style>