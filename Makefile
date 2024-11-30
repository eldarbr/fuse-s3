S3_CLIENT_IMPL_PATH=cgo-s3-client

CCDEFS = -D_FILE_OFFSET_BITS=64

all:

.PHONY: build
build: libs3-client
	mkdir -p bin
	gcc $(CCDEFS) -Wall -Werror -Wextra --std=c11 -o bin/s3 main.c -I./$(S3_CLIENT_IMPL_PATH) -L./$(S3_CLIENT_IMPL_PATH) -ls3-client

.PHONY: libs3-client
libs3-client:
	make -C $(S3_CLIENT_IMPL_PATH) build

.PHONY: clean
clean:
	make -C $(S3_CLIENT_IMPL_PATH) clean
	rm -rf bin/

.PHONY: fmt
fmt:
	find . \( -name "*.c" -o -name "*.h" \) -print0 | xargs -0 clang-format -i --verbose --style=Google
	make -C $(S3_CLIENT_IMPL_PATH) fmt
