package nausf

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/models"
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
	GenerateUPUData(ctx context.Context, supi models.Supi, info UpuInfo) (UpuSecurityInfo, models.ProblemDetails)
}
