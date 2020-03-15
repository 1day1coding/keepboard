package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	expected := []string{"a", "b", "c"}

	e := echo.New()
	h := New(func() ([]string, error) {
		return expected, nil
	})

	t.Run("/ 헬스체크 테스트", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		assert.NoError(t, h.HealthCheck(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	})

	t.Run("/comments 반환 테스트", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/comments", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		assert.NoError(t, h.GetComments(c))
		assert.Equal(t, http.StatusOK, rec.Code)

		var actual []string
		err := json.Unmarshal(rec.Body.Bytes(), &actual)
		assert.NoError(t, err)
		assert.Equal(t, actual, expected)
	})
}
