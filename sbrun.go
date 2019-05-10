package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
)

func main() {
	// applicationPropertiesBytes, err1 := ioutil.ReadFile("src/main/resources/application.properties")
	// if err1 != nil {
	// 	applicationPropertiesBytes, err2 := ioutil.ReadFile("src/main/resources/application.yaml")
	// 	if err2 != nil {
	// 		fmt.Println("你丫配置文件到底搁哪儿呢")
	// 	}
	// 	// else {
	// 	// 	GetPort(applicationPropertiesBytes)
	// 	// }
	// }
	applicationPropertiesBytes, err1 := ioutil.ReadFile("src/main/resources/application.properties")
    if err1 != nil {
		applicationPropertiesBytes, err2 := ioutil.ReadFile("src/main/resources/application.yaml")
		if err2 != nil {
			fmt.Println("你丫配置文件到底搁哪儿呢")
		} else {
			//netstat -lnp|grep 111111
			ExeCommand("netstat", "-lnp", "|grep " + GetPort(applicationPropertiesBytes))
		}
    } else {
			ExeCommand("netstat", "-lnp", "|grep " + GetPort(applicationPropertiesBytes))
    }
}

func GetPort(applicationPropertiesBytes []byte) string{
	applicationProperties := string(applicationPropertiesBytes)
	portStrRe, _ := regexp.Compile("port.*?([1-9]\\d*)")
	portStr := portStrRe.FindString(applicationProperties)
	portRe, _ := regexp.Compile("[1-9]\\d*")
	port := portRe.FindString(portStr)
	return port
}

func ExeCommand(commandName string, arg ...string) bool {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(commandName, arg...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return true
}