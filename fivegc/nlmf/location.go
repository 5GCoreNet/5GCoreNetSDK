package nlmf

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	"github.com/5GCoreNet/5GCoreNetSDK/internal/header"
	openapicommon "github.com/5GCoreNet/openapi/openapi_CommonData"
	nlmfocation "github.com/5GCoreNet/openapi/openapi_Nlmf_Location"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	locationRouterGroup             = "/nlmf-loc/v1"
	cancelLocationEndpoint          = "/cancel-location"
	determineLocationEndpoint       = "/determine-location"
	locationContextTransferEndpoint = "/location-context-transfer"
)

type Location interface {
	fivegc.CommonInterface
	// CancelLocation cancels a location request.
	CancelLocation(context.Context, nlmfocation.CancelLocData) (openapicommon.ProblemDetails, fivegc.RedirectResponse, CancelLocationStatusCode)
	// DetermineLocation determines the location of a UE.
	DetermineLocation(context.Context, nlmfocation.InputData) (nlmfocation.LocationData, openapicommon.ProblemDetails, fivegc.RedirectResponse, DetermineLocationStatusCode)
	// LocationContextTransfer transfers the location context of a UE.
	LocationContextTransfer(context.Context, nlmfocation.LocContextData) (openapicommon.ProblemDetails, fivegc.RedirectResponse, LocationContextTransferStatusCode)
}

type CancelLocationStatusCode fivegc.StatusCode

const (
	// CancelLocationStatusNoContent is the status code for the response when the location request is successfully cancelled.
	CancelLocationStatusNoContent         CancelLocationStatusCode = CancelLocationStatusCode(fivegc.StatusNoContent)
	CancelLocationStatusTemporaryRedirect CancelLocationStatusCode = CancelLocationStatusCode(fivegc.StatusTemporaryRedirect)
	CancelLocationStatusPermanentRedirect CancelLocationStatusCode = CancelLocationStatusCode(fivegc.StatusPermanentRedirect)
)

type DetermineLocationStatusCode fivegc.StatusCode

const (
	// DetermineLocationStatusOK is the status code for a successful response.
	DetermineLocationStatusOK                DetermineLocationStatusCode = DetermineLocationStatusCode(fivegc.StatusOK)
	DetermineLocationStatusNoContent         DetermineLocationStatusCode = DetermineLocationStatusCode(fivegc.StatusNoContent)
	DetermineLocationStatusTemporaryRedirect DetermineLocationStatusCode = DetermineLocationStatusCode(fivegc.StatusTemporaryRedirect)
	DetermineLocationStatusPermanentRedirect DetermineLocationStatusCode = DetermineLocationStatusCode(fivegc.StatusPermanentRedirect)
)

type LocationContextTransferStatusCode fivegc.StatusCode

const (
	// LocationContextTransferStatusNoContent is the status code for the response when the location context transfer is successful.
	LocationContextTransferStatusNoContent         LocationContextTransferStatusCode = LocationContextTransferStatusCode(fivegc.StatusNoContent)
	LocationContextTransferStatusTemporaryRedirect LocationContextTransferStatusCode = LocationContextTransferStatusCode(fivegc.StatusTemporaryRedirect)
	LocationContextTransferStatusPermanentRedirect LocationContextTransferStatusCode = LocationContextTransferStatusCode(fivegc.StatusPermanentRedirect)
)

func attachLocationHandler(router *gin.RouterGroup, l Location) {
	group := router.Group(locationRouterGroup)
	{
		group.POST(cancelLocationEndpoint, func(c *gin.Context) {
			var req nlmfocation.CancelLocData
			if err := c.ShouldBindJSON(&req); err != nil {
				problemDetails := l.Error(c, err)
				c.JSON(int(*problemDetails.Status), problemDetails)
				return
			}
			problemDetails, redirectResponse, status := l.CancelLocation(c, req)
			switch status {
			case CancelLocationStatusNoContent:
				c.JSON(int(status), nil)
			case CancelLocationStatusTemporaryRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(int(status), redirectResponse)
			case CancelLocationStatusPermanentRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(int(status), redirectResponse)
			default:
				c.JSON(int(status), problemDetails)
			}
			return
		})
		group.POST(determineLocationEndpoint, func(c *gin.Context) {
			var req nlmfocation.InputData
			if err := c.ShouldBindJSON(&req); err != nil {
				problemDetails := l.Error(c, err)
				c.JSON(int(*problemDetails.Status), problemDetails)
				return
			}
			res, problemDetails, redirectResponse, status := l.DetermineLocation(c, req)
			switch status {
			case DetermineLocationStatusOK:
				c.JSON(int(status), res)
			case DetermineLocationStatusNoContent:
				c.JSON(int(status), nil)
			case DetermineLocationStatusTemporaryRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(int(status), redirectResponse)
			case DetermineLocationStatusPermanentRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(int(status), redirectResponse)
			default:
				c.JSON(int(status), problemDetails)
			}
			return
		})
		group.POST(locationContextTransferEndpoint, func(c *gin.Context) {
			var req nlmfocation.LocContextData
			if err := c.ShouldBindJSON(&req); err != nil {
				problemDetails := l.Error(c, err)
				c.JSON(int(*problemDetails.Status), problemDetails)
				return
			}
			problemDetails, redirectResponse, status := l.LocationContextTransfer(c, req)
			switch status {
			case LocationContextTransferStatusNoContent:
				c.JSON(int(status), nil)
			case LocationContextTransferStatusTemporaryRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(int(status), redirectResponse)
			case LocationContextTransferStatusPermanentRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(int(status), redirectResponse)
			default:
				c.JSON(int(status), problemDetails)
			}
			return
		})
	}
}

// LocationClient is a client for the NLMF Location service.
type LocationClient struct {
	client *nlmfocation.APIClient
}

// NewLocationClient creates a new client for the NLMF Location service.
func NewLocationClient(cfg fivegc.ClientConfiguration) *LocationClient {
	openapiCfg := &nlmfocation.Configuration{
		Host:             cfg.Host,
		Scheme:           cfg.Scheme,
		DefaultHeader:    cfg.DefaultHeader,
		UserAgent:        cfg.UserAgent,
		Debug:            cfg.Debug,
		Servers:          []nlmfocation.ServerConfiguration{},
		OperationServers: make(map[string]nlmfocation.ServerConfigurations),
		HTTPClient:       cfg.HTTPClient,
	}
	for _, server := range cfg.Servers {
		openapiServer := nlmfocation.ServerConfiguration{
			URL:         server.URL,
			Description: server.Description,
			Variables:   make(map[string]nlmfocation.ServerVariable),
		}
		for name, variable := range server.Variables {
			openapiServer.Variables[name] = nlmfocation.ServerVariable{
				Description:  variable.Description,
				DefaultValue: variable.DefaultValue,
				EnumValues:   variable.EnumValues,
			}
		}
		openapiCfg.Servers = append(openapiCfg.Servers, openapiServer)
	}
	for name, servers := range cfg.OperationServers {
		openapiServers := make(nlmfocation.ServerConfigurations, len(servers))
		for i, server := range servers {
			openapiServers[i] = nlmfocation.ServerConfiguration{
				URL:         server.URL,
				Description: server.Description,
				Variables:   make(map[string]nlmfocation.ServerVariable),
			}
			for name, variable := range server.Variables {
				openapiServers[i].Variables[name] = nlmfocation.ServerVariable{
					Description:  variable.Description,
					DefaultValue: variable.DefaultValue,
					EnumValues:   variable.EnumValues,
				}
			}
		}
		openapiCfg.OperationServers[name] = openapiServers
	}
	return &LocationClient{
		client: nlmfocation.NewAPIClient(openapiCfg),
	}
}

// LocationContextTransfer returns location context transfer request
func (l LocationClient) LocationContextTransfer(ctx context.Context) nlmfocation.ApiLocationContextTransferRequest {
	return l.client.LocationContextTransferApi.LocationContextTransfer(ctx)
}

// LocationContextTransferExecute executes the location context transfer request
func (l LocationClient) LocationContextTransferExecute(r nlmfocation.ApiLocationContextTransferRequest) (*http.Response, error) {
	return r.Execute()
}

// DetermineLocation returns determine location request
func (l LocationClient) DetermineLocation(ctx context.Context) nlmfocation.ApiDetermineLocationRequest {
	return l.client.DetermineLocationApi.DetermineLocation(ctx)
}

// DetermineLocationExecute executes the determine location request
func (l LocationClient) DetermineLocationExecute(r nlmfocation.ApiDetermineLocationRequest) (*nlmfocation.LocationData, *http.Response, error) {
	return r.Execute()
}

// CancelLocation returns cancel location request
func (l LocationClient) CancelLocation(ctx context.Context) nlmfocation.ApiCancelLocationRequest {
	return l.client.CancelLocationApi.CancelLocation(ctx)
}

// CancelLocationExecute executes the cancel location request
func (l LocationClient) CancelLocationExecute(r nlmfocation.ApiCancelLocationRequest) (*http.Response, error) {
	return r.Execute()
}
