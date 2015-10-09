# Go lang TDD kurs

Install go lang:
```
brew install go
```

Hjelp deg selv litt med noen ENVs i ```~/.bashrc```:

```bash

# .bashrc : 
export GITHUB_USER="mitt_github_brukernavn"
export GOPATH="~/src/gopath"
export GOROOT="/usr/local/opt/go/libexec"
alias gogo="cd $GOPATH/src/github.com/${GITHUB_USER}" 


# sett opp go kommandoen på PATH:
# You may wish to add the GOROOT-based install location to your PATH:
export PATH="$PATH:/usr/local/opt/go/libexec/bin:$GOPATH/bin"


$ source ~/.bashrc
$ mkdir -p ${GOPATH}/github.com/${GITHUB_USER}
$ gogo # skal flytte deg til din github brukers mappe under Go sources.

$ go get github.com/${GITHUB_USER}/${REPO}
```

Simple:
```
$ go test -v .../.
```

Fancy:

```
$ go get github.com/smartystreets/goconvey
goconvey
```

