package cmd

import (
	"fmt"
	"strconv"
)

func parseTodoID(value string) (int, error) {
	id, err := strconv.Atoi(value)
	if err != nil || id < 1 {
		return 0, fmt.Errorf("invalid todo ID %q", value)
	}

	return id, nil
}
