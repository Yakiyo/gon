package archives

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

// Extract files from a zip file into `dest` directory
func Unzip(file *os.File, dest string) error {
	r, err := zip.OpenReader(file.Name())
	if err != nil {
		return err
	}
	for _, f := range r.File {
		if err := extract(f, dest); err != nil {
			return err
		}
	}
	return nil
}

// extract individual files
func extract(f *zip.File, dest string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	path := filepath.Join(dest, f.Name)

	// if its a dir, make the directory and exit early
	if f.FileInfo().IsDir() {
		os.MkdirAll(path, f.Mode())
		return nil
	}
	// otherwise we move to writing the file

	os.MkdirAll(filepath.Dir(path), f.Mode())
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, rc)
	if err != nil {
		return err
	}
	log.Debug("Extracting file", "file", f.Name)
	return nil
}
