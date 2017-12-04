package util

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type MyCSVFilterReader struct {
	reader io.Reader
}

func NewMyCSVFilterReader(reader io.Reader) *MyCSVFilterReader {
	p := new(MyCSVFilterReader)
	p.reader = reader
	return p
}

func (this *MyCSVFilterReader) Read(p []byte) (int, error) {
	n, err := this.reader.Read(p)
	if n > 0 {
		for i := 0; i < n; i++ {
			log.Infof("%v", p[i])
			if p[i] == '\t' || p[i] == ' ' {
				p[i] = ' '
			}
		}
	}
	log.Infof("%v", p)
	return n, err
}

func CSVNumbers(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		log.Panicf("Unable to open %s. Reason: %v", filename, err)
	}
	r := csv.NewReader(NewMyCSVFilterReader(f))
	r.Comma = ' '
	r.Comment = '#'
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	numbers := make([][]int, 0)
	for _, row := range records {
		n := make([]int, 0)
		for _, col := range row {
			if strings.Contains(col, " ") {
				log.Panicf("'%s' : Not just numbers", col)
			}
			i, err := strconv.Atoi(col)
			if err != nil {
				log.Panic(err)
			}
			n = append(n, i)
		}
		numbers = append(numbers, n)
	}
	return numbers
}
