package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettyPrintHTML(t *testing.T) {
	// Create an HTML node for testing
	node := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{Key: "class", Val: "container"},
			{Key: "id", Val: "myDiv"},
		},
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: "Hello, World!",
		},
	}

	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Call the prettyPrintHTML function
	prettyPrintHTML(node, &buf)

	// Verify the result
	expected := "<div class='container' id='myDiv'>Hello, World!</div>"
	if buf.String() != expected {
		t.Errorf("Unexpected result:\nExpected:\n%s\n\nGot:\n%s", expected, buf.String())
	}
}

func TestOutline(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a sample HTML
		htmlStr := "<html><head></head><body><h1>Title</h1><p>Paragraph</p><img src='image.jpg'/></body></html>"
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(htmlStr))
	}))
	defer server.Close()

	// Call the outline function
	err := outline(server.URL)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestOutline_Error(t *testing.T) {
	// Call the outline function with an invalid URL
	err := outline("http://invalid-url")
	if err == nil {
		t.Fatal("Expected an error, but got nil")
	}
}

type mockResponseWriter struct{}

func (w *mockResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (w *mockResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w *mockResponseWriter) WriteHeader(statusCode int) {}
