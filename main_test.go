package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	TempFile = "temp.txt"
)

func Test_makeDirIfNotExists(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"Creates a dir successfully", Dir, false},
		{"Creates a test dir successfully", "./testfiles/", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := makeDirIfNotExists(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("makeDirIfNotExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createTLSCert(t *testing.T) {
	tests := []struct {
		name     string
		certFile string
		keyFile  string
		port     string
		wantErr  bool
	}{
		{"Creates TLS cert successfully", CertFile, KeyFile, Port, false},
		{"Failed to create due to empty certFile", "", KeyFile, Port, true},
		{"Failed to create due to empty keyFile", CertFile, "", Port, true},
		{"Failed to create due to empty port", CertFile, KeyFile, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ln, err := createTLSCert(tt.certFile, tt.keyFile, tt.port)
			if ln == nil && ((err != nil) != tt.wantErr) {
				t.Errorf("createTLSCert() ln = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListFilesHandler(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"Gets file lists successfully (if there are files in the folder)", URL},
		{"Failed to get file lists due to incorrect URL", "/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creates gin router & http test
			respRecorder, router, _ := createRouterAndWriter(t)
			router.GET(URL, ListFilesHandler)
			// Creates a request
			request, err := http.NewRequest(http.MethodGet, tt.url, nil)
			if err != nil {
				t.Errorf("http.NewRequest() Error: %v", err)
			}
			router.ServeHTTP(respRecorder, request)
		})
	}
}

func TestDownloadFileHandler(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		wantCode int
	}{
		{"Downloads temp file successfully", URL, http.StatusOK},
		{"Downloads fail due to incorrect URL", "/", http.StatusNotFound},
		{"Downloads fail due to the file doesn't exist", URL, http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creates gin router & http test
			respRecorder, router, ctx := createRouterAndWriter(t)
			router.GET(URLWithFilename, DownloadFileHandler)
			filename, err := createTempFile(Dir, "downloads", tt.wantCode)
			if err != nil {
				t.Errorf("createTempFile() Error: " + err.Error())
			}
			// Sets context params - filename & performs the request
			ctx.Params = []gin.Param{{Key: Param, Value: filename}}
			request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", tt.url, filename), nil)
			if err != nil {
				t.Errorf("http.NewRequest() Error: " + err.Error())
			}
			router.ServeHTTP(respRecorder, request)
			assert.Equal(t, respRecorder.Code, tt.wantCode)
			// Deletes the temp file
			if err = os.Remove(Dir + filename); err != nil && tt.wantCode == http.StatusOK {
				t.Error(http.StatusInternalServerError, err)
			}
		})
	}
}

func TestUploadFileHandler(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		fieldName string
		url       string
		wantCode  int
	}{
		{"Uploads temp file successfully", TempFile, FieldName, URL, http.StatusOK},
		{"Uploads fail due to incorrect URL", TempFile, FieldName, "/", http.StatusNotFound},
		{"Uploads fail due to incorrect field name", TempFile, "image", URL, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creates gin router & http test
			respRecorder, router, _ := createRouterAndWriter(t)
			router.POST(URL, UploadFileHandler)
			// Performs the request
			request, err := newFileUploadRequest(tt.filename, tt.fieldName, tt.url)
			if err != nil {
				t.Errorf("newFileUploadRequest() Error: " + err.Error())
			}
			router.ServeHTTP(respRecorder, request)
			assert.Equal(t, respRecorder.Code, tt.wantCode)
			// Deletes the temp file
			if err = os.Remove(Dir + tt.filename); err != nil {
				t.Error(http.StatusInternalServerError, err)
			}
		})
	}
}

func TestDeleteFileHandler(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		wantCode int
	}{
		{"Deletes temp file successfully", URL, http.StatusOK},
		{"Deletes fail due to incorrect URL", "/", http.StatusNotFound},
		{"Deletes fail due to the file doesn't exist", URL, http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creates gin router & http test
			respRecorder, router, ctx := createRouterAndWriter(t)
			router.DELETE(URLWithFilename, DeleteFileHandler)
			filename, err := createTempFile(Dir, "downloads", tt.wantCode)
			if err != nil {
				t.Errorf("createTempFile() Error: " + err.Error())
			}
			// Sets context params - filename & performs the request
			ctx.Params = []gin.Param{{Key: Param, Value: filename}}
			request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", tt.url, filename), nil)
			if err != nil {
				t.Errorf("http.NewRequest() Error: " + err.Error())
			}
			router.ServeHTTP(respRecorder, request)
			assert.Equal(t, respRecorder.Code, tt.wantCode)
		})
	}
}

func Test_setupRouter(t *testing.T) {
	go func() {
		if _, err := setupRouter(CertFile, KeyFile, ":3001"); err != nil {
			t.Errorf("setupRouter() error = %v", err)
		}
	}()
}

// createRouterAndWriter creates gin router & http test record
func createRouterAndWriter(test *testing.T) (*httptest.ResponseRecorder, *gin.Engine, *gin.Context) {
	// Runs a test create the storage folder to verify that the test is performed without error
	Test_makeDirIfNotExists(test)
	// Creates gin router & http test record
	gin.SetMode(gin.TestMode)
	respRecorder := httptest.NewRecorder()
	ctx, router := gin.CreateTestContext(respRecorder)
	return respRecorder, router, ctx
}

// createTempFile creates a random temporary file in the received path and returns its name
func createTempFile(path, pattern string, statusCode int) (string, error) {
	if statusCode != http.StatusOK {
		return "", nil
	}
	if path == "" {
		path = os.TempDir()
	}
	tempFile, err := ioutil.TempFile(path, pattern)
	if err != nil {
		return "", fmt.Errorf("ioutil.TempFile() Error: %v", err)
	}
	if _, err = tempFile.WriteString("Text to append\n"); err != nil {
		return "", fmt.Errorf("File.WriteString() Error: %v", err)
	}
	return strings.Split(tempFile.Name(), "/")[2], nil
}

// newFileUploadRequest creates a file and transfers it using FormFile and creates a request for the Post function
func newFileUploadRequest(filename, fieldName, url string) (*http.Request, error) {
	// Creates file
	file, err := os.Create(Dir + filename)
	if err != nil {
		return nil, fmt.Errorf("os.Create() Error: %v", err)
	}
	defer file.Close()

	// Creates a temp file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Transfers it using FormFile
	part, err := writer.CreateFormFile(fieldName, filename)
	if err != nil {
		return nil, fmt.Errorf("writer.CreateFormFile() Error: %v", err)
	}
	if _, err = io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("io.Copy Error: %v", err)
	}
	defer writer.Close()

	// Creates a request
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return request, fmt.Errorf("http.NewRequest() Error: %v", err)
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request, nil
}
