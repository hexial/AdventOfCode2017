package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	NewDiskDefrag("flqrgnkx").Debug(8)
	assert.Equal(t, 8108, NewDiskDefrag("flqrgnkx").used)
	assert.Equal(t, 8106, NewDiskDefrag("oundnydw").used)
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 1242, NewDiskDefrag("flqrgnkx").CalcRegions())
	assert.Equal(t, 1164, NewDiskDefrag("oundnydw").CalcRegions())
}
