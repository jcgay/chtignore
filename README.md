# chtignore

Print `.gitignore` template from [https://github.com/github/gitignore](https://github.com/github/gitignore) in standard output.

## Installation

Clone the repository, then

    go install
    
## Usage

    chtignore Java

will output:

```
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

## Build Status

[![Build Status](https://travis-ci.org/jcgay/chtignore.svg)](https://travis-ci.org/jcgay/chtignore)