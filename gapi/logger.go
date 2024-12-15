package gapi

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

func GrpcLogger(ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {

	startTime := time.Now()

	result, err := handler(ctx, req)

	duration := time.Since(startTime)

	statusCode := codes.Unknown

	if err != nil {
		statusCode = status.Code(err)
	}

	//differ type of log depending on handler result
	logger := log.Info()

	if err != nil {
		logger = logger.Err(err)
	}

	logger.
		Str("protocol", "grpc").        // add the protocol to the log
		Str("method", info.FullMethod). // add the invoked method to the log
		Dur("duration", duration).      // add execution time to the log
		Int("status_code", int(statusCode)).
		Str("status_text", statusCode.String()). // add the status code
		Msg("Received a GRPC request")

	return result, err
}

//in order to track status code of http calls (in logger, MUST override the response writer, then pass it in to the handler.ServeHTTP(w, r))

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (rec *ResponseRecorder) WriteHeader(code int) {
	rec.StatusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *ResponseRecorder) Write(b []byte) (int, error) {
	rec.Body = b
	return rec.ResponseWriter.Write(b)
}

func HTTPLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()

		rec := &ResponseRecorder{ResponseWriter: w, StatusCode: http.StatusOK}

		handler.ServeHTTP(rec, r)

		duration := time.Since(startTime)

		logger := log.Info()

		if rec.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("body", rec.Body)
		}

		logger.
			Str("protocol", "http"). // add the protocol to the log
			Str("method", r.Method). // add the invoked method to the log
			Str("path", r.RequestURI).
			Dur("duration", duration).                           // add execution time to the log
			Str("status_text", http.StatusText(rec.StatusCode)). // add the status code
			Int("status_code", rec.StatusCode).
			Msg("Received a HTTP request")

	})
}
