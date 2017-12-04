package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func FileAsLineArray(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}
	r := bufio.NewReader(f)
	l := make([]string, 0)
	for {
		if s, err := r.ReadString('\n'); err == nil {
			l = append(l, strings.Split(s, "\n")[0])
		} else {
			return l
		}
	}
}

func FileAsWordArray(filename string) [][]string {
	l := make([][]string, 0)
	for _, s := range FileAsLineArray(filename) {
		l = append(l, strings.Split(s, " "))
	}
	return l
}

func FileAsWordNumberArray(filename string) [][]int {
	l1 := make([][]int, 0)
	for _, s1 := range FileAsWordArray(filename) {
		l2 := make([]int, 0)
		for _, s2 := range s1 {
			n, err := strconv.Atoi(s2)
			if err != nil {
				log.Panic(err)
			}
			l2 = append(l2, n)
		}
		l1 = append(l1, l2)
	}
	return l1
}
