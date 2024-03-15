package models

type ExecutionResponse struct {
	Logs          string      `json:"logs"`
	VariableTrace []string    `json:"variableTrace"`
	Result        interface{} `json:"result"`
}