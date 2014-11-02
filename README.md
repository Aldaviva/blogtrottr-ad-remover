blogtrottr-ad-remover
=====================

Remove ads from Blogtrottr emails

## Compiling

1. Install [Go](https://golang.org/dl/), [Git](http://git-scm.com/downloads), (so you can clone this repo) and [Mercurial](http://mercurial.selenic.com/downloads) (so you can install repo dependencies).
2. `git clone https://github.com/Aldaviva/blogtrottr-ad-remover.git && cd blogtrottr-ad-remover`
3. Set your `GOPATH` environment variable to be the absolute path of the current working directory.
4. Download dependencies: `go get aldaviva.com/blogtrottr-ad-remover`
5. Compile: `go install aldaviva.com/blogtrottr-ad-remover`

The compiled binary can be found in `bin/`
