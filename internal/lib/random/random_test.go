package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			name:   "size = 1",
			length: 1,
		},
		{
			name:   "size = 5",
			length: 1,
		},
		{
			name:   "size = 10",
			length: 1,
		},
		{
			name:   "size = 20",
			length: 1,
		},
		{
			name:   "size = 30",
			length: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := GenerateString(tt.length)
			str2 := GenerateString(tt.length)

			assert.Len(t, str1, tt.length)
			assert.Len(t, str2, tt.length)

			assert.NotEqual(t, str1, str2)
		})
	}
}
