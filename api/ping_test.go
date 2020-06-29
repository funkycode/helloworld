package api

import (
	"testing"

	"github.com/funkycode/helloworld/api/errors"
)

const (
	emptyRequest           = ""
	emptyNameRequst        = "{\"name\":\"\"}"
	failedToParseRequest   = "{zzz}"
	validRequest           = "{\"name\":\"world\"}"
	responseToValidRequest = "Hello, world"
)

func TestPingPong(t *testing.T) {
	if pong := pingPong([]byte(emptyRequest)); pong.Error != (&errors.EmptyRequest{}).Error() {
		t.Errorf("Empty reequest error was not returned when it should")
	}
	if pong := pingPong([]byte(emptyNameRequst)); pong.Error != (&errors.EmptyNameNotAllowed{}).Error() {
		t.Errorf("Empty name request error was not returned when it should")
	}
	if pong := pingPong([]byte(failedToParseRequest)); pong.Error != (&errors.ParseError{}).Error() {
		t.Errorf("Failed to parse error was not returned when it should")
	}
	if pong := pingPong([]byte(validRequest)); pong.Greeting != responseToValidRequest {
		t.Errorf("Failed to generate proper pong")
	}
}
