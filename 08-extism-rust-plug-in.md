# Create an Extism plug-ins
> - run the plugins with the Extism CLI
> - use the host functions provided by the PDK


## Create an Extism Rust plug-in
> Writing and reading files

```bash
cargo new --lib 08-extism-rust-plug-in --name hello_extism_rust
cd 08-extism-rust-plug-in

rustup target add wasm32-wasi # if needed
```

Restart the rust-analyzer (Cmd+Shipt+p: rust analyzer - restart server)

In the generated `Cargo.toml` file, be sure to include:

```toml
[lib]
crate_type = ["cdylib"]
```
> ref: https://doc.rust-lang.org/reference/linkage.html

```bash
cargo add extism-pdk
cargo add serde
cargo add serde_json
```
> check the `Cargo.toml` file

## Update the source code

You need to update the source code of `/src/lib.rs` with the following code:

```rust
#![no_main]

use extism_pdk::*;
use serde::Serialize;

#[derive(Serialize)]
struct Output {
    pub message: String,
    pub from: String,
}

#[plugin_fn]
pub fn hello(input: String) -> FnResult<Json<Output>> {

    let msg: String = "👋 Hello ".to_string() + &input;

    let output = Output { message: msg , from: "🦀 Rust".to_string()};
    
    Ok(Json(output))
}
```

## Build 

```bash
# build
cargo clean
cargo build --release --target wasm32-wasi #--offline
# ls -lh *.wasm
ls -lh ./target/wasm32-wasi/release/*.wasm
```

## Run the plug-in

```bash
extism call ./target/wasm32-wasi/release/hello_extism_rust.wasm \
  hello \
  --input "👩 Jane Doe" \
  --wasi
echo ""

extism call ./target/wasm32-wasi/release/hello_extism_rust.wasm \
  hello \
  --input "👨 John Doe" \
  --wasi
echo ""

```

## Use the config object and display information

- To display information, you can use `info!` macro: 
  - https://github.com/extism/rust-pdk/blob/main/src/macros.rs#L11
  - https://github.com/extism/rust-pdk/blob/main/examples/http.rs#L7
- To use the config object, have a look to https://github.com/extism/rust-pdk/blob/main/src/config.rs#L22


Build again and run it again with these new flags:

```bash
  --log-level "info" \
  --set-config '{"text":"Hello I am Jane Doe 😊"}' \
```
