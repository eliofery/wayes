package wayes

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ValidatorMock implements the Validater interface.
type ValidatorMock struct {
	isError bool
}

// Struct validates the structure of the provided data against predefined rules.
func (v *ValidatorMock) Struct(_ any) error {
	if v.isError {
		return errors.New("test error")
	}

	return nil
}

// TestWayesMethods_success test suite verifies the functionality of various HTTP methods,
// including OPTIONS, HEAD, GET, POST, PUT, PATCH, and DELETE.
func TestWayesMethods_success(t *testing.T) {
	cases := []struct {
		name         string
		method       string
		path         string
		handler      Handler
		exceptedBody string
	}{
		{
			name:   "Options request",
			method: "OPTIONS",
			path:   "/test-options",
			handler: func(ctx Ctx) error {
				return ctx.Write("Test options response")
			},
			exceptedBody: "Test options response",
		},
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

	rt := New()
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, test.path, nil)
			require.NoError(t, err)

			switch test.method {
			case "OPTIONS":
				rt.Options(test.path, test.handler)
			case "HEAD":
				rt.Head(test.path, test.handler)
			case "GET":
				rt.Get(test.path, test.handler)
			case "POST":
				rt.Post(test.path, test.handler)
			case "PATCH":
				rt.Patch(test.path, test.handler)
			case "PUT":
				rt.Put(test.path, test.handler)
			case "DELETE":
				rt.Delete(test.path, test.handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)

			assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"))
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, test.exceptedBody, rr.Body.String())
		})
	}
}

// TestWayesMethods_errors test suite verifies the functionality of various HTTP methods,
// including OPTIONS, HEAD, GET, POST, PUT, PATCH, and DELETE.
func TestWayesMethods_errors(t *testing.T) {
	cases := []struct {
		name          string
		method        string
		path          string
		handler       Handler
		exceptedError string
	}{
		{
			name:   "Head request",
			method: "HEAD",
			path:   "/test-head",
			handler: func(ctx Ctx) error {
				return errors.New("test head error")
			},
			exceptedError: "test head error",
		},
		{
			name:   "Get request",
			method: "GET",
			path:   "/test-get",
			handler: func(ctx Ctx) error {
				return errors.New("test get error")
			},
			exceptedError: "test get error",
		},
		{
			name:   "Post request",
			method: "POST",
			path:   "/test-post",
			handler: func(ctx Ctx) error {
				return errors.New("test post error")
			},
			exceptedError: "test post error",
		},
		{
			name:   "Patch request",
			method: "PATCH",
			path:   "/test-patch",
			handler: func(ctx Ctx) error {
				return errors.New("test patch error")
			},
			exceptedError: "test patch error",
		},
		{
			name:   "Put request",
			method: "PUT",
			path:   "/test-put",
			handler: func(ctx Ctx) error {
				return errors.New("test put error")
			},
			exceptedError: "test put error",
		},
		{
			name:   "Delete request",
			method: "DELETE",
			path:   "/test-delete",
			handler: func(ctx Ctx) error {
				return errors.New("test delete error")
			},
			exceptedError: "test delete error",
		},
	}

	rt := New()
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, test.path, nil)
			require.NoError(t, err)

			switch test.method {
			case "HEAD":
				rt.Head(test.path, test.handler)
			case "GET":
				rt.Get(test.path, test.handler)
			case "POST":
				rt.Post(test.path, test.handler)
			case "PATCH":
				rt.Patch(test.path, test.handler)
			case "PUT":
				rt.Put(test.path, test.handler)
			case "DELETE":
				rt.Delete(test.path, test.handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)

			assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"))
			assert.Equal(t, http.StatusInternalServerError, rr.Code)
			assert.Contains(t, rr.Body.String(), test.exceptedError)
		})
	}
}

// TestWayesGroup tests the functionality of router groups.
func TestWayesGroup(t *testing.T) {
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

	rt := New()
	groupUrl := "/group"
	group := rt.Group(groupUrl)

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, fmt.Sprintf("%s%s", groupUrl, test.path), nil)
			require.NoError(t, err)

			switch test.method {
			case "HEAD":
				group.Head(test.path, test.handler)
			case "GET":
				group.Get(test.path, test.handler)
			case "POST":
				group.Post(test.path, test.handler)
			case "PATCH":
				group.Patch(test.path, test.handler)
			case "PUT":
				group.Put(test.path, test.handler)
			case "DELETE":
				group.Delete(test.path, test.handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)

			assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"))
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, test.exceptedBody, rr.Body.String())
		})
	}
}

// TestWayesUse_success tests the successful addition of middleware to the router.
func TestWayesUse_success(t *testing.T) {
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

	rt := New()
	rt.Use(func(ctx Ctx) error { return ctx.Next() })
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, test.path, nil)
			require.NoError(t, err)

			switch test.method {
			case "HEAD":
				rt.Head(test.path, test.handler)
			case "GET":
				rt.Get(test.path, test.handler)
			case "POST":
				rt.Post(test.path, test.handler)
			case "PATCH":
				rt.Patch(test.path, test.handler)
			case "PUT":
				rt.Put(test.path, test.handler)
			case "DELETE":
				rt.Delete(test.path, test.handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)

			assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"))
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, test.exceptedBody, rr.Body.String())
		})
	}
}

// TestWayesUse_error tests the error addition of middleware to the router.
func TestWayesUse_error(t *testing.T) {
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
	expectedErr := "middleware error"

	rt := New()
	rt.Use(func(ctx Ctx) error {
		err := errors.New(expectedErr)
		return ctx.Status(http.StatusForbidden).SendError(err)
	})

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, test.path, nil)
			require.NoError(t, err)

			switch test.method {
			case "HEAD":
				rt.Head(test.path, test.handler)
			case "GET":
				rt.Get(test.path, test.handler)
			case "POST":
				rt.Post(test.path, test.handler)
			case "PATCH":
				rt.Patch(test.path, test.handler)
			case "PUT":
				rt.Put(test.path, test.handler)
			case "DELETE":
				rt.Delete(test.path, test.handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)

			assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"))
			assert.Equal(t, http.StatusForbidden, rr.Code)
			assert.Contains(t, rr.Body.String(), expectedErr)
		})
	}
}

// TestWayesCombine tests the error addition of middleware to the router.
func TestWayesCombine(t *testing.T) {
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
	cases2 := []struct {
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

	rt := New()
	v1 := "/v1"
	groupV1 := rt.Group(v1)

	rt2 := New()
	v2 := "/v2"
	groupV2 := rt2.Group(v2)

	for i, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			switch test.method {
			case "HEAD":
				groupV1.Head(test.path, test.handler)
				groupV2.Head(cases2[i].path, cases2[i].handler)
			case "GET":
				groupV1.Get(test.path, test.handler)
				groupV2.Get(cases2[i].path, cases2[i].handler)
			case "POST":
				groupV1.Post(test.path, test.handler)
				groupV2.Post(cases2[i].path, cases2[i].handler)
			case "PATCH":
				groupV1.Patch(test.path, test.handler)
				groupV2.Patch(cases2[i].path, cases2[i].handler)
			case "PUT":
				groupV1.Put(test.path, test.handler)
				groupV2.Put(cases2[i].path, cases2[i].handler)
			case "DELETE":
				groupV1.Delete(test.path, test.handler)
				groupV2.Delete(cases2[i].path, cases2[i].handler)
			default:
				t.Fatalf("unsupported HTTP method: %s", test.method)
			}
		})
	}

	rt.Combine(rt2.Mux())
	for i, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(cases2[i].method, fmt.Sprintf("%s%s", v2, cases2[i].path), nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			rt.Mux().ServeHTTP(rr, req)

			assert.Equal(t, "text/plain; charset=utf-8", rr.Header().Get("Content-Type"))
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, test.exceptedBody, rr.Body.String())
		})
	}
}
