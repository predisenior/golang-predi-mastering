package book

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)


func TestGetBook(t *testing.T){
    e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/admin/book")
    assert.Equal(t, 200, rec.Code, "OK response is expected")
}


var (

	bookJSON = `{"title":"apa saja","AUthor":"siapa saja","publisher":"tukang kopi"}`
)
func TestCreateBook(t *testing.T){
 e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	err := e.NewContext(req, rec)
	// h := &handler{bookJSON}

	// Assertions
	if err!=nil {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, bookJSON, rec.Body.String())
	}

    response := httptest.NewRecorder()
    assert.Equal(t, 201, response.Code, "OK response is expected")
}


func TestDetailBook(t *testing.T){
    e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/admin/book")
    assert.Equal(t, 200, rec.Code, "OK response is expected")
}
