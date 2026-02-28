#include <libmd5-go.h>
#include <stdlib.h>
#include <stdio.h>

int main(int argc,char **argv){
  printf("Compiled libmd5-go version: %s\n", libmd5_go_Version);
  printf("Loaded libmd5-go version: %s\n", libmd5_go__version());
  printf("go runtime version used: %s\n", libmd5_go__version_go());
  return 0;
}