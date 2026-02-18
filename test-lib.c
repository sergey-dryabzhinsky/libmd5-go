#include <libmd5-go.h>
#include <stdlib.h>
#include <stdio.h>

int main(int argc,char **argv){
  printf("Compiled libmd5-go version: %s\n", libmd5_go_Version);
  printf("Loaded libmd5-go version: %s\n", libmd5_go__version());
  printf("go runtime version used: %s\n", libmd5_go__version_go());
  char* dgst;
  // Generic varant
  dgst = libmd5_go__MD5_hexdigest("123");
  printf("md5(123):%s\n", dgst);
  libmd5_go__FreeResult(dgst);
  // NoThread Safe variant
  libmd5_go_nts__MD5_init();
  int updated = *libmd5_go_nts__MD5_update("123");
  printf("md5 updated?:%d\n", updated);
  dgst = libmd5_go_nts__MD5_finish();
  printf("md5(123):%s\n", dgst);
  libmd5_go__FreeResult(dgst);
  dgst = NULL; // Best practice to avoid dangling pointers
  return 0;
}