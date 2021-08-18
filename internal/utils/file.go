package utils

import (
	"fmt"
	utils "github.com/ozonva/ova-plan-api/internal/utils/errors"
	"github.com/spf13/afero"
	"io"
)

type SimpleFileHandler interface {
	io.Reader
	Name() string
}

type ReadFileCallback func(file SimpleFileHandler) error

var AppFs = afero.NewOsFs()

func closeFile(file io.Closer) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func ReadFiles(callback ReadFileCallback, paths ...string) error {
	for _, path := range paths {
		err := ReadFile(path, callback)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFile(path string, callback ReadFileCallback) error {
	file, err := AppFs.Open(path)
	if err != nil {
		return utils.NewReadFileError(path, err)
	}
	defer closeFile(file)

	callbackError := callback(file)

	if callbackError != nil {
		return fmt.Errorf("an error occured while callback called: \n\t%w\n", callbackError)
	}
	return nil
}
