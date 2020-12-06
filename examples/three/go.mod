module github.com/matthewapeters/dali/examples/three

go 1.14

require (
	github.com/lucasb-eyer/go-colorful v1.0.3
	github.com/matthewapeters/dali v0.1.1
)

replace (
	github.com/matthewapeters/dali v0.1.1 => ../../
	github.com/zserge/lorca v0.1.9 => ../../vendor/github.com/zserge/lorca
)
