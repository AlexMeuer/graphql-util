package hasura

import (
	"encoding/json"
	"net/http"
)

type JSONer interface {
	JSON(code int, body interface{})
}

type BinderJSON interface {
	BindJSON(interface{}) error
}

// Calls JSON on the given JSONer, using the Error's code as the code parameter
// and the Error itself as the body parameter.
//
// If the Error has no code, then http.StatusInternalServerError is used instead.
// (This function was written to support cleaner code when using gin-gonic.)
func JSON(j JSONer, err *Error) {
	if err.Code == noErrCode {
		j.JSON(http.StatusInternalServerError, err)
	} else {
		j.JSON(err.Code, err)
	}
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

// UnmarshalHasuraChangeEvent unmarshals a Hasura Event payload into either/both of the given
// new/old structs and returns the operation type that caused the event.
// The BinderJSON "b" argument is used for initial decoding but it not used in conjuction with the interface "i".
// The reason for this is because this function was originally written for the gin web framework (specifically
// the `context.BindJSON(..)` function.)
func UnmarshalHasuraChangeEvent(b BinderJSON, n, o interface{}) (op string, err error) {
	var body struct {
		Event struct {
			Op   string `json:"op"`
			Data struct {
				Old json.RawMessage `json:"old"`
				New json.RawMessage `json:"new"`
			} `json:"data"`
		} `json:"event"`
	}
	if err = b.BindJSON(&body); err != nil {
		return
	}
	op = body.Event.Op
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

// UnmarshalHasuraScheduledEvent unmarshals a Hasura Event payload into the given interface.
// This is intended for use with cron and one-off scheduled events, NOT for events triggered by
// changes to the database (use UnmarshalHasuraChangeEvent for that).
func UnmarshalHasuraScheduledEvent(b BinderJSON, data interface{}) error {
	var body struct {
		Payload json.RawMessage `json:"payload"`
	}
	if err := b.BindJSON(&body); err != nil {
		return err
	}
	return json.Unmarshal(body.Payload, data)
}
