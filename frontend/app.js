import { createApp, h } from 'vue'
import { createInertiaApp } from '@inertiajs/vue3'

import "./scss/app.scss"; // needed for HMR hot reload

createInertiaApp({
  resolve: (name) => import(`./Pages/${name}.vue`),
  setup({ el, App, props, plugin }) {
    createApp({ render: () => h(App, props) })
      .use(plugin)
      .mount(el)
  },
})
