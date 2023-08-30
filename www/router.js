const Home = () => import('./src/Home.vue');
const Login = () => import('./src/Login.vue');
const Register = () => import('./src/Register.vue');
import { createRouter, createWebHashHistory } from 'vue-router'
const routes = [
    { path: '/', component: Home, name: "home" },
    { path: '/login', component: Login, name: "login" },
    { path: '/register', component: Register, name: "register" },
]

const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHashHistory(),
    routes, // `routes: routes` 的缩写
})

export {
    router
}