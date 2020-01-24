package ext

import "encoding/base64"

type Encoding struct {
	cipher []byte
	clear []byte
}

func NewEncoding() *Encoding {
	return &Encoding{}
}

func (this *Encoding)Encrypt(data, secret []byte)(buf []byte, err error) {
	if this.cipher, err = Encrypt(data, secret); err != nil {
		return
	}

	buf = []byte(base64.RawStdEncoding.EncodeToString(this.cipher))
	return
}

func (this *Encoding)Decrypt(data, secret []byte)(buf []byte, err error) {
	if this.clear, err = base64.RawStdEncoding.DecodeString(string(data)); err != nil {
		return
	}
	return Decrypt(this.clear, secret)
}