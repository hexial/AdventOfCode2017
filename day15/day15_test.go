package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 588, Duel(NewGenerator(16807, 65), NewGenerator(48271, 8921), 40000000))
	assert.Equal(t, 619, Duel(NewGenerator(16807, 591), NewGenerator(48271, 393), 40000000))
	assert.Equal(t, 309, Duel(NewGeneratorMultiple(16807, 65, 4), NewGeneratorMultiple(48271, 8921, 8), 5000000))
	assert.Equal(t, 290, Duel(NewGeneratorMultiple(16807, 591, 4), NewGeneratorMultiple(48271, 393, 8), 5000000))
}
