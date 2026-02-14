#include "libmd5-go.h"
#include <stdlib.h>
#include <stdio.h>
#include <time.h>

int main(int argc,char **argv){
  printf("libmd5-go version: %s\n", libmd5_go__version());
  char* dgst;
  time_t now, curt;
  long wait=60; //test 60 sec
  long ops =0;
  time(&now);
  time(&curt);
  while( (long)curt - (long)now < wait) {
    time(&curt);
    dgst = libmd5_go__MD5_digest("123");
    libmd5_go__FreeResult(dgst);
    ops ++;
  }
  dgst = libmd5_go__MD5_hexdigest("123");
  printf("md5(123):%s\nOps:%ld; op/s:%6.3f\n", dgst, ops, ops/1.0/wait);
  libmd5_go__FreeResult(dgst);
  dgst = NULL; // Best practice to avoid dangling pointers
  return 0;
}