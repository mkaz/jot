# Customize Make
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
.RECIPEPREFIX = >
# end

build:
> echo "Building Jot..."
> go build ./cmd/jot
.PHONY: build

multiarch: build
> VERSION=`jot -version | sed -e 's/jot v//'`
> echo "Building multiarch for $(VERSION)..."
> gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -output "{{.Dir}}-$(VERSION)-{{.OS}}/{{.Dir}}"
> for arch in linux darwin windows; do
>   tar cf jot-$(VERSION)-$(arch).tar jot-$(VERSION)-$(arch)
>   gzip jot-$(VERSION)-$(arch).tar
>   rm -rf jot-$(VERSION)-$(arch)
> done
.PHONY: multiarch
