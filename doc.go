/*
Package wayes provides a convenient wrapper around the standard router of the [pkg/net/http.ServeMux] package.
Describing routes is as simple as in favorite frameworks like Fiber, Gin, and others.

Includes [Wayes] that serves as a wrapper for the standard [pkg/net/http.ServeMux] package.

Includes [Ctx] that allows passing context. Context to middleware and handlers for request-specific data.

Demonstrates how to define routes and handlers using the wayes package.
It initializes a new wayes and creates a route group for user-related endpoints.
Within the user group, it defines various HTTP methods (GET, POST, PUT, PATCH, DELETE, OPTIONS and HEAD)
that can be used to interact with user data.

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
*/
package wayes
