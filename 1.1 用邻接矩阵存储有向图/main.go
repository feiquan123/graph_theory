package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	inputFile      = "./testdata/input.txt"
	array          = [100][100]int{}
	vCount, eCount = 0, 0
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		a, b := 0, 0
		fmt.Sscan(scanner.Text(), &a, &b)
		if a == 0 && b == 0 {
			break
		}
		if i == 0 {
			vCount, eCount = a, b
			continue
		}
		array[a-1][b-1] = 1
	}

	buf := bytes.NewBuffer(nil)
	for i := 0; i < vCount; i++ {
		sum := 0
		for j := 0; j < vCount; j++ {
			sum += array[i][j]
		}
		buf.WriteString(strconv.Itoa(sum))
		if i != vCount-1 {
			buf.WriteByte(' ')
		} else {
			buf.WriteByte('\n')
		}
	}

	for i := 0; i < vCount; i++ {
		sum := 0
		for j := 0; j < eCount; j++ {
			sum += array[j][i]
		}
		buf.WriteString(strconv.Itoa(sum))
		if i != vCount-1 {
			buf.WriteByte(' ')
		} else {
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
}
