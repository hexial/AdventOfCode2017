package day02

import (
	"AdventOfCode2017/util"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func calcPartOne(filename string) int {
	//log.Infof("*****************")
	numbers := util.CSVNumbers(filename)
	sum := 0
	for _, row := range numbers {
		min := -1
		max := -1
		for _, num := range row {
			if min == -1 || min > num {
				min = num
			}
			if max == -1 || max < num {
				max = num
			}
		}
		sum += max - min
	}
	return sum
}

func TestPartOne(t *testing.T) {
	//assert.Equal(t, 18, calcPartOne("input.sample.txt"), "oops")
	log.Infof("Answer: %d", calcPartOne("input.txt"))
}

func calcPartTwo(filename string) int {
	numbers := util.CSVNumbers(filename)
	sum := 0
	for _, row := range numbers {
		for i := 0; i < len(row); i++ {
			for j := 0; j < len(row); j++ {
				if j != i && row[i] > row[j] {
					if row[i]%row[j] == 0 {
						sum += int(row[i] / row[j])
					}
				}
			}
		}
	}
	return sum
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 9, calcPartTwo("input.2.sample.txt"), "oops")
	log.Infof("Answer: %d", calcPartTwo("input.2.txt"))
}
