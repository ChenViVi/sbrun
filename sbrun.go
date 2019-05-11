package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	pipe "github.com/b4b4r07/go-pipe"
)

func main() {
	pid := ""
	if pid == ""{
		fmt.Println("0")
	} else {
		fmt.Println("1")
	}
	applicationPropertiesBytes, err1 := ioutil.ReadFile("src/main/resources/application.properties")
    if err1 != nil {
		applicationPropertiesBytes, err2 := ioutil.ReadFile("src/main/resources/application.yaml")
		if err2 != nil {
			fmt.Println("你丫配置文件到底搁哪儿呢")
		} else {
			run(applicationPropertiesBytes)
		}
    } else {
			run(applicationPropertiesBytes)
    }
}

func run(applicationPropertiesBytes []byte){
	port := GetPort(applicationPropertiesBytes)
	pid := GetPid(port)
	if pid != ""{
		ExeCommand(false, "kill", "-s", "9", pid)
	}
	ExeCommand(true, "nohup", "mvn", "spring-boot:run")
	fmt.Println("fuck")
	// runPid := GetPid(port)
	// if runPid == "" {
	// 	fmt.Println("运行失败")
	// } else {
	// 	fmt.Println("运行成功")
	// }
}

//netstat -lnp|grep port
func GetPid(port string) string{
	var pidBytes bytes.Buffer
	if err := pipe.Command(&pidBytes,
		exec.Command("netstat", "-lnp"),
		exec.Command("grep", port),
	); err != nil {
		fmt.Println(err)
		return "";
	}
	pidStr := pidBytes.String()
	pidStrRe, _ := regexp.Compile("([1-9]\\d*)/java")
	pidStr = pidStrRe.FindString(pidStr)
	pidRe, _ := regexp.Compile("[1-9]\\d*")
	pid := pidRe.FindString(pidStr)
	return pid
}

func GetPort(applicationPropertiesBytes []byte) string{
	applicationProperties := string(applicationPropertiesBytes)
	portStrRe, _ := regexp.Compile("port.*?([1-9]\\d*)")
	portStr := portStrRe.FindString(applicationProperties)
	portRe, _ := regexp.Compile("[1-9]\\d*")
	port := portRe.FindString(portStr)
	return port
}

func ExeCommand(wait bool, commandName string, arg ...string) {
	cmd := exec.Command(commandName, arg...)
	fmt.Println(cmd.Args)
	cmd.Start()
	if wait {
		cmd.Wait()
	}
}