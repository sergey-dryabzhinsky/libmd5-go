# libmd5-go
Last version: 0.0.8

Pure md5 function nothing else. Exported from golang runtime. And some examples of code.

## Requirements
To compile library and examples you will need to install
- c99 capable compiler (for now only gcc >=4.8 supported)
- openssl >=1.0 ( libcrypto )
- golang >=1.6

## Prepare system
Instructions for Debian/Ubuntu like systems:
```
apt install build-essential tar xz-utils m4 sed
apt install libssl-dev
apt install golang-go
```

## Compile

If your version of go not default:
```
GOROOT=/path/to/golang/installdir GO=go-x.y make lib
```

## Reason
Not found library with only one functionality, so I tryed make it by myself.

## Result
Slightly faster file reading than `md5sum` command.
But more memory consuming.

## API
- **libmd5_go__version**(void): Returns doted version string.

  return: `* char`

  *since*: 0.0.1

- **libmd5_go__version_go**(void): Returns version string of go rutime used.

  return: `* char`

  *since*: 0.0.4

- **libmd5_go__MD5_hexdigest**(char* text): Return hexed-string with md5 digest.

  *Deprecated*

  params:
  - text (`char *`): input string

  return: `char *` String with hexed digest

  *since*: 0.0.2

- **libmd5_go_nts__MD5_getLastErrorCode**(void): int

  *Not Thread Safe*

  Returns code of last error happened.

  params:
  - none

  return: `int` internal code number

  *since*: 0.0.8

- **libmd5_go_nts__MD5_getErrorDescription**(int errno):

  *Not Thread Safe*

  Returns description of error by its code number.

  params:
  - `int` errno: error internal code number.

  return: `char*` internal error description.

  *since*: 0.0.8

- **libmd5_go_nts__MD5_init**(void): void

  *Not Thread Safe*

 (Re)Initialize md5 context.

  params:
  - none

  return:

  *since*: 0.0.7

- **libmd5_go_nts__MD5_finish**(int and_flush):

  *Not Thread Safe*

  Closes md5 context. Returns md5 digest as hex-digits string.

  params:
  - and_flush: `int`: set internal hasher object to null so it need to be inited again. Values 0/1 accepted.
    *since*: 0.0.8

  return: `char *`:  String with hexed digest, if error occured - empty string.

  *since*: 0.0.7

- **libmd5_go_nts__MD5_finishDefault**(void):

  *Not Thread Safe*

  Closes md5 context. Returns md5 digest as hex-digits string.

  *alias*: libmd5_go_nts__MD5_finish(1)

  params:
  -none

  return: `char *`:  String with hexed digest, if error occured - empty string.

  *since*: 0.0.8

- **libmd5_go_ts__MD5_finish**(int and_flush)

  *Thread Safe*

  Closes md5 context. Returns md5 digest as hex-digits string.

  params:
  - and_flush: `int`: set internal hasher object to null so it need to be inited again. Values 0/1 accepted.
    *since*: 0.0.8

  return: `char *`:  String with hexed digest, if error occured - empty string.

- **libmd5_go_ts__MD5_finishDefault**(void):

  *Thread Safe*

  Closes md5 context. Returns md5 digest as hex-digits string.

  *alias*: libmd5_go_ts__MD5_finish(1)

  params:
  -none

  return: `char *`:  String with hexed digest, if error occured - empty string.

  *since*: 0.0.8

  *since*: 0.0.7

- **libmd5_go_ts__MD5_init**(void): void

  *Thread Safe*

 (Re)Initialize md5 context.

  params:
  - none

  return: none

  *since*: 0.0.7

- **libmd5_go_nts__MD5_update**(char* inputText): Returns 0/1

  *Not Thread Safe*

  params:
  - inputText (`char *`): input full path to file as string

  return: `int` 1 - if md5 updated; 0 - if not updated or some error happens

  *since*: 0.0.7

- **libmd5_go_ts__MD5_update**(char* inputText): Returns 0/1

  *Thread Safe*

  params:
  - inputText (`char *`): input full path to file as string

  return: `int` 1 - if md5 updated; 0 - if not updated or some error happens

  *since*: 0.0.7

- **libmd5_go_nts__MD5File_update**(char* fullPath): Returns 0/1

  *Not Thread Safe*

  params:
  - fullPath (`char *`): input full path to file as string

  return: `int` 1 - if md5 updated and file readable; 0 - if not updated or file not readable

  *since*: 0.0.7

- **libmd5_go_ts__MD5File_update**(char* fullPath): Returns 0/1

  *Thread Safe*

  params:
  - fullPath (`char *`): input full path to file as string

  return: `int` 1 - if md5 updated and file readable; 0 - if not updated or file not readable

  *since*: 0.0.7

- **libmd5_go_ts__MD5File_update**(char* fullPath): Returns 0/1

  *Thread Safe* possibly

  params:
  - fullPath (`char *`): input full path to file as string

  return: `int` 1 - if md5 updated and file readable; 0 - if not updated or file not readable

- **libmd5_go__MD5File_hexdigest**(char* fullPath): Return hexed-string with md5 digest of contents of the file.

  *Deprecated*

  params:
  - fullPath (`char *`): input full path to file as string

  return: `char *` String with hexed digest, if error occured - empty string returns

  *since*: 0.0.5

- **libmd5_go__MD5_digest**(char* text): Return non hexed(byte)-string with md5 digest.

  *Deprecated*

  params:
  - text (`char *`): input string

  return: `char *` Bytes array, char string with digest

  *since*: 0.0.2

- **libmd5_go__FreeResult**(char* ptr): frees memory allocated for md5 degest earlier.

  return: `void`

  *since*: 0.0.3
