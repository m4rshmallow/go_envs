package go_envs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	EnvMap    = Env{}
	f         *os.File
	e         error
	envRegExp = regexp.MustCompile("(.+)=(.+)")
)

func openEnvFiles(filenames ...string) {
	for _, filename := range filenames {
		if f, e = os.Open(filename); e != nil {
			notifyKO(e.Error())
			continue
		}
		defer f.Close()

		readAndStoreEnv(f)
	}
}

func storeEnv(str string) (key, value string) {
	match := envRegExp.FindStringSubmatch(str)

	if len(match) <= 1 {
		return
	}

	if len(match) == 2 {
		fmt.Printf("%s doesn't has value\n", match[1])
		return
	}

	key = match[1]
	value = strings.TrimSpace(match[2])

	EnvMap.set(key, value)

	return key, value
}

func readAndStoreEnv(file *os.File) {
	read := bufio.NewReader(file)

	for {
		str, e := read.ReadString('\n')
		if e == io.EOF {
			setUpEnv(strings.TrimSpace(str))
			break
		}
		if e != nil {
			notifyKO(e.Error())
			break
		}

		setUpEnv(strings.TrimSpace(str))
	}

	notifyOK(fmt.Sprintf("%s - successfully loaded", file.Name()))
}

func setUpEnv(str string) {
	k, v := storeEnv(strings.TrimSpace(str))
	exportEnv(k, v)
}

func exportEnv(k, v string) {
	if e := os.Setenv(k, v); e != nil {
		panic(e)
	}
}

func Init(files ...string) {
	if len(files) == 0 {
		files = []string{".env"}
	}

	openEnvFiles(files...)
}
