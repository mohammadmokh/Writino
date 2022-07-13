package filesystem

import (
	"bufio"
	"os"
)

func (f FsStore) SaveImage(image []byte, path string) error {

	var fo *os.File
	var err error

	fulpath := f.BasePath + "/" + path
	// if file dosn't exists we will create it
	if fo, err = os.OpenFile(fulpath, os.O_APPEND|os.O_WRONLY, os.ModeAppend); err != nil {

		fo, err = os.Create(fulpath)
		if err != nil {
			return err
		}
	}

	fw := bufio.NewWriter(fo)

	_, err = fw.Write(image)
	return err
}

func (f FsStore) DeleteImage(path string) error {

	err := os.Remove(f.BasePath + "/" + path)
	return err
}
