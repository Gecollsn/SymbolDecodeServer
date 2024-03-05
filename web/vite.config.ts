import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'

export default defineConfig(({ command, mode, isSsrBuild }) => {
    return {
        server: {
            host: "0.0.0.0",
            port: 7000
        },
        plugins: [vue(),vueJsx()],
        resolve: {
            extensions: ['.vue'],
        },
    }
})