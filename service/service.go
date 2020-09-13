package service

import (
	"geektrust/models"
	"geektrust/utils"
	"strings"
)

type Service interface {
	DecodeMessage(message string) string
	VerifyKingdom(message string) bool
}

type KingdomService struct {
	Name models.Kingdom
	Emblem models.Emblem
}

func GetKingdomService(kingdom models.Kingdom, emblem models.Emblem) *KingdomService {
	return &KingdomService{
		Name:   kingdom,
		Emblem: emblem,
	}
}

func (kingdom *KingdomService) DecodeMessage(message string) string {
	secretKey := len(kingdom.Emblem)

	decodedMessage := utils.RotateAntiClockwise(message, int32(secretKey))

	return decodedMessage
}

func (kingdom *KingdomService) VerifyKingdom(message string) bool {
	runes := []rune(strings.ToLower(string(kingdom.Emblem)))

	count := 0
	for _, rune := range runes {
		if strings.Contains(message, string(rune)) {
			message = strings.Replace(message, string(rune), "", 1)
			count ++
		}
	}

	if count >= len(kingdom.Emblem) {
		return true
	}

	return false
}