package main

import (
	"fmt"

	"github.com/Yashrajkanade/bookings/internal/config"
	"github.com/go-chi/chi"

	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed
	default:
		t.Errorf("%s",fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}
