build:
	go build -o ./bin/main ./src/cmd

run: generate build
	./bin/main

generate:
	templ generate
	# find templates -name '*_templ.go' -exec mv -t ./src/generated/templates/ -- {} +