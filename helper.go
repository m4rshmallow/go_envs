package envs

import "fmt"

func notifyOK(text string) {
	fmt.Printf("[ + ] %s\n", text)
}

func notifyKO(err string) {
	fmt.Printf("[ - ] %s\n", err)
}
