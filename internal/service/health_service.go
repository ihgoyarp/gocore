package service

import "goc/internal/domain"

func CheckHealth() string {
	return domain.HealthMessage()
}
