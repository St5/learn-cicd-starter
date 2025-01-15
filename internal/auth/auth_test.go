package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	got, err := GetAPIKey(http.Header{
		"Authorization": []string{"ApiKey MyApiKey"},
	})
	if err != nil {
		t.Errorf("GetAPIKey() error = %v", err)
		return
	}
	want := "MyApiKey1"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAPIKey() = %v, want %v", got, want)
	}
}
