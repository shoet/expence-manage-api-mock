package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func DeserializeFromJSON(path string, dest any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed ReadFile %q: %v", path, err)
	}
	if err := json.Unmarshal(b, &dest); err != nil {
		return fmt.Errorf("failed JSON unmarshal %q: %v", path, err)
	}
	return nil
}
