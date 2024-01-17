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

<!-- TODO -->

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
  --log-level "info" \
  --set-config '{"text":"Hello I am Jane Doe 😊"}' \
  --wasi
echo ""

extism call ./target/wasm32-wasi/release/hello_extism_rust.wasm \
  hello \
  --input "👨 John Doe" \
  --set-config '{"text":"Hello I am John Doe 👋"}' \
  --log-level "info" \
  --wasi
echo ""

```