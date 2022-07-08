all: frontend/public/bee.wasm

frontend/public/bee.wasm: bee
	cp $< $@

bee: *.go **/*.go
	GOOS=js GOARCH=wasm go build