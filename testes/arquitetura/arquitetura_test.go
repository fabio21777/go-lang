package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")

	}
	if 1 != 1 {
		t.Error("Teste falhou")
	}
}

/*
Comandos legais

go test -v
go test -v -run TestDependente
go test -cover
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out
*/
