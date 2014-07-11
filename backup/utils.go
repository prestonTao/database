package database

import (
	"os"
	"strings"
)

// Creates a file and writes the content into it.
func CreateAndWrite(filename, content string) int {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		Err("util", "CreateAndWrite", err)
		return CannotCreateFile
	}
	_, err = file.WriteString(content)
	if err != nil {
		Err("util", "CreateAndWrite", err)
		return CannotCreateFile
	}
	return OK
}

// Removes a line's occurances from a file.
func RemoveLine(filename, line string) int {
	// Open and read the file.
	file, err := os.Open(filename)
	if err != nil {
		Err("util", "RemoveLine", err)
		return CannotReadFile
	}
	fi, err := file.Stat()
	if err != nil {
		Err("util", "RemoveLine", err)
		return CannotReadFile
	}
	buffer := make([]byte, fi.Size)
	_, err = file.Read(buffer)
	if err != nil {
		Err("util", "RemoveLine", err)
		return CannotReadFile
	}
	file.Close()
	// Re-open the file and overwrite it.
	file, err = os.OpenFile(filename, os.O_WRONLY+os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		Err("util", "RemoveLine", err)
		return CannotReadFile
	}
	lines := strings.Split(string(buffer), "\n")
	for _, content := range lines {
		if strings.TrimSpace(content) != strings.TrimSpace(line) {
			_, err = file.WriteString(content + "\n")
			if err != nil {
				Err("util", "RemoveLine", err)
				return CannotWriteFile
			}
		}
	}
	return OK
}

// Returns a string which is the original string trimmed to the desired length.
// Trailing spaces are added if the string's length is too short.
// Otherwise, the string is truncated from right to the desired length.
func TrimLength(str string, length int) (trimmed string) {
	lengthDiff := length - len(str)
	if lengthDiff > 0 {
		trimmed = str + strings.Repeat(" ", lengthDiff)
	} else {
		trimmed = str[:length]
	}
	return trimmed
}

// Returns file name (without extension) and extension of a file name.
func FilenameParts(filename string) (name, extension string) {
	dotIndex := strings.LastIndex(filename, ".")
	if dotIndex == -1 || dotIndex == len(filename)-1 {
		name = filename
		extension = ""
	} else {
		name = filename[:dotIndex]
		extension = filename[dotIndex+1:]
	}
	return
}
