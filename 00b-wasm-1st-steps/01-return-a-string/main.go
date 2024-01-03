package main

import (
	"unsafe"
)

//export helloWorld
func helloWorld() uint64 {

	message := "👋 Hello World 🌍 Webassembly is amazing 💜"

	positionAndLengthOfTheMessage := copyBufferToMemory([]byte(message))

	return positionAndLengthOfTheMessage
}

func main() {}

// copyBufferToMemory copies a buffer to the WebAssembly memory buffer
func copyBufferToMemory(buffer []byte) uint64 {

	// create a pointer on the buffer
	bufferPtr := &buffer[0]
	unsafePtr := uintptr(unsafe.Pointer(bufferPtr))

	pos := uint32(unsafePtr)
	length := uint32(len(buffer))

	// [shift left + or]
	//  pos                              length
	// 110101⏺000000000000000000000000⏺10110100

	return (uint64(pos) << uint64(32)) | uint64(length)
}

