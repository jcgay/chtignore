# chtignore

Print `.gitignore` template from [https://github.com/github/gitignore](https://github.com/github/gitignore) in standard output.

## Installation

### Binaries

#### Darwin (Apple Mac)

 * [chtignore\_1.0.0\_darwin\_amd64.zip](https://dl.bintray.com//content/jcgay/tools/chtignore_1.0.0_darwin_amd64.zip?direct)

#### Linux

 * [chtignore\_1.0.0\_amd64.deb](https://dl.bintray.com//content/jcgay/tools/chtignore_1.0.0_amd64.deb?direct)
 * [chtignore\_1.0.0\_linux\_amd64.tar.gz](https://dl.bintray.com//content/jcgay/tools/chtignore_1.0.0_linux_amd64.tar.gz?direct)

#### MS Windows

 * [chtignore\_1.0.0\_windows\_amd64.zip](https://dl.bintray.com//content/jcgay/tools/chtignore_1.0.0_windows_amd64.zip?direct)

### From source, clone the repository, then

    go install
    
## Usage

    chtignore Java

will output:

```
# Java
*.class

# Mobile Tools for Java (J2ME)
.mtj.tmp/

# Package Files #
*.jar
*.war
*.ear

# virtual machine crash logs, see http://www.java.com/en/download/help/error_hotspot.xml
hs_err_pid*
```

Use multiple arguments to get separate templates at once:

    chtignore Go JetBrains

Use `list` to discover available templates:

    chtignore list

## Build

### Status

[![Build Status](https://travis-ci.org/jcgay/chtignore.svg?branch=master)](https://travis-ci.org/jcgay/chtignore)

### Release

- Configure Bintray deployment in `.goxc.local.json`:

```json
{
    "ConfigVersion": "0.9",
    "TaskSettings": {
        "bintray": {
            "apikey": "12d312314235afe56090932ea13234"
        }
    }
}
```

- run `goxc default bintray`

### List available tasks

    goxc -h tasks

