package handler

import (
	"fmt"
	"testing"

	"gofr.dev/pkg/gofr"

)

func TestDashboard(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, nil)

	resp, err := Dashboard(ctx)
	if err != nil {
		t.Errorf("FAILED, got error: %v", err)
	}

	expected := "Hello students! Welcome to the placement dashboard"

	got := fmt.Sprintf("%v", resp)
	if got != expected {
		t.Errorf("FAILED, Expected: %v, Got: %v", expected, got)
	}
}