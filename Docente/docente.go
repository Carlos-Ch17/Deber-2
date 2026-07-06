import "DocenteEvaluacion/docente"

import (
	"strings"
	"unicode"
)

type Docente struct {
	ID           string
	Nombre       string
	Email        string
	Departamento string
	Cargo        string
	evaluaciones []string
}

// Constructor público
func NuevoDocente(id, nombre, email, depto, cargo string) *Docente {
	return &Docente{
		ID:           id,
		Nombre:       nombre,
		Email:        email,
		Departamento: depto,
		Cargo:        cargo,
		evaluaciones: []string{},
	}
}

// Métodos públicos

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

// Métodos privados

func (d *Docente) validarEmail() bool {
	email := strings.TrimSpace(d.Email)

	return strings.Contains(email, "@") &&
		strings.Contains(email, ".")
}

func (d *Docente) normalizarNombre() {
	nombre := strings.ToLower(strings.TrimSpace(d.Nombre))

	palabras := strings.Fields(nombre)

	for i, palabra := range palabras {
		r := []rune(palabra)

		if len(r) > 0 {
			r[0] = unicode.ToUpper(r[0])
		}

		palabras[i] = string(r)
	}

	d.Nombre = strings.Join(palabras, " ")
}

func (d *Docente) agregarEvaluacionInterna(idEvaluacion string) {
	d.evaluaciones = append(d.evaluaciones, idEvaluacion)
}
