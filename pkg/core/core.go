package core

import (
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
	defer f.Close()
	logrus.Debugf("Created file: %s", f.Name())
}
