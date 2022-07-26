# GoGin - CRUD Server
![Coverage](https://img.shields.io/badge/Coverage-71.2%25-brightgreen)

## Gin Golang
[Gin](https://github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter

### Advantages:
1. Simple code - can save the developer a lot of time in large projects
2. High performance
3. Error management
4. Easy JSON authentication
5. Gin has a mode test

### Disadvantages:
1. Not flexible in development
2. Can add more complexity to your project and slow down the development time, the infrastructure now depends on the package that other people maintain

### Alternative Solutions:
1. [net/http:](https://github.com/golang/go) it's easier to use and can handle more cases
2. [fasthttp:](https://github.com/valyala/fasthttp) was designed for some high-performance edge cases
3. [echo:](https://github.com/labstack/echo) supports HTTP/2 for faster performance and an overall better user experience and has automatic TLS certificates
4. [fiber:](https://github.com/gofiber/fiber) built on top of the Fasthttp HTTP engine for Go, which is the fastest HTTP engine for Go


## Details of the application
Backend Golang application that has the following routes (using gin):
* ***PUT /files -*** uploads a file
* ***GET /files/:fileName -*** to download the file
* ***GET /files -*** returns a JSON array with the list of files stored
* ***DELETE /files/:fileName -*** to delete the file


## Running Steps
### Step 1 - Build the docker image:
```
docker build -t ginserver .
```

### Step 2 - Run the docker image:
```
docker run -it --rm -p 3000:3000 ginserver
```

#### For example:
```
[baria@ ~]$ sudo docker run -it --rm -p 3000:3000 ginserver 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /files                    --> main.ListFilesHandler (3 handlers)
[GIN-debug] GET    /files/:fileName          --> main.DownloadFileHandler (3 handlers)
[GIN-debug] POST   /files                    --> main.UploadFileHandler (3 handlers)
[GIN-debug] DELETE /files/:fileName          --> main.DeleteFileHandler (3 handlers)
[GIN-debug] Listening and serving HTTP on listener what's bind with address@[::]:3000
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN] 2022/07/20 - 04:13:34 | 200 |     237.023µs |      172.17.0.1 | GET      "/files"
```


## Unit Tests Output
```bash
=== RUN   Test_makeDirIfNotExists
=== RUN   Test_makeDirIfNotExists/Creates_a_dir_successfully
=== RUN   Test_makeDirIfNotExists/Creates_a_test_dir_successfully
--- PASS: Test_makeDirIfNotExists (0.00s)
    --- PASS: Test_makeDirIfNotExists/Creates_a_dir_successfully (0.00s)
    --- PASS: Test_makeDirIfNotExists/Creates_a_test_dir_successfully (0.00s)
=== RUN   Test_createTLSCert
=== RUN   Test_createTLSCert/Creates_TLS_cert_successfully
=== RUN   Test_createTLSCert/Failed_to_create_due_to_empty_certFile
Cannot load TLS certificate from certFile="", keyFile="certs/ssl.key": open : no such file or directory
=== RUN   Test_createTLSCert/Failed_to_create_due_to_empty_keyFile
Cannot load TLS certificate from certFile="certs/ssl.crt", keyFile="": open : no such file or directory
=== RUN   Test_createTLSCert/Failed_to_create_due_to_empty_port
--- PASS: Test_createTLSCert (0.00s)
    --- PASS: Test_createTLSCert/Creates_TLS_cert_successfully (0.00s)
    --- PASS: Test_createTLSCert/Failed_to_create_due_to_empty_certFile (0.00s)
    --- PASS: Test_createTLSCert/Failed_to_create_due_to_empty_keyFile (0.00s)
    --- PASS: Test_createTLSCert/Failed_to_create_due_to_empty_port (0.00s)
=== RUN   TestListFilesHandler
=== RUN   TestListFilesHandler/#00
=== RUN   TestListFilesHandler/#00/Creates_a_dir_successfully
=== RUN   TestListFilesHandler/#00/Creates_a_test_dir_successfully
=== RUN   TestListFilesHandler/#01
=== RUN   TestListFilesHandler/#01/Creates_a_dir_successfully
=== RUN   TestListFilesHandler/#01/Creates_a_test_dir_successfully
--- PASS: TestListFilesHandler (0.00s)
    --- PASS: TestListFilesHandler/#00 (0.00s)
        --- PASS: TestListFilesHandler/#00/Creates_a_dir_successfully (0.00s)
        --- PASS: TestListFilesHandler/#00/Creates_a_test_dir_successfully (0.00s)
    --- PASS: TestListFilesHandler/#01 (0.00s)
        --- PASS: TestListFilesHandler/#01/Creates_a_dir_successfully (0.00s)
        --- PASS: TestListFilesHandler/#01/Creates_a_test_dir_successfully (0.00s)
=== RUN   TestDownloadFileHandler
=== RUN   TestDownloadFileHandler/Downloads_temp_file_successfully
=== RUN   TestDownloadFileHandler/Downloads_temp_file_successfully/Creates_a_dir_successfully
=== RUN   TestDownloadFileHandler/Downloads_temp_file_successfully/Creates_a_test_dir_successfully
=== RUN   TestDownloadFileHandler/Downloads_fail_due_to_incorrect_URL
=== RUN   TestDownloadFileHandler/Downloads_fail_due_to_incorrect_URL/Creates_a_dir_successfully
=== RUN   TestDownloadFileHandler/Downloads_fail_due_to_incorrect_URL/Creates_a_test_dir_successfully
=== RUN   TestDownloadFileHandler/Downloads_fail_due_to_the_file_doesn't_exist
=== RUN   TestDownloadFileHandler/Downloads_fail_due_to_the_file_doesn't_exist/Creates_a_dir_successfully
=== RUN   TestDownloadFileHandler/Downloads_fail_due_to_the_file_doesn't_exist/Creates_a_test_dir_successfully
--- PASS: TestDownloadFileHandler (0.01s)
    --- PASS: TestDownloadFileHandler/Downloads_temp_file_successfully (0.00s)
        --- PASS: TestDownloadFileHandler/Downloads_temp_file_successfully/Creates_a_dir_successfully (0.00s)
        --- PASS: TestDownloadFileHandler/Downloads_temp_file_successfully/Creates_a_test_dir_successfully (0.00s)
    --- PASS: TestDownloadFileHandler/Downloads_fail_due_to_incorrect_URL (0.00s)
        --- PASS: TestDownloadFileHandler/Downloads_fail_due_to_incorrect_URL/Creates_a_dir_successfully (0.00s)
        --- PASS: TestDownloadFileHandler/Downloads_fail_due_to_incorrect_URL/Creates_a_test_dir_successfully (0.00s)
    --- PASS: TestDownloadFileHandler/Downloads_fail_due_to_the_file_doesn't_exist (0.00s)
        --- PASS: TestDownloadFileHandler/Downloads_fail_due_to_the_file_doesn't_exist/Creates_a_dir_successfully (0.00s)
        --- PASS: TestDownloadFileHandler/Downloads_fail_due_to_the_file_doesn't_exist/Creates_a_test_dir_successfully (0.00s)
=== RUN   TestUploadFileHandler
=== RUN   TestUploadFileHandler/Uploads_temp_file_successfully
=== RUN   TestUploadFileHandler/Uploads_temp_file_successfully/Creates_a_dir_successfully
=== RUN   TestUploadFileHandler/Uploads_temp_file_successfully/Creates_a_test_dir_successfully
=== RUN   TestUploadFileHandler/Uploads_fail_due_to_incorrect_URL
=== RUN   TestUploadFileHandler/Uploads_fail_due_to_incorrect_URL/Creates_a_dir_successfully
=== RUN   TestUploadFileHandler/Uploads_fail_due_to_incorrect_URL/Creates_a_test_dir_successfully
=== RUN   TestUploadFileHandler/Uploads_fail_due_to_incorrect_field_name
=== RUN   TestUploadFileHandler/Uploads_fail_due_to_incorrect_field_name/Creates_a_dir_successfully
=== RUN   TestUploadFileHandler/Uploads_fail_due_to_incorrect_field_name/Creates_a_test_dir_successfully
--- PASS: TestUploadFileHandler (0.00s)
    --- PASS: TestUploadFileHandler/Uploads_temp_file_successfully (0.00s)
        --- PASS: TestUploadFileHandler/Uploads_temp_file_successfully/Creates_a_dir_successfully (0.00s)
        --- PASS: TestUploadFileHandler/Uploads_temp_file_successfully/Creates_a_test_dir_successfully (0.00s)
    --- PASS: TestUploadFileHandler/Uploads_fail_due_to_incorrect_URL (0.00s)
        --- PASS: TestUploadFileHandler/Uploads_fail_due_to_incorrect_URL/Creates_a_dir_successfully (0.00s)
        --- PASS: TestUploadFileHandler/Uploads_fail_due_to_incorrect_URL/Creates_a_test_dir_successfully (0.00s)
    --- PASS: TestUploadFileHandler/Uploads_fail_due_to_incorrect_field_name (0.00s)
        --- PASS: TestUploadFileHandler/Uploads_fail_due_to_incorrect_field_name/Creates_a_dir_successfully (0.00s)
        --- PASS: TestUploadFileHandler/Uploads_fail_due_to_incorrect_field_name/Creates_a_test_dir_successfully (0.00s)
=== RUN   TestDeleteFileHandler
=== RUN   TestDeleteFileHandler/Deletes_temp_file_successfully
=== RUN   TestDeleteFileHandler/Deletes_temp_file_successfully/Creates_a_dir_successfully
=== RUN   TestDeleteFileHandler/Deletes_temp_file_successfully/Creates_a_test_dir_successfully
=== RUN   TestDeleteFileHandler/Deletes_fail_due_to_incorrect_URL
=== RUN   TestDeleteFileHandler/Deletes_fail_due_to_incorrect_URL/Creates_a_dir_successfully
=== RUN   TestDeleteFileHandler/Deletes_fail_due_to_incorrect_URL/Creates_a_test_dir_successfully
=== RUN   TestDeleteFileHandler/Deletes_fail_due_to_the_file_doesn't_exist
=== RUN   TestDeleteFileHandler/Deletes_fail_due_to_the_file_doesn't_exist/Creates_a_dir_successfully
=== RUN   TestDeleteFileHandler/Deletes_fail_due_to_the_file_doesn't_exist/Creates_a_test_dir_successfully
--- PASS: TestDeleteFileHandler (0.00s)
    --- PASS: TestDeleteFileHandler/Deletes_temp_file_successfully (0.00s)
        --- PASS: TestDeleteFileHandler/Deletes_temp_file_successfully/Creates_a_dir_successfully (0.00s)
        --- PASS: TestDeleteFileHandler/Deletes_temp_file_successfully/Creates_a_test_dir_successfully (0.00s)
    --- PASS: TestDeleteFileHandler/Deletes_fail_due_to_incorrect_URL (0.00s)
        --- PASS: TestDeleteFileHandler/Deletes_fail_due_to_incorrect_URL/Creates_a_dir_successfully (0.00s)
        --- PASS: TestDeleteFileHandler/Deletes_fail_due_to_incorrect_URL/Creates_a_test_dir_successfully (0.00s)
    --- PASS: TestDeleteFileHandler/Deletes_fail_due_to_the_file_doesn't_exist (0.00s)
        --- PASS: TestDeleteFileHandler/Deletes_fail_due_to_the_file_doesn't_exist/Creates_a_dir_successfully (0.00s)
        --- PASS: TestDeleteFileHandler/Deletes_fail_due_to_the_file_doesn't_exist/Creates_a_test_dir_successfully (0.00s)
=== RUN   Test_setupRouter
--- PASS: Test_setupRouter (0.00s)
PASS

coverage: 71.2% of statements in ./...
```
