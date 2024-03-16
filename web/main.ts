import {createApp} from 'vue'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import Curtain from './Curtain.vue'

let curtain = createApp(Curtain)

curtain.use(ElementPlus)
curtain.mount('#container')