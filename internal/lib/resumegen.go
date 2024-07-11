package lib

import (
	"io"

	"github.com/AdamAlberty/resumegen/internal/types"
	"github.com/AdamAlberty/resumegen/internal/utils"
	"github.com/go-pdf/fpdf"
)

func GenerateResumePDF(data *types.ResumeData, writer io.Writer) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("Inter", "", "assets/Inter-Regular.ttf")
	pdf.AddUTF8Font("Inter", "B", "assets/Inter-SemiBold.ttf")
	pdf.AddUTF8Font("Inter", "I", "assets/Inter-Regular.ttf")

	// Font, margins - Config
	pdf.SetFont("Inter", "", 16)
	pdf.SetMargins(20, 15, 20)
	width, _ := pdf.GetPageSize()
	pdf.AddPage()

	// Metadata
	pdf.SetTitle(data.PersonalInfo.Name+" - Resumé", true)
	pdf.SetAuthor(data.PersonalInfo.Name, true)

	// Name
	pdf.SetFontStyle("B")
	pdf.Text(width/2-pdf.GetStringWidth(data.PersonalInfo.Name)/2, pdf.GetY(), data.PersonalInfo.Name)
	pdf.SetFontStyle("")

	// Personal info
	pdf.SetTextColor(100, 99, 100)
	pdf.SetFontSize(11)
	personalInfo := data.PersonalInfo.Email + " | " + data.PersonalInfo.Phone + " | " + data.PersonalInfo.Location + " | " + data.PersonalInfo.Website + " | "
	pdf.Text((width-pdf.GetStringWidth(personalInfo+"Li | Gh"))/2, pdf.GetY()+8, personalInfo)

	pdf.SetX((width-pdf.GetStringWidth(personalInfo+"Li | Gh"))/2 + pdf.GetStringWidth(personalInfo))

	pdf.CellFormat(pdf.GetStringWidth("Li | "), 13.4, "Li | ", "", 0, "L", false, 0, data.PersonalInfo.LinkedIn)
	pdf.CellFormat(pdf.GetStringWidth("Gh"), 13.4, "Gh", "", 0, "L", false, 0, data.PersonalInfo.GitHub)

	pdf.SetY(pdf.GetY() + 20)
	utils.SectionHeader(pdf, "WORK EXPERIENCE")
	for i, exp := range data.WorkExperience {
		if i != 0 {
			pdf.SetY(pdf.GetY() + 5)
		}
		utils.WorkEntry(pdf, exp)
	}

	pdf.SetY(pdf.GetY() + 10)
	utils.SectionHeader(pdf, "EDUCATION")

	for i, edu := range data.Education {
		if i != 0 {
			pdf.SetY(pdf.GetY() + 8)
		}
		utils.EducationEntry(pdf, edu)
	}

	pdf.SetY(pdf.GetY() + 10)
	utils.SectionHeader(pdf, "SKILLS, CERTIFICATIONS & INTERESTS")
	pdf.SetTextColor(100, 99, 100)
	pdf.SetY(pdf.GetY() - 4)

	utils.ListSection(pdf, "Skills", data.Skills)
	utils.ListSection(pdf, "Certifications", data.Certifications)
	utils.ListSection(pdf, "Interests", data.Interests)
	utils.ListSection(pdf, "Languages", data.Languages)

	err := pdf.Output(writer)
	return err
}
