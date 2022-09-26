package nausf

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/models/common"
	"github.com/5GCoreNet/5GCoreNetSDK/models/subscription"
)

type UpuSecurityInfo struct {
	UpuMacIausf string
	counterUpu  string
}

type UpuData struct {
	secPacket        string
	defaultConfNssai []string
}

type UpuInfo struct {
	UpuDataList []UpuData
}

type UpuProtectionHandler interface {
	GenerateUPUData(ctx context.Context, supi subscription.Supi, info UpuInfo) (UpuSecurityInfo, common.ProblemDetails)
}
