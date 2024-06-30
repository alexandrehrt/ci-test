package main

import "testing"

func TestMain(t *testing.T) {
	total := Soma(20, 20)

	if total != 40 {
		t.Errorf("Resultado incorreto.\nResultado: %d\nEsperado: %d", total, 40)
	}
}
