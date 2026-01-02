import {fileURLToPath} from "node:url";

import {defineConfig} from "vite";
import vue2 from "@vitejs/plugin-vue2";

const srcPath = fileURLToPath(new URL("./src", import.meta.url));

export default defineConfig(({mode}) => ({
  plugins: [vue2()],
  resolve: {
    alias: {
      "@": srcPath
    }
  },
  css: {
    preprocessorOptions: {
      scss: {
        quietDeps: true,
        silenceDeprecations: ["legacy-js-api", "import"]
      }
    }
  },
  build: {
    chunkSizeWarningLimit: 1000,
    minify: mode !== "development"
  }
}));
