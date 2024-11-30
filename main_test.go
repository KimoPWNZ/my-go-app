package main  

import (  
	"net/http"  
	"net/http/httptest"  
	"testing"  
)  

func TestHandler(t *testing.T) {  
	req := httptest.NewRequest("GET", "/", nil)  
	w := httptest.NewRecorder()  
	handler(w, req)  

	res := w.Result()  
	if res.StatusCode != http.StatusOK {  
		t.Errorf("Expected status OK; got %v", res.Status)  
	}  
}
