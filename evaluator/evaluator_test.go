package evaluator

import (
	"testing"

	"github.com/ahmadjavaidwork/coffee-int/lexer"
	"github.com/ahmadjavaidwork/coffee-int/object"
	"github.com/ahmadjavaidwork/coffee-int/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expechan int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expechan {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expechan)
		return false
	}
	return true
}