package gin_rest

import (
	"testing"

	"github.com/gin-gonic/gin"
)

type ATest struct {
}

func (a *ATest) Get(ctx *gin.Context) {
	ctx.String(200, "")
}

func (a *ATest) Post(ctx *gin.Context) {
	ctx.String(200, "")
}

func (a *ATest) Delete(ctx *gin.Context) {
	ctx.String(200, "")
}

func (a *ATest) Put(ctx *gin.Context) {
	ctx.String(200, "")
}

func TestRouterEngine(t *testing.T) {
	gin.SetMode(gin.TestMode)
	g := gin.Default()
	RouterEngine(g, "/test", &ATest{})
	for _, route := range g.Routes() {
		t.Logf("{%s %s %s}\n", route.Method, route.Path, route.Handler)
	}
}
