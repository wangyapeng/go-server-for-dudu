import vue from '@vitejs/plugin-vue'
const path = require('path')
export default {
    // 配置选项
    root: '.',
    server: {
        port: 3030,
    },
    alias: {
        '~bootstrap': path.resolve(__dirname, 'node_modules/bootstrap'),
        '@': path.resolve(__dirname, 'src')
    },
    plugins: [vue()],
}