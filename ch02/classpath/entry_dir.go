﻿/*
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

//将path转换为绝对路径,表示目录的类路径
func newDirEntry (path string) *DirEntry {
	absDir err := filepath.Abs(path) //将path转换为绝对路径
	if err != nil { //如果转换过程出现错误，则调用panic()函数终止程序执行
		panic (err) 
	}
	return &DirEntry{absDir} //创建DirEntry实例并返回
}

//加载类文件
func (self * DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className) //将目录和类名拼成完整的路径
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

//返回目录的字符串表示
func (self * DirEntry) Strings() string {
	return self.absDir
}
