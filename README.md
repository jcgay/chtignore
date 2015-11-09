# chtignore

Print `.gitignore` template from [https://github.com/github/gitignore](https://github.com/github/gitignore) in standard output.

## Installation

Clone the repository, then

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

Use multiple arguments to get templates at once:

    chtignore Go JetBrains

Use `list` to discover available templates:

    chtignore list

## Build

### Status

[![Build Status](https://travis-ci.org/jcgay/chtignore.svg)](https://travis-ci.org/jcgay/chtignore)

### Release

- Configure Bintray deployment in `.goxc.local.json`:

```
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

