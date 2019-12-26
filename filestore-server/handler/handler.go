package handler

import (
	"io"
	"io/ioutil"
	"net/http"
)

//UploadHandler：处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) { //函数外部调用，需要大写开头
	if r.Method == "GET" {
		//返回上传html页面
		data, err := ioutil.ReadFile("./statci/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error1111")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收用户上传文件流并存储到本地目录
	}
}
