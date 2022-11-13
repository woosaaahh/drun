package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var (
		prog string = filepath.Base(os.Args[0])
		err  error
	)

	log.SetPrefix(prog + " | ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("USAGE: %s program [arg...]\n", prog)
		os.Exit(1)
	}

	cmd := exec.Command(args[0], args[1:]...)
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s (pid:%d)", strings.Join(args, " "), cmd.Process.Pid)
}
