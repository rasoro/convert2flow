package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"code.sajari.com/docconv"
)

type ScannedLines struct {
	Lines []string
}

func NewScannedLines() ScannedLines {
	return ScannedLines{
		Lines: make([]string, 0),
	}
}

type DocxFile struct {
	ScannedLines ScannedLines
	Scanner      *bufio.Scanner
	Messages     []string
}

func NewDocxFile() DocxFile {
	return DocxFile{}
}

type TxtFile struct {
	ScannedLines ScannedLines
	Scanner      *bufio.Scanner
	Messages     []string
}

func NewTxtFile() TxtFile {
	return TxtFile{}
}

func (d *DocxFile) ReadFile(filePath string) {
	res, err := docconv.ConvertPath(filePath)
	if err != nil {
		log.Fatal(err)
	}

	f := res.Body
	d.Scanner = bufio.NewScanner(strings.NewReader(f))
	d.ScannedLines = GetLines(d.Scanner)
	d.Messages = ParseLinesToMsgs(d.ScannedLines)
}

func (t *TxtFile) ReadFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	t.Scanner = bufio.NewScanner(f)
	t.ScannedLines = GetLines(t.Scanner)
	t.Messages = ParseLinesToMsgs(t.ScannedLines)
}

func GetLines(scan *bufio.Scanner) ScannedLines {
	scannedLines := NewScannedLines()
	for scan.Scan() {
		scannedLines.Lines = append(scannedLines.Lines, scan.Text())
	}
	return scannedLines
}

func ParseLinesToMsgs(sl ScannedLines) []string {
	messages := make([]string, 0)
	accumulatedText := ""
	for i, line := range sl.Lines {
		textPart := line
		if textPart != "" {
			accumulatedText = accumulatedText + textPart
			if i < (len(sl.Lines) - 1) {
				nextLine := sl.Lines[i+1]
				if nextLine == "" {
					messages = append(messages, accumulatedText)
					accumulatedText = ""
				} else {
					accumulatedText = accumulatedText + "\n"
				}
			} else {
				messages = append(messages, accumulatedText)
				accumulatedText = ""
			}
		}
	}
	return messages
}

func WriteJsonToFile(json string) {
	f, err := os.Create("flow_import.json")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(json)
	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
