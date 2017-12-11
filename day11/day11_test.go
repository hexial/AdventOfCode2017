package day11

import (
	"AdventOfCode2017/util"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	origo := new(Hexagon)
	assert.Equal(t, 3, NewHexagon("ne,ne,ne").Distance(origo), "oops")
	assert.Equal(t, 0, NewHexagon("ne,ne,sw,sw").Distance(origo), "oops")
	assert.Equal(t, 2, NewHexagon("ne,ne,s,s").Distance(origo), "oops")
	assert.Equal(t, 3, NewHexagon("se,sw,se,sw,sw").Distance(origo), "oops")
	log.Infof("My answer: %d", NewHexagon(util.FileAsLineArray("input")[0]).Distance(origo))
}
