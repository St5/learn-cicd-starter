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
	want := "MyApiKey"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAPIKey() = %v, want %v", got, want)
	}
}

func TestGetAPIKeyNoAuthHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{
	})
	if err == nil {
		t.Errorf("Expected error = %v", err)
		return
	}
}

func TestGetAPIKeyNoApiPrefix(t *testing.T) {
	_, err := GetAPIKey(http.Header{
		"Authorization": []string{"MyApiKey"},
	})
	if err == nil {
		t.Errorf("Expected error = %v", err)
		return
	}
}


