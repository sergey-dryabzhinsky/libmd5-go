#include <libmd5-go.h>
#include <stdlib.h>
#include <stdio.h>

int main(int argc,char **argv){
  printf("Compiled libmd5-go version: %s\n", libmd5_go_Version);
  printf("Loaded libmd5-go version: %s\n", libmd5_go__version());
  printf("go runtime version used: %s\n", libmd5_go__version_go());
  char* dgst;
  // NoThread Safe variant
  libmd5_go_nts__MD5_init();
  int updated = libmd5_go_nts__MD5_update("123");
  printf("md5 updated?:%d\n", updated);
  dgst = libmd5_go_nts__MD5_finish(1);
  printf("md5(123):%s\n", dgst);
  libmd5_go__FreeResult(dgst);
  dgst = NULL; // Best practice to avoid dangling pointers
  int errno = libmd5_go_nts__getLastErrorCode();
  printf("last error code:%d\n", errno);
  updated = libmd5_go_nts__MD5_update("123");
  printf("md5 updated?:%d\n", updated);
  if (updated)
	printf("Not inited update test FAIL\n");
  else
	printf("Not inited update test SUCCESS\n");
  errno = libmd5_go_nts__getLastErrorCode();
  printf("last error code:%d\n", errno);
  printf("last error:%s\n", libmd5_go_nts__getErrorDescription(errno));
  return 0;
}