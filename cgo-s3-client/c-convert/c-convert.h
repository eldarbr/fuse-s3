#ifndef S3_CGO_C_CONVERT
#define S3_CGO_C_CONVERT

char **names_buffer_alloc(int cnt_files, int total_chars);
void names_buffer_add(char **, int, char *);
void names_buffer_all_add(char **buff, int names_cnt, char *names_mono);

#endif
