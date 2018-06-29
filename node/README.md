# Getting Started

- **`express-parse`** - NPM Package containing the APIs and binary.
- **`math-grammar`** - Example implementation

# Setup

**Tools:**
- `node` - https://nodejs.org/
- `yarn` - `npm i -g yarn`
- `verdaccio` - `npm i -g verdaccio` (then: `verdaccio &`)

**Steps:**

1. Publish `express-parse` locally
```
cd express-parse
yarn
yarn build
npm publish --registry=http://localhost:4873
```
2. Build the example grammar
```
cd ../math-grammar
yarn
yarn build
```
3. Compile the example source
```
yarn test
```

# Edit and Test

### Testing changes to `express-parse`

1. Rebuild and publish `express-parse`
> *Context: `node/express-parse`*
```
yarn build && npm unpublish --force && npm publish --force
```
2. Clean old version and re-test `math-grammar`
> *Context: `node/math-grammar`*
```
rm -rf yarn.lock node_modules && yarn cache clean && yarn && yarn start
```

### Testing changes to `math-grammar`

1. Rebuild and run `math-grammar`
```
yarn start
```
