package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stderr, "", 0)

func main() {
	process(os.Args[1:], os.Stdout)
}

func process(args []string, output io.Writer) {
	if len(args) == 0 {
		missingArgument()
	}

	candidate := args[0]
	if candidate == "" {
		missingArgument()
	}

	fmt.Fprint(output, tryGetTemplate(candidate))
}

func tryGetTemplate(template string) string {
	resp := get(fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/master/%s.gitignore", template))
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		resp = get(fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/master/Global/%s.gitignore", template))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatal(err)
	}

	return string(body)
}

func get(url string) (resp *http.Response) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Fatal(err)
	}

	return resp
}

func missingArgument() {
	logger.Fatal("Mandatory argument missing, use: chtignore <template>")
}
