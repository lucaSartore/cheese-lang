package expressions

type Context struct {
	Functions map[string]Function
	Variables map[string]Variable
}

func MakeContext() Context {
	return Context{
		Functions: make(map[string]Function),
		Variables: make(map[string]Variable),
	}
}

func (c *Context) AddFunction(f Function) {
	c.Functions[f.Name] = f
}

func (c *Context) GetFunction(name string) (*Function, bool) {
	f, ok := c.Functions[name]
	return &f, ok
}

func (c *Context) SetVariable(name string, value VariableContainer) bool {
	v, ok := c.Variables[name]
	if !ok {
		return false
	}
	v.Value = value
	c.Variables[name] = v
	return true
}

func (c *Context) AddVariable(t Variable) {
	c.Variables[t.Name] = t
}

func (c *Context) GetVariable(name string) (*Variable, bool) {
	v, ok := c.Variables[name]
	return &v, ok
}

func GetVariable(localContext *Context, globalContext *Context, name string) (*Variable, bool) {
	v, ok := localContext.GetVariable(name)
	if !ok {
		v, ok = globalContext.GetVariable(name)
	}
	return v, ok
}

func GetFunction(localContext *Context, globalContext *Context, name string) (*Function, bool) {
	f, ok := localContext.GetFunction(name)
	if !ok {
		f, ok = globalContext.GetFunction(name)
	}
	return f, ok
}

func SetVariable(localContext *Context, globalContext *Context, name string, newValue VariableContainer) bool {
	ok := localContext.SetVariable(name, newValue)
	if !ok {
		ok = globalContext.SetVariable(name, newValue)
	}
	return ok
}
