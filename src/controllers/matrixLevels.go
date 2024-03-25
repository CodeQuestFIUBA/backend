package controllers

import (
	"fmt"
	"net/http"

	"codequest/src/models"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"rogchap.com/v8go"
)

func ExecuteSearchMatrixLevel() gin.HandlerFunc {
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
			const anchoMatriz = 10;
			const altoMatriz = 6;
			const matriz = new Array(altoMatriz);
			
			for (var i=0; i<anchoMatriz; i++) {
				matriz[i] = new Array(altoMatriz);
				matriz[i].fill(0);
			}
			
			const keyX = Math.floor(Math.random() * anchoMatriz);
			const keyY = Math.floor(Math.random() * altoMatriz);

			matriz[keyY][keyX] = "llave";

			const coordenadasAcceso = [];

			const handler = {
        get: function(target, prop, receiver) {
					if (!isNaN(prop)) {
							const coordenadaX = parseInt(prop);
							const fila = target[prop];
							console.log("fila: " + prop + ", tipo de fila: " + typeof fila);
							return new Proxy(fila, {
									get: function(target, prop) {
											if (!isNaN(prop)) {
													const coordenadaY = parseInt(prop);
													trace.save('Accediendo a la coordenada [' + coordenadaX + ', ' + coordenadaY + ']');
											}
											return Reflect.get(...arguments);
									}
							});
					}
					return Reflect.get(...arguments);
			}
    };

			var matrixProxy = new Proxy(matriz, handler);

			var userFunction = `

		executeUserFunctionCode := "\nuserFunction(matrixProxy, anchoMatriz, altoMatriz)\n"

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

		resultMatrix, _ := ctx.RunScript("matriz", "main.js")

		fmt.Printf("Value: %s\n\n", resultMatrix)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{Logs: executionLogs, Result: resultMatrix, VariableTrace: executionTrace}})
	}
}
