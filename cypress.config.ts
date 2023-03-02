import { defineConfig } from "cypress";
defaultCommandTimeout: 10000

export default defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
});
