package main

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/elgs/filesync/config"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("CPUs: ", runtime.NumCPU())
	input := args()
	done := make(chan bool)
	if len(input) >= 1 {
		start(input[0], done)
	}
	<-done
}

func start(configFile string, done chan bool) {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(configFile, " not found")
		go func() {
			done <- false
		}()
		return
	}
	json, _ := simplejson.NewJson(b)
	mode := json.Get("mode").MustString("server")
	if mode == "server" {
		config.StartServer(configFile)
	} else if mode == "client" {
		config.StartClient(configFile, done)
	}
}

func args() []string {
	ret := []string{}
	if len(os.Args) <= 1 {
		ret = append(ret, "gsync.json")
	} else {
		for i := 1; i < len(os.Args); i++ {
			ret = append(ret, os.Args[i])
		}
	}
	return ret
}
