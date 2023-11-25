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
	if !confirmOverwrite(fileName, flagStore.Force) {
		return
	}

	// Create necessary directories
	createDirectories(fileName)

	// Create or overwrite the file
	f := createOrUpdateFile(fileName, flagStore)
	defer closeFile(f) // Close the file when this function is done

	// Set file mode
	setFileMode(f, os.FileMode(flagStore.Mode))

	// Write content to the file
	writeContentToFile(f, flagStore.Content)
}

// confirmOverwrite prompts the user for confirmation before overwriting an existing file
func confirmOverwrite(fileName string, forceOverwrite bool) bool {
	if !forceOverwrite {
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
	}
	return true
}

// createDirectories creates necessary directories for the file
func createDirectories(fileName string) {
	if err := os.MkdirAll(filepath.Dir(fileName), 0770); err != nil {
		logrus.Panic(err)
	}
	if filepath.Dir(fileName) != "." {
		logrus.Debugf("Created directory: %s", filepath.Dir(fileName))
	} else {
		logrus.Debug("Using the current directory")
	}
}

// createOrUpdateFile creates a new file or overwrites an existing file
func createOrUpdateFile(fileName string, flagStore model.Flagstore) *os.File {
	f, err := os.Create(fileName)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Debugf("Created file: %s", f.Name())
	return f
}

// setFileMode sets the mode of the file
func setFileMode(f *os.File, mode os.FileMode) {
	if err := f.Chmod(mode); err != nil {
		logrus.Panic(err)
	}
	logrus.Debug("Changed file mode to: ", mode)
}

// witeContentToFile writes content to the file
func writeContentToFile(f *os.File, content string) {
	if content != "" {
		f.Write([]byte(fmt.Sprintf("%s\n", content)))
		logrus.Debug("Wrote content to file: ", content)
	}
}

// closeFile closes the file
func closeFile(f *os.File) {
	defer f.Close()
}
