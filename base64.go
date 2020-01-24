package ext

import "encoding/base64"

type Encoding struct {
	// contains filtered or unexported fields
}

func NewEncoding() *Encoding {
	return &Encoding{}
}

func (this *Encoding)Encrypt(data, secret []byte)(buf []byte, err error) {
	var (
		cipherText []byte
	)
	if cipherText, err = Encrypt(data, secret); err != nil {
		return
	}

	buf = []byte(base64.RawStdEncoding.EncodeToString(cipherText))
	return
}

func (this *Encoding)Decrypt(data, secret []byte)(buf []byte, err error) {
	var (
		cipherText []byte
	)
	if cipherText, err = base64.RawStdEncoding.DecodeString(string(data)); err != nil {
		return
	}
	return Decrypt(cipherText, secret)
}