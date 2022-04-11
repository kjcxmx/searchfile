package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"searchfile/model"
	"searchfile/search"
	"searchfile/tool"
)

func usage() {
	fmt.Fprintf(os.Stderr, `SearchFile version: 1.0.1	By kjcxmx@163.com
Usage: SearchFile [-d /home / C:/] [-n] [-s std/json/soj] [-h help]

eg：
  Unix：
	./SearchFile -d /home -n main.go -s json
  Windows：
	./SearchFile.exe -d C:/ -n main.go -s soj -o /home/find.txt
Options:
`)
	flag.PrintDefaults()
}

func main() {
	srcdir := flag.String("d", "/home", "start from this dir, default /home or C:/")
	target := flag.String("n", "main.go", "search file name, default main.go")
	stdout := flag.String("s", "std", "stdout：std /stdjson：json /stdout and stdjso：soj, default std")
	output := flag.String("o", "./result.txt", "output file path and filename, default ./result.txt")

	flag.Parse()
	flag.Usage = usage

	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

	res := make(chan model.FileInfo)
	result := model.Result{SrcDir: *srcdir, Target: *target, StartTime: time.Now()}

	if runtime.GOOS == "windows" {
		result.SrcDir = "C:/ D:/ E:/ F:/ G:/"
	}

	go func() {
		for v := range res {
			if *stdout == "std" || *stdout == "soj" {
				v.PrintFmt()
			}
			result.FileInfos = append(result.FileInfos, v)
		}
	}()

	if err := search.SearchFile(result.SrcDir, result.Target, res); err != nil {
		fmt.Println(err)
	}

	result.EndTime = time.Now()

	if *stdout == "json" || *stdout == "soj" {
		fmt.Println(result.ToJson())
	}

	if *output != "" {
		tool.ToTxt(*output, result)
	}
}
