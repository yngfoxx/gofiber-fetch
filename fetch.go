package fetch

import (
	"errors"
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type RequestHeader struct {
	Method        string      `json:"Method,omitempty"`
	Body          interface{} `json:"Body,omitempty"`
	Header        []string    `json:"Header"`
	Authorization string      `json:"Authorization,omitempty"`
	Agent         *fiber.Agent
}

type Respnse struct {
	Status int
	Data   interface{}
	Error  error
}

func Method(method string) *RequestHeader {
	return &RequestHeader{
		Method: method,
	}
}

func (header *RequestHeader) SetAuthorization(auth string) *RequestHeader {
	header.Authorization = auth
	return header
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (header *RequestHeader) FiberFetch(url string) *Respnse {

	// FastHttp agent
	agent := (map[bool]*fiber.Agent{
		true:  header.Agent,
		false: fiber.AcquireAgent(),
	})[header.Agent != nil]

	// Prepare request
	req := agent.Request()
	req.Header.SetMethod(header.Method)
	req.SetRequestURI(url)

	// body
	if header.Body != nil && !stringInSlice(header.Method, []string{"GET", "HEAD"}) {
		payload, _ := json.Marshal(header.Body)
		agent.Body(payload)
	}
	// authorization
	if header.Authorization != "" {
		agent.Set("Authorization", header.Authorization)
	}
	// headers
	if len(header.Header) > 0 {
		for _, v := range header.Header {
			h := strings.Split(v, "=")
			if len(h) == 2 {
				agent.Set(h[0], h[1])
			}
		}
	}

	// Execute
	if err := agent.Parse(); err != nil {
		return &Respnse{Error: err}
	}

	// Parse
	code, body, errs := agent.Bytes()
	if len(errs) > 0 {
		var errStr string
		for _, v := range errs {
			errStr += v.Error() + "; "
		}
		if code == 0 {
			code = 500
		}
		return &Respnse{Status: code, Error: errors.New(errStr)}
	}

	if code >= 400 && body != nil {
		return &Respnse{Status: code, Error: fmt.Errorf(string(body))}
	}

	return &Respnse{Status: code, Data: body}
}
