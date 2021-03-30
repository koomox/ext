package ext

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

func FileMD5sum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("GetFileMD5 OpenFile Err:%v", err.Error())
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("GetFileMD5 io.Copy Err:%v", err.Error())
	}

	sum := h.Sum(nil)
	return hex.EncodeToString(sum), nil
}

func FileSHA256sum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("GetFileSHA256 Open File Err:%v", err.Error())
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("GetFileSHA256 io.Copy Err:%v", err.Error())
	}

	sum := h.Sum(nil)
	return hex.EncodeToString(sum), nil
}
