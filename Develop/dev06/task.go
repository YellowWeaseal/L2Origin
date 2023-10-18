package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func cut(fields string, delimiter string, separated bool, input *os.File, output *os.File) {
	scanner := bufio.NewScanner(input)
	field := strings.Split(fields, ",")
	var fieldsP []int
	for _, v := range field {
		strInt, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("error fields input")
		}
		fieldsP = append(fieldsP, strInt)
	}

	for scanner.Scan() {
		str := scanner.Text()

		if separated {
			if !strings.Contains(str, delimiter) {
				continue
			}
		}

		strD := strings.Split(str, delimiter)

		var outputLine string

		for i, v := range fieldsP {
			if i > 0 {
				outputLine += " "
			}

			if v >= 0 && v < len(strD) {
				outputLine += strD[v]
			}
		}

		fmt.Fprintln(output, outputLine)
	}
}

func main() {
	fields := flag.String("f", "0,2", "fields")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", true, "separated")

	flag.Parse()

	cut(*fields, *delimiter, *separated, os.Stdin, os.Stdout)
}
