package nlmf

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	"github.com/5GCoreNet/5GCoreNetSDK/internal/header"
	openapinlmflocationclient "github.com/5GCoreNet/client-openapi/Nlmf_Location"
	openapinlmflocationserver "github.com/5GCoreNet/server-openapi/Nlmf_Location"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Location interface {
	// Error returns a problem details, it is used to handle errors when unmarshalling the request.
	Error(ctx context.Context, err error) openapinlmflocationserver.ProblemDetails
	// CancelLocation cancels a location request.
	CancelLocation(context.Context, openapinlmflocationserver.CancelLocData) (openapinlmflocationserver.ProblemDetails, fivegc.RedirectResponse, fivegc.StatusCode)
	// DetermineLocation determines the location of a UE.
	DetermineLocation(context.Context, openapinlmflocationserver.InputData) (openapinlmflocationserver.LocationData, openapinlmflocationserver.ProblemDetails, fivegc.RedirectResponse, fivegc.StatusCode)
	// LocationContextTransfer transfers the location context of a UE.
	LocationContextTransfer(context.Context, openapinlmflocationserver.LocContextData) (openapinlmflocationserver.ProblemDetails, fivegc.RedirectResponse, fivegc.StatusCode)
}

func attachLocationHandler(router *gin.RouterGroup, l Location) {
	group := router.Group("/nlmf-loc/v1")
	{
		group.POST("/cancel-location", func(c *gin.Context) {
			var req openapinlmflocationserver.CancelLocData
			if err := c.ShouldBindJSON(&req); err != nil {
				problemDetails := l.Error(c, err)
				c.JSON(int(problemDetails.Status), problemDetails)
				return
			}
			problemDetails, redirectResponse, status := l.CancelLocation(c, req)
			switch status {
			case fivegc.StatusNoContent:
				c.JSON(status.ToInt(), nil)
			case fivegc.StatusTemporaryRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(status.ToInt(), redirectResponse)
			case fivegc.StatusPermanentRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(status.ToInt(), redirectResponse)
			default:
				c.JSON(status.ToInt(), problemDetails)
			}
			return
		})
		group.POST("/determine-location", func(c *gin.Context) {
			var req openapinlmflocationserver.InputData
			if err := c.ShouldBindJSON(&req); err != nil {
				problemDetails := l.Error(c, err)
				c.JSON(int(problemDetails.Status), problemDetails)
				return
			}
			res, problemDetails, redirectResponse, status := l.DetermineLocation(c, req)
			switch status {
			case fivegc.StatusOK:
				c.JSON(status.ToInt(), res)
			case fivegc.StatusNoContent:
				c.JSON(status.ToInt(), nil)
			case fivegc.StatusTemporaryRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(status.ToInt(), redirectResponse)
			case fivegc.StatusPermanentRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(status.ToInt(), redirectResponse)
			default:
				c.JSON(status.ToInt(), problemDetails)
			}
			return
		})
		group.POST("/location-context-transfer", func(c *gin.Context) {
			var req openapinlmflocationserver.LocContextData
			if err := c.ShouldBindJSON(&req); err != nil {
				problemDetails := l.Error(c, err)
				c.JSON(int(problemDetails.Status), problemDetails)
				return
			}
			problemDetails, redirectResponse, status := l.LocationContextTransfer(c, req)
			switch status {
			case fivegc.StatusNoContent:
				c.JSON(status.ToInt(), nil)
			case fivegc.StatusTemporaryRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(status.ToInt(), redirectResponse)
			case fivegc.StatusPermanentRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(status.ToInt(), redirectResponse)
			default:
				c.JSON(status.ToInt(), problemDetails)
			}
			return
		})
	}
}

// LocationClient is a client for the Nlmf_location service.
type LocationClient struct {
	client *openapinlmflocationclient.APIClient
}

// NewLocationClient creates a new client for the Nlmf_location service.
func NewLocationClient(cfg fivegc.ClientConfiguration) *LocationClient {
	openapiCfg := &openapinlmflocationclient.Configuration{
		Host:             cfg.Host,
		Scheme:           cfg.Scheme,
		DefaultHeader:    cfg.DefaultHeader,
		UserAgent:        cfg.UserAgent,
		Debug:            cfg.Debug,
		Servers:          []openapinlmflocationclient.ServerConfiguration{},
		OperationServers: make(map[string]openapinlmflocationclient.ServerConfigurations),
		HTTPClient:       cfg.HTTPClient,
	}
	for _, server := range cfg.Servers {
		openapiServer := openapinlmflocationclient.ServerConfiguration{
			URL:         server.URL,
			Description: server.Description,
			Variables:   make(map[string]openapinlmflocationclient.ServerVariable),
		}
		for name, variable := range server.Variables {
			openapiServer.Variables[name] = openapinlmflocationclient.ServerVariable{
				Description:  variable.Description,
				DefaultValue: variable.DefaultValue,
				EnumValues:   variable.EnumValues,
			}
		}
		openapiCfg.Servers = append(openapiCfg.Servers, openapiServer)
	}
	for name, servers := range cfg.OperationServers {
		openapiServers := make(openapinlmflocationclient.ServerConfigurations, len(servers))
		for i, server := range servers {
			openapiServers[i] = openapinlmflocationclient.ServerConfiguration{
				URL:         server.URL,
				Description: server.Description,
				Variables:   make(map[string]openapinlmflocationclient.ServerVariable),
			}
			for name, variable := range server.Variables {
				openapiServers[i].Variables[name] = openapinlmflocationclient.ServerVariable{
					Description:  variable.Description,
					DefaultValue: variable.DefaultValue,
					EnumValues:   variable.EnumValues,
				}
			}
		}
		openapiCfg.OperationServers[name] = openapiServers
	}
	return &LocationClient{
		client: openapinlmflocationclient.NewAPIClient(openapiCfg),
	}
}

// LocationContextTransfer returns location context transfer request
func (l LocationClient) LocationContextTransfer(ctx context.Context) openapinlmflocationclient.ApiLocationContextTransferRequest {
	return l.client.LocationContextTransferApi.LocationContextTransfer(ctx)
}

// LocationContextTransferExecute executes the location context transfer request
func (l LocationClient) LocationContextTransferExecute(r openapinlmflocationclient.ApiLocationContextTransferRequest) (*http.Response, error) {
	return r.Execute()
}

// DetermineLocation returns determine location request
func (l LocationClient) DetermineLocation(ctx context.Context) openapinlmflocationclient.ApiDetermineLocationRequest {
	return l.client.DetermineLocationApi.DetermineLocation(ctx)
}

// DetermineLocationExecute executes the determine location request
func (l LocationClient) DetermineLocationExecute(r openapinlmflocationclient.ApiDetermineLocationRequest) (*openapinlmflocationclient.LocationData, *http.Response, error) {
	return r.Execute()
}

// CancelLocation returns cancel location request
func (l LocationClient) CancelLocation(ctx context.Context) openapinlmflocationclient.ApiCancelLocationRequest {
	return l.client.CancelLocationApi.CancelLocation(ctx)
}

// CancelLocationExecute executes the cancel location request
func (l LocationClient) CancelLocationExecute(r openapinlmflocationclient.ApiCancelLocationRequest) (*http.Response, error) {
	return r.Execute()
}
