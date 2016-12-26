BINARY=bookshelf

VERSION=`git describe --tags`
BUILD=`date +%Y%m%d%H%M`
SOURCES:=$(shell find . -name "*.go")

LDFLAGS=-ldflags "-w -s -X main.ver=${VERSION} -X main.date=${BUILD}"

build: ${SOURCES}
	go build ${LDFLAGS} -o ${BINARY}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean
