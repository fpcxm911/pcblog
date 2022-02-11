package database

import (
	"crypto/md5"
	"fmt"
	"io"
)

func EncryptForDb(decryptedPw string, username string) (encryptedPw string) {
	h := md5.New()
	io.WriteString(h, decryptedPw)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	salt1 := "@#$%"
	salt2 := "^&*()"
	io.WriteString(h, salt1)
	io.WriteString(h, username)
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)
	encryptedPw = fmt.Sprintf("%x", h.Sum(nil))
	return
}
