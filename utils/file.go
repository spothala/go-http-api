package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// WriteToFile - Write data to given input file
func WriteToFile(data string, file string) error {
	return WriteBytesToFile([]byte(data), file)
}

// WriteBytesToFile - Write data to given input file
func WriteBytesToFile(data []byte, file string) error {
	return ioutil.WriteFile(file, data, 777)
}

// ReadBytesFromFile - Reads file in byte format
func ReadBytesFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

// ReadFromFile - Reads file in string format
func ReadFromFile(file string) (string, error) {
	data, err := ReadBytesFromFile(file)
	return string(data), err
}

// DeleteFile - Deletes the file at given path
func DeleteFile(file string) error {
	return os.Remove(file)
}

// CopyFile copies the contents from src to dst atomically.
// If dst does not exist, CopyFile creates it with permissions perm.
// If the copy fails, CopyFile aborts and dst is preserved.
func CopyFile(dst, src string, perm os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	tmp, err := ioutil.TempFile(filepath.Dir(dst), "")
	if err != nil {
		return err
	}
	_, err = io.Copy(tmp, in)
	if err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return err
	}
	if err = tmp.Close(); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	if err = os.Chmod(tmp.Name(), perm); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	return os.Rename(tmp.Name(), dst)
}

// GetHomeDir - Gets home directory of the current user
func GetHomeDir() (homeDir string, err error) {
	hmDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return hmDir, nil
}

// Exists - Checks whether the given file path exists or not
func Exists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

// DownloadFromURL - Downloads the remote file in a given URL
func DownloadFromURL(url string) (int64, error) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	var output *os.File
	var err error
	if Exists(fileName) {
		output, err = os.Open(fileName)
	} else {
		output, err = os.Create(fileName)
	}
	if err != nil {
		return 0, err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return 0, err
	}
	return n, nil
}
