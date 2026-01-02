# Vite Migration Plan (Vue 2)

Goal: replace Vue CLI 5 with Vite while keeping Vue 2.x and TypeScript.

## 1) Prep

- Commit or stash current changes.
- Ensure Node 20 is active: `nvm use 20.19.6`.
- Remove Vue CLI build dependencies from `package.json`:
  - `@vue/cli-service`, `@vue/cli-plugin-*`
  - `vue-template-compiler`
  - `@fullhuman/vue-cli-plugin-purgecss`

## 2) Install Vite + Vue 2 plugin

- Add Vite and plugin packages:
  - `vite`
  - `@vitejs/plugin-vue2`
  - `@vitejs/plugin-legacy` (only if needed for older browsers in `browserslist`)
- Add Vue 2 compiler:
  - `vue-template-compiler` should be replaced with `vue-template-compiler@2.x` only if a dependency still needs it
  - otherwise prefer `@vue/compiler-sfc` for Vue 2: not required; Vue 2 plugin handles SFC

## 3) Update scripts

Replace Vue CLI commands with Vite scripts in `package.json`:

- `"serve": "vite"`
- `"build": "vite build"`
- `"lint": "eslint --ext .ts,.vue src"`

Add optional preview:

- `"preview": "vite preview"`

## 4) Add Vite config

Create `vite.config.ts` (or `.js`) with:

- Vue 2 plugin setup.
- Aliases to match `@` -> `src`.
- Define `process.env` or use `import.meta.env` and update code as needed.
- Configure CSS (Sass) and PostCSS.
- Configure build target to match current `browserslist`.

## 5) Port Vue CLI config

Translate settings from `vue.config.js`:

- Webpack performance hints do not map directly; use `build.chunkSizeWarningLimit`.
- If `chainWebpack` is setting `optimization.minimize`, use:
  - `build.minify: true/false` based on `NODE_ENV`.

## 6) PostCSS / PurgeCSS

Move purgecss from the Vue CLI plugin to `postcss.config.js`:

- Keep current `@fullhuman/postcss-purgecss` config.
- Ensure PostCSS 8 compatible version is used.
- Optionally gate by `NODE_ENV === "production"` as it is now.

## 7) Environment variables

Vite uses `VITE_` prefix. For any `process.env.*` usage:

- Rename to `VITE_*` in `.env` files.
- Update code to `import.meta.env.VITE_*`.
- If server-side code still expects `process.env`, add a small compatibility shim or use `define` in Vite config.

## 8) Fix asset handling

- Replace `require()` of assets with `new URL("./path", import.meta.url)` where needed.
- Update any webpack-specific loaders.

## 9) Verify TypeScript + ESLint

- Ensure `tsconfig.json` matches Vite expectations (module `ESNext`, target `ES2018+`).
- Add `vite/client` types if using `import.meta.env`.
- Keep `eslint-plugin-vue` and `@typescript-eslint` versions compatible.

## 10) Validate build

- `yarn`
- `yarn serve`
- `yarn build`
- Fix any runtime or build errors.

## 11) Cleanup

- Remove `vue.config.js` and any Vue CLI-specific docs/scripts.
- Update README or dev docs with Vite commands.

## Notes / Risks

- Some Vue 2 ecosystem libraries expect webpack-specific globals; those may need shims.
- If legacy browser support is required, configure `@vitejs/plugin-legacy`.
- If `vue-cli-plugin-uikit` is required, replace with manual UIkit import/config.
