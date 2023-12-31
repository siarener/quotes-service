// server/server_test.go
package server

import (
	"context"
	"testing"

	pb "github.com/siarener/quotes-service/protos/quotespb"
	"github.com/stretchr/testify/assert"
)

func TestGetQuote(t *testing.T) {
	// Create an instance of your server
	s := &quoteServer{} // Make sure to adjust this line based on your actual server initialization

	// Create a sample QuoteRequest
	req := &pb.QuoteRequest{
		// Set any relevant fields in your QuoteRequest
	}

	// Call the GetQuote method
	res, err := s.GetQuote(context.Background(), req)

	// Check for errors
	if err != nil {
		t.Fatalf("Error calling GetQuote: %v", err)
	}

	// Verify the result
	expectedQuote := "The only true wisdom is in knowing you know nothing. - Socrates"
	if res.Quote != expectedQuote {
		t.Errorf("Expected quote %q, but got %q", expectedQuote, res.Quote)
	}
}

func TestStoreQuote(t *testing.T) {

	mockService := &quoteServer{}
	storeQuoteReq := &pb.StoreQuoteRequest{Quote: "Testing is better than believing."}

	emptyResponse, err := mockService.StoreQuote(context.Background(), storeQuoteReq)

	assert.NoError(t, err, "Expected no error while storing a quote")
	assert.NotNil(t, emptyResponse, "Expected a non-nil response for StoreQuote")
}
