#include "c-convert.h"

#include <stdlib.h>
#include <string.h>

char **names_buffer_alloc(int filenames, int total_names_length) {
  /*
   * The Buffer structure:
   * ptrs to the filenames beginning
   * nullptr
   * filename\0
   * filename\0
   * ...
   *
   *
   * The structure requires this ammount of memory:
   * sizeof(char*) * filenames <- for the pointers
   * 1 nullptr
   * sizeof(char) * (total_names_length + filenames) <- for the names
   *                                                    including nullterm
   */

  const size_t total_size = sizeof(char *) * (filenames + 1) +
                            sizeof(char) * (total_names_length + filenames);

  char **buff = malloc(total_size);
  if (buff == NULL) {
    return NULL;
  }

  memset(buff, 0, total_size);

  return buff;
}

void names_buffer_add(char **buff, int total_names_cap, char *new_name) {
  char *new_content_start = NULL;

  if (*buff) {       // if anythin is in the buffer
    while (*buff) {  // skip it.
      ++(buff);
    }
    new_content_start = *(buff - 1) + strlen(*(buff - 1)) + 1;
  } else {
    new_content_start = (char *)(buff + total_names_cap + 1);
  }

  strcpy(new_content_start, new_name);
  *buff = new_content_start;
}

void names_buffer_all_add(char **buff, int names_cnt, char *names_mono) {
  char *new_content_start = (char *)(buff + names_cnt + 1);

  for (int i = 0; i < names_cnt; ++i) {
    strcpy(new_content_start, names_mono);
    *buff = new_content_start;
    ++buff;
    size_t written_name_bytes = strlen(names_mono) + 1;
    new_content_start += written_name_bytes;
    names_mono += written_name_bytes;
  }
}
