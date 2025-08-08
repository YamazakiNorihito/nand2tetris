package symboltable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolTable(t *testing.T) {
	t.Run("Define should register variables in class scope when kind is STATIC or FIELD", func(t *testing.T) {
		st := New()
		st.Define("x", "int", STATIC)
		st.Define("y", "boolean", FIELD)
		st.Define("z", "char", FIELD)

		assert.Equal(t, STATIC, st.KindOf("x"), "should return STATIC for static variable")
		assert.Equal(t, FIELD, st.KindOf("y"), "should return FIELD for field variable")
		assert.Equal(t, FIELD, st.KindOf("z"), "should return FIELD for field variable")
		assert.Equal(t, "int", st.TypeOf("x"), "should return correct type for static variable")
		assert.Equal(t, "boolean", st.TypeOf("y"), "should return correct type for field variable")
		assert.Equal(t, "char", st.TypeOf("z"), "should return correct type for field variable")
		assert.Equal(t, 0, st.IndexOf("x"), "should return correct index for static variable")
		assert.Equal(t, 0, st.IndexOf("y"), "should return correct index for field variable")
		assert.Equal(t, 1, st.IndexOf("z"), "should return correct index for field variable")
		assert.Equal(t, 1, st.VarCount(STATIC), "should count static variables")
		assert.Equal(t, 2, st.VarCount(FIELD), "should count field variables")
	})

	t.Run("Define should register variables in subroutine scope when kind is ARG or VAR", func(t *testing.T) {
		st := New()
		st.Define("a", "int", ARG)
		st.Define("b", "boolean", VAR)
		st.Define("c", "char", VAR)

		assert.Equal(t, ARG, st.KindOf("a"), "should return ARG for argument variable")
		assert.Equal(t, VAR, st.KindOf("b"), "should return VAR for var variable")
		assert.Equal(t, VAR, st.KindOf("c"), "should return VAR for var variable")
		assert.Equal(t, "int", st.TypeOf("a"), "should return correct type for argument variable")
		assert.Equal(t, "boolean", st.TypeOf("b"), "should return correct type for var variable")
		assert.Equal(t, "char", st.TypeOf("c"), "should return correct type for var variable")
		assert.Equal(t, 0, st.IndexOf("a"), "should return correct index for argument variable")
		assert.Equal(t, 0, st.IndexOf("b"), "should return correct index for var variable")
		assert.Equal(t, 1, st.IndexOf("c"), "should return correct index for var variable")
		assert.Equal(t, 1, st.VarCount(ARG), "should count argument variables")
		assert.Equal(t, 2, st.VarCount(VAR), "should count var variables")
	})

	t.Run("Define should panic when redefining a variable in any scope", func(t *testing.T) {
		st := New()
		st.Define("foo", "int", STATIC)
		assert.Panics(t, func() {
			st.Define("foo", "boolean", FIELD)
		}, "should panic when redefining a variable in class scope")

		st = New()
		st.Define("bar", "int", ARG)
		assert.Panics(t, func() {
			st.Define("bar", "boolean", VAR)
		}, "should panic when redefining a variable in subroutine scope")
	})

	t.Run("Define should panic when redefining a variable with the same name as class-scope variable in subroutine scope", func(t *testing.T) {
		st := New()
		st.Define("x", "int", STATIC)
		st.Reset()
		assert.Panics(t, func() {
			st.Define("x", "String", VAR)
		}, "should panic when redefining variable x in subroutine scope")
	})

	t.Run("Reset should clear subroutine scope variables only", func(t *testing.T) {
		st := New()
		st.Define("a", "int", ARG)
		st.Define("b", "boolean", VAR)
		st.Reset()
		assert.Equal(t, NONE, st.KindOf("a"), "should return NONE after reset for ARG")
		assert.Equal(t, NONE, st.KindOf("b"), "should return NONE after reset for VAR")
		assert.Equal(t, 0, st.VarCount(ARG), "should return 0 after reset for ARG")
		assert.Equal(t, 0, st.VarCount(VAR), "should return 0 after reset for VAR")

		// class scope should remain
		st.Define("c", "int", STATIC)
		assert.Equal(t, STATIC, st.KindOf("c"), "should keep class scope after reset")
	})

	t.Run("KindOf, TypeOf, IndexOf should return NONE, empty, -1 for not found variable", func(t *testing.T) {
		st := New()
		assert.Equal(t, NONE, st.KindOf("notfound"), "should return NONE for unknown variable")
		assert.Equal(t, "", st.TypeOf("notfound"), "should return empty string for unknown variable")
		assert.Equal(t, -1, st.IndexOf("notfound"), "should return -1 for unknown variable")
	})
}
