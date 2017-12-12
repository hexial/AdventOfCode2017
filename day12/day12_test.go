package day12

import (
	"AdventOfCode2017/util"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPartOneAndTwo(t *testing.T) {
	assert.Equal(t, 6, len(NewVillage(util.FileAsLineArray("sample")).GetProgram(0).Group()), "oops")
	log.Infof("My answer: %d", len(NewVillage(util.FileAsLineArray("input")).GetProgram(0).Group()))
	log.Infof("My answer part two: %d", len(NewVillage(util.FileAsLineArray("input")).Group()))
}
