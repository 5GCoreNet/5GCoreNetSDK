package nausf

import "context"

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
	GenerateUPUData(ctx context.Context, supi string, info UpuInfo) (UpuSecurityInfo, ProblemDetails)
}
