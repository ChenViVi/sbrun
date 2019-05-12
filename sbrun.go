package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"io"
	"os/exec"
	"regexp"
)

func main() {
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
		exec.Command("kill", "-s", "9", pid).Start()
	}
	exec.Command( "nohup", "mvn", "spring-boot:run").Start()
}

func GetPid(port string) string{
	c1 := exec.Command("netstat", "-lnp")
	c2 := exec.Command("grep", port)
	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r
	var pidBytes bytes.Buffer
	c2.Stdout = &pidBytes
	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
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