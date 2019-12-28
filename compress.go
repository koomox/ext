package ext

import (
	"archive/tar"
	"compress/gzip"
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
			err := createDir(filename)
			if err != nil {
				return completeFile, err
			}
		} else { // file
			file, err := createFile(filename)
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
			f, err := createFile(fn)
			if err != nil {
				return err
			}
			io.Copy(f, tr)

			return nil
		}
	}

	return errors.New("decompress failed")
}

func createDir(s string) (err error) {
	p := strings.TrimSuffix(s, "/")
	exist, err := PathExist(p)
	if err != nil {
		return
	}
	if !exist {
		return os.MkdirAll(p, 0755)
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	return os.Create(name)
}