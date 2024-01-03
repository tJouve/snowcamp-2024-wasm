---
marp: true
---
# S01E01

## WebAssembly (Wasm), 🐣 baby steps (today in the browser 🌍)
### Part 2

> - Repo: https://gitlab.com/k33g-twitch/20231102-wasm-01/01-wasm-2nd-steps
> - 👀 `README.md` => Open it with Docker Development Environment

---

# Agenda

- Host functions?
- (Simple) Demo
- Display Host function
- "Display" Demo
- Strings and functions: the easy way 🎉
  - Demo with **GoLang**
  - Demo with **TinyGo**
---
# Btw

- No debug with WebAssembly

---

# Host functions?

- A function defined in the Host application
- For The Wasm program, it’s used as an import function

---
# (Simple) Demo

- JavaScript side: `yo()`
- GoLang Wasm side: 
  ```golang
  func helloWorld() {
    yo() // <- this is JavaScript
  }
  ```
---
# (Simple) Demo

- `03-wasm-go`

---
# Display Host function

## No `Println` in Wasm 🤬
### `hostDisplay: (pos, size) => {}` callable by the Wasm module


---
# "Display" Demo

- `04-wasm-go-string-param`

---
## You can do it with every language that targets WebAssembly

---
# Strings and functions: the easy way 🎉

## Golang support of Wasm & JavaScript

### Get `wasm_exec.js`
```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

### Build
```bash
GOOS=js GOARCH=wasm go build -o main.wasm
```

---
# IDE support (VSCode)

> `.vscode/settings.json`
```json
{
    "go.toolsEnvVars": {
      "GOOS": "js",
      "GOARCH": "wasm",
    },
  }
```
---
# Demo with GoLang

- `05-wasm-go-the-easy-way`

---
# Strings and functions: the easy way 🎉

## TinyGo 🥰 support of Wasm & JavaScript

### Get `wasm_exec.js`
```bash
cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" .
```

### Build
```bash
tinygo build --no-debug -o main.wasm -target wasm main.go
```

---

# Demo with TinyGo

- `06-wasm-tinygo-the-easy-way`


---

# To read

- https://blog.suborbital.dev/foundations-wasm-in-golang-is-fantastic

---

# Next time(s):

- Go deeper with the Wasm support with GoLang
- You can do it with Rust (or another language 🤔)
