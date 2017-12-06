package day04

import (
	"AdventOfCode2017/util"
	"sort"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func DoPartOne(filename string, anagram bool) int {
	lines := util.FileAsWordArray(filename)
	var count int
	for _, line := range lines {
		if CheckPartOne(line, anagram) {
			count++
		}
	}
	return count
}

func CheckPartOne(in []string, anagram bool) bool {
	m := make(map[string]bool)
	log.Infof("*****************")
	for _, s := range in {
		if anagram {
			r := strings.NewReader(s)
			sl := make([]string, len(s))
			for i := 0; i < len(s); i++ {
				b, _ := r.ReadByte()
				sl[i] = string(b)
			}
			sort.Strings(sl)
			var ss string
			for i := 0; i < len(s); i++ {
				ss += sl[i]
			}
			s = ss
		}
		if _, ok := m[s]; ok {
			log.Infof("%s: false", s)
			return false
		}
		log.Infof("%s: true", s)
		m[s] = true
	}
	return true
}

func TestPartOne(t *testing.T) {

	//
	//

	assert.Equal(t, 2, DoPartOne("input.part.one.sample.txt", false), "oops")
	log.Infof("My answer: %d", DoPartOne("input.part.one.txt", false))
}

func TestPartTwo(t *testing.T) {

	//
	//

	assert.Equal(t, 3, DoPartOne("input.part.two.sample.txt", true), "oops")
	log.Infof("My answer: %d", DoPartOne("input.part.two.txt", true))
}
