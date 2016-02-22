package genstr

import (
	crand "crypto/rand"
	"encoding/base64"
	"io"
	mrand "math/rand"
	"time"
)

var (
	lunRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	lnRunes  = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	random   = mrand.New(mrand.NewSource(time.Now().UnixNano()))
)

func RandomLUNStr(n int) string {
	return randomStr(n, lunRunes)
}

func RandomLNStr(n int) string {
	return randomStr(n, lnRunes)
}

func randomStr(n int, s []rune) string {
	m2 := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		m2 = append(m2, byte(s[random.Int()%len(s)]))
	}
	return string(m2)
}

func SecureRandomStr(n int) string {
	k := make([]byte, n)
	if _, err := io.ReadFull(crand.Reader, k); err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(k)
}
