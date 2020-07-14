package utils

import (
	"archive/zip"
	"io"
	"os"
)

// FileToZipNew --
func FileToZipNew(srcFile, zipName string) (err error) {
	fr, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer fr.Close()
	fi, err := fr.Stat()
	if err != nil {
		return err
	}
	fh, err := zip.FileInfoHeader(fi)
	if err != nil {
		return err
	}
	fh.Method = zip.Deflate

	fw, err := os.OpenFile(zipName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer fw.Close()

	zw := zip.NewWriter(fw)
	gw, err := zw.CreateHeader(fh)
	defer zw.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(gw, fr)
	if err != nil {
		return err
	}
	err = zw.Flush()
	if err != nil {
		return err
	}
	return nil
}
