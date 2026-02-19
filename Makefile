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
VERSION?=0.0.7
goVERSION?=$(shell $(GO) version | cut -d' ' -f3)
GOLDFLAGS?=-ldflags="-s -w" -ldflags "-X main.VERSION=$(VERSION)"
#VERSION=$(shell grep 'const VERSION' $(LIBNAME).go | cut -d= -f2|tr -d '"')
ifeq (1,$(DEBUG))
$(info $(LIBNAME) version:$(VERSION))
$(info go version:$(goVERSION))
endif
MACHINE?=$(shell uname -m)
ifeq (x86_64,$(MACHINE))
MACHINE=amd64
#$(info amd64)
endif
TARFLAGS?=-v --xz
TARNAME?=$(LIBNAME)-$(goVERSION)_$(MACHINE).tar.xz
$(info Tar name will be: $(TARNAME))
INSTALL_ROOT?=./tmp
DIST_DIR?=./dist/
INSTALL_VERBOSE=
ifeq (1,$(VERBOSE))
INSTALL_VERBOSE=-v
endif
PREFIX?=/usr/local
INCLUDES_DIR?=include/libmd5-go
MULTIARCH?=$(shell uname -m)-linux-gnu
LIBS_DIR?=lib/$(MULTIARCH)

vet:
	$(GO) vet $(LIBNAME).go

$(LIBNAME)$(LIBEXT): constants.h
	$(GO) build $(GOFLAGS) $(GOLDFLAGS) -o $(LIBNAME)$(LIBEXT) -buildmode=c-shared $(LIBNAME).go

$(ldLIBNAME).pc:
	m4 \
 -DVERSION=$(VERSION) \
 -DMULTIARCH=$(MULTIARCH) \
 -DPREFIX=$(PREFIX) \
 -DLIBS_DIR=$(LIBS_DIR) \
 -DINCLUDES_DIR=$(INCLUDES_DIR) \
	 $(ldLIBNAME).pc.in > $(ldLIBNAME).pc

constants.h:
	sed -e 's#VERSION#$(VERSION)#g' \
	constants.h.in > constants.h

lib: $(LIBNAME)$(LIBEXT) $(ldLIBNAME).pc constants.h

lib-link: lib
	test ! -e $(LIBNAME)$(LIBEXT).$(VERSION) && mv $(LIBNAME)$(LIBEXT) $(LIBNAME)$(LIBEXT).$(VERSION)
	test ! -e $(LIBNAME)$(LIBEXT) && ln -snf $(LIBNAME)$(LIBEXT).$(VERSION) $(LIBNAME)$(LIBEXT)
	touch lib-link

test-lib: lib-link
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-lib -I. -L. -l$(ldLIBNAME) test-lib.c

test-lib-speed: lib-link
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-lib-speed -I. -L. -l$(ldLIBNAME) test-lib-speed.c

test-lib-file: lib
	$(CC) $(CFLAGS) $(LDFLAGS) -o test-lib-file -I. -L. $(ldLIBNAME) test-lib-file.c

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
	rm -f  $(LIBNAME)$(LIBEXT)* $(LIBNAME).h constants.h $(ldLIBNAME).pc lib-link
	rm -f test-lib test-lib-speed test-crypto-speed test-lib-file
	rm -rf tmp dist

all: test_lib

install: lib lib-link
	install $(INSTALL_VERBOSE) -d $(INSTALL_ROOT)/$(PREFIX)/$(INCLUDES_DIR)
	install $(INSTALL_VERBOSE) $(LIBNAME).h $(INSTALL_ROOT)/$(PREFIX)/$(INCLUDES_DIR)
	install $(INSTALL_VERBOSE) -d $(INSTALL_ROOT)/$(PREFIX)/$(LIBS_DIR)/pkgconfig
	install $(INSTALL_VERBOSE) $(ldLIBNAME).pc $(INSTALL_ROOT)/$(PREFIX)/$(LIBS_DIR)/pkgconfig
	install $(INSTALL_VERBOSE) $(LIBNAME)$(LIBEXT)* $(INSTALL_ROOT)/$(PREFIX)/$(LIBS_DIR)

tar: install
	mkdir -p $(DIST_DIR)
	cd $(INSTALL_ROOT) && tar $(TARFLAGS) -cf ../$(DIST_DIR)/$(TARNAME) .
