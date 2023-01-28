package nlmf

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	"github.com/5GCoreNet/5GCoreNetSDK/internal/header"
	openapinlmfbroadcastclient "github.com/5GCoreNet/client-openapi/Nlmf_Broadcast"
	openapinlmfbroadcastserver "github.com/5GCoreNet/server-openapi/Nlmf_Broadcast"
	"github.com/gin-gonic/gin"
)

type Broadcast interface {
	// Error returns a problem details, it is used to handle errors when unmarshalling the request.
	Error(ctx context.Context, err error) openapinlmfbroadcastserver.ProblemDetails
	// CipherKeyData returns a cipher response data, a problem details, a redirect response and a status code.
	CipherKeyData(context.Context, openapinlmfbroadcastserver.CipherRequestData) (openapinlmfbroadcastserver.CipherResponseData, openapinlmfbroadcastserver.ProblemDetails, fivegc.RedirectResponse, fivegc.StatusCode)
}

func attachBroadcastHandler(router *gin.RouterGroup, b Broadcast) {
	group := router.Group("/nlmf-broadcast/v1")
	{
		group.POST("/cipher-key-data", func(c *gin.Context) {
			var req openapinlmfbroadcastserver.CipherRequestData
			if err := c.ShouldBindJSON(&req); err != nil {
				problemDetails := b.Error(c, err)
				c.JSON(int(problemDetails.Status), problemDetails)
				return
			}
			res, problemDetails, redirectResponse, status := b.CipherKeyData(c, req)
			switch status {
			case fivegc.StatusOK:
				c.JSON(status.ToInt(), res)
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

// BroadcastClient is a client for the Nlmf_Broadcast service.
type BroadcastClient struct {
	client *openapinlmfbroadcastclient.APIClient
}

// NewBroadcastClient creates a new client for the Nlmf_Broadcast service.
func NewBroadcastClient(cfg fivegc.ClientConfiguration) *BroadcastClient {
	openapiCfg := &openapinlmfbroadcastclient.Configuration{
		Host:             cfg.Host,
		Scheme:           cfg.Scheme,
		DefaultHeader:    cfg.DefaultHeader,
		UserAgent:        cfg.UserAgent,
		Debug:            cfg.Debug,
		Servers:          []openapinlmfbroadcastclient.ServerConfiguration{},
		OperationServers: make(map[string]openapinlmfbroadcastclient.ServerConfigurations),
		HTTPClient:       cfg.HTTPClient,
	}
	for _, server := range cfg.Servers {
		openapiServer := openapinlmfbroadcastclient.ServerConfiguration{
			URL:         server.URL,
			Description: server.Description,
			Variables:   make(map[string]openapinlmfbroadcastclient.ServerVariable),
		}
		for name, variable := range server.Variables {
			openapiServer.Variables[name] = openapinlmfbroadcastclient.ServerVariable{
				Description:  variable.Description,
				DefaultValue: variable.DefaultValue,
				EnumValues:   variable.EnumValues,
			}
		}
		openapiCfg.Servers = append(openapiCfg.Servers, openapiServer)
	}
	for name, servers := range cfg.OperationServers {
		openapiServers := make(openapinlmfbroadcastclient.ServerConfigurations, len(servers))
		for i, server := range servers {
			openapiServers[i] = openapinlmfbroadcastclient.ServerConfiguration{
				URL:         server.URL,
				Description: server.Description,
				Variables:   make(map[string]openapinlmfbroadcastclient.ServerVariable),
			}
			for name, variable := range server.Variables {
				openapiServers[i].Variables[name] = openapinlmfbroadcastclient.ServerVariable{
					Description:  variable.Description,
					DefaultValue: variable.DefaultValue,
					EnumValues:   variable.EnumValues,
				}
			}
		}
		openapiCfg.OperationServers[name] = openapiServers
	}
	return &BroadcastClient{
		client: openapinlmfbroadcastclient.NewAPIClient(openapiCfg),
	}
}

// CipheringKeyData returns a cipher request.
func (c *BroadcastClient) CipheringKeyData(ctx context.Context) openapinlmfbroadcastclient.ApiCipheringKeyDataRequest {
	return c.client.RequestCipheringKeyDataApi.CipheringKeyData(ctx)
}

// CipheringKeyDataExecute executes a cipher request.
func (c *BroadcastClient) CipheringKeyDataExecute(r openapinlmfbroadcastclient.ApiCipheringKeyDataRequest) (*openapinlmfbroadcastclient.CipherResponseData, error) {
	resp, _, err := r.Execute()
	return resp, err
}
