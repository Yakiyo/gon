package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// exists returns whether the given file or directory exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// read a file's content to string
func ReadFile(path string) (string, error) {
	if !PathExists(path) {
		return "", fmt.Errorf("no file exists at path %v", path)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// create directory at path if it doesnt exist
func EnsureDir(path string) error {
	if PathExists(path) {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}

// create file at path if it doesnt exist
func EnsureFile(path string, content string) error {
	if PathExists(path) {
		return nil
	}
	parent := filepath.Dir(path)
	err := EnsureDir(parent)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), os.ModePerm)
}
