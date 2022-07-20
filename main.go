package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

const (
	FieldName       = "file"
	Port            = ":3000"
	URL             = "/files"
	Dir             = "./files/"
	Param           = "fileName"
	CertFile        = "/etc/ssl/certs/ssl.crt"
	KeyFile         = "/etc/ssl/certs/ssl.key"
	URLWithFilename = "/files/:fileName"
)

func main() {
	if err := makeDirIfNotExists(Dir); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := setupRouter(CertFile, KeyFile, Port); err != nil {
		fmt.Println(err)
		return
	}
}

// makeDirIfNotExists checks if the destination directory exists, if not, creates it
func makeDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

// setupRouter setups the server and the routers according to the HTTP requests
func setupRouter(certFile, keyFile, port string) (*gin.Engine, error) {
	// Setups the server and the routers according to the HTTP requests
	router := gin.Default()
	router.GET(URL, ListFilesHandler)
	router.GET(URLWithFilename, DownloadFileHandler)
	router.POST(URL, UploadFileHandler)
	router.DELETE(URLWithFilename, DeleteFileHandler)

	// Creates tls certificate
	ln, err := createTLSCert(certFile, keyFile, port)
	if err != nil || ln == nil {
		return router, err
	}
	// Starts server with https/ssl enabled on http://localhost:Port
	log.Fatal(router.RunListener(*ln))
	return router, nil
}

// createTLSCert creates tls certificate
func createTLSCert(certFile, keyFile, port string) (*net.Listener, error) {
	// Creates tls certificate
	certs, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		fmt.Printf("Cannot load TLS certificate from certFile=%q, keyFile=%q: %s\n", certFile, keyFile, err)
		return nil, err
	}
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	// Creates custom listener
	ln, err := tls.Listen("tcp", port, &tls.Config{
		RootCAs:      rootCAs,
		Certificates: []tls.Certificate{certs},
	})
	if err != nil {
		return &ln, err
	}
	return &ln, nil
}

// ListFilesHandler returns a JSON array with the list of files stored on the server
func ListFilesHandler(ctx *gin.Context) {
	// Reads the directory named and returns a list of directory entries sorted by filename
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		ctx.String(fiber.StatusNotFound, "ReadDir() Error: "+err.Error())
		return
	}

	type File struct {
		FileName string `json:"filename"`
		Format   string `json:"format"`
		Size     int64  `json:"size"`
		ModDate  string `json:"moddate"`
	}
	// For each file - we will save the values according to the File struct
	var data []File
	for _, file := range files {
		searchFormat := strings.Split(file.Name(), ".")
		data = append(data, File{
			FileName: file.Name(),
			Format:   searchFormat[len(searchFormat)-1],
			Size:     file.Size(),
			ModDate:  file.ModTime().Format("2006-01-02"),
		})
	}
	ctx.JSON(http.StatusOK, data)
}

// DownloadFileHandler downloads the received file as a parameter from the URL
func DownloadFileHandler(ctx *gin.Context) {
	// Gets the parameter from the URL
	filename := ctx.Param(Param)
	path := filepath.Join(Dir, filename)
	// Downloads the file
	ctx.FileAttachment(path, path)
	ctx.String(http.StatusOK, filename+" uploaded successfully!")
}

// UploadFileHandler uploads a file to a specific destination
func UploadFileHandler(ctx *gin.Context) {
	file, err := ctx.FormFile(FieldName)
	if err != nil {
		ctx.String(fiber.StatusBadRequest, "ctx.FormFile() Error: "+err.Error())
		return
	}
	// Uploads the file
	if err = ctx.SaveUploadedFile(file, Dir+file.Filename); err != nil {
		ctx.String(fiber.StatusBadRequest, "ctx.SaveUploadedFile() Error: "+err.Error())
	}
	ctx.String(fiber.StatusOK, file.Filename+" uploaded successfully!")
}

// DeleteFileHandler deletes the received file as a parameter from the URL
func DeleteFileHandler(ctx *gin.Context) {
	// Gets the parameter from the URL
	filename := ctx.Param(Param)

	// Checks if the file exists - open & close the file
	file, err := os.Open(filepath.Join(Dir, filename))
	if err != nil {
		ctx.String(fiber.StatusNotFound, "os.Open() Error: "+err.Error())
	}
	defer file.Close()

	// Removes the file
	if err = os.Remove(Dir + filename); err != nil {
		ctx.String(fiber.StatusNotFound, "os.Remove() Error: "+err.Error())
	}
	ctx.String(fiber.StatusOK, filename+" deleted successfully!")
}
