package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)



func TestGetUser(t *testing.T){
    e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")
    assert.Equal(t, 200, rec.Code, "OK response is expected")
}

func TestCreateUser(t *testing.T){
    e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")
    assert.Equal(t, 200, rec.Code, "OK response is expected")
}

func TestLoginUser(t *testing.T){
    e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")
    assert.Equal(t, 200, rec.Code, "OK response is expected")

}
