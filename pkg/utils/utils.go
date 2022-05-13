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

func ParseID(id string) (int64, error) {
	dbId, err := strconv.Atoi(id) // FIXME error handling
	if err != nil {
		fmt.Println("Error while parsing.", err)
	}
	parsedId := int64(dbId)
	return parsedId, err
}
