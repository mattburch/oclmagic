package main

import (
	"bytes"
	"fmt"
	"github.com/docopt/docopt.go"
	"io"
	"log"
	"os"
)

func ByteToString(r io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String()
}

func main() {
	arguments, err := docopt.Parse(usage, nil, true, "oclmagic v0.1", false)
	if err != nil {
		log.Fatal("Error parsing usage. Error: ", err.Error())
	}

	hashcat := arguments["--hashcat"].(string)
	args := []string{
		"-m=5500",
		`/tmp/ntlm`,
		"--debug-mode=4",
		"-a=3",
		`?a?a?a?a?a?a?a`,
	}
	fmt.Printf("%#v\n", hashcat)
	fmt.Printf("%#v\n", args)

	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}

	pid, err := os.StartProcess(hashcat, args, &procAttr)
	if err != nil {
		log.Fatalf("Start Failed: %v", err)
	}

	_, err = pid.Wait()

}
