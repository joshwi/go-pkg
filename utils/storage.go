package utils

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joshwi/go-pkg/logger"
)

var write_validation_1 = regexp.MustCompile(`^(\/[a-zA-Z0-9\-\_]{1,20})+$`)
var write_validation_2 = regexp.MustCompile(`^[a-zA-Z0-9\-\_]{0,20}\.(csv|txt|json)$`)
var eof_validation = regexp.MustCompile(`(?i)(\/[a-zA-Z0-9\-\_]+\.\w+$)`)

//Scan a directory for files and subfolders
func Scan(directory string) ([]string, error) {

	output := []string{}

	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		rel_path := strings.ReplaceAll(path, directory, "")
		output = append(output, rel_path)
		return err
	})

	if err != nil {
		logger.Logger.Error().Str("directory", directory).Str("status", "Failed").Err(err).Msg("Scan")
		return nil, err
	} else {
		logger.Logger.Info().Str("directory", directory).Str("status", "Success").Msg("Scan")
	}

	return output, nil
}

func Copy(source string, target string) error {
	srcFile, err := os.Open(source)
	defer srcFile.Close()

	_, err = os.Stat(target)
	if os.IsNotExist(err) {
		// Creates any directories in the path that don't exist
		err = os.MkdirAll(path.Dir(target), 0755)
		if err != nil {
			logger.Logger.Error().Str("source", source).Str("destination", target).Str("status", "Failed").Err(err).Msg("Copy")
			return err
		}
	}

	destFile, err := os.Create(target) // creates if file doesn't exist
	defer destFile.Close()

	// Move the file to new location
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		logger.Logger.Error().Str("source", source).Str("destination", target).Str("status", "Failed").Err(err).Msg("Copy")
		return err
	}

	logger.Logger.Info().Str("source", source).Str("destination", target).Str("status", "Success").Msg("Copy")

	return nil
}

// Move a file to a new directory
func Move(source string, destination string) error {

	// Check if the file path exists
	_, err := os.Stat(destination)
	if os.IsNotExist(err) {
		// Creates any directories in the path that don't exist
		err = os.MkdirAll(path.Dir(destination), 0755)
		if err != nil {
			logger.Logger.Error().Str("source", source).Str("destination", destination).Str("status", "Failed").Err(err).Msg("Move")
			return err
		}
	}

	// Move the file to new location
	err = os.Rename(source, destination)
	if err != nil {
		logger.Logger.Error().Str("source", source).Str("destination", destination).Str("status", "Failed").Err(err).Msg("Move")
		return err
	}

	logger.Logger.Info().Str("source", source).Str("destination", destination).Str("status", "Success").Msg("Move")

	return nil

}

//Read contents of a file
func Read(filename string) ([]byte, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Logger.Error().Str("file", filename).Str("status", "Failed").Err(err).Msg("Read")
		return nil, err
	} else {
		logger.Logger.Info().Str("file", filename).Str("status", "Success").Msg("Read")
	}

	return data, nil

}

//Write contents to a file
func Write(filename string, data []byte, mode int) error {

	// Check if file already exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// Creates any directories that don't exist
		err = os.MkdirAll(filepath.Dir(filename), 0755)
		if err != nil {
			logger.Logger.Error().Str("file", filename).Str("status", "Failed").Err(err).Msg("Write")
			return err
		}
		// Creates file
		_, err = os.Create(filename)
		if err != nil {
			logger.Logger.Error().Str("file", filename).Str("status", "Failed").Err(err).Msg("Write")
			return err
		}
	}

	// Writes byte data to the file
	err = ioutil.WriteFile(filename, data, os.FileMode(mode))
	if err != nil {
		logger.Logger.Error().Str("file", filename).Str("status", "Failed").Err(err).Msg("Write")
		return err
	} else {
		logger.Logger.Info().Str("file", filename).Str("status", "Success").Msg("Write")
	}

	return nil

}
