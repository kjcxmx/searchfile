
### Introduction

SearchFile is a simple and efficient local file retrieval tool developed based on golang.


### Installation

```
go get github.com/kjcxmx/searchfile
```


### Usage

```
root@root:#/github.com/kjcxmx/searchfile/target$ ./SearchFile_unix_X86
```

```
SearchFile version: 1.0.1       By kjcxmx@163.com
Usage: SearchFile [-d /home / C:/] [-n] [-s std/json/soj] [-h help]

eg：
  Unix：
        ./SearchFile -d /home -n main.go -s json
  Windows：
        ./SearchFile.exe -d C:/ -n main.go -s soj -o /home/find.txt
Options:
  -d string
        start from this dir, default /home or C:/ (default "/home")
  -n string
        search file name, default main.go (default "main.go")
  -o string
        output file path and filename, default ./result.txt (default "./result.txt")
  -s string
        stdout：std /stdjson：json /stdout and stdjso：soj, default std (default "std")
```

### Create FuncCall
```
	res := make(chan model.FileInfo)
	result := model.Result{}

	go func() {
		for v := range res {
			result.FileInfos = append(result.FileInfos, v)
		}
	}()

	if err := search.SearchFile(result.SrcDir, result.Target, res); err != nil {
		fmt.Println(err)
	}
      fmt.Printf("%+v\n",result)
```

