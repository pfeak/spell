package src

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestLCSMap_Train(t *testing.T) {
	var (
		f         *os.File
		lineBytes []byte
		line      string

		err error
	)

	lcsMap := NewLCSMap("", "")

	f, err = os.Open("../data/demo.log")
	if err != nil {
		t.Fatal(err)
	}

	r := bufio.NewReader(f)
	for {
		lineBytes, err = r.ReadBytes('\n')
		line = strings.TrimSpace(string(lineBytes))
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}

		_, err = lcsMap.Train(line, 0.01)
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, object := range lcsMap.lcsObjects {
		fmt.Println(object.lcsSeq, object.getPosition())
	}
	// Output:
	// [<*> is <*> pen] [0 2]
	// [<*> am <*>] [0 2]
	// [i am grey and black] []
	// [<*> "logId <*>] [0 2]
	// [<*> "message" <*>] [0 2]
	// [<*> FC020000067245 devId <*> logId <*>] [0 3 5]
}
