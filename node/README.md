# Getting Started

- **`express-parse`** - NPM Package containing the APIs and binary.
- **`math-grammar`** - Example implementation

# Using `express-parse`

## CLI
```
$ express-parse -f x.expr -g grammyGrams/ [-s program]

Options:
  --help                 Show help                                     [boolean]
  --version              Show version number                           [boolean]
  -g, --grammar          Directory containing grammar definitions
                                                             [string] [required]
  -f, --in-file          Source file                         [string] [required]
  -s, --start-collector  Starting token            [string] [default: "program"]
```


## APIs

The following are exported from the `express-parse` module:
- **`default`** `function` - Executes lexer, passes to a tree reducer, and returns the resulting tree.
- **`lexer`** `function` - Executes lexer and returns the raw collectors and tokens (rarely necessary; just use `default()`)
- **`Collector`** `class` - Base class for all collectors
- **`Token`** `class` - Base class for all tokens

## Collectors

> **NOTICE:** This section needs to be expanded.

To create a collector, extend the exported `Collector` class and set the `type` and `expected` static variables.

- **`type`** `string` - Represents the collector name even after it's reduced
- **`expected`** `array<array<Token||Collector>>` - Array of arrays containing tokens and collectors. This defines your expected tokens and collectors to be parsed. The first dimension defines the individual possible streams, while the second dimension defines each token or collector that is expected for that stream.

## Tokens

> **NOTICE:** This section needs to be expanded.

To create a token, extend the exported `Token` class and set the `type` and `expected` static variables.

- **`type`** `string` - Represents the token name even after it's reduced
- **`expected`** `array<string||RegExp>` - Array containing the possible expected values for this particular token. Each can either be a string or a Regular Expression.

# Contributing

## TODOs

- Values should be represented as Buffers
- Bug exists that allows Tokens to indicate they're satisfied even when they only contain partials of their expected matches
- Handle whitespace
	- Implementation: Pass *everything* to Token; if it kicks it back and that character is whitespace, continue ignoring all characters until we find a non-whitespace, and then start on the next token.

## Setup

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
3. Test compiling the example source
```
yarn test
```

You should get something like this for the input `5+5/8`:
```json
{
	"type": "MATHSTMT",
	"tokens": [
		{ "type": "DIGIT", "value": "5" },
		{ "type": "MATHOP", "value": "+" },
		{
			"type": "MATHSTMT",
			"tokens": [
				{ "type": "DIGIT", "value": "5" },
				{ "type": "MATHOP", "value": "/" },
				{ "type": "DIGIT", "value": "8" }
			]
		}
	]
}
```

## Edit and Test

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
