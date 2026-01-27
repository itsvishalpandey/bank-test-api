package controller

import (
	config "bank-test-api/Config"
	model "bank-test-api/Models"
	"database/sql"
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *sql.DB

func ReadCSV() ([]model.BankMaster, error) {

	// Read CSV File
	file, err := os.Open("ifsc.csv")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	headerIndex := make(map[string]int)
	for i, h := range headers {
		headerIndex[strings.ToLower(strings.TrimSpace(h))] = i
	}

	var result []model.BankMaster
	for {
		records, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if len(records) < len(headers) {
			continue
		}

		data := model.BankMaster{
			Bank:   records[headerIndex["bank"]],
			Ifsc:   records[headerIndex["ifsc"]],
			Micr:   records[headerIndex["micr"]],
			Branch: records[headerIndex["branch"]],
		}

		result = append(result, data)
	}

	return result, nil
}

func GetIfscCode(c *gin.Context) {
	data, err := ReadCSV()
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Status": "Success",
		"Data": data})
}

func PostIFSC(c *gin.Context) {
	data, err := ReadCSV()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(data, 100).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"Status": "Data has been stored successfully.",
		"rows":   len(data),
	})
}
