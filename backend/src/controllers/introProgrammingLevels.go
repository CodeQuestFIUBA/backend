package controllers

import (
	"fmt"
	"net/http"

	"codequest/src/models"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"rogchap.com/v8go"
)

func IntroFunctions() gin.HandlerFunc {
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
			var caminoCorrecto = Math.floor(Math.random() * 3) + 1;
			var tomoMapa = false;
			var completado = false;

			const resultado = {
				caminoCorrecto: caminoCorrecto,
				tomoMapaCorrectamente: false,
				caminoElegido: 0,
				completado: false
			};

			var recogerMapa = function() {
				tomoMapa = true;
			}

			var buscarCamino = function() {
				resultado.tomoMapaCorrectamente = tomoMapa;
				return caminoCorrecto;
			}

			var elegirCamino = function(camino) {
				resultado.caminoElegido = camino;
				resultado.completado = camino === caminoCorrecto && resultado.tomoMapaCorrectamente;
			}
			
			var userFunction = `

		executeUserFunctionCode := "\nuserFunction()"

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

		resultExercise, _ := ctx.RunScript("resultado", "main.js")

		fmt.Printf("Value: %s\n\n", resultExercise)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{
				Logs:          executionLogs,
				Result:        resultExercise,
				VariableTrace: executionTrace}})
	}
}

func IntroIfElse() gin.HandlerFunc {
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
			var direccion = Math.floor(Math.random() * 2);
			
			var getPosicionPrimerGuardia = function() {
				if (direccion === 0) {
					return "IZQUIERDA";
				} else {
					return "DERECHA";
				}
			}

			var getPosicionSegundoGuardia = function() {
				if (direccion === 0) {
					return "DERECHA";
				} else {
					return "IZQUIERDA";
				}
			}

			var resultado = {
				posiciones: [],
				posicionesGuardias: [getPosicionPrimerGuardia(), getPosicionSegundoGuardia()],
			};

			var moverALaIzquierda = function() {
				resultado.posiciones.push("IZQUIERDA");
			}

			var moverALaDerecha = function() {
				resultado.posiciones.push("DERECHA");
			}
			
			var userFunction = `

		executeUserFunctionCode := "\nuserFunction()"

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

		resultExercise, _ := ctx.RunScript("resultado", "main.js")

		fmt.Printf("Value: %s\n\n", resultExercise)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{
				Logs:          executionLogs,
				Result:        resultExercise,
				VariableTrace: executionTrace}})
	}
}

func IntroFor() gin.HandlerFunc {
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
			var copias = Math.floor(Math.random() * 4) + 1;
			if (copias === 1) {
				copias = 2;
			}
			
			var resultado = {
				copias: copias,
				copiasUsuario: 0,
			};

			var getCantidadDeCopias = function() {
				return copias;
			}

			var generarCopia = function() {
				resultado.copiasUsuario++;
			}
			
			var userFunction = `

		executeUserFunctionCode := "\nuserFunction()"

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

		resultExercise, _ := ctx.RunScript("resultado", "main.js")

		fmt.Printf("Value: %s\n\n", resultExercise)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{
				Logs:          executionLogs,
				Result:        resultExercise,
				VariableTrace: executionTrace}})
	}
}

func IntroWhile() gin.HandlerFunc {
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
			function generarArray() {
				const x = Math.floor(Math.random() * 3) + 3;
				const array = [];
				const numEsclavos = Math.floor(Math.random() * (x - 2)) + 1;
				const totalElementos = x + numEsclavos;

				for (let i = 0; i < x; i++) {
					array.push("GUARDIA");
				}

				for (let i = 0; i < numEsclavos; i++) {
					let posicion;
					do {
						posicion = Math.floor(Math.random() * (totalElementos - 2)) + 1;
					} while (array[posicion] === "ESCLAVO");
					array.splice(posicion, 0, "ESCLAVO");
				}

				if (array[0] !== "GUARDIA") {
					array[0] = "GUARDIA";
				}
				if (array[array.length - 1] !== "GUARDIA") {
					array[array.length - 1] = "GUARDIA";
				}
				return array;
			}

			let indice = 0;

			const resultado = {
				cola: generarArray(),
				colaDelUsuario: [],
			};

			const quedanGuardias = function() {
				return resultado.cola.length > indice;
			}

			const esGuardia = function() {
				return resultado.cola[indice] === "GUARDIA";
			}

			const atacar = function() {
				indice++;
				resultado.colaDelUsuario.push("GUARDIA");
			}

			const liberar = function() {
				indice++;
				resultado.colaDelUsuario.push("ESCLAVO");
			}
			
			var userFunction = `

		executeUserFunctionCode := "\nuserFunction()"

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

		resultExercise, _ := ctx.RunScript("resultado", "main.js")

		fmt.Printf("Value: %s\n\n", resultExercise)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{
				Logs:          executionLogs,
				Result:        resultExercise,
				VariableTrace: executionTrace}})
	}
}
