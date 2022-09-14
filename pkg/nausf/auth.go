package nausf

import "context"

type AuthHandler interface {
	Authentication(ctx context.Context) error
	AkaConfirmation(ctx context.Context) error
	EapSession(ctx context.Context) error
}
