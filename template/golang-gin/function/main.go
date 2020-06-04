package function

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func readSecret(name string) string {
	sname := filepath.Join("/var/openfaas/secrets", name)
	txt, err := ioutil.ReadFile(sname)
	txt = bytes.Trim(txt, "\n")
	if err != nil {
		return ""
	}
	return string(txt)
}

func init() {
	if runtime.GOOS == "windows" {
		log.Println("not support windows")
		return
	}

	//entry
	main()
}

func main() {
	//read your config
	user := readSecret("user")
	config := os.Getenv("config")
	_ = user
	_ = config

	f, err := NewFunctionApp(&RuntimeConfig{
		//TODO: set your config params
	})
	if err != nil {
		log.Panic("[main.go::main] NewFynctionApp", err)
	}
	app = f
}
