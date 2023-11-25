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
	f, err := os.Create(fileName)
	if err != nil {
		logrus.Panic(err)
	}
	if err := f.Chmod(os.FileMode(flagStore.Mode)); err != nil {
		logrus.Panic(err)
	}
	if flagStore.Content != "" {
		f.Write([]byte(fmt.Sprintf("%s\n", flagStore.Content)))
	}
	defer f.Close()
	logrus.Debugf("Created file: %s", f.Name())
}
