package day09

import (
	"AdventOfCode2017/util"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	assert.Equal(t, 1, ParseStreamString("{}", false), "oops")
	assert.Equal(t, 6, ParseStreamString("{{{}}}", false), "oops")
	assert.Equal(t, 5, ParseStreamString("{{},{}}", false), "oops")
	assert.Equal(t, 16, ParseStreamString("{{{},{},{{}}}}", false), "oops")
	assert.Equal(t, 1, ParseStreamString("{<a>,<a>,<a>,<a>}", false), "oops")
	assert.Equal(t, 9, ParseStreamString("{{<ab>},{<ab>},{<ab>},{<ab>}}", false), "oops")
	assert.Equal(t, 9, ParseStreamString("{{<!!>},{<!!>},{<!!>},{<!!>}}", false), "oops")
	assert.Equal(t, 3, ParseStreamString("{{<a!>},{<a!>},{<a!>},{<ab>}}", false), "oops")
	log.Infof("My answer: %d", ParseStreamString(util.FileAsLineArray("input.1.txt")[0], false))
}

func TestPartTwo(t *testing.T) {
	log.Infof("My answer: %d", ParseStreamString(util.FileAsLineArray("input.1.txt")[0], true))
}
