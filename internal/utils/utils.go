package utils

import (
	"github.com/AdamAlberty/resumegen/internal/types"
	"github.com/go-pdf/fpdf"
)

func SectionHeader(pdf *fpdf.Fpdf, text string) {
	width, _ := pdf.GetPageSize()
	left, _, _, _ := pdf.GetMargins()

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFontStyle("B")
	pdf.Text(pdf.GetX(), pdf.GetY(), text)
	pdf.SetFontStyle("")
	pdf.Line(left, pdf.GetY()+2, width-left, pdf.GetY()+2)

	pdf.SetY(pdf.GetY() + 8)
}

func WorkEntry(pdf *fpdf.Fpdf, exp types.WorkExperience) {
	left, _, _, _ := pdf.GetMargins()
	width, _ := pdf.GetPageSize()

	pdf.SetFontStyle("B")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFontSize(11)
	pdf.Text(pdf.GetX(), pdf.GetY(), exp.Company)
	pdf.Text(width-left-pdf.GetStringWidth(exp.Location), pdf.GetY(), exp.Location)
	pdf.SetFontStyle("")

	pdf.SetTextColor(100, 99, 100)
	pdf.SetY(pdf.GetY() + 5)

	fontSize, _ := pdf.GetFontSize()
	pdf.SetFontStyle("I")
	pdf.Text(pdf.GetX(), pdf.GetY(), exp.Role)
	pdf.Text(width-left-pdf.GetStringWidth(exp.TimeFrame), pdf.GetY(), exp.TimeFrame)

	pdf.SetFont("Inter", "", fontSize)
	pdf.SetY(pdf.GetY() + 3)
	for _, bullet := range exp.Bullets {
		pdf.Text(left, pdf.GetY()+3.5, "•")

		pdf.SetX(left + 5)
		pdf.MultiCell(width-left*2-5, 5, bullet, "", "", false)

		pdf.SetY(pdf.GetY() + 1)
	}
}

func EducationEntry(pdf *fpdf.Fpdf, edu types.Education) {
	left, _, _, _ := pdf.GetMargins()
	width, _ := pdf.GetPageSize()

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFontStyle("B")
	pdf.SetFontSize(11)
	pdf.Text(pdf.GetX(), pdf.GetY(), edu.Institution)
	pdf.Text(width-left-pdf.GetStringWidth(edu.Location), pdf.GetY(), edu.Location)
	pdf.SetFontStyle("")

	pdf.SetTextColor(100, 99, 100)
	pdf.SetY(pdf.GetY() + 5)

	pdf.SetFontStyle("I")
	pdf.Text(pdf.GetX(), pdf.GetY(), edu.Degree)
	graduationYear := "Graduation date: " + edu.GraduationYear
	pdf.Text(width-left-pdf.GetStringWidth(graduationYear), pdf.GetY(), graduationYear)
	pdf.SetFontStyle("")

	pdf.SetY(pdf.GetY() + 3)
	for _, bullet := range edu.Bullets {
		pdf.Text(left, pdf.GetY()+3.5, "•")

		pdf.SetX(left + 5)
		pdf.MultiCell(width-left*2-5, 5, bullet, "", "", false)

		pdf.SetY(pdf.GetY() + 1)
	}
}

func ListSection(pdf *fpdf.Fpdf, listName string, items string) {
	left, _, _, _ := pdf.GetMargins()
	width, _ := pdf.GetPageSize()

	pdf.SetFontStyle("B")
	pdf.SetTextColor(0, 0, 0)
	pdf.Text(pdf.GetX(), pdf.GetY()+3.6, listName+": ")
	pdf.SetX(left + pdf.GetStringWidth(listName+": "))
	pdf.SetTextColor(100, 99, 100)
	pdf.SetFontStyle("")
	pdf.MultiCell(width-left*2-pdf.GetStringWidth(listName+": "), 5, items, "", "", false)

	pdf.SetY(pdf.GetY() + 1)
}
