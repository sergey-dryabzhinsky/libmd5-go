#include "libmd5-go.h"
#include <stdlib.h>
#include <stdio.h>

int main(int argc,char **argv){
  printf("libmd5-go version: %s\n", libmd5_go__version());
  printf("go runtime version used: %s\n", libmd5_go__version_go());
  printf("input args number:%d\n",argc);

  libmd5_go_nts__MD5_init();
  printf("md5-file(non-existent): ...\n");
  int updated = libmd5_go_nts__MD5File_update("non-existent");
  printf("md5-file(%s): updated:%d\n","non-existent", updated);
  if (updated) {
    printf("test FAILED\n");
    return 1;
  } else {
    printf("test SUCCESS\n");
    int errno = libmd5_go_nts__getLastErrorCode();
    printf("last error code: %d\n", errno);
    printf("last error desc: %s\n", libmd5_go_nts__getErrorDescription(errno));
  }
  return 0;
}