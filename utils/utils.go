package utils

import (
	"encoding/json"
	"github/SanuKumar/go_ecom/types"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {

	}
	var payload types.RegisterUserPayload
	err := json.NewDecoder(r.Body).Decode(payload)
}
