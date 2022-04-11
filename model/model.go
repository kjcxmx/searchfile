package model

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	DATAFMT = "2006-01-02 15:04:05.000"
)

type Result struct {
	SrcDir    string
	Target    string
	FileInfos []FileInfo
	StartTime time.Time
	EndTime   time.Time
}

type FileInfo struct {
	Dir      string
	Name     string
	Size     int64
	Mode     uint32
	FindTime time.Time
	ModTime  time.Time
}

func (this *FileInfo) Print() {
	if this == nil {
		return
	}
	fmt.Printf("Dir：%v\tMode：%v\tSize：%v\tModTime：%v\tFindTime：%v\n", this.Dir, this.Mode, this.Size, this.ModTime, this.FindTime)
}

func (this *FileInfo) PrintFmt() {
	if this == nil {
		return
	}
	fmt.Printf("Dir：%v\tMode：%v\tSize：%v\tModTime：%v\tFindTime：%v\n", this.Dir, this.Mode, this.Size, this.ModTime.Format(DATAFMT), this.FindTime.Format(DATAFMT))
}

func (this *FileInfo) ToJson() (string, error) {
	if this == nil {
		return "", fmt.Errorf("fileinfo is null")
	}

	b, err := json.Marshal(this)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (res *Result) ToJson() (string, error) {
	if res == nil {
		return "", fmt.Errorf("result is null")
	}

	b, err := json.Marshal(res)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
