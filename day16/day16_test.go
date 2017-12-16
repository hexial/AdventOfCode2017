package day16

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	//d := NewDance("sample", 5)
	//d.Move("s1")
	//d.State()

	log.Infof("%d", Fix(100, 42))

	assert.Equal(t, "baedc", NewDance("sample", 5).ExecutePartOne())
	assert.Equal(t, "kgdchlfniambejop", NewDance("input", 16).ExecutePartOne())
	assert.Equal(t, "fjpmholcibdgeakn", NewDance("input", 16).ExecutePartTwo())
}
