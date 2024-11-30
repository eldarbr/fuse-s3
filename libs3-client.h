#ifndef LIBS3_CLIENT_
#define LIBS3_CLIENT_

extern char* Auth(char* iam_url, char* username, char* password);
extern char** ListFiles(char* s3_url, char* token, char* bucket_name);

#endif
