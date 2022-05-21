package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ScanInt() int64 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return i
}

func ScanText() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()
	return t
}
