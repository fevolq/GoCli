package model

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

var content []string
var threadNum int
var wg sync.WaitGroup

func Execute(filePath string) {
	readFile(filePath)
	doCli()
}

func readFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 表示文件末尾
			break
		}
		if str == "\r\n" {
			continue
		} else if len(str) > 0 {
			str = strings.TrimRight(str, "\r\n")
			content = append(content, str)
		}
	}
}

func splitStr(str string) (result []string) {
	return strings.Split(str, " ")
}

func doCli() {
	firstCmd := splitStr(content[0])
	if firstCmd[0] != "THREAD" {
		threadNum = 1
	} else {
		num, err := strconv.Atoi(firstCmd[len(firstCmd)-1])
		if err != nil {
			threadNum = 1
			panic(err)
		} else {
			threadNum = num
			content = content[1:]
		}
	}

	for i := 0; i < threadNum; i++ {
		wg.Add(1)
		go func() {
			cmd := exec.Command("cmd")
			for _, str := range content {
				log.Println("doCli", str)
				strList := splitStr(str)
				for _, str := range strList {
					log.Println("doCli:", str)
				}
				doCmd(cmd, strList)
			}

			cmd.Stdout = os.Stdout
			cmd.Run()

			wg.Done()
		}()
	}
	wg.Wait()
}

func doCmd(cmd *exec.Cmd, strList []string) {
	if len(strList) != 2 {
		return
	}
	switch {
	case true:
		log.Println("doCmd", strList[1])
		fallthrough
	case strList[0] == "WORKDIR":
		os.Chdir(strList[1])
	case strList[0] == "CMD":
		cmd = exec.Command("cmd", "/c", strList[1])
	default:
		panic("参数配置异常")
	}
}
