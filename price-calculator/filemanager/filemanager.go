package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (filemanager FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(filemanager.InputFilePath)
	if err != nil {
		return nil, errors.New("Opening file failed")
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 10)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Reading file failed")
	}

	file.Close()
	return lines, nil
}

func (filemanager FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(filemanager.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON")
	}
	file.Close()

	// writer := bufio.NewWriter(file)
	// json, err := json.Marshal(data)
	// if err != nil {
	// 	file.Close()
	// 	return errors.New("Failed to convert data to JSON")
	// }
	// writer.Write(json)

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
