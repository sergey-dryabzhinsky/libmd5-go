package main

// #include <constants.h>
// #include <stdlib.h>
import "C"
import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"hash"
	"os"
	"sync"
	"runtime"
	"unsafe"
)

var VERSION string
var commonL sync.Mutex
var commonHasher hash.Hash

//export libmd5_go_nts__MD5_init
func libmd5_go_nts__MD5_init() {
	if commonHasher != nil {
		commonHasher = nil
	}
	commonHasher = md5.New()            // Creates a new hash.Hash object
}

//export libmd5_go_ts__MD5_init
func libmd5_go_ts__MD5_init() {
	commonL.Lock()
	if commonHasher != nil {
		commonHasher = nil
	}
	commonHasher = md5.New()            // Creates a new hash.Hash object
	commonL.Unlock()
}

//export libmd5_go_nts__MD5_update
func libmd5_go_nts__MD5_update(inputText *C.char) *C.int {
	goText := C.GoString(inputText)
	if commonHasher == nil {
		result := C.int(0)
		return &result
	}
	io.WriteString(commonHasher, goText) // Writes the string data to the hasher
	result := C.int(1)
	return &result
}

//export libmd5_go_ts__MD5_update
func libmd5_go_ts__MD5_update(inputText *C.char) *C.int {
	goText := C.GoString(inputText)
	commonL.Lock()
	if commonHasher == nil {
		result := C.int(0)
		return &result
	}
	io.WriteString(commonHasher, goText) // Writes the string data to the hasher
	commonL.Unlock()
	result := C.int(1)
	return &result
}

//export libmd5_go_nts__MD5_finish
func libmd5_go_nts__MD5_finish() *C.char {
	// Get the final hash as a byte slice. Passing nil appends the hash to an empty slice.
	hashInBytes := commonHasher.Sum(nil)

	// Convert the byte slice to a hex string
	gohexDigest := hex.EncodeToString(hashInBytes)
	return C.CString(gohexDigest)
}

//export libmd5_go_ts__MD5_finish
func libmd5_go_ts__MD5_finish() *C.char {
	// Get the final hash as a byte slice. Passing nil appends the hash to an empty slice.
	commonL.Lock()
	hashInBytes := commonHasher.Sum(nil)
	commonL.Unlock()

	// Convert the byte slice to a hex string
	gohexDigest := hex.EncodeToString(hashInBytes)
	return C.CString(gohexDigest)
}

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
