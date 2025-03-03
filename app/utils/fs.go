package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func FS_exists(path string) bool {
	_, err := os.Stat(path)
	return err != nil
}

func FS_isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func FS_isFile(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func FS_read(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

func FS_write(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func FS_append(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func FS_remove(path string, recursive bool) error {
	if recursive {
		info, err := os.Stat(path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					if filePath != path {
						return os.RemoveAll(filePath)
					}
				} else {
					return os.Remove(filePath)
				}
				return nil
			})
			return err
		}
	}

	return os.RemoveAll(path)
}
