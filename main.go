//程序入口
package main

import (
	"Pan_Project/filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("File to start server,err:%s", err.Error())
		}
}