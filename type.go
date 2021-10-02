package go_envs

import (
	"fmt"
)

type Env map[string]string

func (e Env) Get(n string) string {
	value, ok := e[n]
	if !ok {
		panic(fmt.Sprintf("%s doesn't exists\n", n))
	}

	return value
}

func (e *Env) set(n, v string) {
	(*e)[n] = v
}
