# go-module-template
[![go.dev reference](https://pkg.go.dev/badge/github.com/soypat/go-module-template)](https://pkg.go.dev/github.com/soypat/go-module-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/soypat/go-module-template)](https://goreportcard.com/report/github.com/soypat/go-module-template)
<!--[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) -->
[![codecov](https://codecov.io/gh/soypat/go-module-template/branch/main/graph/badge.svg)](https://codecov.io/gh/soypat/go-module-template)

Go module template with instructions on how to make your code importable and setting up codecov CI.

## First steps
Fix `go.mod` file by replacing `github.com/YOURUSER/YOURREPONAME` with your corresponding project repository link.

## Setting up codecov CI
This instructive will allow for tests to run on pull requests and pushes to your repository.

1. Create an account on [codecov.io](https://app.codecov.io/)

2. Setup repository on codecov and obtain the CODECOV_TOKEN token, which is a string of base64 characters.

3. Open up the github repository for this project and go to `Settings -> Secrets and variables -> Actions`. Once there create a New Repository Secret. Name it `CODECOV_TOKEN` and copy paste the token obtained in the previous step in the `secret` input box. Click "Add secret".


