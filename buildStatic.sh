#!/bin/bash
# (c) Xavier Gandillot - 2019

echo "Building statically-linked go files"
echo "This takes longer than the usual 'go build'"
echo "Please, be patient ..."
# disable use of C from GO, forcing to recompile librairies, declare os and architecture
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64

############ flags explanation set to build statically linked binaries ########

### go build flags ###
# -a
#       force rebuilding of packages that are already up-to-date.
# -ldflags '[pattern=]arg list'
#       arguments to pass on each go tool link invocation.
# -tags 'tag list'
#	      a space-separated list of build tags to consider satisfied during the
#	      build. For more information about build tags, see the description of
#	      build constraints in the documentation for the go/build package.

#### go tool link flags ###
# -extldflags flags
#     	Set space-separated flags to pass to the external linker (gcc)
# -s
#	      Omit the symbol table and debug information.
# -w
#     	Omit the DWARF symbol table.

### gcc ldflags ###
# -static
#       On systems that support dynamic linking, this overrides -pie and prevents
#       linking with the shared libraries.
#       On other systems, this option has no effect.

### undertanding netgo vs netcgo (or cgo)  tags
#
# When the net package is compiled, it choose to use a "pure go DNS resolver"
# (netGo) or a cgo based resolver (netcgo), that calls C librairy routines.
# "netgo" is normally the default, but for some strange reason, it does not work
# if the netgo constraint tag is not specified ...
########################################################################
go build -a -tags netgo -ldflags='-w -s -extldflags "-static"' -o bin/main.static  ./main
echo "Done."
