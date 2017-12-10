package day10

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	assert.Equal(t, 12, NewKnotHash("sample", 5).result1(), "oops")
	log.Infof("My answer: %d", NewKnotHash("input", 256).result1())
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, "a2582a3a0e66e6e86e3812dcb672a272", NewKnotHashPartTwo("", 256).denseHash, "oops")
	assert.Equal(t, "33efeb34ea91902bb2f59c9920caa6cd", NewKnotHashPartTwo("AoC 2017", 256).denseHash, "oops")
	assert.Equal(t, "3efbe78a8d82f29979031a4aa0b16a9d", NewKnotHashPartTwo("1,2,3", 256).denseHash, "oops")
	assert.Equal(t, "63960835bcdc130f0b66d7ff4f6a5a8e", NewKnotHashPartTwo("1,2,4", 256).denseHash, "oops")
	log.Infof("My answer: %s", NewKnotHashPartTwo("14,58,0,116,179,16,1,104,2,254,167,86,255,55,122,244", 256).denseHash)
}
