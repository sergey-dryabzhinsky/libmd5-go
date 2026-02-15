.PHONY: all lib test_lib clean install

GO?=go
CC?=gcc
#linux
CFLAGS?=-O2
LDFLAGS?=-Wl,-s
GOLDFLAGS?=-ldflags="-s -w" -trimpath
LIBEXT?=.so
LIBNAME=libmd5-go
VERSION=0.0.5
#VERSION=$(shell grep 'const VERSION' $(LIBNAME).go | cut -d= -f2|tr -d '"')

INSTALL_ROOT?=/
PREFIX?=/usr/local
INCLUDES_DIR?=/include
LIBS_DIR?=/lib

$(LIBNAME)$(LIBEXT):
	$(GO) build -v -a $(GOLDFLAGS) -o $(LIBNAME)$(LIBEXT) -buildmode=c-shared $(LIBNAME).go

lib: $(LIBNAME)$(LIBEXT)
	mv $(LIBNAME)$(LIBEXT) $(LIBNAME)$(LIBEXT).$(VERSION)
	ln -snf $(LIBNAME)$(LIBEXT).$(VERSION) $(LIBNAME)$(LIBEXT)

test_lib: lib
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-lib  ./$(LIBNAME)$(LIBEXT) test-lib.c
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-lib-speed  ./$(LIBNAME)$(LIBEXT) test-lib-speed.c
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-lib-file  ./$(LIBNAME)$(LIBEXT) test-lib-file.c
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-crypto-speed  -lcrypto test-crypto-speed.c

clean:
	rm -f  $(LIBNAME)$(LIBEXT)* $(LIBNAME).h
	rm -f test-lib test-lib-speed test-crypto-speed test-lib-file

all: test_lib