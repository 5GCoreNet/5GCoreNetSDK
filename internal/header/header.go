package header

import (
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	"github.com/gin-gonic/gin"
)

// BindRedirectHeader binds the redirect header to the gin context
func BindRedirectHeader(c *gin.Context, header fivegc.RedirectHeader) {
	c.Header("Location", header.Location)
	c.Header("3gpp-Sbi-Target-Nf-Id", header.SbiTarget)
}
