package ext

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io"
)

type Encoding struct {
	cipher []byte
	clear []byte
}

func NewEncoding() *Encoding {
	return &Encoding{}
}

func (this *Encoding)Compress(data []byte)(buf []byte, err error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	if _, err = w.Write(data); err != nil {
		return
	}
	if err = w.Close(); err != nil {
		return
	}

	buf = b.Bytes()
	return
}

func (this *Encoding)DeCompress(data []byte)(buf []byte, err error) {
	var (
		r io.ReadCloser
		buffer bytes.Buffer
	)
	b := bytes.NewReader(data)
	if r, err = zlib.NewReader(b); err != nil {
		return
	}
	if _, err = io.Copy(&buffer, r); err != nil {
		return
	}
	if err = r.Close(); err != nil {
		return
	}

	buf = buffer.Bytes()
	return
}

func (this *Encoding)Encrypt(data, secret []byte)(buf []byte, err error) {
	if this.cipher, err = this.Compress(data); err != nil {
		return
	}
	if this.cipher, err = Encrypt(this.cipher, secret); err != nil {
		return
	}

	buf = []byte(base64.RawStdEncoding.EncodeToString(this.cipher))
	return
}

func (this *Encoding)Decrypt(data, secret []byte)(buf []byte, err error) {
	if this.clear, err = base64.RawStdEncoding.DecodeString(string(data)); err != nil {
		return
	}
	if this.clear, err = Decrypt(this.clear, secret); err != nil {
		return
	}

	return this.DeCompress(this.clear)
}