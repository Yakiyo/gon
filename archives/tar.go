package archives

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

// extract files from tar `file` to directory `dest`
func Untar(file *os.File, dest string) error {
	tr := tar.NewReader(file)
	var target string
	for {
		header, err := tr.Next()

		// no more files
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		} else if header == nil {
			continue
		}

		target = filepath.Join(dest, header.Name)

		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
	return nil
}
