package wayes

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCtxResponseRequest tests the response and request handling functionality.
func TestCtxResponseRequest(t *testing.T) {
	cases := []struct {
		name         string
		method       string
		path         string
		handler      Handler
		exceptedBody string
	}{
		{
			name:   "Head request",
			method: "HEAD",
			path:   "/test-head",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test head response")
			},
			exceptedBody: "Test head response",
		},
		{
			name:   "Get request",
			method: "GET",
			path:   "/test-get",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test get response")
			},
			exceptedBody: "Test get response",
		},
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test post response")
			},
			exceptedBody: "Test post response",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test patch response")
			},
			exceptedBody: "Test patch response",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test put response")
			},
			exceptedBody: "Test put response",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test delete response")
			},
			exceptedBody: "Test delete response",
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, test.path, nil)
			rr := httptest.NewRecorder()
			ctx := NewCtx(nil, rr, req)

			assert.Equal(t, rr, ctx.Response())
			assert.Equal(t, req, ctx.Request())
		})
	}
}

// TestCtxLocals tests the setting of local values in the context.
func TestCtxLocals(t *testing.T) {
	cases := []struct {
		name         string
		method       string
		path         string
		handler      Handler
		exceptedBody string
	}{
		{
			name:   "Head request",
			method: "HEAD",
			path:   "/test-head",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test head response")
			},
			exceptedBody: "Test head response",
		},
		{
			name:   "Get request",
			method: "GET",
			path:   "/test-get",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test get response")
			},
			exceptedBody: "Test get response",
		},
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test post response")
			},
			exceptedBody: "Test post response",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test patch response")
			},
			exceptedBody: "Test patch response",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test put response")
			},
			exceptedBody: "Test put response",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test delete response")
			},
			exceptedBody: "Test delete response",
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, test.path, nil)
			rr := httptest.NewRecorder()
			ctx := NewCtx(nil, rr, req)

			// Set value to context.
			ctx.Locals("test", 123)

			// Get value from context.
			value := ctx.Locals("test")

			assert.Equal(t, 123, value)
		})
	}
}

// TestCtxHeader tests the handling of HTTP headers.
func TestCtxHeader(t *testing.T) {
	cases := []struct {
		name         string
		method       string
		path         string
		handler      Handler
		exceptedBody string
	}{
		{
			name:   "Head request",
			method: "HEAD",
			path:   "/test-head",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test head response")
			},
			exceptedBody: "Test head response",
		},
		{
			name:   "Get request",
			method: "GET",
			path:   "/test-get",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test get response")
			},
			exceptedBody: "Test get response",
		},
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test post response")
			},
			exceptedBody: "Test post response",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test patch response")
			},
			exceptedBody: "Test patch response",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test put response")
			},
			exceptedBody: "Test put response",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test delete response")
			},
			exceptedBody: "Test delete response",
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, test.path, nil)
			rr := httptest.NewRecorder()
			ctx := NewCtx(nil, rr, req)

			expected := "123"

			// Set value to header.
			ctx.Set("test", expected)

			// Get value from context.
			value := ctx.Get("test")
			assert.Equal(t, expected, value)

			defaultValue := ctx.Get("undefined", expected)
			assert.Equal(t, defaultValue, value)

			undefined := ctx.Get("undefined")
			assert.Equal(t, "", undefined)
		})
	}
}

// TestCtxDecode_success tests the successful decoding of data.
func TestCtxDecode_success(t *testing.T) {
	cases := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "Get request",
			method: "GET",
			path:   "/test-get",
		},
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
		},
	}

	rt := New()
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			data := Response{
				Success: true,
				Message: "test message",
			}

			jsonData, err := json.Marshal(data)
			require.NoError(t, err)

			handler := func(ctx Ctx) error {
				var response Response
				err := ctx.Decode(&response)
				require.NoError(t, err)

				assert.Equal(t, data, response)

				return nil
			}

			switch test.method {
			case "GET":
				rt.Get(test.path, handler)
			case "POST":
				rt.Post(test.path, handler)
			case "PATCH":
				rt.Patch(test.path, handler)
			case "PUT":
				rt.Put(test.path, handler)
			case "DELETE":
				rt.Delete(test.path, handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			req, err := http.NewRequest(test.method, test.path, bytes.NewBuffer(jsonData))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)
		})
	}
}

// TestCtxDecode_errorInvalidBody tests the error decoding of data.
func TestCtxDecode_errorInvalidBody(t *testing.T) {
	cases := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "Get request",
			method: "GET",
			path:   "/test-get",
		},
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
		},
	}

	rt := New()
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			data := []byte("")

			jsonData, err := json.Marshal(data)
			require.NoError(t, err)

			handler := func(ctx Ctx) error {
				var response Response
				err = ctx.Decode(&response)
				assert.Error(t, err)
				assert.Equal(t, errInvalidBody, err)

				return nil
			}

			switch test.method {
			case "GET":
				rt.Get(test.path, handler)
			case "POST":
				rt.Post(test.path, handler)
			case "PATCH":
				rt.Patch(test.path, handler)
			case "PUT":
				rt.Put(test.path, handler)
			case "DELETE":
				rt.Delete(test.path, handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			req, err := http.NewRequest(test.method, test.path, bytes.NewBuffer(jsonData))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)
		})
	}
}

// TestCtxDecode_errorEmptyBody tests the error decoding of data.
func TestCtxDecode_errorEmptyBody(t *testing.T) {
	cases := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
		},
	}

	rt := New(&ValidatorMock{})
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			handler := func(ctx Ctx) error {
				var response Response
				err := ctx.Validate(&response)
				assert.Error(t, err)
				assert.Equal(t, errEmptyBody, err)

				return nil
			}

			switch test.method {
			case "POST":
				rt.Post(test.path, handler)
			case "PATCH":
				rt.Patch(test.path, handler)
			case "PUT":
				rt.Put(test.path, handler)
			case "DELETE":
				rt.Delete(test.path, handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			req, err := http.NewRequest(test.method, test.path, bytes.NewBuffer([]byte("")))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)
		})
	}
}

// TestCtxValidate_success tests the successful validation of data.
func TestCtxValidate_success(t *testing.T) {
	cases := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
		},
	}

	rt := New()
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			type Request struct {
				Name string `json:"name" validate:"required"`
			}

			reqData := Request{
				Name: "test",
			}

			jsonData, err := json.Marshal(reqData)
			require.NoError(t, err)

			handler := func(ctx Ctx) error {
				var data Request
				err = ctx.Validate(&data)
				require.NoError(t, err)

				assert.Equal(t, reqData, data)

				return nil
			}

			switch test.method {
			case "POST":
				rt.Post(test.path, handler)
			case "PATCH":
				rt.Patch(test.path, handler)
			case "PUT":
				rt.Put(test.path, handler)
			case "DELETE":
				rt.Delete(test.path, handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			req, err := http.NewRequest(test.method, test.path, bytes.NewBuffer(jsonData))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)
		})
	}
}

// TestCtxValidate_error tests the error validation of data.
func TestCtxValidate_error(t *testing.T) {
	cases := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
		},
	}

	rt := New(&ValidatorMock{true})
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			type Request struct {
				Name string `json:"name" validate:"required"`
			}

			var reqData Request

			jsonData, err := json.Marshal(reqData)
			require.NoError(t, err)

			handler := func(ctx Ctx) error {
				var data Request
				err = ctx.Validate(&data)
				require.Error(t, err)

				assert.Equal(t, reqData, data)

				return err
			}

			switch test.method {
			case "POST":
				rt.Post(test.path, handler)
			case "PATCH":
				rt.Patch(test.path, handler)
			case "PUT":
				rt.Put(test.path, handler)
			case "DELETE":
				rt.Delete(test.path, handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			req, err := http.NewRequest(test.method, test.path, bytes.NewBuffer(jsonData))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)
		})
	}
}

// TestCtxEncode tests the encoding of data.
func TestCtxEncode(t *testing.T) {
	cases := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "Get request",
			method: "GET",
			path:   "/test-get",
		},
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
		},
	}
	data := Response{
		Success: true,
		Message: "test message",
		Data: Map{
			"foo": "bar",
			"baz": float64(123),
		},
	}

	rt := New()
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			handler := func(ctx Ctx) error {
				err := ctx.JSON(data)

				assert.NoError(t, err)

				return err
			}

			switch test.method {
			case "GET":
				rt.Get(test.path, handler)
			case "POST":
				rt.Post(test.path, handler)
			case "PATCH":
				rt.Patch(test.path, handler)
			case "PUT":
				rt.Put(test.path, handler)
			case "DELETE":
				rt.Delete(test.path, handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			req, err := http.NewRequest(test.method, test.path, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)

			var resData Response
			err = json.Unmarshal(rr.Body.Bytes(), &resData)
			require.NoError(t, err)

			resData.Data = Map(resData.Data.(map[string]interface{}))

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, data, resData)
		})
	}
}

// TestCtxWrite tests the writing of data.
func TestCtxWrite(t *testing.T) {
	cases := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "Get request",
			method: "GET",
			path:   "/test-get",
		},
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
		},
	}
	data := "Test write response"

	rt := New()
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			handler := func(ctx Ctx) error {
				err := ctx.Write(data)

				assert.NoError(t, err)

				return err
			}

			switch test.method {
			case "GET":
				rt.Get(test.path, handler)
			case "POST":
				rt.Post(test.path, handler)
			case "PATCH":
				rt.Patch(test.path, handler)
			case "PUT":
				rt.Put(test.path, handler)
			case "DELETE":
				rt.Delete(test.path, handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			req, err := http.NewRequest(test.method, test.path, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, data, rr.Body.String())
		})
	}
}

// TestCtxWrite_noContent tests the writing of data.
func TestCtxWrite_noContent(t *testing.T) {
	rt := New()
	rt.Options("/*", func(ctx Ctx) error {
		return ctx.SendStatus(http.StatusNoContent)
	})

	getUrl := "/test"
	rt.Get("/test", func(ctx Ctx) error {
		return ctx.Status(http.StatusOK).Write("Test write response")
	})

	req, err := http.NewRequest("OPTIONS", getUrl, nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	rt.Mux().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}
