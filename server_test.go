package simpleWeb

import (
	"testing"
)

func TestHTTPServer_Start(t *testing.T) {
	s := NewHttpServer()
	err := s.Start(":8081")
	if err != nil {
		panic(err)
	}
}
