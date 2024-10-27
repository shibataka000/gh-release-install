package external

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendMap(t *testing.T) {
	tests := []struct {
		name string
		m1   map[string]string
		m2   map[string]string
		m3   map[string]string
	}{
		{
			name: "AppendMap",
			m1: map[string]string{
				"a": "1",
				"b": "2",
			},
			m2: map[string]string{
				"a": "3",
				"c": "4",
			},
			m3: map[string]string{
				"a": "3",
				"b": "2",
				"c": "4",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			m3 := appendMap(tt.m1, tt.m2)
			require.Equal(tt.m3, m3)
		})
	}
}
