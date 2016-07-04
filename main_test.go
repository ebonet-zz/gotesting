package gotesting

import "testing"

func TestSumWrong(t *testing.T) {

	if Sum(3, 5) != 8 {
		t.Errorf("Soma incorreta")
	}
}

func TestSum(t *testing.T) {

	threshold := 0.001

	if s := Sum(3, 5); (s < 8-threshold) || (s > 8+threshold) {
		t.Errorf("Value out of the threshold")
	}

}

func TestCreateMap(t *testing.T) {

	m := CreateMap("teste", 1234)
	valor, ok := m["teste"]

	if !ok {
		t.Errorf("Chave teste não encontrada")
	}

	if valor != 1234 {
		t.Errorf("Valor incorreto")
	}

	if _, ok := m["chaveincorreta"]; ok {
		t.Errorf("Chave inesperada no mapa")
	}

}
