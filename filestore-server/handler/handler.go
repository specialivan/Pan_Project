//处理上传接口
package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
		file,head,err:=r.FormFile("file")
		if err != nil{
			fmt.Println("Failed to get data,err:%s\n",err.Error())
			return
		}
		defer file.Close()
		newFile,err:=os.Create("d:\123"+head.Filename)
		if err != nil{
			fmt.Println("Faile to create file,err:%s\n",err.Error())
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
			if err != nil{
				fmt.Println("Failed to data into file,err:%s\n",err.Error())
				return
			}

		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}
}

//新建一个上传成功的跳转页面
func UploadSucHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w ,"Upload Success!")
}