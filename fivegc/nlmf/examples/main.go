package main

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc/nlmf"
	openapicommon "github.com/5GCoreNet/openapi/openapi_CommonData"
	nlmfbroadcast "github.com/5GCoreNet/openapi/openapi_Nlmf_Broadcast"
	"log"
)

type MyBroadcast struct {
}

func (m MyBroadcast) Error(ctx context.Context, err error) openapicommon.ProblemDetails {
	return openapicommon.ProblemDetails{
		Type:               fivegc.ToString("error"),
		Title:              fivegc.ToString("error"),
		Status:             fivegc.ToInt32(int32(fivegc.StatusInternalServerError)),
		Detail:             fivegc.ToString(err.Error()),
		Instance:           fivegc.ToString("fake_instance"),
		Cause:              fivegc.ToString("unknown"),
		InvalidParams:      nil,
		SupportedFeatures:  fivegc.ToString(""),
		AccessTokenError:   &openapicommon.AccessTokenErr{},
		AccessTokenRequest: &openapicommon.AccessTokenReq{},
		NrfId:              fivegc.ToString("1234567890"),
	}
}

func (m MyBroadcast) CipherKeyData(ctx context.Context, data nlmfbroadcast.CipherRequestData) (nlmfbroadcast.CipherResponseData, openapicommon.ProblemDetails, fivegc.RedirectResponse, nlmf.CypherResponseStatusCode) {
	// Your code here ...
	return nlmfbroadcast.CipherResponseData{}, openapicommon.ProblemDetails{}, fivegc.RedirectResponse{}, nlmf.CypherResponseStatusCodeOK
}

func main() {
	m := MyBroadcast{}
	nlmfServer := nlmf.NewServer(":8080", "/v1/", log.Default())
	nlmfServer.AttachBroadcast(m)
	nlmfServer.Start()
	// Your code here ...
	nlmfServer.Stop()
}
