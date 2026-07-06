import "DocenteEvaluacion/docente"

import (
	"strings"

	"DocenteEvaluacion/evaluacion"
)

type Docente struct {
	ID           string
	Nombre       string
	Email        string
	Departamento string
	Cargo        string

	evaluaciones []evaluacion.Evaluacion
}

// =======================
// Constructor
// =======================

func NuevoDocente(id, nombre, email, depto, cargo string) *Docente {
	return &Docente{
		ID:           id,
		Nombre:       nombre,
		Email:        email,
		Departamento: depto,
		Cargo:        cargo,
		evaluaciones: []evaluacion.Evaluacion{},
	}
}

// =======================
// Métodos públicos
// =======================

func (d *Docente) GetID() string {
	return d.ID
}

func (d *Docente) GetNombre() string {
	return d.Nombre
}

func (d *Docente) GetEmail() string {
	return d.Email
}

func (d *Docente) GetDepartamento() string {
	return d.Departamento
}

func (d *Docente) GetCargo() string {
	return d.Cargo
}

// Agregar evaluación (público)
func (d *Docente) AgregarEvaluacion(eval evaluacion.Evaluacion) {
	if d.validarEvaluacion(eval) {
		d.evaluaciones = append(d.evaluaciones, eval)
	}
}

// Obtener solo evaluaciones finalizadas
func (d *Docente) ObtenerEvaluaciones() []evaluacion.Evaluacion {

	var finalizadas []evaluacion.Evaluacion

	for _, e := range d.evaluaciones {
		if e.EstaFinalizada() {
			finalizadas = append(finalizadas, e)
		}
	}

	return finalizadas
}

// =======================
// Métodos privados
// =======================

// validar email simple
func (d *Docente) validarEmail() bool {
	email := strings.TrimSpace(d.Email)
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// normalizar nombre
func (d *Docente) normalizarNombre() {
	d.Nombre = strings.Title(strings.ToLower(strings.TrimSpace(d.Nombre)))
}

// validar evaluación
func (d *Docente) validarEvaluacion(eval evaluacion.Evaluacion) bool {

	if eval.ID == "" {
		return false
	}

	if eval.DocenteID != d.ID {
		return false
	}

	return true
}

// agregar evaluación interna (opcional según práctica)
func (d *Docente) agregarEvaluacionInterna(id string) {
	// ejemplo simple
	d.evaluaciones = append(d.evaluaciones, evaluacion.Evaluacion{
		ID:        id,
		DocenteID: d.ID,
	})
}