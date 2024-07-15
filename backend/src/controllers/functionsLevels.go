package controllers

import (
	"fmt"
	"net/http"

	"codequest/src/models"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"rogchap.com/v8go"
)

func FunctionsIntro() gin.HandlerFunc {
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
			var userFunction;
			function shuffleArray(array) {
				for (let i = array.length - 1; i > 0; i--) {
					const j = Math.floor(Math.random() * (i + 1));
					[array[i], array[j]] = [array[j], array[i]];
				}
				return array;
			}
			var armas = ['sai', 'katana', 'ninjaku', 'whip']
			var pedido = shuffleArray([...armas]);
			var result = {
				armas: pedido,
				pedido: {}
			};
			var executeLevel = () => {
				for (let i = 0; i < pedido.length; i++) {
					var msgs = [];
					for (let j = 0; j < armas.length; j++) {
						var msg = userFunction(pedido[i], armas[j]);
						msgs.push(msg.trim().toUpperCase());
					}
					result.pedido[pedido[i]] = msgs;
				}
			}
			
			userFunction = `

		executeUserFunctionCode := "\nexecuteLevel()"

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

		resultExercise, _ := ctx.RunScript("result", "main.js")

		fmt.Printf("Value: %s\n\n", resultExercise)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{
				Logs:          executionLogs,
				Result:        resultExercise,
				VariableTrace: executionTrace}})
	}
}

func FunctionsIfElse() gin.HandlerFunc {
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
			var userFunction;
			function shuffleArray(array) {
				for (let i = array.length - 1; i > 0; i--) {
					const j = Math.floor(Math.random() * (i + 1));
					[array[i], array[j]] = [array[j], array[i]];
				}
				return array;
			}
			var armas = ['sai', 'katana', 'ninjaku', 'whip']
			var pedido = shuffleArray([...armas]);
			var result = {
				pedidos: pedido,
				armas: []
			};
			var executeLevel = () => {
				for (let i = 0; i < pedido.length; i++) {
					result.armas.push(userFunction(pedido[i]));
				}
			}
			
			userFunction = `

		executeUserFunctionCode := "\nexecuteLevel()"

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

		resultExercise, _ := ctx.RunScript("result", "main.js")

		fmt.Printf("Value: %s\n\n", resultExercise)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{
				Logs:          executionLogs,
				Result:        resultExercise,
				VariableTrace: executionTrace}})
	}
}

func FunctionsObject() gin.HandlerFunc {
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
			var userFunction;
			var armas = ['sai', 'katana', 'ninjaku', 'whip']
			const copias = Math.floor(Math.random() * 2) + 2;
			var arma = armas[Math.floor(Math.random() * armas.length)];
			var result = {
				enemy_result: {
					total: copias,
					arma: arma
				},
				user_result: {
					total: 0,
					arma: ''
				}
			};
			var executeLevel = () => {
				const responseObject = userFunction(copias, arma);
				result.user_result.total = responseObject.total;
				result.user_result.arma = responseObject.arma;
			}
			
			userFunction = `

		executeUserFunctionCode := "\nexecuteLevel()"

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

		resultExercise, _ := ctx.RunScript("result", "main.js")

		fmt.Printf("Value: %s\n\n", resultExercise)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{
				Logs:          executionLogs,
				Result:        resultExercise,
				VariableTrace: executionTrace}})
	}
}
