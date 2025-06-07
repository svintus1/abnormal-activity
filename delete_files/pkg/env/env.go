package env

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Parse(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	env := make(map[string]string)
	scanner := bufio.NewScanner(file)

	var currentKey string
	var currentValue strings.Builder
	inMultiline := false

	for scanner.Scan() {
		line := scanner.Text()

		if inMultiline {
			if strings.HasSuffix(line, "`") {
				currentValue.WriteString(line[:len(line)-1])
				env[currentKey] = currentValue.String()
				inMultiline = false
				currentKey = ""
				currentValue.Reset()
			} else {
				currentValue.WriteString(line + "\n")
			}
			continue
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		if parts := strings.SplitN(trimmed, "=", 2); len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])

			switch {
			case strings.HasPrefix(val, "`"):
				val = val[1:]
				if strings.HasSuffix(val, "`") {
					env[key] = val[:len(val)-1]
				} else {
					inMultiline = true
					currentKey = key
					currentValue.WriteString(val + "\n")
				}

			case strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") && len(val) >= 2:
				env[key] = val[1 : len(val)-1]

			default:
				if commentIdx := strings.Index(val, "#"); commentIdx != -1 {
					val = strings.TrimSpace(val[:commentIdx])
				}
				env[key] = val
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return env, nil
}

func PrintConfigVars(config map[string]string) {
	fmt.Println("Конфигурационные переменные и их значения:")

	keys := make([]string, 0, len(config))
	for key := range config {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s = %s\n", key, config[key])
	}
}
