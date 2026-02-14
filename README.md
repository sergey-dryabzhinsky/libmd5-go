# libmd5-go

Pure md5 function nothing else. Exported from golang runtime. And some examples of code.

## Requirements
To compile library and examples you will need to install
- c99 capable compiler (for now only gcc >=4.8 supported)
- openssl >=1.0 ( libcrypto )
- golang >=1.6

## Prepare system
Instructions for Debian/Ubuntu like systems:
```
apt install libssl-dev
apt install golang-go
```

## Compile

If your version of go not default:
```
GOROOT=/path/to/golang/installdir make all GO=go-x.y
```

## Reson
Not found library with only one function, so I tryed by myself.