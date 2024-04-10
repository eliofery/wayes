package wayes

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	// errEmptyBody represents an error indicating an empty request body.
	errEmptyBody = errors.New("empty body")

	// errInvalidBody represents an error indicating an invalid request body.
	errInvalidBody = errors.New("invalid body")
)

// Map represents a map of key-value pairs.
type Map map[string]any

// Response represents a structure for a response.
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// Ctx represents a structure for a context.
type Ctx interface {
	// Response returns the underlying [http.ResponseWriter] associated with the context.
	Response() http.ResponseWriter

	// Request returns the underlying [http.Request] associated with the context.
	Request() *http.Request

	// Locals sets or retrieves values associated with the context using the provided key.
	Locals(key any, value ...any) any

	// Status sets the status code for the response.
	Status(status int) Ctx

	// Get returns the header value for the given key.
	Get(key string, defaultValue ...string) string

	// Set sets the header value for the given key.
	Set(key, value string)

	// ContentType sets the Content-Type header for the response.
	ContentType(value string)

	// Decode decodes the request body into the provided data.
	Decode(data any) error

	// Validate decodes and validates the request body into the provided data.
	Validate(data any) error

	// Encode encodes the provided data into the response body.
	Encode(data any) error

	// Write sends a plain text response message to the user.
	Write(message string) error

	// JSON sends a json object response message to the user.
	JSON(data any) error

	// Next executes the next handler in the chain.
	Next() error

	// SendStatus sends an HTTP status code to the user.
	SendStatus(code int) error

	// SendError creates and returns an error with the specified message.
	SendError(message error) error
}

// ctx represents a structure that implements the [Ctx] interface.
type ctx struct {
	validator Validater
	response  http.ResponseWriter
	request   *http.Request
	status    int
}

// NewCtx creates a new instance of [Ctx].
func NewCtx(validator Validater, w http.ResponseWriter, r *http.Request) Ctx {
	return &ctx{
		validator: validator,
		response:  w,
		request:   r,
		status:    http.StatusOK,
	}
}

// Response returns the underlying [http.ResponseWriter] associated with the context.
func (c *ctx) Response() http.ResponseWriter {
	return c.response
}

// Request returns the underlying [http.Request] associated with the context.
func (c *ctx) Request() *http.Request {
	return c.request
}

// Locals sets or retrieves values associated with the context using the provided key.
// If only the key is provided, it retrieves the value associated with that key.
// If key and value are provided, it sets the value associated with the key and returns the value.
func (c *ctx) Locals(key any, value ...any) any {
	if len(value) == 0 {
		return c.Request().Context().Value(key)
	}

	loadCtx := context.WithValue(c.Request().Context(), key, value[0])
	c.request = c.Request().WithContext(loadCtx)

	return value[0]
}

// Status sets the status code for the response.
func (c *ctx) Status(code int) Ctx {
	c.status = code
	return c
}

// Get returns the header value for the given key.
func (c *ctx) Get(key string, defaultValue ...string) string {
	header := c.response.Header().Get(key)

	if len(header) == 0 {
		if len(defaultValue) == 0 {
			return ""
		}

		return defaultValue[0]
	}

	return header
}

// Set sets the header value for the given key.
func (c *ctx) Set(key, value string) {
	c.response.Header().Set(key, value)
}

// ContentType sets the Content-Type header for the response.
func (c *ctx) ContentType(value string) {
	c.Set("Content-Type", value)
}

// Decode decodes the request body into the provided data.
func (c *ctx) Decode(data any) error {
	if err := json.NewDecoder(c.request.Body).Decode(data); err != nil {
		c.Status(http.StatusBadRequest)

		if (c.request.Method == http.MethodPost ||
			c.request.Method == http.MethodPut ||
			c.request.Method == http.MethodPatch ||
			c.request.Method == http.MethodDelete) &&
			errors.Is(err, io.EOF) {
			return errEmptyBody
		}

		return errInvalidBody
	}

	return nil
}

// Validate decodes and validates the request body into the provided data.
func (c *ctx) Validate(data any) error {
	if err := c.Decode(data); err != nil {
		return err
	}

	if c.validator != nil {
		if err := c.validator.Struct(data); err != nil {
			c.Status(http.StatusBadRequest)
			return err
		}
	}

	return nil
}

// Encode encodes the provided data into the response body.
func (c *ctx) Encode(data any) error {
	encoder := json.NewEncoder(c.response)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

// Write sends a plain text response message to the user.
func (c *ctx) Write(message string) error {
	c.ContentType("text/plain; charset=utf-8")
	c.response.WriteHeader(c.status)

	if c.status == http.StatusNoContent {
		return nil
	}

	if _, err := c.response.Write([]byte(message)); err != nil {
		return err
	}

	return nil
}

// JSON sends a json object response message to the user.
func (c *ctx) JSON(data any) error {
	c.ContentType("application/json")
	c.response.WriteHeader(c.status)

	return c.Encode(data)
}

// Next calls the next handler in the chain.
func (c *ctx) Next() error {
	return nil
}

// SendStatus sends a plain text response message to the user.
func (c *ctx) SendStatus(code int) error {
	c.Status(code)

	return c.Write(http.StatusText(code))
}

// SendError creates and returns an error with the specified message.
func (c *ctx) SendError(message error) error {
	c.response.WriteHeader(c.status)

	return message
}
