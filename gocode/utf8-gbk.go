package main

import (
	"fmt"
	"path/filepath"
	"os"
	"path"
	"github.com/go-ini/ini"
	"time"
	"log"
	"net/http"
	"runtime"
	"sync"
	"github.com/axgle/mahonia"
)

func HttpDownload()  {
	defer wg.Done()
	var updir string
	var gbkdir string
	fmt.Println("httprestapi server running..")
	fmt.Println("中文")
	for{
		pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil{
			fmt.Println("filepath Abs Failed:%v",err)
			time.Sleep(time.Second*20)
			continue
		}
		inipath := path.Join(pwd, "config.dat")
		//fmt.Println(inipath)
		cfg, err := ini.Load(inipath)
		if err != nil{
			fmt.Println("Failed to read file: %v", err)
			time.Sleep(time.Second*20)
			continue
		}
		updir = cfg.Section("updateclient").Key("save").String()
		fmt.Println(updir)
		if updir == ""{
			fmt.Println("updateclient download paht not set")
			time.Sleep(time.Second*20)
			continue
		}
		enc := mahonia.NewDecoder("GBK")
		gbkdir = enc.ConvertString(updir)
		fmt.Println(gbkdir)
		_, err = os.Stat(gbkdir)
		if err != nil{
			fmt.Printf("File '%v' does not exist,err:%v", gbkdir,err)
			log.Fatal("Download Dir not exist")
		}
		break
	}
	port := ""
	if len(os.Args) == 2{
		port = ":"+os.Args[1]
	}else{
		port = ":9416"
	}
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(gbkdir))))
}
var wg sync.WaitGroup
func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(1)
	go HttpDownload()
	wg.Wait()
}