package util

import "testing"

func TestStringToTime(t *testing.T) {
	convertedTime := StringToTime("2019-02-12T10:10:10")

	if convertedTime.Year() != 2019 {
		t.Errorf("Ano esperado é 2019. Resultado: %v", convertedTime.Year())
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Mês esperado é 02. Resultado: %v", convertedTime.Month())
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("Hora esperada é 10. Resultado: %v", convertedTime.Hour())
	}
}
