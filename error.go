package iris_handler_error

import (
	"github.com/kataras/iris/v12"
)

// StatusMovedPermanently = 301
// RFC 7231, 6.4.2
func MovedPermanently(ctx iris.Context) {
	message := ctx.Values().Get("LoggerMessage")
	if message == nil {
		message = ctx.Translate("301")
		// You can use this function to Set and Get local values
		// that can be used to share information between handlers and middleware.
		ctx.Values().Set("LoggerMessage", message)
	}
	ctx.StatusCode(iris.StatusMovedPermanently)
	ctx.JSON(iris.Map{
		"code":    iris.StatusMovedPermanently,
		"success": false,
		"message": message,
	})
}

func BadRequest(ctx iris.Context) {
	message := ctx.Values().Get("LoggerMessage")
	if message == nil {
		message = ctx.Translate("400")
		ctx.Values().Set("LoggerMessage", message)
	}
	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(iris.Map{
		"code":    iris.StatusBadRequest,
		"success": false,
		"message": message,
	})
}

func Unauthorized(ctx iris.Context) {
	message := ctx.Values().Get("LoggerMessage")
	if message == nil {
		message = ctx.Translate("401")
		ctx.Values().Set("LoggerMessage", message)
	}
	// StatusUnauthorized = 401
	// RFC 7235, 3.1
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.JSON(iris.Map{
		"code":    iris.StatusUnauthorized,
		"success": false,
		"message": message,
	})
}

func Forbidden(ctx iris.Context) {
	message := ctx.Values().Get("LoggerMessage")
	if message == nil {
		message = ctx.Translate("403")
		ctx.Values().Set("LoggerMessage", message)
	}
	ctx.StatusCode(iris.StatusForbidden)
	ctx.JSON(iris.Map{
		"code":    iris.StatusForbidden,
		"success": false,
		"message": message,
	})
}

func NotFound(ctx iris.Context) {
	message := ctx.Values().Get("LoggerMessage")
	if message == nil {
		message = ctx.Translate("404")
		ctx.Values().Set("LoggerMessage", message)
	}
	//StatusNotFound = 404
	ctx.StatusCode(iris.StatusNotFound)
	ctx.JSON(
		iris.Map{
			"code":    iris.StatusNotFound,
			"success": false,
			"message": message,
			"data": iris.Map{
				"request_path": ctx.RequestPath(true),
			},
		})
}

func MethodNotAllowed(ctx iris.Context) {
	message := ctx.Values().Get("LoggerMessage")
	if message == nil {
		message = ctx.Translate("405")
		ctx.Values().Set("LoggerMessage", message)
	}
	ctx.StatusCode(iris.StatusMethodNotAllowed)
	ctx.JSON(iris.Map{
		"code":    iris.StatusMethodNotAllowed,
		"success": false,
		"message": message,
	})
}

func Conflict(ctx iris.Context) {
	message := ctx.Values().Get("LoggerMessage")
	if message == nil {
		message = ctx.Translate("409")
		ctx.Values().Set("LoggerMessage", message)
	}
	ctx.StatusCode(iris.StatusConflict)
	ctx.JSON(iris.Map{
		"code":    iris.StatusConflict,
		"success": false,
		"message": message,
	})
}

func InternalServerError(ctx iris.Context) {
	message := ctx.Values().Get("LoggerMessage")
	if message == nil {
		message = ctx.Translate("500")
		ctx.Values().Set("LoggerMessage", message)
	}
	ctx.StatusCode(iris.StatusInternalServerError)
	ctx.JSON(iris.Map{
		"code":    iris.StatusInternalServerError,
		"success": false,
		"message": message,
	})
}
