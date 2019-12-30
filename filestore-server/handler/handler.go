//处理上传接口
package handler

import (
	"io"
	"io/ioutil"
	"net/http"
)

//UploadHandler：处理文件上传
//函数外部调用，需要大写开头
//ResponseWriter用于向用户返回数据，Request用于接收用户请求对象指针
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./filestore-server/static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method== "POST" {

	}
}
