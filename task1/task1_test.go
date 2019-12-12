package task1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	r, e := StrToInt("123")
	assert.Equal(t, 123, r, e)
	assert.Nil(t, e, r)
}

func TestInt2(t *testing.T) {
	r, e := StrToInt("01")
	assert.Equal(t, 1, r, e)
	assert.Nil(t, e, r)
}

func TestFloat(t *testing.T) {
	r, e := StrToInt("1.0")
	assert.NotNil(t, e, r)
}

func TestText(t *testing.T) {
	r, e := StrToInt("z11a")
	assert.NotNil(t, e, r)
}

func TestSpaces(t *testing.T) {
	r, e := StrToInt("1 2 3")
	assert.NotNil(t, e, r)
}
