package utils

import "go/format"

func FormatGo(src string) (string, error) {
	dest, err := format.Source([]byte(src))
	if err != nil {
		return "", err
	}
	return string(dest), nil
}
