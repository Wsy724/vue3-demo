import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  // base: './', //打包的相对路径
  plugins: [
    vue(),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },

  server: {
    proxy: {
      // 匹配以/user开头的请求
      '/user': {
        // 目标后端接口地址
        target: 'http://127.0.0.1:3300',
        changeOrigin: true, // 开启跨域代理（必须！）
      }
    }
  }
})

