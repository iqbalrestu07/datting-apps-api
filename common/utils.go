package common

import (
	"encoding/json"
	"log"
	"runtime"
	"strconv"
	"strings"
)

func LogErrorWithLine(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("[%s:%d] Error: %v\n", file, line, err)
	}
}

func LogPretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("error:", err)
	}
	log.Print(string(b) + "\n")
}

func IntSliceToCommaSeparatedString(slice []int64) string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(result, ",")
}
