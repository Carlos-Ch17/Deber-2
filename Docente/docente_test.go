import "DocenteEvaluacion/docente"

import "testing"

// ===============================
// TEST: validar email (privado)
// ===============================

func TestValidarEmailPrivado(t *testing.T) {

	casos := []struct {
		email  string
		valido bool
	}{
		{"juan@correo.com", true},
		{"ana@gmail.com", true},
		{"correo.com", false},
		{"correo@", false},
		{"", false},
	}

	for _, c := range casos {

		d := NuevoDocente("1", "Juan", c.email, "Sistemas", "Profesor")

		if d.validarEmail() != c.valido {
			t.Errorf("Error en email: %s", c.email)
		}
	}
}

// ===============================
// TEST: normalizar nombre (privado)
// ===============================

func TestNormalizarNombrePrivado(t *testing.T) {

	casos := []struct {
		entrada string
		salida  string
	}{
		{"juan perez", "Juan Perez"},
		{"JUAN PEREZ", "Juan Perez"},
		{"mArIa lOpEz", "Maria Lopez"},
	}

	for _, c := range casos {

		d := NuevoDocente("1", c.entrada, "test@test.com", "Sistemas", "Profesor")

		d.normalizarNombre()

		if d.Nombre != c.salida {
			t.Errorf("Esperado %s pero obtuvo %s", c.salida, d.Nombre)
		}
	}
}

// ===============================
// TEST: agregar evaluación interna
// ===============================

func TestAgregarEvaluacionInterna(t *testing.T) {

	d := NuevoDocente("1", "Juan", "juan@correo.com", "Sistemas", "Profesor")

	d.agregarEvaluacionInterna("EV1")
	d.agregarEvaluacionInterna("EV2")

	if len(d.evaluaciones) != 2 {
		t.Error("No se agregaron evaluaciones correctamente")
	}
}