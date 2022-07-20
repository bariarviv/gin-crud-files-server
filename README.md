# Gofiber - CRUD Server
![Coverage](https://img.shields.io/badge/Coverage-71.2%25-brightgreen)
![Go](https://github.com/bariarviv/.........../workflows/CI/badge.svg)

## Gin Golang
[Gin](https://github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter.

## Details of the application
Backend Golang application that has the following routes (using gin):
* PUT /files - uploads a file.
* GET /files/:fileName - to download the file.
* GET /files - returns a JSON array with the list of files stored.
* DELETE /files/:fileName - to delete the file.

## Running Steps
### Step 1 - Install Gin:



1. To install Gin package, you need to install Go and set your Go workspace first.
2. The first need Go installed (version 1.13+ is required), then you can use the below Go command to install Gin.
    ```
      go get -u github.com/gin-gonic/gin
    ```
Step 2 - Running locally:
To run the server, open a terminal, go to the project folder and use the command:
    ```
        go run .
    ```

[//]: # (Note: If you want to build a new docker image yourself &#40;Dockerfile file attached to the repository&#41; you will need to change the proxy back to: "proxy": "http://server-service:8080",.)


# Unit tests output

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