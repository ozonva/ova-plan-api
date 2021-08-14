package utils

type ReadFileError struct {
	filePath string
	err      error
}

func (fileError *ReadFileError) Error() string {
	return "Read file error: " + fileError.err.Error()
}

func (fileError *ReadFileError) Unwrap() error {
	return fileError.err
}

func NewReadFileError(filePath string, err error) error {
	return &ReadFileError{filePath: filePath, err: err}
}
