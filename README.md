# hello-go
Repo and structure to help developers learn Go easily!

This repo is meant to help developers get started with learning and using Go. It provides both example syntax and operations in the `hello.go` file, some sample test cases in `hello_test.go`, and finally a structure with which to run all these Go programs within a docker container.

## Prerequisites

Please make sure to have Docker up and running on your host machine. You also must be able to execute bash commands via a command line interface.

Please note: you do not necessarily have to have Go installed on your local to get started!

## Getting started

To get started, simply run `bash run_docker.sh`. This will kick off the process to download a Golang docker image, create the appropriate mount points and bash functions, and then opens up a bash session for you to run your Go programs.

Please note: the `run_docker.sh` script creates a bind mount to your present working directory, and maps it to /go/src/app/. This helps the go runtime have access to all your directories/files in this repo, but it also requires that you run this script from the root of the `hello-go` repository.

## Runtime and aliases

Please see the `.bashrc` file for the bash functions and other settings that are getting applied to the Golang docker image at runtime. These are applied before the bash session is started in the /go/src/app/ directory.

The following bash functions are currently supported:
* `build` - This function helps the user automatically download all dependencies for the project, as well as build the project. This is a required step for the `run` command below.
* `test` - This function helps the user run Go's test module on the entire project. Thus, this will run all unit test cases that can be found within the project structure. Note that these tests are run with a `-v` option, so that the user sees details around which test cases succeeded and failed, along with high-level test timings.
* `run` - This function helps the user run the project executable. This executable gets built by the `build` command above (though the user should feel free to leverage `go install` and `go build` as they see fit as well.)

## Further improvements

We intend to also include in the future a sample Dockerfile that helps developers build compiled Go containers that automatically execute their built projects at runtime. This image build process would fetch dependencies, install the project, and run all the tests at build-time.

Another area needing improvement is `hello.go` and `hello_test.go` - these files only contain a rather small sample of the syntax and capabilities that Go provides. We'd like to continue to expand these in the future.

Please feel free to fork and create a PR with changes if you see an area you'd like to help with!

## Thanks!