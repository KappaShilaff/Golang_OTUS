package stringpkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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