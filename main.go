package main

import (
	"bufio"
	"fmt"
	"geektrust/models"
	"geektrust/service"
	"geektrust/utils"
	"log"
	"os"
	"strings"
)

func main() {

	kingdomMap := utils.GetKingdomMap()

	args := os.Args[1:]

	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := make([]string, 0)
	output := make([]string, 0)
	output = append(output, string(models.KingdomSpace))
	count := 0
	inputMap := make(map[models.Kingdom]int)
	for scanner.Scan() {
		input = strings.Split(strings.ToUpper(scanner.Text()), " ")
		message := ""
		for index, v := range input {
			if index == 0 { //message will be after first whitespace
				continue
			}

			message += v
		}

		if _, ok := inputMap[models.Kingdom(input[0])]; ok { //ignore duplicate kingdom
			continue
		} else {
			inputMap[models.Kingdom(input[0])] = 0
		}

		if _, ok := kingdomMap[models.Kingdom(input[0])]; ok {

			//create kingdom service
			kingdomService := service.GetKingdomService(models.Kingdom(input[0]),
				kingdomMap[models.Kingdom(input[0])])

			decodedMessage := kingdomService.DecodeMessage(strings.ToLower(message))

			isRule := kingdomService.VerifyKingdom(decodedMessage)
			if isRule {
				output = append(output, string(kingdomService.Name))
				count++
			}
		}
	}

	if count >= 3 {
		fmt.Println(strings.Join(output, " "))
	} else {
		fmt.Println("NONE")
	}
}
