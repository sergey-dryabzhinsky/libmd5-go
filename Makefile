#!/usr/bin/make -f
.PHONY: all lib test_lib clean install
VERBOSE?=1
DEBUG?=1

GO?=go
CC?=gcc
#linux
CFLAGS?=-O2
GOFLAGS?=-a
ifeq (1,$(DEBUG))
CFLAGS+=-Wall -Werror
endif
ifeq (1,$(VERBOSE))
CFLAGS+=-v
GOFLAGS+=-v
endif
LDFLAGS?=-Wl,-s
LIBEXT?=.so
LIBNAME=libmd5-go
ldLIBNAME=md5-go
VERSION?=0.0.6
GOLDFLAGS?=-ldflags="-s -w" -ldflags "-X main.VERSION=$(VERSION)"
#VERSION=$(shell grep 'const VERSION' $(LIBNAME).go | cut -d= -f2|tr -d '"')
ifeq (1,$(DEBUG))
$(info libmd5-go version:$(VERSION))
endif

INSTALL_ROOT?=./tmp/
PREFIX?=/usr/local
INCLUDES_DIR?=include
MULTIARCH?=$(shell uname -m)-linux-gnu
LIBS_DIR?=lib/$(MULTIARCH)

vet:
	$(GO) vet $(LIBNAME).go

$(LIBNAME)$(LIBEXT):
	$(GO) build $(GOFLAGS) $(GOLDFLAGS) -o $(LIBNAME)$(LIBEXT) -buildmode=c-shared $(LIBNAME).go

$(ldLIBNAME).pc:
	m4 \
 -DVERSION=$(VERSION) \
 -DMULTIARCH=$(MULTIARCH) \
 -DPREFIX=$(PREFIX) \
 -DLIBS_DIR=$(LIBS_DIR) \
 -DINCLUDES_DIR=$(INCLUDES_DIR) \
	 $(ldLIBNAME).pc.in > $(ldLIBNAME).pc

lib: $(LIBNAME)$(LIBEXT) $(ldLIBNAME).pc

lib-link: lib
	test ! -r $(LIBNAME)$(LIBEXT).$(VERSION) && mv $(LIBNAME)$(LIBEXT) $(LIBNAME)$(LIBEXT).$(VERSION)
	test ! -e $(LIBNAME)$(LIBEXT) && ln -snf $(LIBNAME)$(LIBEXT).$(VERSION) $(LIBNAME)$(LIBEXT)
	touch lib-link

test-lib: lib-link
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-lib -L. -l$(ldLIBNAME) test-lib.c

test-lib-speed: lib-link
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-lib-speed -L. -l$(ldLIBNAME) test-lib-speed.c

test-lib-file: lib
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-lib-file -L. $(ldLIBNAME) test-lib-file.c

test-crypto-speed: lib
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-crypto-speed  -lcrypto test-crypto-speed.c

tests: \
 test-lib \
 test-lib-speed \
 test-lib-file \
 test-crypto-speed
	 export LD_LIBRARY_PATH=.
	./test-lib
	./test-lib-file
	md5sum LICENSE
	./test-lib-speed
	./test-crypto-speed

clean:
	rm -f  $(LIBNAME)$(LIBEXT)* $(LIBNAME).h $(ldLIBNAME).pc lib-link
	rm -f test-lib test-lib-speed test-crypto-speed test-lib-file

all: test_lib