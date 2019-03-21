# ~/.bashrc: executed by bash(1) for non-login shells.

# Script replacing default .bashrc, defining shell functions for within the go serving as the runtime

# Build function that both downloads dependencies and installs your application
build () {
  go get -d -t -v ./... && go install -v ./...
}

# Test function that helps run go's test module
test () {
  go test -v
}

# Run function that helps run your built application
# NOTE: Can only run after you have already built and installed your application
run () {
  app
}
