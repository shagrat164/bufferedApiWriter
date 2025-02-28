package buffer

import (
	"fmt"
	"testing"

	"github.com/shagrat164/bufferedApiWriter/internal/model"
	"github.com/stretchr/testify/require"
)

func TestBuffer(t *testing.T) {
	t.Run("empty buffer", func(t *testing.T) {
		buf := NewBuffer()
		_, hasData := buf.Get()

		require.Equal(t, false, hasData)
	})

	t.Run("record to buffer", func(t *testing.T) {
		buf := NewBuffer()

		for i := 0; i < 2; i++ {
			buf.Put(model.Fact{
				Value: fmt.Sprintf("%d", i),
			})
		}

		_, hasData := buf.Get()
		require.Equal(t, true, hasData)

		_, hasData = buf.Get()
		require.Equal(t, true, hasData)

		_, hasData = buf.Get()
		require.Equal(t, false, hasData)
	})
}
