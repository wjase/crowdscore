package apis

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// RestStatus obj for json status
type RestStatus struct {
	Success bool
	Message string
}

// WriteRestStatus write the status as a rest obj
func WriteRestStatus(success bool, msg string, w *http.ResponseWriter) {
	WriteRestResponse(&RestStatus{Success: success, Message: msg}, w)
}

// WriteRestResponse write the specified object as json
func WriteRestResponse(toWrite interface{}, w *http.ResponseWriter) (err error) {
	toWriteBytes, _ := json.Marshal(&toWrite)
	n, err := fmt.Fprintf(*w, string(toWriteBytes))
	switch {
	case err != nil:
		return err
	case n != len(toWriteBytes):
		return errors.New("Didn't write all of response" + string(toWriteBytes))
	default:
		return nil
	}
}
