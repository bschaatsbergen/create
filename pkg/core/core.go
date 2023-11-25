package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bschaatsbergen/create/pkg/model"
	"github.com/sirupsen/logrus"
)

// CreateFile is the entry point to create a file with the given name and options
func CreateFile(fileName string, flagStore model.Flagstore) {
	// Check for file existence and confirmation
	if !ConfirmOverwrite(fileName) {
		return
	}

	// Create necessary directories
	CreateDirectories(fileName)

	// Create or overwrite the file
	f := CreateOrUpdateFile(fileName, flagStore)
	defer CloseFile(f) // Close the file when this function is done

	// Set file mode
	SetFileMode(f, os.FileMode(flagStore.Mode))

	// Write content to the file
	WriteContentToFile(f, flagStore.Content)
}

// ConfirmOverwrite prompts the user for confirmation before overwriting an existing file
func ConfirmOverwrite(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("File '%s' already exists. Do you want to overwrite it? (yes/no): ", fileName)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := scanner.Text()

		if answer != "yes" {
			logrus.Warn("File creation aborted.")
			return false
		}
	}
	return true
}

// CreateDirectories creates necessary directories for the file
func CreateDirectories(fileName string) {
	if err := os.MkdirAll(filepath.Dir(fileName), 0770); err != nil {
		logrus.Panic(err)
	}
	if filepath.Dir(fileName) != "." {
		logrus.Debugf("Created directory: %s", filepath.Dir(fileName))
	} else {
		logrus.Debug("Using the current directory")
	}
}

// CreateOrUpdateFile creates a new file or overwrites an existing file
func CreateOrUpdateFile(fileName string, flagStore model.Flagstore) *os.File {
	f, err := os.Create(fileName)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Debugf("Created file: %s", f.Name())
	return f
}

// SetFileMode sets the mode of the file
func SetFileMode(f *os.File, mode os.FileMode) {
	if err := f.Chmod(mode); err != nil {
		logrus.Panic(err)
	}
	logrus.Debug("Changed file mode to: ", mode)
}

// WriteContentToFile writes content to the file
func WriteContentToFile(f *os.File, content string) {
	if content != "" {
		f.Write([]byte(fmt.Sprintf("%s\n", content)))
		logrus.Debug("Wrote content to file: ", content)
	}
}

// CloseFile closes the file
func CloseFile(f *os.File) {
	defer f.Close()
}
