package utils

import (
	"net"
	"strconv"
)

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func IsValidPort(portStr string) bool {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return false
	}
	return port >= 1 && port <= 65535
}

func CountChars(s string, char rune) int {
	count := 0
	for _, r := range s {
		if r == char {
			count++
		} else {
			break
		}
	}
	return count
}

func ContainsInSlice(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func IsValidService(validServices []string, service string) bool {
	return ContainsInSlice(validServices, service)
}

func AppendUniqueLines(existing, toAdd []string) []string {
	lineSet := make(map[string]struct{}, len(existing))
	for _, line := range existing {
		lineSet[line] = struct{}{}
	}
	for _, line := range toAdd {
		if _, exists := lineSet[line]; !exists {
			existing = append(existing, line)
			lineSet[line] = struct{}{}
		}
	}
	return existing
}
