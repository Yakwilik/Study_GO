package actions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum(t *testing.T) {
	var input string = "1+2+3"
	var output float64 = 6
	resExpression, _ := Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestSub(t *testing.T) {
	var input string = "10-10-1001"
	var output float64 = -1001
	resExpression, _ := Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestCombinationSubSum(t *testing.T) {
	var input string = "2+2-4-11+20"
	var output float64 = 9
	resExpression, _ := Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "1+1+1-1-1-1+2+2-4"
	output = 0
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "0+1*2"
	output = 2
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestMultiply(t *testing.T) {
	var input string = "6*3*5"
	var output float64 = 90
	resExpression, _ := Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "6*3*5+4"
	output = 94
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}
func TestCombinationSubSumMul(t *testing.T) {
	var input string = "2+2*2"
	var output float64 = 6
	resExpression, _ := Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2+2*2-2-2*2+2-2"
	output = 0
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestDivide(t *testing.T) {
	var input string = "300/25/4"
	var output float64 = 3
	resExpression, _ := Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "300/25/4+41"
	output = 44
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "300/25/4+41"
	output = 44
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "100-300/25/4-41"
	output = 56
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "300/25*5"
	output = 60
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "10/2*4/5"
	output = 4
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestCombination(t *testing.T) {
	var input string = "2+2*10"
	var output float64 = 22
	resExpression, _ := Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2+2*2/2"
	output = 4
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2/2-2"
	output = -1
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2+2*2/2+2+2"
	output = 8
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2/2+2"
	output = 3
	resExpression, _ = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestBadInput(t *testing.T) {
	var input string = "123+21323dsagfa"
	_, error_ := Calculate(input)
	require.NotNil(t, error_, "this is error", input)

	input = "123/21323dsagfa"
	_, error_ = Calculate(input)
	require.NotNil(t, error_, "this is error", input)

	input = "123*21323dsagfa"
	_, error_ = Calculate(input)
	require.NotNil(t, error_, "this is error", input)

	input = "1+123-21323dsagfa"
	_, error_ = Calculate(input)
	require.NotNil(t, error_, "this is error", input)

	input = "1-123*21323dsagfa"
	_, error_ = Calculate(input)
	require.NotNil(t, error_, "this is error", input)

	input = "1*123*21323dsagfa"
	_, error_ = Calculate(input)
	require.NotNil(t, error_, "this is error", input)

	input = "1/123-21323dsagfa"
	_, error_ = Calculate(input)
	require.NotNil(t, error_, "this is error", input)
}
