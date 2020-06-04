package conv

import (
	"fmt"
	"io/ioutil"
)

func ReadNotes(dir string) error {
	_, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read %v: %w", dir, err)
	}

	return nil
}
