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

func ExecuteBinarySearchLevel() gin.HandlerFunc {
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
			function generateUniqueSortedRandomArray(x) {
				const uniqueNumbers = new Set();

				while (uniqueNumbers.size < x) {
					uniqueNumbers.add(Math.floor(Math.random() * 100));
				}

				const arr = Array.from(uniqueNumbers).sort((a, b) => a - b);

				return arr;
			}
			var array = generateUniqueSortedRandomArray(14);
			var jutsuPosition = Math.floor(Math.random() * 7) + 7
			var jutsu = array[jutsuPosition];
			const coordenadasAcceso = [];
			const expectedSolution = {
				positions: array,
				jutsu: jutsu,
				jutsuPosition: jutsuPosition
			};

			var probarJutsu = function(pos) {
				coordenadasAcceso.push(pos);
				return array[pos] === jutsu;
			}

			var esMenor = function(pos) {
				return array[pos] < jutsu;
			}

			var esMayor = function(pos) {
				return array[pos] > jutsu;
			}

			var userFunction = `

		executeUserFunctionCode := "\nuserFunction(array)"

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

		expectedResult, _ := ctx.RunScript("expectedSolution", "main.js")
		resultArray, _ := ctx.RunScript("coordenadasAcceso", "main.js")

		fmt.Printf("Value: %s\n\n", resultArray)

		c.JSON(http.StatusOK, models.StandardResponse{Code: http.StatusOK, Message: "Request body executed",
			Data: models.ExecutionResponse{
				Logs:           executionLogs,
				Result:         resultArray,
				VariableTrace:  executionTrace,
				ExpectedResult: expectedResult}})
	}
}

func VectorMaxLevel() gin.HandlerFunc {
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
			function getPociones() {
				const pociones = [];
				while (pociones.length < 6) {
					const value = Math.floor(Math.random() * 9) + 1;
					if (!pociones.includes(value)) {
						pociones.push(value);
					}
				}
				for (let i = pociones.length - 1; i > 0; i--) {
					const j = Math.floor(Math.random() * (i + 1));
					[pociones[i], pociones[j]] = [pociones[j], pociones[i]];
				}
				return pociones;
			}
			function indiceDelMaximo(array) {
				let indiceMaximo = 0;
				for (let i = 1; i < array.length; i++) {
					if (array[i] > array[indiceMaximo]) {
						indiceMaximo = i;
					}
				}
				return indiceMaximo;
			}
			var pociones = getPociones();
			var maxIndex = indiceDelMaximo(pociones);
			var result = {
				pociones: pociones,
				maxIndex: maxIndex,
				index: 0
			};
			var executeLevel = () => {
				var index = userFunction(result.pociones);
				if (typeof index === 'number') {
					result.index = index;
				} else {
					result.index = -1;
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

func VectoLinealSearchLevel() gin.HandlerFunc {
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
			function getPergaminos() {
				const pergaminos = [
					"Sombras del Caos",
					"Llama Eterna",
					"Viento Cortante",
					"Tormenta Silenciosa",
					"Golpes Fantasmales",
					"Rayo Inmortal",
					"Espíritu del Dragón"
				];

				const posicionAleatoria = 6 + Math.floor(Math.random() * 3);
				pergaminos.splice(posicionAleatoria, 0, "Ataques devastadores");

				return pergaminos;
			}

			const pergaminos = getPergaminos();
			const posicionPergamino = pergaminos.indexOf("Ataques devastadores");

			var result = {
				pergaminos: pergaminos,
				posicionPergamino: posicionPergamino,
				index: 0
			};
			var executeLevel = () => {
				var index = userFunction(result.pergaminos);
				if (typeof index === 'number') {
					result.index = index;
				} else {
					result.index = -1;
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