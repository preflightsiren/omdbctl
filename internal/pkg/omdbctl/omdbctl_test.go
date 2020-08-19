package omdbctl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsID(t *testing.T) {
	var tests = []struct {
		input string
		IsID  bool
	}{
		{
			input: "foo",
			IsID:  false,
		},
		{
			input: "tt1608290",
			IsID:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			assert.Equal(t, test.IsID, IsID(test.input))
		})
	}

}
