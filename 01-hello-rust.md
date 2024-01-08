# 01-hello-rust

## Create a new project with source code

```bash
cargo new 01-hello-rust --name hello

# the name `01-hello` cannot be used as a package name, the name cannot start with a digit
# If you need a package name to not match the directory name, consider using --name flag.
```

Restart the rust-analyzer (Cmd+Shipt+p: rust analyzer - restart server)

## Build

```bash
cd 01-hello-rust
cargo build --target wasm32-wasi
# the wasm file is generated into:
# target/wasm32-wasi/debug/hello.wasm
#cargo build --target=wasm32-wasi --release
```

> if needed (already installed): `rustup target add wasm32-wasi``

## Run

```bash
wasmtime target/wasm32-wasi/debug/hello.wasm
wasmedge target/wasm32-wasi/debug/hello.wasm
wazero run target/wasm32-wasi/debug/hello.wasm
```

## Change the source code

```rust
use std::io;

fn main() {
    println!("👋 hello, please type your name:");
    
    let mut user_input = String::new();
    let stdin = io::stdin();
    let _ = stdin.read_line(&mut user_input);

    println!("🤗 thank you {} ", user_input);
}
```
> Re-build and re-run
