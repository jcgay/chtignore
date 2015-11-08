package main

import (
	"bytes"
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
