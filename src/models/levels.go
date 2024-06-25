package models

type SubLevel struct {
	Title string `json:"title"`
	Key   string `json:"key"`
	Order int    `json:"order"`
}

type Level struct {
	Order     int        `json:"order"`
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

var Levels = []Level{
    {
        Order: 0,
        Title: "Introducción a la algoritmia",
        Key:   "introduction",
        SubLevels: []SubLevel{
            {Title: "if", Key: "0", Order: 0},
            {Title: "Varios if", Key: "1", Order: 1},
        },
    },
    {
        Order: 1,
        Title: "Bases de la programación",
        Key:   "bases",
        SubLevels: []SubLevel{
            {Title: "Variables", Key: "0", Order: 0},
            {Title: "Ifs, operadores", Key: "1", Order: 1},
            {Title: "For, operadores", Key: "2", Order: 2},
            {Title: "While, operadores", Key: "3", Order: 3},
        },
    },
    {
        Order: 2,
        Title: "Funciones y procedimientos",
        Key:   "funciones_y_operadores",
        SubLevels: []SubLevel{
            {Title: "Armas", Key: "0", Order: 0},
            {Title: "Caminito", Key: "1", Order: 1},
            {Title: "Pensamiento lateral", Key: "2", Order: 2},
        },
    },
}
