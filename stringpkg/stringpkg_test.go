package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordsMap(t *testing.T) {
	assert.Equal(t, tenWordsMap("kek lol kek lol kek 1 he 2 he 2 he 2 he 2 he 3 5 3 5 3 5 3 5 3 5 3 5 5 4 4 4 4 4 4 4 4 6 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7 7 7 7"), []string{"7", "6", "4", "5", "3", "he", "2", "kek", "lol", "1"})
	assert.Equal(t, tenWordsMap("da o o da da da kek"), []string{"da", "o", "kek"})
}

func TestWords(t *testing.T) {
	assert.Equal(t, tenWords("kek lol kek lol kek 1 he 2 he 2 he 2 he 2 he 3 5 3 5 3 5 3 5 3 5 3 5 5 4 4 4 4 4 4 4 4 6 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7 7 7 7"), []string{"7", "6", "4", "5", "3", "he", "2", "kek", "lol", "1"})
	assert.Equal(t, tenWords("da o o da da da kek"), []string{"da", "o", "kek"})
}

func TestString(t *testing.T) {
	assert.Equal(t, StringPkg("asd"), "asd")
	assert.Equal(t, StringPkg("a4b\\4\\5c0d5e\\52qwe\\4\\5"), "aaaab45ddddde55qwe45")
	assert.Equal(t, StringPkg(""), "")
	assert.Equal(t, StringPkg("asd4"), "asdddd")
	assert.Equal(t, StringPkg("a4sd4"), "aaaasdddd")
	assert.Equal(t, StringPkg("qwe\\4\\5"), "qwe45")
	assert.Equal(t, StringPkg("qwe\\45"), "qwe44444")
	assert.Equal(t, StringPkg("qwe\\\\5"), "qwe\\\\\\\\\\")
	assert.Equal(t, StringPkg("asd0"), "as")
}

func TestItoaPositive(t *testing.T) {
	assert.Equal(t, Itoa(5), "5")
	assert.Equal(t, Itoa(0), "0", "string 0")
	assert.Equal(t, Itoa(15), "15", "string 15")
	assert.Equal(t, Itoa(9223372036854775807), "9223372036854775807", "string 9223372036854775807")
}

func TestItoaIsNegative(t *testing.T) {
	assert.Equal(t, Itoa(-5), "-5", "string -5")
	assert.Equal(t, Itoa(-15), "-15", "string -15")
	assert.Equal(t, Itoa(-9223372036854775808), "-9223372036854775808", "string -9223372036854775808")
}

func BenchmarkMap(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		tenWordsMap("kek lol kek lol kek 1 he 2 he 2 he 2 he 2 he 3 5 3 5 3 5 3 5 3 5 3 5 5 4 4 4 4 4 4 4 4 6 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7 7 7 7")
		tenWordsMap("da o o da da da kek")
	}
}

func BenchmarkNoMap(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		tenWords("kek lol kek lol kek 1 he 2 he 2 he 2 he 2 he 3 5 3 5 3 5 3 5 3 5 3 5 5 4 4 4 4 4 4 4 4 6 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7 7 7 7")
		tenWords("da o o da da da kek")
	}
}