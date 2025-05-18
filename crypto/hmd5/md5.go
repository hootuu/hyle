package hmd5

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
)

func Join(prefix string, split string, paras ...string) string {
	if len(paras) == 0 {
		return ""
	}
	h := md5.New()
	h.Write([]byte(prefix))
	for i, s := range paras {
		if i > 0 {
			h.Write([]byte(split))
		}
		h.Write([]byte(s))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MD5(str string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, str)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func MD5Bytes(str string) []byte {
	hash := md5.New()
	_, _ = io.WriteString(hash, str)
	return hash.Sum(nil)
}

func Is(str string) bool {
	matched, err := regexp.MatchString("^[0-9a-fA-F]{32}$", str)
	if err != nil {
		return false
	}
	return matched
}
