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
			Alignment: &excelize.Alignment{
				Vertical: "center",
				WrapText: true,
			},
		},
	)
	if err != nil {
		return &excelize.File{}, fmt.Errorf("failed creating array style | %s", err.Error())
	}

	f.SetColWidth("Sheet1", "A", "A", float64(21))
	f.SetColWidth("Sheet1", "B", "B", float64(18))
	f.SetColWidth("Sheet1", "C", "D", float64(14))
	f.SetColWidth("Sheet1", "E", "E", float64(18))
	f.SetRowHeight("Sheet1", 7, float64(40))

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
	f.SetCellValue("Sheet1", "E7", "RESTITUTION (date & sign. réceptionnaire)")

	if err := DownloadFile("https://www.delubac.com/images/logo_Banque_Delubac_Cie_RVB_70.jpg", "tmp.jpg"); err != nil {
		fmt.Printf("%s\n", fmt.Errorf("notice: failed downloading image | %s", err.Error()))
	} else if err := f.AddPicture("Sheet1", "D3", `tmp.jpg`, `{"lock_aspect_ratio": true, "x_scale":0.75, "y_scale":0.75}`); err != nil {
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

	var a8 string
	var b8 string
	switch borrow.Type {
	case "MOBILE":
		a8 = fmt.Sprintf("Mobile : %s", borrow.Mobile.ModelName)
		b8 = fmt.Sprintf("IMEI : %s", borrow.Mobile.ImeiNumber)
	case "PORTABLE":
		a8 = fmt.Sprintf("Portable : %s", borrow.Laptop.ModelName)
		b8 = fmt.Sprintf("Série : %s", borrow.Laptop.SerialNumber)
	case "CASQUE":
		a8 = fmt.Sprintf("Casque : %s", borrow.Headphones.ModelName)
		b8 = fmt.Sprintf("Série : %s", borrow.Headphones.Serial)
	case "AIRBOX":
		a8 = fmt.Sprintf("Airbox : %s", borrow.Airbox.LineNumber)
		b8 = fmt.Sprintf("IMEI : %s", borrow.Airbox.ImeiNumber)
	case "CHARGEUR":
		a8 = "Chargeur PC portable"
		b8 = fmt.Sprintf(borrow.Charger.ModelName)
	}
	f.SetCellValue("Sheet1", "A8", a8)
	f.SetCellValue("Sheet1", "B8", b8)
	f.SetCellValue("Sheet1", "C8", fmt.Sprintf("Le %02d/%02d/%04d", borrow.Date.Day(), int(borrow.Date.Month()), borrow.Date.Year()))
	return nil
}
