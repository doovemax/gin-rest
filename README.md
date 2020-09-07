# gin-rest

#### 介绍
golang gin restful
将stuct 结构体方法注册称为路由
函数命名规则必须是以下名字:
1. Get
2. Post
3. Put
4. Delete
5. Head
6. Options
7. Patch

#### Example:
```go
package main

import (
    "github.com/gin-gonic/gin"
    gin_rest "gitee.com/doove/gin-rest"
)

type ATest struct {
}

func (a *ATest) Get(ctx *gin.Context) {
	ctx.String(200, "")
}

func (a *ATest) Post(ctx *gin.Context) {
	ctx.String(200, "")
}

func main(){
	gin.SetMode(gin.TestMode)
	g := gin.Default()
	gin_rest.RouterEngine(g, "/test", &ATest{})
	for _, route := range g.Routes() {
		t.Logf("{%s %s %s}\n", route.Method, route.Path, route.Handler)
	}
}
```

