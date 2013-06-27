package engine

import (
	"code.google.com/p/mx3/cuda"
	"code.google.com/p/mx3/data"
	"log"
	"reflect"
)

// An excitation, typically field or current,
// can be defined region-wise plus extra mask*multiplier terms.
// This way, arbitrarily complex excitations can be constructed.
type excitation struct {
	v          VectorParam // Region-based excitation
	extraTerms []mulmask   // add extra mask*multiplier terms
	autosave
}

type mulmask struct {
	mul  func() float64
	mask *data.Slice
}

func (e *excitation) init(m *data.Mesh, name, unit string) {
	e.v = vectorParam(name+"_param", unit, nil)
	e.autosave = newAutosave(3, name, unit, m)
}

func (e *excitation) addTo(dst *data.Slice) {
	if !e.v.zero {
		cuda.RegionAddV(dst, e.v.Gpu(), regions.Gpu())
	}
	for _, t := range e.extraTerms {
		cuda.Madd2(dst, dst, t.mask, 1, float32(t.mul()))
	}
}

func (e *excitation) IsZero() bool {
	return e.v.zero && len(e.extraTerms) == 0
}

func (e *excitation) GetGPU() (*data.Slice, bool) {
	buf := cuda.GetBuffer(e.v.NComp(), e.v.Mesh())
	cuda.Zero(buf)
	e.addTo(buf)
	return buf, true
}

func (e *excitation) Get() (q *data.Slice, recycle bool) {
	return e.GetGPU()
}

// Add an extra maks*multiplier term to the excitation.
func (e *excitation) Add(mask *data.Slice, mul func() float64) {
	e.extraTerms = append(e.extraTerms, mulmask{mul, assureGPU(mask)})
}

func (e *excitation) GetVec() []float64 {
	if len(e.extraTerms) != 0 {
		log.Fatal(e.Name(), " is space-dependent, cannot be used as value")
	}
	return e.v.GetVec()
}

func (e *excitation) SetValue(v interface{})  { e.v.SetValue(v) }
func (e *excitation) Eval() interface{}       { return e }
func (e *excitation) Type() reflect.Type      { return reflect.TypeOf(new(excitation)) }
func (e *excitation) InputType() reflect.Type { return reflect.TypeOf([3]float64{}) }