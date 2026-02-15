#include "libmd5-go.h"
#include <stdlib.h>
#include <stdio.h>

int main(int argc,char **argv){
  printf("libmd5-go version: %s\n", libmd5_go__version());
  printf("go runtime version used: %s\n", libmd5_go__version_go());
  char* dgst;
  printf("md5-file(LICENSE): ...\n");
  dgst = libmd5_go__MD5File_hexdigest("LICENSE");
  printf("md5-file(LICENSE):%s\n", dgst);

  printf("md5-file(non-existent): ...\n");
  dgst = libmd5_go__MD5File_hexdigest("non-existent");
  printf("md5-file(non-existent):%s\n", dgst);
  libmd5_go__FreeResult(dgst);
  dgst = NULL; // Best practice to avoid dangling pointers
  return 0;
}