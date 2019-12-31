//处理上传接口
package handler

import (
	"Pan_Project/filestore-server/meta"
	"Pan_Project/filestore-server/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

		fileMeta := meta.FileMeta{
			FileName:head.Filename,
			Location:"d:\123"+head.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile,err:=os.Create(fileMeta.Location)
		if err != nil{
			fmt.Println("Faile to create file,err:%s\n",err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
			if err != nil{
				fmt.Println("Failed to data into file,err:%s\n",err.Error())
				return
			}
		newFile.Seek(0,0)
		fileMeta.FileSha1 =util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)
		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}
}

//新建一个上传成功的跳转页面
func UploadSucHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w ,"Upload Success!")
}

// GetFileMetaHandler:获取文件源信息
func GetFileMetaHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data,err := json.Marshal(fMeta)
	if err != nil{
		//结构体转换到json过程失败
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}