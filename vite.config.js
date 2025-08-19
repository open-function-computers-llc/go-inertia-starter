import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { viteStaticCopy } from 'vite-plugin-static-copy';
import path from 'path';

export default defineConfig(({command}) => ({
  base: command === 'serve' ? '' : '/dist/',
  root: 'frontend',
  server: command === 'serve'
    ? {}
    : { middlewareMode: true}
  ,
  plugins: [
    vue(),
    viteStaticCopy({
      targets: [
        { src: '../server/index.html', dest: '.' },
      ]
    })
  ],
  build: {
    outDir: '../dist',
    emptyOutDir: true,
    manifest: true,
    rollupOptions: {
      input: {
        scss: "frontend/scss/app.scss",
        main: "frontend/app.js",
      }
    },
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "frontend")
    }
  },
}));
