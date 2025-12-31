package service

import "goc/internal/domain"

func Echo(message string) domain.EchoResponse {
	return domain.EchoResponse{
		Echo: message,
	}
}
