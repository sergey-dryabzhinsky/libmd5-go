#include <stdlib.h>
#include <stdio.h>
#include <time.h>
#include <string.h>
#include <openssl/md5.h>

int main() {
    unsigned char digest[MD5_DIGEST_LENGTH];
    const char *string = "123";
    MD5_CTX context;
    int i;
  time_t now, curt;
  long wait=60; //test 60 sec
  long ops =0;
  time(&now);
  time(&curt);

    while( (long)curt - (long)now < wait) {
    time(&curt);
//    dgst = libmd5_go__MD5("123");
//  dgst = MD5("123");

    // Initialize the MD5 context
    MD5_Init(&context);

    // Update the hash with the string data
    MD5_Update(&context, string, strlen(string));

    // Finalize the hash calculation
    MD5_Final(digest, &context);

    ops ++;
  }
    printf(" md5(%s):", string);
    // Print the digest in hexadecimal format
    for (i = 0; i < MD5_DIGEST_LENGTH; ++i) {
        printf("%02x", digest[i]);
    }
  printf("\nOps:%ld; op/s:%6.3f\n", ops, ops/1.0/wait);

    return 0;
}
