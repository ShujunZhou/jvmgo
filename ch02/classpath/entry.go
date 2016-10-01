/*
	Entry接口表示类路径
*/

package classpath

import "os"
import "strings"

/*
	常量pathSeparator是string类型，存放路径的分隔符
*/
const pathListAeparator = string(os.PathListSeparator)

type Entry interface {
	/*
		readClass()方法负责寻找和加载class文件。
		方法的参数是class文件的相对路径，路径之间用斜线（/）分隔，
		文件名有.class后缀；
		返回值是读取到的字节数据、最终定位到class文件的Entry,以及错误信息
	*/
	readClass(className string) ([]byte, Entry, error) 
	String() string //用于返回变量的字符串表示
}

/*
	newEntry()函数根据参数创建不同类型的Entry实例
*/
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
	   strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	 }

	 return newDirEntry(path)
}