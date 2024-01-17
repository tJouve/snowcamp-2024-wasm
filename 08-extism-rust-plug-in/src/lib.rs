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

    info!("message: {}", msg);

    let my_text = config::get("text")?.unwrap_or("🎉 tada".to_string());

    info!("my text from config: {}", my_text);

    let output = Output { message: msg , from: "🦀 Rust".to_string()};
    
    Ok(Json(output))
}

