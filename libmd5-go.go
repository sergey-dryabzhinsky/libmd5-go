package main

// #include <stdlib.h>
import "C"
import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"runtime"
	"unsafe"
)

const VERSION = "0.0.5"

//export libmd5_go__MD5_hexdigest
func libmd5_go__MD5_hexdigest(inputText *C.char) *C.char {
	goText := C.GoString(inputText)
	hasher := md5.New()            // Creates a new hash.Hash object
	io.WriteString(hasher, goText) // Writes the string data to the hasher

	// Get the final hash as a byte slice. Passing nil appends the hash to an empty slice.
	hashInBytes := hasher.Sum(nil)

	// Convert the byte slice to a hex string
	gohexDigest := hex.EncodeToString(hashInBytes)
	return C.CString(gohexDigest)
}

//export libmd5_go__MD5File_hexdigest
func libmd5_go__MD5File_hexdigest(fullPath *C.char) *C.char {
	goFullPath := C.GoString(fullPath)

	// Open the file
	file, err := os.Open(goFullPath)
	if err != nil {
		return C.CString("")
	}
	// Ensure the file is closed after the function returns
	defer file.Close()

	hash := md5.New() // Creates a new hash.Hash object

	// Copy the file content to the hash object.
	// The hash object implements the io.Writer interface.
	if _, err := io.Copy(hash, file); err != nil {
		return C.CString("")
	}

	// Get the final hash as a byte slice. Passing nil appends the hash to an empty slice.
	hashInBytes := hash.Sum(nil)

	// Convert the byte slice to a hex string
	gohexDigest := hex.EncodeToString(hashInBytes)
	return C.CString(gohexDigest)
}

//export libmd5_go__MD5_digest
func libmd5_go__MD5_digest(inputText *C.char) *C.char {
	goText := C.GoString(inputText)
	hasher := md5.New()            // Creates a new hash.Hash object
	io.WriteString(hasher, goText) // Writes the string data to the hasher

	// Get the final hash as a byte slice. Passing nil appends the hash to an empty slice.
	hashInBytes := hasher.Sum(nil)

	// Convert the byte slice to a Go string first.
	// This creates a copy of the data.
	goString := string(hashInBytes)

	// Convert the Go string to a C-style string (*C.char).
	// C.CString makes another copy and ensures null-termination.
	return C.CString(goString)
}

//export libmd5_go__FreeResult
func libmd5_go__FreeResult(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}

//export libmd5_go__version
func libmd5_go__version() *C.char {
	return C.CString(VERSION)
}

//export libmd5_go__version_go
func libmd5_go__version_go() *C.char {
	return C.CString(runtime.Version())
}

func main() {}
