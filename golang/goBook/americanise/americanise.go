package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var americaniseFile = "british-american.txt"

func main() {
	infile, outfile, err := getArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	in, out := os.Stdin, os.Stdout

	if infile != "" {
		in, _ = os.Open(infile)
		defer in.Close()
	}

	if outfile != "" {
		out, _ = os.Create(outfile)
		defer out.Close()
	}

	if err := americanise(in, out); err != nil {
		log.Fatal(err)
	}
}

func getArgs() (string, string, error) {
	var infile, outfile string
	if len(os.Args) > 1 {
		infile = os.Args[1]
		if len(os.Args) > 2 {
			outfile = os.Args[2]
		}
	}
	return infile, outfile, nil
}

func americanise(in io.Reader, out io.Writer) (err error) {
	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	var replacer func(string) string
	if replacer, err = makerReplacerFunc(americaniseFile); err != nil {
		return err
	}

	wordRx := regexp.MustCompile("[A-Za-z]+")
	eof := false
	var line string
	for !eof {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return err
		}
		line = wordRx.ReplaceAllStringFunc(line, replacer)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}

func makerReplacerFunc(file string) (func(string) string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	repalceHash := make(map[string]string)
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			repalceHash[fields[0]] = fields[1]
		}

	}
	return func(word string) string {
		if usword, found := repalceHash[word]; found {
			return usword
		}
		return word
	}, nil
}
