import "DocenteEvaluacion/docente"

import "testing"

func TestValidarEmailPrivado(t *testing.T) {

	pruebas := []struct {
		email  string
		valido bool
	}{
		{"juan@correo.com", true},
		{"ana@gmail.com", true},
		{"correo.com", false},
		{"correo@", false},
		{"", false},
	}

	for _, prueba := range pruebas {

		d := NuevoDocente(
			"1",
			"Juan",
			prueba.email,
			"Sistemas",
			"Profesor",
		)

		if d.validarEmail() != prueba.valido {
			t.Errorf("Error con el email %s", prueba.email)
		}
	}
}

func TestNormalizarNombrePrivado(t *testing.T) {

	pruebas := []struct {
		entrada string
		salida  string
	}{
		{"juan perez", "Juan Perez"},
		{"JUAN PEREZ", "Juan Perez"},
		{"mArIa lOpEz", "Maria Lopez"},
	}

	for _, prueba := range pruebas {

		d := NuevoDocente(
			"1",
			prueba.entrada,
			"correo@correo.com",
			"Sistemas",
			"Profesor",
		)

		d.normalizarNombre()

		if d.Nombre != prueba.salida {
			t.Errorf("Esperado %s pero obtuvo %s", prueba.salida, d.Nombre)
		}
	}
}

func TestAgregarEvaluacionInterna(t *testing.T) {

	d := NuevoDocente(
		"1",
		"Juan",
		"juan@correo.com",
		"Sistemas",
		"Profesor",
	)

	d.agregarEvaluacionInterna("EV001")
	d.agregarEvaluacionInterna("EV002")

	if len(d.evaluaciones) != 2 {
		t.Error("No se agregaron correctamente las evaluaciones")
	}

	if d.evaluaciones[0] != "EV001" {
		t.Error("Primera evaluación incorrecta")
	}

	if d.evaluaciones[1] != "EV002" {
		t.Error("Segunda evaluación incorrecta")
	}
}
