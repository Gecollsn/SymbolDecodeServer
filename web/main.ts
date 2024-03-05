import {createApp} from 'vue'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import Curtain from './Symparse.vue'

Curtain.use(ElementPlus)
createApp(Curtain).mount('#container')