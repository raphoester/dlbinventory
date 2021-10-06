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
	count := len(borrow.Airboxes) + len(borrow.Chargers) + len(borrow.Headphones) + len(borrow.Laptops) + len(borrow.Mobiles)
	if count != 1 {
		return entity.Borrow{}, fmt.Errorf("too many equipments are mentionned")
	}
	if borrow.Date.Before(time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)) {
		return entity.Borrow{}, fmt.Errorf("timestamp is too old (%s)", borrow.Date)
	}

	if len(borrow.Airboxes) == 1 {
		borrow.Type = "AIRBOX"
	} else if len(borrow.Chargers) == 1 {
		borrow.Type = "CHARGEUR"
	} else if len(borrow.Headphones) == 1 {
		borrow.Type = "CASQUE"
	} else if len(borrow.Laptops) == 1 {
		borrow.Type = "PORTABLE"
	} else if len(borrow.Mobiles) == 1 {
		borrow.Type = "MOBILE"
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
