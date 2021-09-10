import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Vue.use(Buefy)

Wails.Init(() => {
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
