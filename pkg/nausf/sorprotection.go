package nausf

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/models"
)

type SorProtectionHandler interface {
	Sor(ctx context.Context, supi models.Supi) error
}
