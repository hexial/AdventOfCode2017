package day13

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPartOneAndTwo(t *testing.T) {
	assert.Equal(t, 1, Steps(1))
	assert.Equal(t, 3, Steps(2))
	assert.Equal(t, 5, Steps(3))
	assert.Equal(t, 7, Steps(4))

	for i := 0; i < 9; i++ {
		if i == 0 || i == 4 || i == 8 {
			assert.Equal(t, true, TestHit(3, i, 0))
		} else {
			assert.Equal(t, false, TestHit(3, i, 0))
		}
	}

	assert.Equal(t, 24, NewFirewall("sample").Severity())
	assert.Equal(t, 2160, NewFirewall("input").Severity())
	//log.Infof("*********************'")
	assert.Equal(t, 10, NewFirewall("sample").Delay())
	log.Infof("My answer: %d", NewFirewall("input").Delay())
}
