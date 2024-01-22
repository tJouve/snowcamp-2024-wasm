# Create a JavaScript plug-in
> https://github.com/extism/js-pdk

Create a directory `09-extism-js-plug-in` with two files: `index.js` and `index.d.ts`:
```bash
mkdir 09-extism-js-plug-in
cd 09-extism-js-plug-in
touch index.js
touch index.d.ts
```

## Content of `index.js`

The program exports a hello function which will take a `data` as a string and return a `result` string:

```javascript
function hello() {
	let data = Host.inputString()

	let text = "💛 Hello " + data

    console.log(text)

	let result = {
		message: text,
        from: "Extism"
	}
    
	Host.outputString(JSON.stringify(result))
}

module.exports = { hello }
```

## Content of `index.d.ts`

Describe the Wasm interface for the plug-in in a typescript module DTS file:

```typescript
declare module 'main' {
    export function hello(): I32;
}
```

## Build the plug-in

```bash
extism-js index.js -i index.d.ts -o hello-js.wasm
```

## Run the plug-in (and call the hello function)

```bash
extism call hello-js.wasm \
  hello \
  --input "👩 Jane Doe" \
  --log-level "info" \
  --wasi
```


