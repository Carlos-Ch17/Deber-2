import "DocenteEvaluacion/docente"

type Evaluable interface {
	GetEstado() string
	EstaFinalizada() bool
	CalcularPromedio() float64
}