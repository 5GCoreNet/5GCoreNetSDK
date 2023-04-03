# Design principals

As 5GCoreNetSDK aims to be a simple and easy to use SDK, we have defined some design principals that we want to follow.

If you want to contribute to 5GCoreNetSDK, please make sure that your code follows these principals.

## Simple and easy to use

This is the main goal of 5GCoreNetSDK. We want to provide a simple and easy to use SDK to build 5G Core Network NFs following the R18 3GPP specifications.

Indeed, some 3GPP specifications are very complex and hard to understand. We want to make it easy to use the 3GPP specifications in code.

## Interface based

5GCoreNetSDK is based on interfaces. This means that you need to implement the interfaces. To do so, you just need to 
create a struct that implements the interface and that's it.

This is very useful because it allows you to provide your own implementation of the Network Functions without bothering about the 3GPP specifications.

Plus, it makes it easy to test your code, because you can mock the interfaces you need.

So every Network Function connectivity layer must be based on an interface.

Let's take an example, consider we want to implement the `Foo` endpoint. 

```go
package foo

import (
	"context"
	"github.com/5GCore/5GCoreNetSDK/fivegc"
)

type Foo interface {
	fivegc.CommonInterface
    Foo(context.Context, string) string, ProblemDetails, fivegc.RedirectResponse, FooStatusCode
}
```

This interface is composed of the `CommonInterface` interface, which is the interface that every endpoint must implement.

`CommonInterface` is here to provide a common interface wrapping error handling (maybe we will add more stuff in the future).

Then, we have the `Foo` method, which is the method that will be called when a request is sent to the `Foo` endpoint.

And that's it, you have implemented the `Foo` endpoint. Not so hard, right?

## Client flow
Interfaces are used for the server part, whereas the client part is following a more classical flow.

Client must follow the following flow to send a request to a service:

1 - Create a new client
2 - Create a new query
3 - Fill the query with the parameters
3 - Execute the query

```go
// Create a new client
client := client.NewClient(cfg)
// Client can be used to create a new query
query := client.NetworkFunctionEndpoint(ctx)
// Query can be used to fill the parameters
query = query.Paramaters(model)
// Query can be used to send a request to the server
response, err := client.NetworkFunctionEndpointExecute(query)
```

This flow is very simple and easy to use. It is also very easy to test, because you can mock the client and the query,
thus we decided to follow this flow for the client part.

## Error handling

As in 3GPP specifications, errors are handled using the `ProblemDetails` struct.

Unmarshalling errors must be handled using the `ProblemDetails` struct as well. This struct is used to provide additional information in an error response.
That's why we decided to add na `Error()` method to all the interfaces by using the `CommonInterface` interface.

```go
type CommonInterface interface {
    // Error returns a problem details, it is used to handle errors when unmarshalling the request.
    Error(ctx context.Context, err error) openapicommon.ProblemDetails
}
```
Error is called when an error occurs during the unmarshalling of the request. It is used to provide additional information about the error.

Thus, SDK users can handle by themselves the errors without effort.

## Mocking

Every interface implemented must be mocked. This is a very elegant way to provide user a way to test their code.

In order to do so, we use the `mockgen` library. This library allows us to generate mocks from interfaces.

Let's take an example, consider we want to mock the `Foo` interface.

```go
package mock // located under foo/mock

//go:generate mockgen -source=../foo.go -destination=foo.go -package=mock
```

This command will generate a `foo.go` file under the `foo/mock` package. This file will contain a mocked `Foo` interface.

## Testing

Every endpoint must be tested against the 3GPP specifications. Thus, we can ensure that our code is compliant with the 3GPP specifications and that we are not breaking anything.

## Documentation

Every endpoint must be documented. This is very important because it allows users to understand how to use the endpoint.

You have to use GoDoc format to document your code. This is very easy to use, and it is very useful. You can find more information about GoDoc [here](https://blog.golang.org/godoc).

Plus, you can add an example of how to use the endpoint. This is very useful because it allows users to better understand how to use the endpoint.