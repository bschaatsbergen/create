package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bschaatsbergen/create/pkg/model"
	"github.com/sirupsen/logrus"
)

// CreateFile creates a file with the given name
func CreateFile(fileName string, flagStore model.Flagstore) {
	if err := os.MkdirAll(filepath.Dir(fileName), 0770); err != nil {
		logrus.Panic(err)
	}
	if filepath.Dir(fileName) != "." {
		logrus.Debugf("Created directory: %s", filepath.Dir(fileName))
	} else {
		logrus.Debug("Using current directory")
	}
	f, err := os.Create(fileName)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Debugf("Created file: %s", f.Name())
	if err := f.Chmod(os.FileMode(flagStore.Mode)); err != nil {
		logrus.Panic(err)
	}
	logrus.Debug("Changed file mode to: ", os.FileMode(flagStore.Mode))
	if flagStore.Content != "" {
		f.Write([]byte(fmt.Sprintf("%s\n", flagStore.Content)))
		logrus.Debug("Wrote content to file: ", flagStore.Content)
	}
	defer f.Close()
}
