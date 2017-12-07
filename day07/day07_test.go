package day07

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	assert.Equal(t, "tknk", Load("input.1.sample.txt").FindBase().Name, "oops")
	log.Infof("My answer: %s", Load("input.1.txt").FindBase().Name)
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 60, Load("input.1.sample.txt").FindBase().Balance(), "oops")
	Load("input.1.txt").FindBase().Balance()
}
