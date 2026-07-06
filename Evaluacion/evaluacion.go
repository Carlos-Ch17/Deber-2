import "DocenteEvaluacion/docente"

import "strings"

type EstadoEvaluacion string

const (
	Borrador   EstadoEvaluacion = "borrador"
	EnProceso  EstadoEvaluacion = "en_proceso"
	Finalizada EstadoEvaluacion = "finalizada"
)

type Evaluacion struct {
	ID          string
	DocenteID   string
	Estado      EstadoEvaluacion
	Puntaje     float64
	comentarios []string
}

func NuevaEvaluacion(id, docenteID string) *Evaluacion {
	return &Evaluacion{
		ID:          id,
		DocenteID:   docenteID,
		Estado:      Borrador,
		Puntaje:     0,
		comentarios: []string{},
	}
}

// Métodos públicos existentes

func (e Evaluacion) GetEstado() string {
	return string(e.Estado)
}

func (e Evaluacion) EstaFinalizada() bool {
	return e.Estado == Finalizada
}

// =======================
// Métodos privados
// =======================

// Calcula el promedio de los puntajes.
// Como la estructura solo tiene un puntaje,
// el promedio corresponde al mismo valor.
func (e *Evaluacion) calcularPromedioInterno() float64 {
	return e.Puntaje
}

func (e *Evaluacion) agregarComentario(comentario string) {
	e.comentarios = append(e.comentarios, comentario)
}

// =======================
// Métodos públicos nuevos
// =======================

func (e *Evaluacion) CalcularPromedio() float64 {
	return e.calcularPromedioInterno()
}

func (e *Evaluacion) AgregarComentarioPublico(comentario string) bool {

	comentario = strings.TrimSpace(comentario)

	if comentario == "" {
		return false
	}

	e.agregarComentario(comentario)

	return true
}