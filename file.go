package ext

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Reader interface {
	ReadLine(string)
}

func FromFile(reader Reader, name string) (err error){
	f, err := os.Open(name)
	if err != nil {
		return
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for  {
		line, err := buf.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		line = strings.Replace(line, "\r", "", -1)
		line = strings.Replace(line, "\n", "", -1)
		reader.ReadLine(line)
		if err == io.EOF {
			return nil
		}
	}

	return
}