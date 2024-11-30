package main

// #cgo LDFLAGS: ${SRCDIR}/c-convert/c-convert.o
// #include "c-convert/c-convert.h"
// #include <stdlib.h>
import "C"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"unsafe"

	"github.com/eldarbr/fuse-s3/cgo-s3-client/model"
)

const (
	defaultClienetTimeout = 20 * time.Second
)

const (
	pathIAMAuth   = "/auth/authenticate"
	pathListFiles = "/api/manage/buckets/%s/files"
)

//export Auth
func Auth(iamURL, username, password *C.char) *C.char {
	if iamURL == nil || username == nil || password == nil {
		return nil
	}

	var reqBody *bytes.Buffer
	{
		bodySlice, encodeErr := json.Marshal(
			model.AuthRequestBody{
				Username: C.GoString(username),
				Password: C.GoString(password),
			},
		)
		if encodeErr != nil {
			return nil
		}

		reqBody = bytes.NewBuffer(bodySlice)
	}

	var client http.Client

	client.Timeout = defaultClienetTimeout

	request, requestErr := http.NewRequest(http.MethodPost, C.GoString(iamURL)+pathIAMAuth, reqBody)
	if requestErr != nil {
		return nil
	}

	response, fetchErr := client.Do(request)
	if fetchErr != nil || response.StatusCode != http.StatusOK {
		return nil
	}

	defer response.Body.Close()

	var responseBody model.AuthResponseBody

	decodeErr := json.NewDecoder(response.Body).Decode(&responseBody)
	if decodeErr != nil {
		return nil
	}

	return C.CString(responseBody.Token)
}

//export ListFiles
func ListFiles(s3URL, token, bucketName *C.char) **C.char {
	if bucketName == nil || token == nil {
		return nil
	}

	var client http.Client

	client.Timeout = defaultClienetTimeout

	request, requestErr := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s"+pathListFiles, C.GoString(s3URL), C.GoString(bucketName)), nil)
	if requestErr != nil {
		return nil
	}

	request.Header.Set("Authorization", C.GoString(token))

	response, fetchErr := client.Do(request)
	if fetchErr != nil || response.StatusCode != http.StatusOK {
		return nil
	}

	defer response.Body.Close()

	var responseBody model.ListFilesResponseBody

	decodeErr := json.NewDecoder(response.Body).Decode(&responseBody)
	if decodeErr != nil {
		return nil
	}

	namesCnt := len(responseBody.Files)

	var (
		builder       strings.Builder
		namesTotalLen int
	)

	for i := range namesCnt {
		bytesWritten, _ := builder.WriteString(responseBody.Files[i].Filename)
		namesTotalLen += bytesWritten
		_ = builder.WriteByte(0)
	}

	buff := C.names_buffer_alloc(C.int(namesCnt), C.int(namesTotalLen))

	C.names_buffer_all_add(buff, C.int(namesCnt), (*C.char)(unsafe.Pointer(unsafe.StringData(builder.String()))))

	return buff
}

func main() {}
