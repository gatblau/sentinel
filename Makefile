#
#    Sentinel - Copyright (c) 2019 by www.gatblau.org
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
#    Unless required by applicable law or agreed to in writing, software distributed under
#    the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
#    either express or implied.
#    See the License for the specific language governing permissions and limitations under the License.
#
#    Contributors to this project, hereby assign copyright in this code to the project,
#    to be licensed under the same terms as the rest of the code.
#

# the name of the Sentinel binary file
BINARY_NAME=sentinel

# the name of the go command to use to build the binary
GO_CMD = go

# the version of the container image to build
IMG_VER = v0.0.1

# the name of the folder where the packaged binaries will be placed after the build
BUILD_FOLDER=build

# build the Sentinel binary in the current platform
build:
	$(GO_CMD) fmt
	export GOROOT=/usr/local/go; export GOPATH=$HOME/go; $(GO_CMD) build -o $(BINARY_NAME) -v

# build the Sentinel docker image
docker-image:
	@docker build -t gatblau/$(BINARY_NAME):$(IMG_VER) .

# deletes images with no tag
docker-clean:
	@docker rmi $(docker images -f dangling=true -q)

# package the terraform provider for all platforms
package:
	go fmt
	$(MAKE) package_linux
	$(MAKE) package_darwin
	$(MAKE) package_windows

# package Sentinel for linux amd64 platform
package_linux:
	export GOROOT=/usr/local/go; export GOPATH=$(HOME)/go; export CGO_ENABLED=0; export GOOS=linux; export GOARCH=amd64; $(GO_CMD) build -o $(BUILD_FOLDER)/$(BINARY_NAME) -v
	zip -mjT $(BUILD_FOLDER)/$(BINARY_NAME)_linux_amd64.zip $(BUILD_FOLDER)/$(BINARY_NAME)

# package Sentinel for MacOS
package_darwin:
	export GOROOT=/usr/local/go; export GOPATH=$(HOME)/go; export CGO_ENABLED=0; export GOOS=darwin; export GOARCH=amd64; $(GO_CMD) build -o $(BUILD_FOLDER)/$(BINARY_NAME) -v
	zip -mjT $(BUILD_FOLDER)/$(BINARY_NAME)_darwin_amd64.zip $(BUILD_FOLDER)/$(BINARY_NAME)

# package Sentinel for Windows
package_windows:
	export GOROOT=/usr/local/go; export GOPATH=$(HOME)/go; export CGO_ENABLED=0; export GOOS=windows; export GOARCH=amd64; $(GO_CMD) build -o $(BUILD_FOLDER)/$(BINARY_NAME) -v
	zip -mjT $(BUILD_FOLDER)/$(BINARY_NAME)_windows_amd64.zip $(BUILD_FOLDER)/$(BINARY_NAME)