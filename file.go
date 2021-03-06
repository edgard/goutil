package goutil

import (
	"fmt"
	"io"
	"os"
)

// IsDir checks if directory exists and is a directory
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// MoveFile copies and delete file from original location
func MoveFile(src string, dst string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if !dfi.Mode().IsRegular() {
			return fmt.Errorf("%s is not a regular file", dst)
		}
		if os.SameFile(sfi, dfi) {
			return err
		}
	}

	err = os.Rename(src, dst)
	if err == nil {
		return nil
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.Open(src)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		in.Close()
		return err
	}
	in.Close()

	return os.Remove(src)
}
