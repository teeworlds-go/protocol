package require

import (
	"reflect"
	"testing"
)

func New(t *testing.T) *Require {
	return &Require{
		t: t,
	}
}

type Require struct {
	t *testing.T
}

func (r *Require) Equal(expected any, actual any) {
	r.t.Helper()
	Equal(r.t, expected, actual)
}

func (r *Require) NoError(err error) {
	r.t.Helper()
	NoError(r.t, err)
}

func (r *Require) NotNil(a any) {
	r.t.Helper()
	NotNil(r.t, a)
}

func Equal(t *testing.T, expected any, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, got: %v", expected, actual)
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func NotNil(t *testing.T, a any) {
	t.Helper()
	if a != nil {
		return
	}
	t.Fatalf("expected not nil, got: %v", a)
}
