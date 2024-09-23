package swagger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComparation(t *testing.T) {
	result := compareJSON(
		[]byte(`{"a": 1, "b": "cc", "c": true, "d": null, "e": ["1", "2", "3"], "f": {"a": 1, "b": "2"}}`),
		[]byte(`{"a": 1, "b": "cc", "c": true, "d": null, "e": ["1", "2", "3"], "f": {"a": 1, "b": "2"}}`),
	)
	require.Equal(t, 0, len(result), result)
	result = compareJSON(
		[]byte(`{"a": [1,2]}`),
		[]byte(`{"a": [1]}`),
	)
	require.Equal(t, 1, len(result), result)
	result = compareJSON(
		[]byte(`{"a": 1}`),
		[]byte(`{"a": [1]}`),
	)
	require.Equal(t, 1, len(result), result)
	result = compareJSON(
		[]byte(`{"a": 1, "b": 2}`),
		[]byte(`{"a": 1, "c": 3}`),
	)
	require.Equal(t, 2, len(result), result)
}
