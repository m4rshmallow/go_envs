package envs

import (
	"fmt"
	"os"
)

type Env map[string]string

func (e Env) Get(n string) string {
	value, ok := e[n]
	if !ok {
		notifyKO(fmt.Sprintf(" %s doesn't exists\n", n))
		os.Exit(1)
	}

	return value
}

func (e *Env) set(n, v string) {
	(*e)[n] = v
}
