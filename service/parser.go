package service

import (
	"encoding/json"
	"fmt"
	"inventory/entity"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func ParseJsonFile(jsonPath string) (entity.Borrow, error) {
	reader, err := os.Open(jsonPath)
	if err != nil {
		return entity.Borrow{}, fmt.Errorf("failed opening file %s : %s", jsonPath, err.Error())
	}
	defer reader.Close()
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return entity.Borrow{}, fmt.Errorf("failed reading file content | %s", err.Error())
	}
	var borrow entity.Borrow
	if err := json.Unmarshal(content, &borrow); err != nil {
		return entity.Borrow{}, fmt.Errorf("failed unmarshalling input | %s", err.Error())
	}
	return borrow, nil
}

func IsBorrowValid(borrow entity.Borrow) (entity.Borrow, error) {
	if borrow.Date.Before(time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)) {
		borrow.Date = time.Now()
	}

	if borrow.Airbox.ImeiNumber != "" {
		borrow.Type = "AIRBOX"
	} else if borrow.Charger.ModelName != "" {
		borrow.Type = "CHARGEUR"
	} else if borrow.Headphones.Serial != "" {
		borrow.Type = "CASQUE"
	} else if borrow.Laptop.SerialNumber != "" {
		borrow.Type = "PORTABLE"
	} else if borrow.Mobile.ImeiNumber != "" {
		borrow.Type = "MOBILE"
	} else {
		borrow.Type = "CUSTOM"
	}
	return borrow, nil
}

func GenerateFileName(borrow entity.Borrow) string {
	fileName := fmt.Sprintf(
		"ATTRIBUTION_%s %s_%s_%s_%04d%02d%02d",
		strings.ToUpper(borrow.Borrower.Lastname),
		borrow.Borrower.Firstname,
		borrow.Type,
		borrow.Borrower.Department,
		borrow.Date.Year(),
		int(borrow.Date.Month()),
		borrow.Date.Day(),
	)
	return fileName
}
