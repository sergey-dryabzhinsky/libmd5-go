package main

// #include "constants.h"
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

const ERRNO_NO_ERROR = 0
const ERRNO_GENERIC_ERROR = 1
const ERRNO_MD5_CTX_NOT_INITED = 100
const ERRNO_OS_FILE_NOT_EXISTS = 404
const ERRNO_OS_FILE_NOT_READABLE = 403

var VERSION string
var lastErrorCode int = ERRNO_NO_ERROR
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

//export libmd5_go_nts__getLastErrorCode
func libmd5_go_nts__getLastErrorCode() C.int {
	result := C.int(lastErrorCode)
	lastErrorCode = ERRNO_NO_ERROR
	return result
}

//export libmd5_go_nts__getLastErrorStr
func libmd5_go_nts__getLastErrorStr() *C.char {
	result := C.CString("No errors happened")
	if (lastErrorCode == ERRNO_NO_ERROR){
		// pass
	}
	if (lastErrorCode == ERRNO_GENERIC_ERROR){
		result = C.CString("Some generic error happend")
	}
	if (lastErrorCode == ERRNO_MD5_CTX_NOT_INITED){
		result = C.CString("Init md5 context first!")
	}
	if (lastErrorCode == ERRNO_OS_FILE_NOT_EXISTS){
		result = C.CString("File by given path not exists! Or not readable!")
	}
	if (lastErrorCode == ERRNO_OS_FILE_NOT_READABLE){
		result = C.CString("File by given path maybe not readable!")
	}
	lastErrorCode = ERRNO_NO_ERROR
	return result
}

//export libmd5_go_nts__MD5_update
func libmd5_go_nts__MD5_update(inputText *C.char) C.int {
	goText := C.GoString(inputText)
	if commonHasher == nil {
		result := C.int(0)
		lastErrorCode = ERRNO_MD5_CTX_NOT_INITED
		return result
	}
	io.WriteString(commonHasher, goText) // Writes the string data to the hasher
	result := C.int(1)
	return result
}

//export libmd5_go_ts__MD5_update
func libmd5_go_ts__MD5_update(inputText *C.char) C.int {
	goText := C.GoString(inputText)
	commonL.Lock()
	if commonHasher == nil {
		commonL.Unlock()
		result := C.int(0)
		lastErrorCode = ERRNO_MD5_CTX_NOT_INITED
		return result
	}
	io.WriteString(commonHasher, goText) // Writes the string data to the hasher
	commonL.Unlock()
	result := C.int(1)
	return result
}

//export libmd5_go_nts__MD5_finish
func libmd5_go_nts__MD5_finish() *C.char {

	if commonHasher == nil {
		result := C.CString("")
		lastErrorCode = ERRNO_MD5_CTX_NOT_INITED
		return result
	}
	// Get the final hash as a byte slice. Passing nil appends the hash to an empty slice.
	hashInBytes := commonHasher.Sum(nil)

	// Convert the byte slice to a hex string
	gohexDigest := hex.EncodeToString(hashInBytes)
	return C.CString(gohexDigest)
}

//export libmd5_go_ts__MD5_finish
func libmd5_go_ts__MD5_finish() *C.char {

	if commonHasher == nil {
		result := C.CString("")
		lastErrorCode = ERRNO_MD5_CTX_NOT_INITED
		return result
	}
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

//export libmd5_go_nts__MD5File_update
func libmd5_go_nts__MD5File_update(fullPath *C.char) C.int {

	if commonHasher == nil {
		result := C.int(0) // @todo: return error code to get error description
		lastErrorCode = ERRNO_MD5_CTX_NOT_INITED
		return result
	}
	goFullPath := C.GoString(fullPath)

	// Open the file
	file, err := os.Open(goFullPath)
	if err != nil {
		if err == os.ErrNotExist {
			lastErrorCode = ERRNO_OS_FILE_NOT_EXISTS
		}
		lastErrorCode = ERRNO_OS_FILE_NOT_READABLE
		result := C.int(0)
		return result
	}
	// Ensure the file is closed after the function returns
	defer file.Close()

	// Copy the file content to the hash object.
	// The hash object implements the io.Writer interface.
	if _, err := io.Copy(commonHasher, file); err != nil {
		lastErrorCode = ERRNO_GENERIC_ERROR
		result := C.int(0)
		return result
	}
	result := C.int(1)
	return result
}

//export libmd5_go_ts__MD5File_update
func libmd5_go_ts__MD5File_update(fullPath *C.char) C.int {

	commonL.Lock()
	if commonHasher == nil {
		commonL.Unlock()
		result := C.int(0)
		lastErrorCode = ERRNO_MD5_CTX_NOT_INITED
		return result
	}
	commonL.Unlock()
	goFullPath := C.GoString(fullPath)

	// Open the file
	file, err := os.Open(goFullPath)
	if err != nil {
		if err == os.ErrNotExist {
			lastErrorCode = ERRNO_OS_FILE_NOT_EXISTS
		}
		lastErrorCode = ERRNO_OS_FILE_NOT_READABLE
		result := C.int(0)
		return result
	}
	// Ensure the file is closed after the function returns
	defer file.Close()

	// Copy the file content to the hash object.
	// The hash object implements the io.Writer interface.
	commonL.Lock()
	if _, err := io.Copy(commonHasher, file); err != nil {
		commonL.Unlock()
		result := C.int(0)
		lastErrorCode = ERRNO_GENERIC_ERROR
		return result
	}
	commonL.Unlock()
	result := C.int(1)
	return result
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
