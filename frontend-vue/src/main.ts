import { createApp, provide, h } from "vue";

import App from "./App.vue";
import router from "./router";
import store from "./store";

import { ApolloClient, InMemoryCache } from "@apollo/client";
import { DefaultApolloClient } from "@vue/apollo-composable";

const defaultApolloClient = new ApolloClient({
  uri: "http://localhost:4000/query",
  cache: new InMemoryCache(),
});

createApp({
  setup() {
    provide(DefaultApolloClient, defaultApolloClient);
  },
  render() {
    return h(App);
  },
})
  .use(store)
  .use(router)
  .mount("#app");
