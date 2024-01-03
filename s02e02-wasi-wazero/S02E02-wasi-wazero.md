---
marp: true
---
# S02E02

## WebAssembly (Wasm), outside the browser with Wasi

> - Repo: https://gitlab.com/k33g-twitch/s02e02-wasi-wazero
> - 👀 `README.md` => Open it with Docker Development Environment

---
# Agenda

- Small demo (reminder)
- `f(string) string`
- Wazero

---
# Wasm/Wasi: some limitations

- Quick reminder: little **demo** `01-first-wasm-program`

> ✋ I use code snippets because I have no memory
> 🥰 I use TinyGo to build the wasi programs
---
# One of the “annoying” limitations

- Only numbers 😮
  - How to pass string arguments to a Wasm function?
  - How to return a string as the result of a Wasm function call?


---
# Solution

## Exchange data with the Shared Memory Buffer


---
![auto](imgs/01.png)

---
![auto](imgs/02.png)

---
![auto](imgs/03.png)

---
![auto](imgs/04.png)

---
# Wasi CLI: DIY 🛠️

- You can develop your own CLI
- But, you need to handle the limitations
  - == Develop all the **“plumbing”**

---
# Wazero Runtime & SDK 🩵🩵🩵

**wazero**: the **zero** dependency WebAssembly runtime for **Go** developers

> - https://wazero.io
> - https://github.com/tetratelabs/wazero/tree/main/examples

---
# Write your 1st CLI 🚀

## Demo `02-wazero`

---
# But, sometimes, you need more

- Make HTTP requests
- Make Redis requests from the Wasm module
- Use MQTT or NATS
- …

---
# Host Function?

- A function defined in the Host application
- For The Wasm program, it’s used as an import function

---
![auto](imgs/05.png)

---
# “Helpers”, but…

- ✋ You need to write your own glue
- For every language you want to support on the Wasm side 😵‍💫

## 🤬 It’s complicated! But…

---
# There is another way (easier) 👀
> The cross-language framework for building with WebAssembly

![auto](imgs/06.svg)
> by **Dylibso**

---
# Next time(s): S02E03

- Extism: présentation
- Demos: SDK +PDK

---
# 📝 Reading

https://k33g.hashnode.dev/series/wazero-first-steps

- Wazero Cookbook - Part One: WASM function & Host application
- Wazero Cookbook - Part Two: Host functions

---

# 👋


