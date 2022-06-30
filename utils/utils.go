package utils

import (
	"io"
	"log"
	"net/http"
)

func Sanitize2D(arr2D [][]string) []string {
	temp := make([]string, 0)
	emptySet := make(map[string]bool)
	for _, val := range arr2D {
		if len(val[1]) > 0 {
			if _, ok := emptySet[val[1]]; !ok {
				emptySet[val[1]] = true
				temp = append(temp, val[1])
			}
		}
	}
	return temp
}

func Sanitize1D(arr1D []string) []string {
	temp := make([]string, 0)
	emptySet := make(map[string]bool)
	for _, val := range arr1D {
		if len(val) > 0 {
			if _, ok := emptySet[val]; !ok {
				emptySet[val] = true
				temp = append(temp, val)
			}
		}
	}
	return temp
}

func FetchHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		return string(body)
	} else {
		log.Fatal("Error: " + resp.Status)
	}
	return ""
}
