import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './registerServiceWorker'
import VueResource from '../node_modules/vue-resource'
import VueCookie from '../node_modules/vue-cookie'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faStar } from '@fortawesome/free-solid-svg-icons'
import { faUndo } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

window.$ = require('jquery')
window.JQuery = require('jquery')

library.add(faStar)
library.add(faUndo)

Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.use(VueResource);
Vue.use(VueCookie);
Vue.use(BootstrapVue);

Vue.config.productionTip = false;
export const EventBus = new Vue();


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
