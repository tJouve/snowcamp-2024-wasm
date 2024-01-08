---
marp: true
---
# S01E00

## WebAssembly (Wasm), 🐣 baby steps (today in the browser 🌍)

> - Repo: https://gitlab.com/k33g-twitch/20231026-wasm-00/00-wasm-1st-steps
> - 👀 `README.md` => Open it with Docker Development Environment

---

# Wasm?

- Code > Bytecode (wasm binary file)
- Binary format for executing code on the Web
- The JavaScript VM is responsible for the execution of the WASM code
- WASM is polyglot
- WASM is safe

---

# Why WASM?

- A complement to JavaScript
- Near-native speeds
- Complex applications in web browsers

---

# The primary qualities of WASM

- Speed, 
- Efficiency, 
- Safe, 
- Versatile, 
- Portable

---

# WebAssembly in the browser is amazing

- https://earth.google.com/
- https://stackblitz.com/edit/node-5xrecy?file=index.js

---

# Simplest Demo

- 00-wasm-c

---

# Wasi

> https://wasi.dev

- Wasi: WebAssembly System Interface
- Interface between 
  - WebAssembly (WASM) code  
  - and a Runtime environment
- Allowing WASM code to be run in various contexts (even the browser) 


---

# Demo

- 01-wasm-go

--- 

# Some Limitations

- Only numbers 😮
- How to pass string arguments to a Wasm function?
- How to return a string as the result of a Wasm function call?


---

# Workaround

## Solution: Shared Memory Buffer

---

![01](pictures/wasm-string-1.png)

---

![02](pictures/wasm-string-2.png)

---

![03](pictures/wasm-string-3.png)

---

![04](pictures/wasm-string-4.png)

---

# Demo "Plumbing"

- 02-wasm-go

---

# To read

- About WASM, WASI and Strings with NodeJS: https://k33g.hashnode.dev/series/wasi-nodejs

---

# Next time: Host Functions
