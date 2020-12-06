module github.com/matthewapeters/dali/examples/three

go 1.15

require (
	github.com/lucasb-eyer/go-colorful v1.0.3
	github.com/matthewapeters/dali v0.1.1
)

replace (
	github.com/matthewapeters/dali => ../../
	github.com/zserge/lorca => ../../vendor/github.com/zserge/lorca
)
