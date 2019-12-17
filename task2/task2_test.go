package task2

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestInt(t *testing.T) {
	r, e := StrToInt("123")
	assert.Equal(t, 123, r, e)
	assert.Nil(t, e, r)
}

func TestFloat(t *testing.T) {
	r, e := StrToInt("1.0")
	assert.Equal(t, 1, r, e)
	assert.Nil(t, e, r)
}

func TestText(t *testing.T) {
	r, e := StrToInt("z11a")
	assert.Error(t, errors.New("expected integer"), e)
	assert.NotNil(t, e, r)
}

func TestSpaces(t *testing.T) {
	r, e := StrToInt("1 2 3")
	assert.Equal(t, 1, r, e)
	assert.Nil(t, e, r)
}

func Benchmark_Task2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToInt(strconv.Itoa(i))
	}
}

func Benchmark_Task1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToInt0(strconv.Itoa(i))
	}
}
