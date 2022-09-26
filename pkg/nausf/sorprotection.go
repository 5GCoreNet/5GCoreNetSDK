package nausf

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/models/subscription"
)

type SorProtectionHandler interface {
	Sor(ctx context.Context, supi subscription.Supi) error
}
