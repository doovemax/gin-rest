package gin_rest

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	GET     = "Get"
	POST    = "Post"
	PUT     = "Put"
	DELETE  = "Delete"
	HEAD    = "Head"
	OPTIONS = "Options"
	PATCH   = "Patch"
)

var methods = []string{
	GET,
	POST,
	PUT,
	DELETE,
	HEAD,
	OPTIONS,
	PATCH,
}

// RouterGroup add the struct method to gin router
func RouterGroup(router *gin.RouterGroup, path string, object interface{}) {
	or := reflect.ValueOf(object)
	for _, method := range methods {
		switch {
		case method == GET && or.MethodByName(method).IsValid():
			router.GET(path, or.MethodByName(GET).Interface().(func(ctx *gin.Context)))
		case method == POST && or.MethodByName(POST).IsValid():
			router.POST(path, or.MethodByName(POST).Interface().(func(ctx *gin.Context)))
		case method == PUT && or.MethodByName(PUT).IsValid():
			router.PUT(path, or.MethodByName(PUT).Interface().(func(ctx *gin.Context)))
		case method == DELETE && or.MethodByName(DELETE).IsValid():
			router.DELETE(path, or.MethodByName(DELETE).Interface().(func(ctx *gin.Context)))
		case method == HEAD && or.MethodByName(HEAD).IsValid():
			router.HEAD(path, or.MethodByName(HEAD).Interface().(func(ctx *gin.Context)))
		case method == OPTIONS && or.MethodByName(OPTIONS).IsValid():
			router.OPTIONS(path, or.MethodByName(OPTIONS).Interface().(func(ctx *gin.Context)))
		case method == PATCH && or.MethodByName(PATCH).IsValid():
			router.PATCH(path, or.MethodByName(PATCH).Interface().(func(ctx *gin.Context)))
		default:
		}
	}
}

func RouterEngine(engine *gin.Engine, path string, object interface{}) {
	or := reflect.ValueOf(object)
	for _, method := range methods {
		switch {
		case method == GET && or.MethodByName(method).IsValid():
			engine.GET(path, or.MethodByName(GET).Interface().(func(ctx *gin.Context)))
		case method == POST && or.MethodByName(POST).IsValid():
			engine.POST(path, or.MethodByName(POST).Interface().(func(ctx *gin.Context)))
		case method == PUT && or.MethodByName(PUT).IsValid():
			engine.PUT(path, or.MethodByName(PUT).Interface().(func(ctx *gin.Context)))
		case method == DELETE && or.MethodByName(DELETE).IsValid():
			engine.DELETE(path, or.MethodByName(DELETE).Interface().(func(ctx *gin.Context)))
		case method == HEAD && or.MethodByName(HEAD).IsValid():
			engine.HEAD(path, or.MethodByName(HEAD).Interface().(func(ctx *gin.Context)))
		case method == OPTIONS && or.MethodByName(OPTIONS).IsValid():
			engine.OPTIONS(path, or.MethodByName(OPTIONS).Interface().(func(ctx *gin.Context)))
		case method == PATCH && or.MethodByName(PATCH).IsValid():
			engine.PATCH(path, or.MethodByName(PATCH).Interface().(func(ctx *gin.Context)))
		default:
		}
	}
}
