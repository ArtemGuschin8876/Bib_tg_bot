package handlers

import (
	"log"
	"os"
)

func SaveAdminDataToFile(data string) error {

	file, err := os.OpenFile("admin_data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(data + "\n"); err != nil {
		log.Printf("Error writing to file: %v", err)
		return err
	}

	log.Printf("Admin data saved to file: %s", data)
	return nil
}
