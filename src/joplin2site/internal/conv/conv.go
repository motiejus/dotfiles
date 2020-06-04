package conv

import (
	"fmt"
	"io/ioutil"
)

func ReadDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read %v: %w", dir, err)
	}

	return nil
}
