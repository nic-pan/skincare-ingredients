package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ParseBodyToMap(r *http.Request) (map[string]json.RawMessage, error) {
	var objmap map[string]json.RawMessage
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), &objmap); err != nil {
			return objmap, err
		}
	}
	return objmap, nil
}

func ParseStringToArray(arrString string) ([]string, error) {
	var arr []string
	err := json.Unmarshal([]byte(arrString), &arr)
	return arr, err
}

func TrimQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

func ParseID(ID string) (uint, error) {
	dbID, err := strconv.Atoi(ID) // FIXME error handling
	if err != nil {
		fmt.Println("Error while parsing.", err)
	}
	// parsedId := int64(dbId)
	parsedID := uint(dbID)
	return parsedID, err
}
