package models

type SubLevel struct {
	Title string `json:"title"`
	Key   string `json:"key"`
}

type Level struct {
	Title     string     `json:"title"`
	Key       string     `json:"key"`
	SubLevels []SubLevel `json:"sublevels"`
}

type LevelResponse struct {
	Title     string         `json:"title"`
	Key       string         `json:"key"`
	SubLevels []SubLevelInfo `json:"sublevels"`
}

type SubLevelInfo struct {
	Key   string `json:"key"`
	Title string `json:"title"`
}

type LevelKeys struct {
	Level     string `json:"level"`
	SubLevel  string `json:"subLevel"`
	Completed bool   `json:"completed"`
}

var Levels = []Level{
	{
		Title: "Introducción a la algoritmia",
		Key:   "introduction",
		SubLevels: []SubLevel{
			{Title: "Cruzando el charco", Key: "0"},
			{Title: "Navegando en bote", Key: "1"},
			{Title: "Infiltrándose", Key: "2"},
			{Title: "Boss", Key: "3"},
		},
	},
	{
		Title: "Bases de la programación",
		Key:   "bases_de_la_programacion",
		SubLevels: []SubLevel{
			{Title: "Armar la mochila", Key: "0"},
			{Title: "Encontrando el camino correcto", Key: "1"},
			{Title: "Esquivando los guardia", Key: "2"},
			{Title: "Clonación", Key: "3"},
			{Title: "Infiltración", Key: "4"},
			{Title: "Boss", Key: "5"},
		},
	},
	{
		Title: "Funciones y procedimientos",
		Key:   "funciones_y_operadores",
		SubLevels: []SubLevel{
			{Title: "Consejos de los maestros", Key: "0"},
			{Title: "Ventajas entre armas", Key: "1"},
			{Title: "La guardia personal", Key: "2"},
			{Title: "Boss", Key: "3"},
		},
	},
	{
		Title: "Vectores",
		Key:   "vectores",
		SubLevels: []SubLevel{
			{Title: "El papiro adecuado", Key: "0"},
			{Title: "La pócima más fuerte", Key: "1"},
			{Title: "Batalla de velocidad", Key: "2"},
			{Title: "La llave encontrada", Key: "3"},
			{Title: "Reparando la escalera", Key: "4"},
			{Title: "Ordenando interruptores", Key: "5"},
			{Title: "Buscando la llave", Key: "6"},
			{Title: "Boss", Key: "7"},
		},
	},
	{
		Title: "Buenas prácticas",
		Key:   "buenas_practicas",
		SubLevels: []SubLevel{
			{Title: "Boss", Key: "0"},
		},
	},
	{
		Title: "Desafios",
		Key:   "desafios",
		SubLevels: []SubLevel{
			{Title: "Ataque en masa", Key: "0"},
			{Title: "Esquivar la trampa", Key: "1"},
			{Title: "Tras la llave", Key: "2"},
			{Title: "Boss", Key: "3"},
		},
	},
}
