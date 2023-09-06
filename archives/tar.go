package archives

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Yakiyo/gom/utils"
	"github.com/charmbracelet/log"
)

// extract files from tar `file` to directory `dest`
func Untar(file *os.File, dest string) error {
	// gzr, err := gzip.NewReader(file)
	// if err != nil {
	// 	fmt.Println("here")
	// 	return err
	// }
	// defer gzr.Close()
	tr := tar.NewReader(file)
	var target string
	var name string
	for {
		fmt.Println("in iter")
		header, err := tr.Next()

		// no more files
		if err == io.EOF {
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
