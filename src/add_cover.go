package main

import (
	"bytes"
	"os"
	"os/exec"
	"fmt"
	"io/ioutil"
	"strings"
)


func Exec(cmd string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	com := exec.Command("zsh", "-c", cmd)
	com.Stdout, com.Stderr = &stdout, &stderr

	err := com.Run()
	return stdout.String(), stderr.String(), err
}

func AddCover(ext, command string) {
	files, err := ioutil.ReadDir(".")

	if err != nil {
		fmt.Println("Cannot read current dir...")
		return
	}

	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && strings.HasSuffix(name, "." + ext) {
			name := file.Name()
			cmd := command + fmt.Sprintf("'%s'", name)
			o, e, err := Exec(cmd)

			if err != nil {
				fmt.Println(e)
				return
			}

			fmt.Println(o)
		}
	}
}

func main() {
	args := os.Args[1:]
	var cmd string
	filetype, covername := args[0], args[1]

	if filetype != "m4a" {
		// Currently just using this with iTunes purchased music
		// However other formats should be supported
		return
	} else {
		cmd = "mp4art --add %s "
		AddCover(filetype, fmt.Sprintf(cmd, covername))
	}
}
