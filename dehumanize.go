package dehumanize

import (
	"regexp"
	"strconv"
)

var bytePattern *regexp.Regexp
var kbytePattern *regexp.Regexp
var mbytePattern *regexp.Regexp
var gbytePattern *regexp.Regexp
var tbytePattern *regexp.Regexp
var pbytePattern *regexp.Regexp
var initErr error

// static initialization, the only way you know if it failed is the test
func init() {
	bytePattern, initErr = regexp.Compile(`^(\d+\.*\d*)[bB]`)
	if initErr != nil {
		return
	}
	kbytePattern, initErr = regexp.Compile(`^(\d+\.*\d*)[kK][bB]`)
	if initErr != nil {
		return
	}
	mbytePattern, initErr = regexp.Compile(`^(\d+\.*\d*)[mM][bB]`)
	if initErr != nil {
		return
	}
	gbytePattern, initErr = regexp.Compile(`^(\d+\.*\d*)[gG][bB]`)
	if initErr != nil {
		return
	}
	tbytePattern, initErr = regexp.Compile(`^(\d+\.*\d*)[tT][bB]`)
	if initErr != nil {
		return
	}
	pbytePattern, initErr = regexp.Compile(`^(\d+\.*\d*)[Pp][bB]`)
	if initErr != nil {
		return
	}
}

// use a regular expression to convert a human readable to an approximate number
func convertWithReg(s string, r *regexp.Regexp, base int64) int64 {
	ms := r.FindStringSubmatch(string(s))
	if ms != nil && len(ms) == 2 {
		size, err := strconv.ParseFloat(ms[1], 64)
		if err == nil {
			return int64(size * float64(base))
		}
	}
	return 0
}

// convert a supported size from human readable to an approximate size, if the base is 0, kb = 1024
// or you can specify it for your needs (eg marketing/metric kbytes of 1000bytes)
// this will return 0 for unconvertable values
func SizeConvert(s string, base int64) (size int64) {
	if base == 0 {
		base = 1024
	}
	size = convertWithReg(s, bytePattern, 1)
	if size != 0 {
		return size
	}
	size = convertWithReg(s, kbytePattern, base)
	if size != 0 {
		return size
	}
	size = convertWithReg(s, mbytePattern, base*base)
	if size != 0 {
		return size
	}
	size = convertWithReg(s, gbytePattern, base*base*base)
	if size != 0 {
		return size
	}
	size = convertWithReg(s, tbytePattern, base*base*base*base)
	if size != 0 {
		return size
	}
	size = convertWithReg(s, pbytePattern, base*base*base*base*base)
	if size != 0 {
		return size
	}
	return 0

}
