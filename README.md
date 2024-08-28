# Recipe Helper Front-End

## TODO:

- [ ] CSS Improvements
  - [ ] Fix buttons on header
  - [ ] Add hover transform (darker) to buttons/links

- [ ] Combine FE and BE into monorepo
  - [ ] Create docker files for both
  - [ ] Use docker-compose to run both
- [ ] Use GitHub Actions to deploy to ... somewhere

- [ ] Extract recipe CRUD functions to a service
- [ ] Add images for recipes
- [ ] Add search bar to recipes view
- [ ] Add fridge/pantry ingredient input
- [ ] Add recipe suggestion based on fridge/pantry ingredients
  - [ ] Use Ollama locally, but what to use when deployed?
- [ ] Add a shopping list view

## Done:
- [x] Remove boilerplate from Vite template
- [x] Convert recipes output to a table component
  - [x] Add buttons for edit and delete to table
  - [x] Click on a row to go to the recipe detail page
- [x] Add a recipe detail page
- [x] Add a recipe deletion function
- [x] Add a recipe edit page
- [x] Add a recipe creation page

## React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type aware lint rules:

- Configure the top-level `parserOptions` property like this:

```js
export default tseslint.config({
  languageOptions: {
    // other options...
    parserOptions: {
      project: ["./tsconfig.node.json", "./tsconfig.app.json"],
      tsconfigRootDir: import.meta.dirname,
    },
  },
});
```

- Replace `tseslint.configs.recommended` to `tseslint.configs.recommendedTypeChecked` or `tseslint.configs.strictTypeChecked`
- Optionally add `...tseslint.configs.stylisticTypeChecked`
- Install [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) and update the config:

```js
// eslint.config.js
import react from "eslint-plugin-react";

export default tseslint.config({
  // Set the react version
  settings: { react: { version: "18.3" } },
  plugins: {
    // Add the react plugin
    react,
  },
  rules: {
    // other rules...
    // Enable its recommended rules
    ...react.configs.recommended.rules,
    ...react.configs["jsx-runtime"].rules,
  },
});
```
