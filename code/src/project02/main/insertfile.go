package temp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name string
	Gender string
	Num   int
}

func (this *Student) Store(stu Student)  {
	file := "src/go_code/project02/temp/001"
	File,err := os.OpenFile(file,os.O_RDWR|os.O_APPEND,0666)
	if err != nil {
		fmt.Println("open file false")
	}
	defer File.Close()
	stux,err := json.Marshal(stu)
	Write := bufio.NewWriter(File)
	Write.WriteString(string(stux))
	Write.Flush()

}

func mainq()  {
	var stu Student
	stu = Student{"Jack","man",2022010539}
	stu.Store(stu)
}