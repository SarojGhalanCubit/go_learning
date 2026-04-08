package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DecodeJSON is a wrapper around json.NewDecoder that provides friendly error messages
func DecodeJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		// 1. Handle Type Mismatch (e.g., sending "" for an int)
		if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
			errMsg := fmt.Sprintf("Field '%s' expects a %s, but you gave '%s'",
				unmarshalErr.Field, unmarshalErr.Type.String(), unmarshalErr.Value)

			WriteError(w, http.StatusBadRequest, "Type Mismatch", errMsg)
			return err
		}

		// 2. Handle Syntax Errors (e.g., missing a comma or bracket)
		if _, ok := err.(*json.SyntaxError); ok {
			WriteError(w, http.StatusBadRequest, "Invalid JSON", "Your JSON syntax is incorrect. Check for missing commas or braces.")
			return err
		}

		// 3. Handle Generic Errors (e.g., empty body)
		WriteError(w, http.StatusBadRequest, "Bad Request", err.Error())
		return err
	}
	return nil
}
