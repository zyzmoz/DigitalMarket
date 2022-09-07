package infra

import (
	"bufio"
	"os"
	"strings"

	"github.com/zyzmoz/DigitalMarket/usecases"
)

func Load(logger usecases.Logger) {
	filepath := ".env"

	file, err := os.Open(filepath)

	if err != nil {
		logger.LogError("%s", err)
	}
	defer file.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		logger.LogError("%s", err)
	}

	for _, line := range lines {
		pair := strings.Split(line, "=")
		os.Setenv(pair[0], pair[1])
	}
}
