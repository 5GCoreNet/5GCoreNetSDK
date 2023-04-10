package header

import (
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestBindRedirectHeader(t *testing.T) {
	ginContext, _ := gin.CreateTestContext(&httptest.ResponseRecorder{})
	BindRedirectHeader(ginContext, fivegc.RedirectHeader{
		Location:  "http://localhost:8080",
		SbiTarget: "1234",
	})

	if ginContext.Writer.Header().Get("Location") != "http://localhost:8080" {
		t.Errorf("Location header not set")
	}

	if ginContext.Writer.Header().Get("3gpp-Sbi-Target-Nf-Id") != "1234" {
		t.Errorf("3gpp-Sbi-Target-Nf-Id header not set")
	}
}
