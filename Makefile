.PHONY: wiki

wiki: tools
	git submodule update --init --recursive
	hugo -s wiki

tools: $GOPATH/bin/hugo

$GOPATH/bin/hugo:
	go install github.com/gohugoio/hugo

clean:
	rm -rf wiki/resources
	rm -rf wiki/public
