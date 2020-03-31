package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func readAllInfoMemory(filename string) (content []byte, err error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	fileInfo, err := fp.Stat()
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, fileInfo.Size())
	_, err = fp.Read(buffer)
	if err != nil {
		return  nil, err
	}
	return buffer, nil
}

func readByBlock(filename string) (content []byte, err error) {
	fp, err := os.Open(filename)
	if err != nil {
		return content, err
	}
	defer fp.Close()

	const blockSize = 12
	buffer := make([]byte, blockSize)
	for {
		bytesread, err := fp.Read(buffer)
		content = append(content, buffer[:bytesread]...)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return nil, err
			}
		}
	}
	return 
}

func readByLine(filename string) (lines [][]byte, err error) {
	fp, err := os.Open(filename)
	if err != nil {
		return  nil, err
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	for {
		line, _, err:= reader.ReadLine()
		fmt.Println(line, err, len(line))
		// line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		} else {
			// l := make([]byte, len(line))
			// copy(l, line)
			// lines = append(lines, l)
			lines = append(lines, line)
		}
		fmt.Println(lines)
	}
	return lines, nil
}

func main() {
	filename := "../config/config.yaml"
	content, err := readByLine(filename)
	for _, v := range content {
		fmt.Printf("%s\n", v)
	}
	fmt.Println(err)
}