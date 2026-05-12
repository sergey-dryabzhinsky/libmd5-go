#include "libmd5-go.h"
#include <stdlib.h>
#include <stdio.h>

int main(int argc,char **argv){
  printf("libmd5-go version: %s\n", libmd5_go__version());
  printf("go runtime version used: %s\n", libmd5_go__version_go());
  char* dgst;
  printf("input args number:%d\n",argc);
if (argc>1){
  dgst = libmd5_go_ts__MD5File_hexdigest(argv[1]);
  printf("hex md5-file(%s)=%s\n",argv[1], dgst);
} else {
  dgst = libmd5_go_ts__MD5File_hexdigest("LICENSE");
  printf("hex md5-file(%s)=%s\n","LICENSE", dgst);
}

  libmd5_go__FreeResult(dgst);
  dgst = NULL; // Best practice to avoid dangling pointers
  return 0;
}