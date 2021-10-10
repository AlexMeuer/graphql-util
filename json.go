package gql

import (
	"encoding/json"
)

type BinderJSON interface {
	BindJSON(interface{}) error
}

// UnmarshalHasuraAction unmarshals Hasura Action Inputs to the given interface and returns the ID of the user.
//
// It is expected that the Hasura Action will be configured with a single parameter type, named "params".
// The contents of "params" will be unmarshalled into the given interface.
// The BinderJSON "b" argument is used for initial decoding but it not used in conjuction with the interface "i".
// The reason for this is because this function was originally written for the gin web framework (specifically
// the `context.BindJSON(..)` function.)
func UnmarshalHasuraAction(b BinderJSON, i interface{}) (string, error) {
	var body struct {
		SessionVars struct {
			UserID string `json:"x-hasura-user-id"`
		} `json:"session_variables"`
		Input struct {
			Params json.RawMessage `json:"params"`
		} `json:"input"`
	}
	if err := b.BindJSON(&body); err != nil {
		return body.SessionVars.UserID, err
	}
	err := json.Unmarshal(body.Input.Params, i)
	return body.SessionVars.UserID, err
}

// UnmarshalHasuraEvent unmarshals a Hasura Event payload into either/both of the given
// new/old structs and returns the operation type that caused the event.
// The BinderJSON "b" argument is used for initial decoding but it not used in conjuction with the interface "i".
// The reason for this is because this function was originally written for the gin web framework (specifically
func UnmarshalHasuraEvent(b BinderJSON, n, o interface{}) (op string, err error) {
	var body struct {
		Event struct {
			Op   string `json:"op"`
			Data struct {
				Old json.RawMessage `json:"old"`
				New json.RawMessage `json:"new"`
			} `json:"data"`
		} `json:"event"`
	}
	op = body.Event.Op
	if err = b.BindJSON(&body); err != nil {
		return
	}
	unmarshal := func(j json.RawMessage, i interface{}) error {
		if i == nil {
			return nil
		}
		return json.Unmarshal(j, i)
	}
	if err = unmarshal(body.Event.Data.New, n); err != nil {
		return
	}
	err = unmarshal(body.Event.Data.Old, o)
	return
}
