package ext

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func GenPassword(pwd string) string {
	return MD5sum(SHA1sum(pwd))
}

func MD5sum(ctx string) string {
	h := md5.New()
	io.WriteString(h, ctx)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func SHA1sum(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func SHA256sum(s string) string {
	h := sha256.New()
	io.WriteString(h, s)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func MD5sumWithFile(pa string) (string, error) {
	f, err := os.Open(pa)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	sum := h.Sum(nil)
	return hex.EncodeToString(sum), nil
}

func SHA256sumWithFile(pa string) (string, error) {
	f, err := os.Open(pa)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	sum := h.Sum(nil)
	return hex.EncodeToString(sum), nil
}
