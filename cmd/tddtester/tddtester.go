package main

import (
	"flag"
	"fmt"
	//"github.com/fsnotify/fsnotify"
	"os"
)

var lang string
var directory string

func main() {

	flag.StringVar(&lang, "lang", "", "programming language")
	flag.StringVar(&dir, "dir", "", "directory to watch")
	flag.StringVar(&command, "cmd", "", "command to run when code files change")
	flag.Parse()

	if lang == "" || dir == "" || cmd == "" {
		fmt.Println("usage: tddtester --lang (go|python|perl|...) --dir . --cmd go test")
		os.Exit(0)
	}

}
