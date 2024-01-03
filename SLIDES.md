---
marp: true
---

# Créer des plugins en Rust, TinyGo, ... 
## pour vos applications Go avec WebAssembly et Wazero et Extism
### SnowCamp 2024 ⛄️

---
# Agenda

- 🚧 work in progress

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

# WebAssembly in the browser is amazing

- https://earth.google.com/
- https://stackblitz.com/edit/node-5xrecy?file=index.js


---

# The primary qualities of WASM

- Speed, 
- Efficiency, 
- Safe, 
- Versatile, 
- Portable

---
# Free WASM from the browser

## Let it Go!

---

# Wasi?

> https://wasi.dev

- Wasi: WebAssembly System Interface
- Interface between 
  - WebAssembly (WASM) code  
  - and a Runtime environment
- Allowing WASM code to be run in various contexts (even the browser) 

---
# Some WASI Use Cases

- CLI applications
- Applications with plug-ins (Zellij, Lapce)
- Database UDF (ScyllaDB, PostgreSQL)
- WebHooks, Filters, … (Webhook Relay, Envoy)
- FaaS (Fermyon cloud, WasmCloud, Shopify, …)
- …
---
# At least, 3 ways to run Wasm programs outside the browser

- WASI Runtimes CLI
- WASI Runtimes SDK
- Ready to use applications with embedded Wasm runtime
  - Spin from Fermyon
  - Wasm Workers Server from Wasm Lab
  - …
---
# WASI Runtimes

- WasmEdge, 
- Wasmtime, 
- Wasmer, 
- Wazero 🩵, 
- NodeJS, 
- … 

---
# Demos and Exercise


