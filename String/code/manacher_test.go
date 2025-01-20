package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_preprocess(t *testing.T) {
	s1 := "abba"
	res1 := preprocess(s1)
	assert.Equal(t, "^#a#b#b#a#$", res1)

	s2 := ""
	res2 := preprocess(s2)
	assert.Equal(t, "^#$", res2)
}
