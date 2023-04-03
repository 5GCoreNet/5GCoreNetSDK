package fivegc

import (
	"context"
	openapicommon "github.com/5GCoreNet/openapi/openapi_CommonData"
)

type CommonInterface interface {
	// Error returns a problem details, it is used to handle errors when unmarshalling the request.
	Error(ctx context.Context, err error) openapicommon.ProblemDetails
}
