# Buildversion

Позволяет маркировать билд текущим git-тегом, датой билда, названием текущей ветки, хэшем коммита и предназначением билда.

## Пример использования

```sh

#!/bin/sh
DATE=`date +%FT%T%z`
VERSION=`git describe --abbrev=0`
BRANCH=`git symbolic-ref --short -q HEAD`
COMMIT=`git log -n 1 --pretty=format:"%h"`
ENV="staging"

BVERS='github.com/cannonflesh/buildversion'
LDFLAGS="-w -s -X $BVERS.BuildTime=$DATE -X $BVERS.Version=$VERSION -X $BVERS.Branch=$BRANCH -X $BVERS.Commit=$COMMIT -X $BVERS.Env=$ENV"

go build -ldflags "${LDFLAGS}" -v

```


## Получение данных

```go 

import (
    "fmt"
    
    "github.com/cannonflesh/buildversion"
)

func print() {
    fmt.Printf(
        "Version: %s, commit: %s, branch: %s, environment: %s, build time: %s", 
        buildversion.Verion, 
        buildversion.Commit, 
        buildversion.Branch, 
        buildversion.Env,
        buildversion.BuildTime,
    )
}

```
