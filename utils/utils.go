package utils

import (
	"geektrust/models"
	"strings"
)

func GetKingdomMap() map[models.Kingdom]models.Emblem {

	kingdomMap := make(map[models.Kingdom]models.Emblem)

	kingdomMap[models.KingdomSpace] = models.EmblemSpace
	kingdomMap[models.KingdomAir] = models.EmblemAir
	kingdomMap[models.KingdomFire] = models.EmblemFire
	kingdomMap[models.KingdomIce] = models.EmblemIce
	kingdomMap[models.KingdomLand] = models.EmblemLand
	kingdomMap[models.KingdomWater] = models.EmblemWater

	return kingdomMap
}

func RotateAntiClockwise(val string, d int32) string {

	result := ""
	runes := []rune(strings.ToLower(val))
	d = d % 26
	for _, rune := range runes {
		//duplicate := d
		if (rune - d) < 97 {
			result = result + string(122 - d - (97 - rune) + 1)
			//d = d % 26
		} else {
			result = result + string(rune-d)
		}
		//d = duplicate
	}

	return strings.ToLower(result)
}




