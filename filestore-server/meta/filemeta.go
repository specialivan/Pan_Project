//文件属性
package meta

//FileMeta：文件源信息结构体
type FileMeta struct{
	FileSha1 string
	FileName string
	FileSize int64
	Location string  //存储位置
	UploadAt string //时间戳
}

//定义对象存储每一个上传文件的源信息
var fileMetas map[string]FileMeta

//初始化map
func init(){
	fileMetas = make(map[string]FileMeta)
}

//UpdateFileMeta:新增/更新文件源信息
func UpdateFileMeta(fmeta FileMeta){
	fileMetas[fmeta.FileSha1]= fmeta
}

//GetFileMeta:通过sha1获取文件源信息对象
func GetFileMeta(fileSha1 string)FileMeta{
	return fileMetas[fileSha1]

}