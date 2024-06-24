#!/bin/bash

mkdir -p tmp
awk '/^```go$/ {p=1}; p; /^```$/ {p=0;print"--- --- ---"}' README.md |
	grep -vE '^```(go)?$' |
	csplit \
		-z -s -\
		'/--- --- ---/' \
		'{*}' \
		--suppress-matched \
		-f tmp/readme_snippet_ -b '%02d.go'

for snippet in ./tmp/readme_snippet_*.go
do
	echo "building $snippet ..."
	go build -v "$snippet" || exit 1
done

