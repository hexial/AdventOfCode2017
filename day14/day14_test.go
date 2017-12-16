package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	assert.Equal(t, 8108, CalcUsed("flqrgnkx"))
	assert.Equal(t, 8106, CalcUsed("oundnydw"))
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 1242, CalcRegions("flqrgnkx"))
	assert.Equal(t, 0, CalcRegions("oundnydw"))
}
