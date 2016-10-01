package main

import "flag"   //用于处理命令行选项
import "fmt"  
import "os"    //os包定义了Args变量，其中存放了传递给命令行的全部参数

type Cmd struct {  //定义Cmd结构体
	helpFlag bool
	versionFlag bool
	cpOption string
	xjreOption string //使用jdk的启动类路径来寻找和加载java标准库中的类
	class string
	args []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", " ", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", " ", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}