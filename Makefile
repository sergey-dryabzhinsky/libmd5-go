.PHONY: all lib test_lib

GO?=go
CC?=gcc
#linux
LIBEXT?=.so
LIBNAME=libmd5-go
VERSION=0.0.1

lib:
	$(GO) build -o $(LIBNAME)$(LIBEXT) -buildmode=c-shared $(LIBNAME).go

test_lib: lib
	$(CC) -o test-lib  ./$(LIBNAME)$(LIBEXT) test-lib.c
	$(CC) -o test-lib-speed  ./$(LIBNAME)$(LIBEXT) test-lib-speed.c
	$(CC) -o test-crypto-speed  -lcrypto test-crypto-speed.c

clean:
	rm -f  $(LIBNAME)$(LIBEXT) $(LIBNAME).h test-lib test-lib-speed test-crypto-speed

all: test_lib