package ext

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

// Decompress .tar.gz file
func DeCompress(tarFile, dest string) ([]string, error) {
	completeFile := make([]string, 0)
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return completeFile, err
	}
	defer srcFile.Close()

	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return completeFile, err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return completeFile, err
			}
		}
		filename := dest + hdr.Name
		if strings.HasSuffix(filename, "/") { // Dir
			_, err := os.Create(filename)
			if err != nil {
				return completeFile, err
			}
		} else { // file
			file, err := os.Create(filename)
			if err != nil {
				return completeFile, err
			}
			io.Copy(file, tr)
		}
		completeFile = append(completeFile, filename)
	}
	return completeFile, nil
}

func DeCompressFile(tarFile, suffix, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if strings.HasSuffix(hdr.Name, suffix) {
			fn := path.Join(dest, suffix)
			f, err := os.Create(fn)
			if err != nil {
				return err
			}
			io.Copy(f, tr)

			return nil
		}
	}

	return errors.New("decompress failed")
}

func CompressWithBase64(b []byte) ([]byte, error) {
	enc := func(data []byte) (buf []byte, err error) {
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

	p, err := enc(b)
	if err != nil {
		return p, err
	}

	buf := make([]byte, base64.RawStdEncoding.EncodedLen(len(p)))
	base64.RawStdEncoding.Encode(buf, p)

	return buf, nil
}

func DeCompressWithBase64(b []byte) ([]byte, error) {
	dec := func(data []byte) (buf []byte, err error) {
		var (
			r      io.ReadCloser
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

	buf := make([]byte, base64.RawStdEncoding.DecodedLen(len(b)))
	if _, err := base64.RawStdEncoding.Decode(buf, b); err != nil {
		return nil, err
	}

	return dec(buf)
}
