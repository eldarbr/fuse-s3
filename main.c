#define FUSE_USE_VERSION 31
#include <fuse.h>
#include <stdio.h>
#include <stdlib.h>

#include "libs3-client.h"

int main(int argc, char **argv) {
  if (argc < 6) return 1;

  char *token = Auth(argv[1], argv[2], argv[3]);

  if (token == NULL) {
    return 1;
  }

  char **filenames = ListFiles(argv[4], token, argv[5]);
  if (filenames != NULL) {
    for (int i = 0; filenames[i] != NULL; ++i) {
      fputs(filenames[i], stdout);
      fputc('\n', stdout);
    }

    free(filenames);
  }

  free(token);

  return 0;
}
