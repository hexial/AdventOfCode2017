package day08

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	assert.Equal(t, 1, NewCPU("input.1.sample.txt").LargestRegisterValue(), "oops")
	log.Infof("My answer: %d", NewCPU("input.1.txt").LargestRegisterValue())
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 10, NewCPU("input.1.sample.txt").alltTimeHigh, "oops")
	log.Infof("My answer: %d", NewCPU("input.1.txt").alltTimeHigh)
}
