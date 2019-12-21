package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInstruction(t *testing.T) {
	ins := newInstruction(21101)
	require.Equal(t, 1, ins.paramOneMode)
	require.Equal(t, 1, ins.paramTwoMode)
	require.Equal(t, 2, ins.paramThreeMode)
}
