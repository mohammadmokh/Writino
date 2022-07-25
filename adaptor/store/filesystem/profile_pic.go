package filesystem

import (
	"bufio"
	"os"
)

func (f FsStore) SaveImage(image []byte, filename string) error {

	var fo *os.File
	var err error

	fulpath := f.BasePath + filename
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

func (f FsStore) DeleteImage(filename string) error {

	err := os.Remove(f.BasePath + filename)
	return err
}
