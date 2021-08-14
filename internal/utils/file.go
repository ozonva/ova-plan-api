package utils

import (
	"fmt"
	utils "github.com/ozonva/ova-plan-api/internal/utils/errors"
	"os"
)

type ReadFileCallback func(file *os.File) error

func closeFile(file *os.File) {
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
	file, err := os.Open(path)
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
