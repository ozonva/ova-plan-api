package utils

import (
	"errors"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"io/ioutil"
	"testing"
)

type fileInfo struct {
	name    string
	content string
}

var badNewError = errors.New("bad news")

var (
	fileOne       = fileInfo{name: "testOne.txt", content: "ONE"}
	fileTwo       = fileInfo{name: "testTwo.txt", content: "TWO"}
	fileNotExists = fileInfo{name: "testNotExist.txt", content: "testNotExist"}
)

func initTest() {
	AppFs = afero.NewMemMapFs()

	afero.WriteFile(AppFs, fileOne.name, []byte(fileOne.content), 0644)
	afero.WriteFile(AppFs, fileTwo.name, []byte(fileTwo.content), 0644)
}

func TestReadFile(t *testing.T) {
	tables := []struct {
		info        fileInfo
		expectedErr error
		callback    ReadFileCallback
	}{
		{fileOne, nil, assertReadFileCallback(t, fileOne)},
		{fileNotExists, fs.ErrNotExist, assertReadFileCallback(t, fileNotExists)},
		{fileTwo, badNewError, readFileCallbackWithError()},
	}

	initTest()
	for _, table := range tables {
		err := ReadFile(table.info.name, table.callback)
		if table.expectedErr == nil {
			assert.NoError(t, err)
		} else {
			assert.True(t, errors.Is(err, table.expectedErr))
		}
	}
}

func TestReadFiles(t *testing.T) {
	dataSlice := make([]byte, 0)

	tables := []struct {
		fileNames      []string
		expectedErr    error
		callback       ReadFileCallback
		expectedResult string
	}{
		{[]string{fileOne.name, fileTwo.name}, nil, readFileToSliceCallback(&dataSlice), "ONETWO"},
		{[]string{fileOne.name, fileNotExists.name}, fs.ErrNotExist, readFileToSliceCallback(&dataSlice), "ONE"},
		{[]string{fileOne.name}, badNewError, readFileCallbackWithError(), ""},
	}

	initTest()

	for _, table := range tables {
		dataSlice = make([]byte, 0)
		err := ReadFiles(table.callback, table.fileNames...)

		if table.expectedErr == nil {
			assert.NoError(t, err)
		} else {
			assert.True(t, errors.Is(err, table.expectedErr))
		}
		assert.Equal(t, table.expectedResult, string(dataSlice))
	}
}

func readFileToSliceCallback(dataSlice *[]byte) ReadFileCallback {
	return func(file SimpleFileHandler) error {
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		*dataSlice = append(*dataSlice, bytes...)
		return nil
	}
}

func assertReadFileCallback(t *testing.T, info fileInfo) ReadFileCallback {
	return func(file SimpleFileHandler) error {
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		assert.Equal(t, info.content, string(bytes))
		return nil
	}
}

func readFileCallbackWithError() ReadFileCallback {
	return func(file SimpleFileHandler) error {
		return badNewError
	}
}
