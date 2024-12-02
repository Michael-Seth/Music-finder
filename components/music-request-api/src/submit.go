package main

import (
	"errors"
	"fmt"
	"music-request-api/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SubmitRequest(writer http.ResponseWriter, request *http.Request, p httprouter.Params, config models.Configuration) (err error) {
	fmt.Println("Submitting request to queue")

	// Get the "query" parameter from the query string
	query := request.URL.Query().Get("query")

	if query == "" {
		http.Error(writer, "Expected 'query' parameter in URL", http.StatusBadRequest)
		return errors.New("missing 'query' parameter")
	}

	fmt.Println("Received: " + query)

	// Parse the query into a message
	message, err := ParseQueryToMessage("/query/" + query)
	if err != nil {
		http.Error(writer, "Failed to parse query", http.StatusBadRequest)
		return err
	}

	// Process the message (e.g., send to RabbitMQ)
	ProduceMessage(message, config.RabbitMq)

	// Return a success response
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Request submitted successfully"))

	return nil
}

func ParseQueryToMessage(queryfull string) (message models.Message, err error) {
	// Ensure query starts with "/query/"
	if len(queryfull) < 7 || queryfull[:7] != "/query/" {
		return models.Message{}, errors.New("invalid query format")
	}

	// Extract the value after "/query/"
	queryVal := queryfull[7:]
	fmt.Println("Parsed query value:", queryVal)

	if queryVal == "" {
		return models.Message{}, errors.New("query value is empty")
	}

	// Return a message object
	return models.Message{SongUri: queryVal}, nil
}
