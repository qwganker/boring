import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  css: {
    preprocessorOptions: {
      scss: { api: 'modern-compiler' },
    }
  },
  server: {
    port: 5173,
    // local dev proxy: forward API requests to backend service
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:7832',
        changeOrigin: true,
        secure: false,
        // keep the /api prefix when forwarding; remove or adjust rewrite if backend expects otherwise
        // rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  }
})
