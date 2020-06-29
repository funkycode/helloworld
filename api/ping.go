package api

import (
	"encoding/json"

	"github.com/funkycode/helloworld/api/errors"
)

type Ping struct {
	Name string `json:"name"`
}

type Pong struct {
	Error    string `json:"error,omitempty"`
	Greeting string `json:"greeting,omitempty"`
}

func pingPong(jsonData []byte) (pong *Pong) {
	pong = &Pong{}
	ping := &Ping{}
	if len(jsonData) == 0 {
		pong.Error = (&errors.EmptyRequest{}).Error()
		return
	}
	switch err := json.Unmarshal(jsonData, ping); {
	case err != nil:
		pong.Error = (&errors.ParseError{}).Error()
	case ping.Name == "":
		pong.Error = (&errors.EmptyNameNotAllowed{}).Error()
	default:
		pong.Greeting = "Hello, " + ping.Name
	}
	return
}

func PingPong(jsonData []byte) ([]byte, error) {
	pong := pingPong(jsonData)
	return json.Marshal(pong)
}
