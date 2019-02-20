package monitorcommon

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var Url = beego.AppConfig.String("monitor::monitorPwd")
var FileName = beego.AppConfig.String("monitor::monitorFile")

//var ITSM_URL = beego.AppConfig.String("itsm::ITSM_URL")

func Test() {
	fmt.Println("test")
	var url = beego.AppConfig.String("monitor::monitorPwd")
	var fileName = beego.AppConfig.String("monitor::monitorFile")
	fmt.Println("%s/%s", url, fileName)

}

func WriteFile(data string) (err error) {
	/*
		if u, err := user.Current(); err == nil {
			fmt.Println("用户ID: " + u.Uid)
			fmt.Println("主组ID: " + u.Gid)
			fmt.Println("用户名: " + u.Username)
			fmt.Println("主组名: " + u.Name)
			fmt.Println("家目录: " + u.HomeDir)
		}
	*/
	var file = fmt.Sprintf("%s/%s", Url, FileName)
	//str := "dkfslfjdsklfjlskjflsjflsjflsjflks"
	err = ioutil.WriteFile(file, []byte(data), 0755)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func LoadFile() (hashList map[string]int, err error) {
	var file = fmt.Sprintf("%s/%s", Url, FileName)
	fmt.Println(file)
	inputFile, inputError := os.Open(file)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()
	//var hashList = make(map[string]int)
	bf := bufio.NewReader(inputFile)
	var index = 1
	for {
		line, isPrefix, err := bf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if isPrefix {
			fmt.Println("Line too long")
		}
		if _, ok := hashList[string(line)]; ok {
			hashList[string(line)]++
		} else {
			if len(string(line)) != 0 {
				hashList[string(line)] = index
				index++
				fmt.Println("???")
			}

		}
	}
	fmt.Println(hashList)
	return
}

func GetItsmStatus() (data string) {
	var ITSM_URL = beego.AppConfig.String("itsm::ITSM_URL")
	var GetChangeInterface = beego.AppConfig.String("itsm::ITSM_GET_CHANGE")
	var url = fmt.Sprintf("%s%s", ITSM_URL, GetChangeInterface)
	//fmt.Println(url)
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Get(url)
	if err != nil {
		// handle error
		data = fmt.Sprintf("%s{url='%s',type='get'} 1", beego.AppConfig.String("monitor::metric"), url)
		return
	}
	data = fmt.Sprintf("%s{url='%s',type='get'} 0", beego.AppConfig.String("monitor::metric"), url)
	defer response.Body.Close()
	return

}

func SelfServiceMonitor() (data string) {

	return
}

func CronMonitor() {
	//var aL [2]string
	data := [2]string{GetItsmStatus(), GetItsmStatus()}
	var dataString string
	for _, value := range data {
		dataString = fmt.Sprintf("%s\n%s", dataString, value)
	}
	err := WriteFile(dataString)
	if err != nil {
		fmt.Println(err)
	}
}
