package utils

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80
	}
	fields := strings.Fields(string(out))
	if len(fields) != 2 {
		return 80
	}
	width, err := strconv.Atoi(fields[1])
	if err != nil {
		return 80
	}
	return width
}
