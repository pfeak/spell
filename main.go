package main

import (
	"bufio"
	"io"
	"os"
	"spell/src"
	"strings"
)

func main() {
	var (
		f         *os.File
		lineBytes []byte
		line      string

		err error
	)

	lcsMap := src.NewLCSMap("", "")

	f, err = os.Open("./data/demo.log")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		lineBytes, err = r.ReadBytes('\n')
		line = strings.TrimSpace(string(lineBytes))
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		_, err = lcsMap.Train(line, 0.01)
		if err != nil {
			panic(err)
		}
	}

	lcsMap.Print()
}
