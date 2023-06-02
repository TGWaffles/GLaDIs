package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BadRequest StatusErrorCode = iota + 400
	Unauthorized
	Forbidden          = 403
	NotFound           = 404
	MethodNotAllowed   = 405
	TooManyRequests    = 429
	GatewayUnavailable = 502
)

type StatusErrorCode int

type StatusError struct {
	Code     StatusErrorCode
	Response *http.Response
}

func (s StatusError) tryFindJsonErrorCode() *DiscordError {
	// try decoding the response body to json map
	jsonMap := map[string]interface{}{}
	err := json.NewDecoder(s.Response.Body).Decode(&jsonMap)
	if err != nil {
		return nil
	}

	// try finding the error code
	code, ok := jsonMap["code"].(int)
	if !ok {
		return nil
	}

	return &DiscordError{
		Code:     ErrorCode(code),
		Response: s.Response,
	}
}

func (s StatusError) Error() string {
	discordError := s.tryFindJsonErrorCode()
	if discordError != nil {
		return discordError.Error()
	}
	switch s.Code {
	case BadRequest:
		return "Bad Request"
	case Unauthorized:
		return "Unauthorized"
	case Forbidden:
		return "Forbidden"
	case NotFound:
		return "Not Found"
	case MethodNotAllowed:
		return "Method Not Allowed"
	case TooManyRequests:
		return "Too Many Requests"
	case GatewayUnavailable:
		return "Gateway Unavailable"
	default:
		if s.Code >= 500 {
			return "Internal Server Error"
		}
		return fmt.Sprintf("Unknown Error (%d)", s.Code)
	}
}
