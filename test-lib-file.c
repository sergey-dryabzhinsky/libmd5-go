#include "libmd5-go.h"
#include <stdlib.h>
#include <stdio.h>

int main(int argc,char **argv){
  printf("libmd5-go version: %s\n", libmd5_go__version());
  printf("go runtime version used: %s\n", libmd5_go__version_go());
  char* dgst;
  printf("input args number:%d\n",argc);
if (argc>1){
  libmd5_go_nts__MD5_init();
  printf("md5-file(%s): ...\n",argv[1]);
  int updated = libmd5_go_nts__MD5File_update(argv[1]);
  printf("md5-file(%s): updated:%d\n",argv[1], updated);
  dgst = libmd5_go_nts__MD5_finish();
  printf("md5-file(%s):%s\n",argv[1], dgst);
} else {
  libmd5_go_nts__MD5_init();
  printf("md5-file(LICENSE): ...\n");
  int updated = libmd5_go_nts__MD5File_update("LICENSE");
  printf("md5-file(%s): updated:%d\n","LICENSE", updated);
  dgst = libmd5_go_nts__MD5_finish();
  printf("md5-file(LICENSE):%s\n", dgst);
}

  libmd5_go_nts__MD5_init();
  printf("md5-file(non-existent): ...\n");
  int updated = libmd5_go_nts__MD5File_update("non-existent");
  printf("md5-file(%s): updated:%d\n","non-existent", updated);
  dgst = libmd5_go_nts__MD5_finish();
  printf("md5-file(non-existent):%s\n", dgst);
  libmd5_go__FreeResult(dgst);
  dgst = NULL; // Best practice to avoid dangling pointers
  return 0;
}