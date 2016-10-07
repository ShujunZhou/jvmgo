package classpath

import (
	"archive/zip"
	"io/ioutil"
	"path/filepath"
	"errors"
)

type ZipEntry struct {
	absPath string
}

//实例化ZipEntry
func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}

//加载压缩包
func (self * ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)  //打开zip包
	if err != nil {
		return null, null, err
	}

	defer r.Close()
	for _, f := range r.File { //遍历zip包中的文件，寻找class文件
		if f.Name == className { //如果能找到，则打开
			rc, err := f.Open()
		}
		
		defer rc.Close()
		data, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, nil, err
		}
		return data, self, nil
	}

	return nil, nil, errors.New("class not found:" + className)
}

func (self * ZipEntry) string() string {
	return self.absPath
}