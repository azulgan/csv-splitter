package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := new(MyArgs)
	args.parseArgs(os.Args[1:])

	r := openReader(args.InputFileName)
	r.Comma = args.Comma
	w := openWriter(args.OutputFileName)
	w.Comma = args.Comma
	// title
	titleRow, err := r.Read()
	if err != nil {
		panic(err)
	}
	w.Write(append(titleRow, "Fraction"))

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err) // or handle it another way
		}
		list := row[len(row)-1]
        arr := strings.Split(list, ",")
        fraction := 1.0/float64(len(arr))
        for _, v := range arr {
        	newRow := row[0:len(row)-1]
			newRow = append(newRow, v)
			newRow = append(newRow, fmt.Sprintf("%f", fraction))
			w.Write(newRow)
		}
	}
	w.Flush()
}

func openReader(fileName string) *csv.Reader {
	csvFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	//defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	return csvReader
}


func openWriter(fileName string) *csv.Writer {
	csvFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	//defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	return csvWriter
}
