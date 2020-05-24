package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const (
	letterBytes   = `123456789afhjkoqrsuvwxyzAFHJKQRSUVWXYZ`
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// http://stackoverflow.com/a/31832326
func RandStringWithLengthLimit(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func Sha1(in string) string {
	hasher := sha1.New()
	hasher.Write([]byte(in))
	out := hasher.Sum(nil)

	return hex.EncodeToString(out)
}

func Sha256(in string) string {
	hasher := sha256.New()
	hasher.Write([]byte(in))
	out := hasher.Sum(nil)

	return hex.EncodeToString(out)
}

// orig: original string
// mask: maskCharacter
// revealLength: how many characters to reveal (from right to left)
//               0 or -1 will not reveal any
// length: if set, will set the result string length
//               0 or -1 will leave the length as the orig string length

func MaskString(orig, mask string, revealLength, length int) string {

	if orig == "" {
		return ""
	}

	if length <= 0 {
		length = len(orig)
	}

	str := ""
	for len(str) < length {
		str = str + mask
	}

	if revealLength == -1 {
		return str
	}

	// fmt.Println(str, length, revealLength, orig, len(orig))
	str = fmt.Sprintf("%s%s", str[:length-revealLength], orig[len(orig)-revealLength:])

	return str
}
