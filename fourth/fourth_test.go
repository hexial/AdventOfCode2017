package fourth

import (
	"bufio"
	"os"
	"sort"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func DoPartOne(filename string, anagram bool) int {
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)
	var count int
	for {
		if s, err := r.ReadString('\n'); err == nil {
			if CheckPartOne(strings.Split(s, "\n")[0], anagram) {
				count++
			}
		} else {
			return count
		}
	}
}

func CheckPartOne(in string, anagram bool) bool {
	m := make(map[string]bool)
	l := strings.Split(in, " ")
	log.Infof("*****************")
	for _, s := range l {
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
