#!/bin/bash

cat > ./00-wasm-c/main.c <<- EOM
// c
EOM


cat > ./00-wasm-c/index.html <<- EOM
<html>
  <head>
    <meta charset="utf-8"/>
    <link rel="stylesheet" href="my.css">
  </head>
    <!-- c -->
    <body>
        <h1>WASM Experiments</h1>
        <h2>👋 Open the developer tools 😃</h2>
        <script> 
            // Load the wasm file
        </script>
  </body>
</html>
EOM

cat > ./01-wasm-go/main.go <<- EOM
package main
// go
// add add (1) and hello (2) function

func main() {}
EOM

cat > ./01-wasm-go/index.html <<- EOM
<html>
  <head>
    <meta charset="utf-8"/>
    <link rel="stylesheet" href="my.css">
  </head>
    <!-- go -->
    <body>
        <h1>WASM Experiments</h1>
        <h2>👋 Open the developer tools 😃</h2>
        <script> 
            // Load the wasm file (3)
        </script>
  </body>
</html>
EOM


cat > ./02-wasm-go/main.go <<- EOM
package main

import (
	"unsafe"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

// readBufferFromMemory (6)

// copyBufferToMemory (7)

// helloWorld (8)

EOM

cat > ./02-wasm-go/index.html <<- EOM
<html>
  <head>
    <meta charset="utf-8"/>
    <link rel="stylesheet" href="my.css">
  </head>
    <!-- go -->
    <body>
        <h1>WASM Experiments</h1>
        <h2>👋 Open the developer tools 😃</h2>
        <script> 
            // Load the wasm file
            let importObject = {
              wasi_snapshot_preview1: {
                fd_write: () => 0,
              }
            }

            WebAssembly.instantiateStreaming(fetch("main.wasm"), importObject) 
              .then(({ instance }) => {
                console.log("📦 Instance", instance)
                // 🖐 Prepare the string argument (9)

                // Copy argument to memory (10)

                // Call the helloWorld TinyGo function (11)

                // Get pos and size of the result (12)

                // Read the memory and (13)
                // "Extract" the buffer 
                
                // Display the string (14)

              })
              .catch(error => {
                console.log("😡 ouch", error)
              })
        </script>
  </body>
</html>
EOM