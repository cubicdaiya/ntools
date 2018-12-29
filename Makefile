all: gip nip

gip: cmd/gip/*.go
	go build cmd/$@/$@.go

nip: cmd/nip/*.go
	go build cmd/$@/$@.go

fmt:
	go fmt ./...

clean:
	rm -rf gip nip
