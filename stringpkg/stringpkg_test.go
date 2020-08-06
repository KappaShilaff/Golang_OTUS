package stringpkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrodsmap(t *testing.T) {
	assert.Equal(t, tenwords_map("kek lol kek lol kek 1 he 2 he 2 he 2 he 2 he 3 5 3 5 3 5 3 5 3 5 3 5 5 4 4 4 4 4 4 4 4 6 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7 7 7 7"), []string{"7", "6", "4", "5", "3", "he", "2", "kek", "lol", "1"})
	assert.Equal(t, tenwords_map("da o o da da da kek"), []string{"da", "o", "kek"})
}

func TestWrods(t *testing.T) {
	assert.Equal(t, tenwords("kek lol kek lol kek 1 he 2 he 2 he 2 he 2 he 3 5 3 5 3 5 3 5 3 5 3 5 5 4 4 4 4 4 4 4 4 6 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7 7 7 7"), []string{"7", "6", "4", "5", "3", "he", "2", "kek", "lol", "1"})
	assert.Equal(t, tenwords("da o o da da da kek"), []string{"da", "o", "kek"})
}

func TestString(t *testing.T) {
	assert.Equal(t, stringpkg("asd"), "asd")
	assert.Equal(t, stringpkg("a4b\\4\\5c0d5e\\52qwe\\4\\5"), "aaaab45ddddde55qwe45")
	assert.Equal(t, stringpkg(""), "")
	assert.Equal(t, stringpkg("asd4"), "asdddd")
	assert.Equal(t, stringpkg("a4sd4"), "aaaasdddd")
	assert.Equal(t, stringpkg("qwe\\4\\5"), "qwe45")
	assert.Equal(t, stringpkg("qwe\\45"), "qwe44444")
	assert.Equal(t, stringpkg("qwe\\\\5"), "qwe\\\\\\\\\\")
	assert.Equal(t, stringpkg("asd0"), "as")
}

func TestItoaPositive(t *testing.T) {
	assert.Equal(t, itoa(5), "5")
	assert.Equal(t, itoa(0), "0", "string 0")
	assert.Equal(t, itoa(15), "15", "string 15")
	assert.Equal(t, itoa(9223372036854775807), "9223372036854775807", "string 9223372036854775807")
}

func TestItoaIsNegative(t *testing.T) {
	assert.Equal(t, itoa(-5), "-5", "string -5")
	assert.Equal(t, itoa(-15), "-15", "string -15")
	assert.Equal(t, itoa(-9223372036854775808), "-9223372036854775808", "string -9223372036854775808")
}

func BenchmarkMap(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		tenwords_map("kek lol kek lol kek 1 he 2 he 2 he 2 he 2 he 3 5 3 5 3 5 3 5 3 5 3 5 5 4 4 4 4 4 4 4 4 6 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7 7 7 7")
		tenwords_map("da o o da da da kek")
	}
}

func BenchmarkNomap(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		tenwords("kek lol kek lol kek 1 he 2 he 2 he 2 he 2 he 3 5 3 5 3 5 3 5 3 5 3 5 5 4 4 4 4 4 4 4 4 6 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7 7 7 7")
		tenwords("da o o da da da kek")
	}
}