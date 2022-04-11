package tool

import (
	"os"

	"searchfile/model"
)

// func GetLogicalDrives() []string {
// 	kernel32 := syscall.MustLoadDLL("kernel32.dll")
// 	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
// 	n, _, _ := GetLogicalDrives.Call()
// 	s := strconv.FormatInt(int64(n), 2)
// 	var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：", "V：", "W：", "X：", "Y：", "Z："}
// 	temp := drives_all[0:len(s)]
// 	var d []string
// 	for i, v := range s {
// 		if v == 49 {
// 			l := len(s) - i - 1
// 			d = append(d, temp[l])
// 		}
// 	}
// 	var drives []string
// 	for i, v := range d {
// 		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
// 	}
// 	return drives
// }

func ToTxt(p string, d model.Result) error {
	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		return err
	}

	res, err := d.ToJson()
	if err != nil {
		return err
	}

	f.Write([]byte(res))

	return nil
}
