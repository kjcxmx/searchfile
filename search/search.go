package search

import (
	"io/ioutil"
	"path"
	"strings"
	"time"

	"searchfile/model"
)

func SearchFile(p, n string, res chan model.FileInfo) (err error) {
	for _, v := range strings.Split(p, " ") {
		if v == "" {
			continue
		}
		err = find(res, v, n)
	}

	if err != nil && strings.Contains(err.Error(), "open") {
		return nil
	}

	return err
}

func find(res chan model.FileInfo, p, n string) error {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return err
	}

	for _, v := range files {
		if v.IsDir() {
			find(res, path.Join(p, v.Name()), n)
		} else {
			if v.Name() == n {
				res <- model.FileInfo{Dir: p, Name: n, FindTime: time.Now(), Size: v.Size(), ModTime: v.ModTime(), Mode: uint32(v.Mode())}
			}
		}
	}
	return nil
}
