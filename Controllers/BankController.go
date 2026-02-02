package controller

import (
	config "bank-test-api/Config"
	model "bank-test-api/Models"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	requiredHeaders := []string{"ifsc", "bank", "micr", "branch"}
	for _, rh := range requiredHeaders {
		if _, ok := headerIndex[rh]; !ok {
			return nil, fmt.Errorf("The file doesn't have column: %s", rh)
		}
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

		// Gives error if any IFSC column has no value
		if strings.TrimSpace(records[headerIndex["ifsc"]]) == "" {
			return nil, fmt.Errorf("IFSC cannot be empty")
		}

		data := model.BankMaster{
			Bank:     records[headerIndex["bank"]],
			Ifsc:     records[headerIndex["ifsc"]],
			Micr:     records[headerIndex["micr"]],
			Branch:   records[headerIndex["branch"]],
			Address:  records[headerIndex["address"]],
			Contact:  records[headerIndex["contact"]],
			City:     records[headerIndex["city"]],
			District: records[headerIndex["district"]],
			State:    records[headerIndex["state"]],
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
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	// To insert the entires batch wise in DB
	// err = config.DB.Transaction(func(tx *gorm.DB) error {
	// 	if err := tx.CreateInBatches(data, 100).Error; err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })

	// To insert the entires batchwise in DB and to also update the already present IFSC row
	err = config.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(
			clause.OnConflict{Columns: []clause.Column{{Name: "ifsc"}},
				DoUpdates: clause.AssignmentColumns([]string{"bank", "micr"}),
			}).CreateInBatches(data, 100).Error
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

func DeleteAllEntry(c *gin.Context) {

	result := config.DB.Exec("DELETE FROM bank_masters")

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{
		"Status": "Data has been deleted successfully.",
		"rows":   result.RowsAffected,
	})

}
