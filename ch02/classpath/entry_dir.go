/*
	DirEntry，表示目录形式的类路径
*/

package classpath

import "io/ioutil"
import "path/filepath"

/*
	DirEntry只有一个字段，用于存放目录的绝对路径
*/
type DirEntry struct {
	absDir string
}

func newDirEntry (path string) *DirEntry {
	absDir err := filepath.Abs(path)
	if err != nil { //如果转换过程出现错误，则调用panic()函数终止程序执行
		panic (err) 
	}
	return &DirEntry{absDir} //创建DirEntry实例并返回
}
