package service

import (
	"fmt"
	_ "image/jpeg"
	"inventory/entity"
	"os"

	"github.com/xuri/excelize/v2"
)

func GenerateExcelTemplate(name string) (*excelize.File, error) {
	f := excelize.NewFile()
	titleStyle, err := f.NewStyle(`{"font":{"bold":true, "underline":"single", "size":16}}`)
	if err != nil {
		return &excelize.File{}, fmt.Errorf("failed creating sheet title style | %s", err.Error())
	}

	coordinatesStyle, err := f.NewStyle(`{"font":{"size":14}}`)
	if err != nil {
		return &excelize.File{}, fmt.Errorf("failed creating coordinates style | %s", err.Error())
	}

	arrayStyle, err := f.NewStyle(
		&excelize.Style{
			Border: []excelize.Border{
				{
					Type:  "top",
					Color: "#000000",
					Style: 2,
				},
				{
					Type:  "bottom",
					Color: "#000000",
					Style: 2,
				},
				{
					Type:  "left",
					Color: "#000000",
					Style: 2,
				},
				{
					Type:  "right",
					Color: "#000000",
					Style: 2,
				},
			},
		},
	)
	if err != nil {
		return &excelize.File{}, fmt.Errorf("failed creating array style | %s", err.Error())
	}

	f.SetColWidth("Sheet1", "A", "E", float64(21))
	f.SetRowHeight("Sheet1", 8, float64(65))

	f.SetCellStyle("Sheet1", "A1", "A1", titleStyle)
	f.SetCellStyle("Sheet1", "A3", "A5", coordinatesStyle)

	if err := f.SetCellStyle("Sheet1", "A7", "E8", arrayStyle); err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet1", "A1", "FICHE DE REMISE D'ÉQUIPEMENTS")
	f.SetCellValue("Sheet1", "A3", "NOM")
	f.SetCellValue("Sheet1", "A4", "PRENOM")
	f.SetCellValue("Sheet1", "A5", "DEPARTEMENT")

	f.SetCellValue("Sheet1", "A7", "EQUIPEMENT")
	f.SetCellValue("Sheet1", "B7", "TYPE et N°")
	f.SetCellValue("Sheet1", "C7", "DATE")
	f.SetCellValue("Sheet1", "D7", "SIGNATURE")
	f.SetCellValue("Sheet1", "E7", "RESTITUTION")

	if err := DownloadFile("https://www.delubac.com/images/logo_Banque_Delubac_Cie_RVB_70.jpg", "tmp.jpg"); err != nil {
		return &excelize.File{}, fmt.Errorf("failed downloading image | %s", err.Error())
	}

	if err := f.AddPicture("Sheet1", "D3", `tmp.jpg`, `{"lock_aspect_ratio": true, "x_scale":0.75, "y_scale":0.75}`); err != nil {
		return &excelize.File{}, fmt.Errorf("failed adding image to excel sheet | %s", err.Error())
	}

	if err := os.Remove("tmp.jpg"); err != nil {
		return &excelize.File{}, fmt.Errorf("failed deleting tmp image | %s", err.Error())
	}
	return f, nil
}

func FillExcelTemplate(f *excelize.File, borrow entity.Borrow) error {
	f.SetCellValue("Sheet1", "B3", borrow.Borrower.Lastname)
	f.SetCellValue("Sheet1", "B4", borrow.Borrower.Firstname)
	f.SetCellValue("Sheet1", "B5", borrow.Borrower.Department)

	return nil
}
