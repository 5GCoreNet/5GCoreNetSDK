package nlmf

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	"github.com/5GCoreNet/5GCoreNetSDK/internal/header"
	openapicommon "github.com/5GCoreNet/openapi/openapi_CommonData"
	openapinlmfbroadcast "github.com/5GCoreNet/openapi/openapi_Nlmf_Broadcast"
	"github.com/gin-gonic/gin"
)

const (
	broadcastRouterGroup = "/nlmf-broadcast/v1"
	cypherKeyEndpoint    = "/cipher-key-data"
)

// Broadcast is the interface that wraps the NLMF Broadcast service.
type Broadcast interface {
	fivegc.CommonInterface
	CipherKeyData(context.Context, openapinlmfbroadcast.CipherRequestData) (openapinlmfbroadcast.CipherResponseData, openapicommon.ProblemDetails, fivegc.RedirectResponse, CypherResponseStatusCode)
}

type CypherResponseStatusCode fivegc.StatusCode

const (
	// CypherResponseStatusCodeOK is the status code for a successful response.
	CypherResponseStatusCodeOK            CypherResponseStatusCode = CypherResponseStatusCode(fivegc.StatusOK)
	CypherResponseStatusTemporaryRedirect CypherResponseStatusCode = CypherResponseStatusCode(fivegc.StatusTemporaryRedirect)
	CypherResponseStatusPermanentRedirect CypherResponseStatusCode = CypherResponseStatusCode(fivegc.StatusPermanentRedirect)
)

func attachBroadcastHandler(router *gin.RouterGroup, b Broadcast) {
	group := router.Group(broadcastRouterGroup)
	{
		group.POST(cypherKeyEndpoint, func(c *gin.Context) {
			var req openapinlmfbroadcast.CipherRequestData
			if err := c.ShouldBindJSON(&req); err != nil {
				problemDetails := b.Error(c, err)
				c.JSON(int(*problemDetails.Status), problemDetails)
				return
			}
			res, problemDetails, redirectResponse, status := b.CipherKeyData(c, req)
			switch status {
			case CypherResponseStatusCodeOK:
				c.JSON(int(status), res)
			case CypherResponseStatusTemporaryRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(int(status), redirectResponse)
			case CypherResponseStatusPermanentRedirect:
				header.BindRedirectHeader(c, redirectResponse.RedirectHeader)
				c.JSON(int(status), redirectResponse)
			default:
				c.JSON(int(status), problemDetails)
			}
			return
		})
	}
}

// BroadcastClient is a client for the NLMF Broadcast service.
type BroadcastClient struct {
	client *openapinlmfbroadcast.APIClient
}

// NewBroadcastClient creates a new client for the NLMF Broadcast service.
func NewBroadcastClient(cfg fivegc.ClientConfiguration) *BroadcastClient {
	openapiCfg := &openapinlmfbroadcast.Configuration{
		Host:             cfg.Host,
		Scheme:           cfg.Scheme,
		DefaultHeader:    cfg.DefaultHeader,
		UserAgent:        cfg.UserAgent,
		Debug:            cfg.Debug,
		Servers:          []openapinlmfbroadcast.ServerConfiguration{},
		OperationServers: make(map[string]openapinlmfbroadcast.ServerConfigurations),
		HTTPClient:       cfg.HTTPClient,
	}
	for _, server := range cfg.Servers {
		openapiServer := openapinlmfbroadcast.ServerConfiguration{
			URL:         server.URL,
			Description: server.Description,
			Variables:   make(map[string]openapinlmfbroadcast.ServerVariable),
		}
		for name, variable := range server.Variables {
			openapiServer.Variables[name] = openapinlmfbroadcast.ServerVariable{
				Description:  variable.Description,
				DefaultValue: variable.DefaultValue,
				EnumValues:   variable.EnumValues,
			}
		}
		openapiCfg.Servers = append(openapiCfg.Servers, openapiServer)
	}
	for name, servers := range cfg.OperationServers {
		openapiServers := make(openapinlmfbroadcast.ServerConfigurations, len(servers))
		for i, server := range servers {
			openapiServers[i] = openapinlmfbroadcast.ServerConfiguration{
				URL:         server.URL,
				Description: server.Description,
				Variables:   make(map[string]openapinlmfbroadcast.ServerVariable),
			}
			for name, variable := range server.Variables {
				openapiServers[i].Variables[name] = openapinlmfbroadcast.ServerVariable{
					Description:  variable.Description,
					DefaultValue: variable.DefaultValue,
					EnumValues:   variable.EnumValues,
				}
			}
		}
		openapiCfg.OperationServers[name] = openapiServers
	}
	return &BroadcastClient{
		client: openapinlmfbroadcast.NewAPIClient(openapiCfg),
	}
}

// CipheringKeyData returns a cipher request.
func (c *BroadcastClient) CipheringKeyData(ctx context.Context) openapinlmfbroadcast.ApiCipheringKeyDataRequest {
	return c.client.RequestCipheringKeyDataApi.CipheringKeyData(ctx)
}

// CipheringKeyDataExecute executes a cipher request.
func (c *BroadcastClient) CipheringKeyDataExecute(r openapinlmfbroadcast.ApiCipheringKeyDataRequest) (*openapinlmfbroadcast.CipherResponseData, error) {
	resp, _, err := r.Execute()
	return resp, err
}
