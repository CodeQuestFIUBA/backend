package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rogchap.com/v8go"
)

func JsExecute() gin.HandlerFunc {
	return func(c *gin.Context) {

		body, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}

		fmt.Printf("Request Body:\n\n%s\n\n", body)

		iso := v8go.NewIsolate()
		defer iso.Dispose()
		ctx := v8go.NewContext(iso)
		defer ctx.Close()
		global := ctx.Global()

		console := v8go.NewObjectTemplate(iso)
		logfn := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
			fmt.Println(info.Args()[0])
			return nil
		})
		console.Set("log", logfn)
		consoleObj, _ := console.NewInstance(ctx)

		global.Set("console", consoleObj)

		_, err = ctx.RunScript(string(body), "main.js")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed executing de JavaScript code -> " + err.Error()})
			return
		}

		val, _ := ctx.RunScript("result", "main.js")

		fmt.Printf("Value: %s\n\n", val)

		c.JSON(http.StatusOK, gin.H{"message": "Request body executed"})
	}
}
