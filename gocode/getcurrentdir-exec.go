package main

import (
	"fmt"
	"path/filepath"
	"os"
	"strings"
	"log"
	"path"
)


func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	var str1, str2 string
	//fmt.Println(getCurrentDirectory())
	str1 = getCurrentDirectory()
	fmt.Println(str1)
	str2 = getParentDirectory(str1)
	fmt.Println(str2)
	path1 := path.Join(str1,"plugin/sysproxy.exe")
	fmt.Println(path1)
	//path, err := exec.LookPath("d:/sysproxy.exe")
	//if err != nil {
	//	log.Fatal("installing fortune is in your future")
	//}
	//fmt.Printf("fortune is available at %s\n", path)
	////cmd := exec.Command(path, "pac file://e:/pac.txt ")
	//cmd := exec.Command(path, "pac","file://e:/pac123.txt")
	//
	////cmd := exec.Command(path, "off")
	//log.Println("Run")
	//err = cmd.Run()
	//log.Println(err)
}