package gapi

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"log"
)

const (
	grpcClientUserAgent = "grpc-client-user-agent"
	userAgentHeader     = "user-agent"
	xForwardedHeader    = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetaData(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("Metadata: %+v\n", md)

		//for grpc
		if userAgents := md.Get(grpcClientUserAgent); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		//for http
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]

		}

		if clientIPs := md.Get(xForwardedHeader); len(clientIPs) > 0 {
			mtdt.ClientIP = clientIPs[0]
		}
	}
	// getting IP
	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()

	}
	return mtdt
}
