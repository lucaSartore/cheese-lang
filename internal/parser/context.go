package parser

type Context struct {
	Functions map[string]Function
	Variables map[string]VariableType
}
