package main
import(
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
func main(){
    dirUrls := ""
	dirSclie :=make([]string,100)
	fmt.Println("请输入jar文件的路径，多个路径以英文逗号分隔开(,)。")
	_, err := fmt.Scanf("%s",&dirUrls)
	fmt.Println("input","=",dirUrls)
	if nil == err {
		file,err_f := os.Create("MANIFEST.MF")
		if err_f != nil {
			panic(err_f)
		}
		file.WriteString("Manifest-Version: 1.0\r\n")
		file.WriteString("Permissions: sandbox\r\n")
		file.WriteString("JavaFX-Version: 8.0\r\n")
		file.WriteString("Class-Path:")
		dirSclie = strings.Split(dirUrls,",")
		for i,dirs := range dirSclie{
			fmt.Println(i,"=",dirs)
			tempdirSlice := strings.Split(dirs,"\\")
            tempdir := tempdirSlice[(len(tempdirSlice)-1)]
			dir_list,e := ioutil.ReadDir(dirs)
			if e != nil{
				fmt.Print("read dir error")
				return
			}
			for j, dname := range dir_list{
				fileName := dname.Name()
				if j != 0 && j%2==0{
					file.WriteString(" \r\n")
				}
				file.WriteString(" "+tempdir+"/"+fileName)
			}
   		}
		file.WriteString(" \r\n")
		file.WriteString("Created-By: MANIFEST.go\r\n")
		file.WriteString("Main-Class: ")
		//此处不能少，MANIFEST.MF文件后面必须空出一行，否则jar运行报错
		file.WriteString("\r\n")
		file.WriteString("\r\n")
		defer file.Close()
	}
   
}