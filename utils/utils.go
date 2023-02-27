package utils

import (
	"fmt"
	"strings"
)

func IsCoinType(t string) bool {
	_, err := ExtractAddressFromType(t)
	if err != nil {
		return false
	}
	return true
}

func ExtractAddressFromType(t string) (string, error) {
	indexStart := strings.Index(t, "::")
	if indexStart == -1 {
		return "", fmt.Errorf("type is invalid")
	}
	item := t[0:indexStart]
	return item, nil
}

func ExtractFromResource(r string) (string, []string, error) {
	indexStart := strings.IndexByte(r, '<')
	if indexStart == -1 {
		return r, []string{}, nil
	}
	indexEnd := strings.LastIndexByte(r, '>')
	if indexEnd == -1 {
		return "", []string{}, fmt.Errorf("resource type is invalid")
	}
	if indexEnd <= indexStart || indexEnd != len(r)-1 {
		return "", []string{}, fmt.Errorf("resource type is invalid")
	}
	m := r[0:indexStart]
	t := r[indexStart+1 : indexEnd]
	//
	tStart, tEnd, tDepth := 0, 0, 0
	types := make([]string, 0)
	for tEnd < len(t) {
		if t[tEnd] == '<' {
			tDepth++
		}
		if t[tEnd] == '>' {
			tDepth--
		}
		if t[tEnd] == ',' && tDepth == 0 {
			types = append(types, t[tStart:tEnd])
			tStart = tEnd + 1
		}
		tEnd++
	}
	if tEnd > tStart {
		types = append(types, t[tStart:tEnd])
	}
	return m, types, nil
}
