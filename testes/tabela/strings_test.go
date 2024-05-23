package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Testando", "Tes", 0},
		{"Testando", "ando", 4},
		{"Testando", "teste", -1},
		{"", "", 0},
		{"Opa", "opa", -1},
	}

	for _, teste := range testes {
		t.Logf("Massa de teste: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)
		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
