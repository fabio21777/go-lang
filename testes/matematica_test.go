package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 7.28
	valorEncontrado := Media(7.2, 9.9, 6.1, 5.9)
	if valorEsperado != valorEncontrado {
		t.Errorf(erroPadrao, valorEsperado, valorEncontrado)
	}
}
