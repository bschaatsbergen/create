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
	if !confirmOverwrite(fileName, flagStore.Force) {
		return
	}

	createDirectories(fileName)

	f := createFile(fileName)
	defer closeFile(f)

	setFileMode(f, os.FileMode(flagStore.Mode))

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

func createFile(fileName string) *os.File {
	f, err := os.Create(fileName)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Debugf("Created file: %s", f.Name())
	return f
}

func setFileMode(f *os.File, mode os.FileMode) {
	if err := f.Chmod(mode); err != nil {
		logrus.Panic(err)
	}
	logrus.Debug("Changed file mode to: ", mode)
}

func writeContentToFile(f *os.File, content string) {
	if content != "" {
		f.Write([]byte(fmt.Sprintf("%s\n", content)))
		logrus.Debug("Wrote content to file: ", content)
	}
}

func closeFile(f *os.File) {
	defer f.Close()
}
