import { createApp } from 'vue'
import './assets/css/style.css'
import App from './App.vue'
import { create, NButton } from 'naive-ui'

const naive = create({
	components: [NButton]
})


const app = createApp(App)
app.use(naive)
app.mount('#app')
