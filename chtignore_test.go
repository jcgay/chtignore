package main

import (
	"bytes"
	"fmt"
	"github.com/assertgo/assert"
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestGetUniqueTemplate(t *testing.T) {
	assert := assert.New(t)
	output := new(bytes.Buffer)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/github/gitignore/master/Java.gitignore",
		httpmock.NewStringResponder(200, "*.class"))

	process([]string{"Java"}, output)

	assert.ThatString(output.String()).IsEqualTo(
		`# Java
*.class
`)
}

func TestGetUniqueGlobalTemplate(t *testing.T) {
	assert := assert.New(t)
	output := new(bytes.Buffer)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/github/gitignore/master/Vagrant.gitignore",
		httpmock.NewStringResponder(404, "Not Found"))
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/github/gitignore/master/Global/Vagrant.gitignore",
		httpmock.NewStringResponder(200, ".vagrant/"))

	process([]string{"Vagrant"}, output)

	assert.ThatString(output.String()).IsEqualTo(
		`# Vagrant
.vagrant/
`)
}

func TestTemplateStartWithUpperCase(t *testing.T) {
	assert := assert.New(t)
	output := new(bytes.Buffer)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/github/gitignore/master/Java.gitignore",
		httpmock.NewStringResponder(200, "*.class"))

	process([]string{"java"}, output)

	assert.ThatString(output.String()).IsEqualTo(
		`# Java
*.class
`)
}

func TestGetMultipleTemplates(t *testing.T) {
	assert := assert.New(t)
	output := new(bytes.Buffer)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/github/gitignore/master/Java.gitignore",
		httpmock.NewStringResponder(200, "*.class"))
	httpmock.RegisterResponder("GET", "https://raw.githubusercontent.com/github/gitignore/master/Go.gitignore",
		httpmock.NewStringResponder(200, "*.o"))

	process([]string{"Java", "Go"}, output)

	assert.ThatString(output.String()).IsEqualTo(
		`# Java
*.class
# Go
*.o
`)
}

func TestListAvailableTemplates(t *testing.T) {
	assert := assert.New(t)
	output := new(bytes.Buffer)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/github/gitignore/contents/",
		httpmock.NewStringResponder(200, `[
  {
    "name": "Go.gitignore",
    "path": "Go.gitignore",
    "sha": "daf913b1b347aae6de6f48d599bc89ef8c8693d6",
    "size": 266,
    "url": "https://api.github.com/repos/github/gitignore/contents/Go.gitignore?ref=master",
    "html_url": "https://github.com/github/gitignore/blob/master/Go.gitignore",
    "git_url": "https://api.github.com/repos/github/gitignore/git/blobs/daf913b1b347aae6de6f48d599bc89ef8c8693d6",
    "download_url": "https://raw.githubusercontent.com/github/gitignore/master/Go.gitignore",
    "type": "file",
    "_links": {
      "self": "https://api.github.com/repos/github/gitignore/contents/Go.gitignore?ref=master",
      "git": "https://api.github.com/repos/github/gitignore/git/blobs/daf913b1b347aae6de6f48d599bc89ef8c8693d6",
      "html": "https://github.com/github/gitignore/blob/master/Go.gitignore"
    }
  },
  {
    "name": "Java.gitignore",
    "path": "Java.gitignore",
    "sha": "32858aad3c383ed1ff0a0f9bdf231d54a00c9e88",
    "size": 189,
    "url": "https://api.github.com/repos/github/gitignore/contents/Java.gitignore?ref=master",
    "html_url": "https://github.com/github/gitignore/blob/master/Java.gitignore",
    "git_url": "https://api.github.com/repos/github/gitignore/git/blobs/32858aad3c383ed1ff0a0f9bdf231d54a00c9e88",
    "download_url": "https://raw.githubusercontent.com/github/gitignore/master/Java.gitignore",
    "type": "file",
    "_links": {
      "self": "https://api.github.com/repos/github/gitignore/contents/Java.gitignore?ref=master",
      "git": "https://api.github.com/repos/github/gitignore/git/blobs/32858aad3c383ed1ff0a0f9bdf231d54a00c9e88",
      "html": "https://github.com/github/gitignore/blob/master/Java.gitignore"
    }
  }
]
`))
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/github/gitignore/contents/Global",
		httpmock.NewStringResponder(200, `[
  {
    "name": "Vagrant.gitignore",
    "path": "Global/Vagrant.gitignore",
    "sha": "a977916f6583710870b00d50dd7fddd6701ece11",
    "size": 10,
    "url": "https://api.github.com/repos/github/gitignore/contents/Global/Vagrant.gitignore?ref=master",
    "html_url": "https://github.com/github/gitignore/blob/master/Global/Vagrant.gitignore",
    "git_url": "https://api.github.com/repos/github/gitignore/git/blobs/a977916f6583710870b00d50dd7fddd6701ece11",
    "download_url": "https://raw.githubusercontent.com/github/gitignore/master/Global/Vagrant.gitignore",
    "type": "file",
    "_links": {
      "self": "https://api.github.com/repos/github/gitignore/contents/Global/Vagrant.gitignore?ref=master",
      "git": "https://api.github.com/repos/github/gitignore/git/blobs/a977916f6583710870b00d50dd7fddd6701ece11",
      "html": "https://github.com/github/gitignore/blob/master/Global/Vagrant.gitignore"
    }
  }
]
`))

	process([]string{"list"}, output)

	assert.ThatString(output.String()).IsEqualTo(fmt.Sprint("Go, Java, Vagrant\n"))
}
