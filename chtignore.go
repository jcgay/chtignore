package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	process(os.Args[1:], os.Stdout)
}

func process(args []string, output io.Writer) {
	if len(args) == 0 {
		log.Fatal("Mandatory argument missing: chtignore Java")
	}

	template := args[0]
	if template == "" {
		log.Fatal("Mandatory argument missing: chtignore Java")
	}

	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/master/%s.gitignore", template))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(output, string(body))
}
