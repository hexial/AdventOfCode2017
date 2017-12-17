package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	//d := NewDance("sample", 5)
	//d.Move("s1")
	//d.State()
	//NewCircularBuffer(3, 9)
	assert.Equal(t, 638, NewCircularBuffer(3, 2017).NextVal())
	assert.Equal(t, 204, NewCircularBuffer(380, 2017).NextVal())
	//assert.Equal(t, 0, NewCircularBuffer(380, 50000000).ValPos1())

}
