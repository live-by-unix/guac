package manager

import (
	"strings"
)

func splitLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func splitFields(s string) []string {
	return strings.Fields(s)
}
