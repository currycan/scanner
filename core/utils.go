package core

import "os"

func IsFileExist(fileName string) (error, bool) {
	_, err := os.Stat(fileName)
	if err == nil {
		return nil, true
	}
	if os.IsNotExist(err) {
		return nil, false
	}
	return err, false
}
