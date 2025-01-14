package semantics

type Environment struct {
	enclosing *Environment
	values    map[string]interface{}
}

func InitEnvironment(enclosing *Environment) *Environment {
	return &Environment{enclosing: enclosing, values: make(map[string]interface{})}
}

func (e *Environment) define(name string, value interface{}) {
	e.values[name] = value
	// fmt.Printf("This is the environment %v", e.values)
}

func (e *Environment) get(name Token) interface{} {
	// fmt.Printf("This is the environment during get %v", e.values)
	if _, ok := e.values[name.Lexeme]; ok {
		return e.values[name.Lexeme]
	}

	if e.enclosing != nil {
		return e.enclosing.get(name)
	}
	panic(&RuntimeError{token: name, message: "Undefined variable '" + name.Lexeme + "'."})
}

func (e *Environment) assign(name Token, value interface{}) {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
		return
	}

	if e.enclosing != nil {
		e.enclosing.assign(name, value)
		return
	}

	panic(&RuntimeError{token: name, message: "Undefined variable '" + name.Lexeme + "'."})
}
