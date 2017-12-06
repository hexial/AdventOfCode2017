package day05

import (
	"AdventOfCode2017/util"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func calc(filename string) int {
	list := util.FileAsNumberArray(filename)
	index := 0
	steps := 0
	for index >= 0 && index < len(list) {
		//log.Infof("%+v [index=%d][%d]", list, index, list[index])
		if list[index] == 0 {
			//
			// Do nothing
			list[index]++
		} else {
			jump := list[index]
			list[index]++
			index += jump

		}
		steps++
	}
	return steps
}

func calcSecond(filename string) int {
	list := util.FileAsNumberArray(filename)
	index := 0
	steps := 0
	for index >= 0 && index < len(list) {
		//log.Infof("%+v [index=%d][%d]", list, index, list[index])
		if list[index] == 0 {
			//
			// Do nothing
			list[index]++
		} else {
			jump := list[index]
			if jump >= 3 {
				list[index]--
			} else {
				list[index]++
			}
			index += jump

		}
		steps++
	}
	return steps
}

func TestPartOne(t *testing.T) {

	//
	//

	assert.Equal(t, 5, calc("input.1.sample.txt"), "oops")
	log.Infof("My answer: %d", calc("input.1.txt"))
}

func TestPartTwo(t *testing.T) {

	//
	//

	assert.Equal(t, 10, calcSecond("input.1.sample.txt"), "oops")
	log.Infof("My answer: %d", calcSecond("input.1.txt"))
}
