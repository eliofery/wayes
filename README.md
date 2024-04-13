# wayes : 🌐 simple router using http.ServeMux

Package **wayes** provides a convenient wrapper around the standard router of the [net/http](https://pkg.go.dev/net/http#ServeMux) package. Describing routes is as simple as in favorite frameworks like Fiber, Gin, and others.

**wayes** only works with Go 1.22+ as it requires the new [net/http](https://pkg.go.dev/net/http#ServeMux) package.

## Installation

```bash
go get github.com/eliofery/wayes
```

## Usage

Example of creating and grouping routes.

```go
// Create a new router with use go-playground/validator.
// Validator should implement the interface wayes.Validater.
router := wayes.New(validator.New())

// Also, you can use the router without employing a validator.
// router := wayes.New()

// Create a route group for endpoints.
users := router.Group("/users")
{
    // Define a GET handler for the "/users/welcome" endpoint.
    // Also, there are other HTTP methods available, such as POST, PUT, PATCH, DELETE, OPTIONS, and HEAD.
    users.Get("/welcome", func(ctx wayes.Ctx) error {
        return ctx.JSON(wayes.Response{
            Success: true,
            Message: "Hi bro",
            Data: wayes.Map{
                "foo": "bar",
                "baz": 123,
            },
        })
    })

    // Define a POST handler for the "/users/welcome" endpoint.
    // Decodes the request body into the provided data and validates it.
    users.Post("/welcome", func(ctx wayes.Ctx) error {
        type User struct {
            Name string `json:"name" validate:"required"`
        }

        var user User
        if err := ctx.Validate(&user); err != nil {
            return ctx.Status(http.StatusBadRequest).SendError(err)
        }

        return ctx.JSON(wayes.Response{
            Success: true,
            Data:    user,
        })
    })
}

// Start the server.
if err := http.ListenAndServe(":8081", router.Mux()); err != nil {
    log.Fatal(err)
}
```

### Middlewares

Example of creating middlewares.

```go
// Cors is a middleware that sets CORS headers.
func Cors() wayes.Handler {
    return func(ctx wayes.Ctx) error {
        ctx.Set("Access-Control-Allow-Origin", "*")
        ctx.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        ctx.Set("Access-Control-Allow-Headers", "Origin, Accept, Authorization, Content-Type, X-CSRF-Token")
        ctx.Set("Access-Control-Expose-Headers", "Link, Content-Length, Access-Control-Allow-Origin")
        ctx.Set("Access-Control-Max-Age", "0")
        
        return ctx.Next()
    }
}

func main() {
    // Create a new router with use go-playground/validator.
    // Validator should implement the interface wayes.Validater.
    router := wayes.New(validator.New())
    
    fooMiddleware := func(ctx wayes.Ctx) error {
        return ctx.Next()
    }
    
    barMiddleware := func(ctx wayes.Ctx) error {
        if rand.IntN(2) == 1 {
            err := errors.New("access denied")
            return ctx.Status(http.StatusForbidden).SendError(err)
        }
        
        return ctx.Next()
    }
    
    // Define middleware for the all routes.
    router.Use(Cors(), fooMiddleware)
    
    // Create a route group for endpoints.
    users := router.Group("/users")
    {
        // Define middleware for the user group.
        users.Use(barMiddleware)
        
        // Define a GET handler for the "/users/welcome" endpoint.
        users.Get("/welcome", func(ctx wayes.Ctx) error {
            return ctx.JSON(wayes.Response{
                Success: true,
                Message: "Hi bro",
                Data: wayes.Map{
                    "foo": "bar",
                    "baz": 123,
                },
            })
        })
    }
    
    // Start the server.
    if err := http.ListenAndServe(":8081", router.Mux()); err != nil {
        log.Fatal(err)
    }
}
```
## Combine routers

Example of creating merged routes.

```go
// Create a new routers.
router := wayes.New(validator.New())
router2 := wayes.New()

// Create a route group for endpoints.
groupV1 := router.Group("/v1")
groupV2 := router2.Group("/v2")

// Define a GET handler for the "/v1/users" endpoint.
groupV1.Get("/users", func(ctx wayes.Ctx) error {
    return ctx.JSON(wayes.Response{
        Success: true,
        Message: "response users version 1",
    })
})

// Define a GET handler for the "/v2/users" endpoint.
groupV2.Get("/users", func(ctx wayes.Ctx) error {
    return ctx.JSON(wayes.Response{
        Success: true,
        Message: "response users version 2",
    })
})

// Combine routers into a single router.
// Now routes created in another router are accessible within a single router.
router.Combine(router2.Mux())

// Start the server.
if err := http.ListenAndServe(":8081", router.Mux()); err != nil {
    log.Fatal(err)
}
```

## Inspiration

I was inspired to write this package by the [http](https://pkg.go.dev/net/http), [fiber](https://github.com/gofiber/fiber) and [gin](https://github.com/gin-gonic/gin).
