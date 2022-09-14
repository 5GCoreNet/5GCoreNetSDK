package nausf

import "context"

type SorProtectionHandler interface {
	Sor(ctx context.Context, supi string) error
}
