package nausf

const (
	ApiVersion  = "v1"
	ApiNameAuth = "nausf-auth"
	ApiNameSor  = "nausf-sorprotection"
	ApiNameUpu  = "nausf-upuprotection"
)

type NausfConfig struct {
	UriScheme string
}

type Nausf struct {
	cfg           *NausfConfig
	auth          AuthHandler
	sorProtection SorProtectionHandler
	upuProtection UpuProtectionHandler
}

func NewNausf(cfg *NausfConfig) *Nausf {
	return &Nausf{
		cfg:           cfg,
		auth:          nil,
		sorProtection: nil,
		upuProtection: nil,
	}
}
