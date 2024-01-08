package main

// #include <stdlib.h>
import "C"

import "unsafe"

func pack(pos uint32, size uint32) uint64 {
	return (uint64(pos) << uint64(32)) | uint64(size) // shift left + or bit operation
}

// copyStringToMemory returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
// The pointer is not automatically managed by TinyGo hence it must be freed by the host.
func copyStringToMemory(s string) (uint32, uint32) {
	size := C.ulong(len(s))
	ptr := unsafe.Pointer(C.malloc(size))
	copy(unsafe.Slice((*byte)(ptr), size), s)
	return uint32(uintptr(ptr)), uint32(size)
}

// readStringFromMemory returns a string from WebAssembly compatible numeric types
// representing its pointer(position) and length.
func readStringFromMemory(position uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(position))), size)
}

//export hello
func hello(strPos uint32, strSize uint32) uint64 {

	value := readStringFromMemory(strPos, strSize)
	message := "👋 Hello " + string(value) + " 😃"

	// copy the value to memory
	// get the position and the size of the string (in memory)
	pos, size := copyStringToMemory(message)

	// return a pointer/size pair packed into a uint64.
	// Note: This uses a uint64 instead of two result values for compatibility with WebAssembly 1.0.
	return pack(pos, size)

}

func main() {}
