package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"unicode"
)

var logger = log.New(os.Stderr, "", 0)

func main() {
	process(os.Args[1:], os.Stdout)
}

func process(args []string, output io.Writer) {
	if len(args) == 0 {
		missingArgument()
	}

	if args[0] == "list" {
		fmt.Fprintln(output, listTemplates())
	} else {
		for _, candidate := range args {
			candidate = upperFirstChar(candidate)
			content := tryGetTemplate(candidate)

			if content != "" {
				fmt.Fprintf(output, "# %s\n", candidate)
				fmt.Fprintln(output, content)
			}
		}
	}
}

func tryGetTemplate(template string) string {
	resp := get(fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/master/%s.gitignore", template))
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		resp = get(fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/master/Global/%s.gitignore", template))
	}

	if resp.StatusCode != 200 {
		logger.Fatal("Cannot find a template for: ", template)
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

func upperFirstChar(str string) string {
	a := []rune(str)
	firstChar := a[0]

	if unicode.IsUpper(firstChar) {
		return str
	}
	a[0] = unicode.ToUpper(firstChar)
	return string(a)
}

func listTemplates() string {
	templates := make([]string, 0)
	templates = getAndAppend(templates, "https://api.github.com/repos/github/gitignore/contents/")
	templates = getAndAppend(templates, "https://api.github.com/repos/github/gitignore/contents/Global")
	sort.Strings(templates)
	return strings.Join(templates, ", ")
}

func getAndAppend(templates []string, url string) []string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Fatal(err)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Fatal(err)
		}

		result := make([]GitIgnoreTemplate, 0)
		if err := json.Unmarshal(body, &result); err != nil {
			logger.Fatal(err)
		}

		for _, template := range result {
			if template.Name != "" && strings.Contains(template.Name, ".gitignore") {
				templates = append(templates, strings.Replace(template.Name, ".gitignore", "", 1))
			}
		}
	} else {
		logger.Fatal("Cannot list templates: %s", resp.StatusCode)
	}

	return templates
}

type GitIgnoreTemplate struct {
	Name string `json:"name"`
}
