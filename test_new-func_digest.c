#include "libmd5-go.h"
#include <stdlib.h>
#include <stdio.h>
#include <time.h>

int main(int argc,char **argv){
  printf("libmd5-go version: %s\n", libmd5_go__version());
  char* dgst;
    dgst = libmd5_go_ts__MD5_digest("123");
  printf("raw md5(123):%s\n", dgst);
  libmd5_go__FreeResult(dgst);
  dgst = NULL; // Best practice to avoid dangling pointers
  return 0;
}