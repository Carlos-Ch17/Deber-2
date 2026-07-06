import "DocenteEvaluacion/docente"

import "testing"

// ===============================
// TEST: método privado promedio
// ===============================

func TestCalcularPromedioInterno(t *testing.T) {

	e := NuevaEvaluacion("1", "DOC1")
	e.Puntaje = 85

	resultado := e.calcularPromedioInterno()

	if resultado != 85 {
		t.Errorf("Se esperaba 85 pero se obtuvo %f", resultado)
	}
}

// ===============================
// TEST: agregar comentario privado
// ===============================

func TestAgregarComentario(t *testing.T) {

	e := NuevaEvaluacion("1", "DOC1")

	e.agregarComentario("Buen trabajo")
	e.agregarComentario("Mejorar redacción")

	if len(e.comentarios) != 2 {
		t.Errorf("Se esperaban 2 comentarios pero hay %d", len(e.comentarios))
	}
}

// ===============================
// TEST: método público promedio
// ===============================

func TestCalcularPromedioPublico(t *testing.T) {

	e := NuevaEvaluacion("1", "DOC1")
	e.Puntaje = 90

	if e.CalcularPromedio() != 90 {
		t.Error("Error en CalcularPromedio()")
	}
}

// ===============================
// TEST: comentario público
// ===============================

func TestAgregarComentarioPublico(t *testing.T) {

	e := NuevaEvaluacion("1", "DOC1")

	ok := e.AgregarComentarioPublico("Excelente evaluación")

	if !ok {
		t.Error("Debería aceptar comentario válido")
	}

	if len(e.comentarios) != 1 {
		t.Error("No se agregó el comentario correctamente")
	}
}