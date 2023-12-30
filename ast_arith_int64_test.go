package z3

import (
	"testing"
)

func TestASTAdd_Int64(t *testing.T) {
	config := NewConfig()
	defer config.Close()

	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int64
	v1 := ctx.Int64(1, ctx.IntSort())
	v2 := ctx.Int64(2, ctx.IntSort())
	v3 := ctx.Int64(3, ctx.IntSort())

	// Add
	raw := v1.Add(v2, v3)

	actual := raw.String()
	if actual != "(+ 1 2 3)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTMul_Int64(t *testing.T) {
	config := NewConfig()
	defer config.Close()

	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int64
	v1 := ctx.Int64(1, ctx.IntSort())
	v2 := ctx.Int64(2, ctx.IntSort())
	v3 := ctx.Int64(3, ctx.IntSort())

	// Mul
	raw := v1.Mul(v2, v3)

	actual := raw.String()
	if actual != "(* 1 2 3)" {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestASTSub_Int64(t *testing.T) {
	config := NewConfig()
	defer config.Close()

	ctx := NewContext(config)
	defer ctx.Close()

	// Create an int64
	v1 := ctx.Int64(1, ctx.IntSort())
	v2 := ctx.Int64(2, ctx.IntSort())
	v3 := ctx.Int64(3, ctx.IntSort())

	// Sub
	raw := v1.Sub(v2, v3)

	actual := raw.String()
	if actual != "(- (- 1 2) 3)" {
		t.Fatalf("bad:\n%s", actual)
	}
}
