package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Print(err)}
}
var addString,addFolder string = "",""

func seeFolder(dirName string){
	//var path string
	files, err := ioutil.ReadDir(dirName)
	CheckErr(err)
	sumFiles:=len(files)
	//fmt.Println("количество файлов в папке - ", sumFiles)
	addSimbol:="├───"
	for _, file := range files {
		sumFiles-=1
		if sumFiles==0{
			addSimbol="└───"
		}
		//если файл - выводим с размером
		if file.IsDir()==false{
			if file.Size()>0{
				fmt.Printf("%s%s%s (%d B)\n",addString,addSimbol, file.Name(), file.Size())
			}else{fmt.Printf("%s%s%s (empty)\n",addString,addSimbol, file.Name())}

			//если папка - выводим название и добавляем отступ
		}else {
			addString+="│	"
			fmt.Printf("%s├───%s\n", addFolder,file.Name())
			addFolder+="│	"
			path:=dirName + "/" + file.Name()
			//inherit+=1
			seeFolder(path)
		}
		if addSimbol=="└───"{
			if len(addString)>=4{
				addString=addString[0:len(addString)-4]
			}
			if len(addFolder)>=4{
				addFolder=addFolder[0:len(addFolder)-4]
				//inherit-=1
			}
		}
	}
}

func main() {
	seeFolder("tree/testdata/")
	//seeFolder("tree/example/tree")
}

