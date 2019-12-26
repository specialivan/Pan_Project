package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler) //路由规则
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("File to start server,err:%s", err.Error())
	}
}