package archives

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Yakiyo/gom/utils"
	"github.com/charmbracelet/log"
)

// extract files from tar `file` to directory `dest`
func Untar(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Unable to open tar file %v", err)
	}
	gzr, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("gzip reader failure %v", err)
	}
	defer gzr.Close()
	tr := tar.NewReader(gzr)
	var target string
	var name string
	for {
		header, err := tr.Next()

		// no more files
		if err == io.EOF {
			log.Debug("EOF")
			return nil
		}
		if err != nil {
			fmt.Println("\n\n\n", err == io.EOF, err.Error())
			return err
		}
		if header == nil {
			continue
		}
		name = header.Name
		target = filepath.Join(dest, name)

		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			err = utils.EnsureDir(target)
			if err != nil {
				return err
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
		log.Debug("extracting file", "file", name)
	}
}
