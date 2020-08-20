package kabusapi_test

import (
	"encoding/json"
	"testing"

	"github.com/hori-ryota/go-kabusapi/kabusapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDate(t *testing.T) {
	t.Run("MarshalJSON", func(t *testing.T) {
		date := kabusapi.NewDate(2020, 8, 21)
		j := map[string]interface{}{
			"Date": date,
		}
		b, err := json.Marshal(j)
		require.NoError(t, err)
		assert.JSONEq(t, `{"Date":20200821}`, string(b))
	})
	t.Run("UnmarshalJSON", func(t *testing.T) {
		j := struct {
			Date kabusapi.Date
		}{}
		err := json.Unmarshal([]byte(`{"Date":20200821}`), &j)
		require.NoError(t, err)
		assert.Equal(t, kabusapi.NewDate(2020, 8, 21), j.Date)
	})
}
