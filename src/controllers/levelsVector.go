package controllers

import (
	"fmt"
	"net/http"

	"codequest/src/models"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"rogchap.com/v8go"
)

func ExecuteVectorLevel() gin.HandlerFunc {
	return func(c *gin.Context) {

		body, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				models.StandardResponse{
					Code:    http.StatusInternalServerError,
					Message: "Failed to read request body",
					Data:    nil})
			return
		}

		color.Set(color.FgYellow)
		fmt.Printf("Request Body:\n%s\n", body)
		color.Unset()

		iso := v8go.NewIsolate()
		defer iso.Dispose()
		ctx := v8go.NewContext(iso)
		defer ctx.Close()
		global := ctx.Global()

		executionLogs := ""

		console := v8go.NewObjectTemplate(iso)
		logfn := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
			executionLogs += info.Args()[0].String() + "\n"
			return nil
		})
		console.Set("log", logfn)
		consoleObj, _ := console.NewInstance(ctx)

		global.Set("console", consoleObj)

		executionTrace := []string{}

		trace := v8go.NewObjectTemplate(iso)
		savefn := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
			executionTrace = append(executionTrace, info.Args()[0].String())
			return nil
		})
		trace.Set("save", savefn)
		traceObj, _ := trace.NewInstance(ctx)

		global.Set("trace", traceObj)

		proxyCode := `
			var array =[];
			var arrayProxy = new Proxy(array, {
  				set: function (target, key, value) {
					trace.save('Setting value ' + value + ' at index ' + key);
      				return Reflect.set(...arguments);
  				}
			});

			var userFunction = `

		executeUserFunctionCode := "\nuserFunction(arrayProxy)"

		codeToExecute := proxyCode + string(body) + executeUserFunctionCode
		_, err = ctx.RunScript(codeToExecute, "main.js")

		color.Set(color.FgRed)
		fmt.Println(executionLogs)
		color.Unset()

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil})
			return
		}

		resultArray, _ := ctx.RunScript("array", "main.js")

		fmt.Printf("Value: %s\n\n", resultArray)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{Logs: executionLogs, Result: resultArray, VariableTrace: executionTrace}})
	}
}
