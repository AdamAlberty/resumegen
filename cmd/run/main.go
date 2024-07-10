package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/AdamAlberty/resumegen/internal/data"
	"github.com/AdamAlberty/resumegen/internal/lib"
	"github.com/AdamAlberty/resumegen/internal/types"
)

func main() {
	var inputPath string
	var outputPath string
	var generate bool

	flag.BoolVar(&generate, "generate", false, "Generates boilerplate JSON file for resumé data")
	flag.StringVar(&inputPath, "input", "data.json", "Specifies resumé data in JSON format (run `generate` to generate boilerplate json file)")
	flag.StringVar(&outputPath, "output", "resume.pdf", "Specifies resumé output path")
	flag.Parse()

	// Generates skeleton JSON data for resumé
	if generate {
		file, err := os.Create("data.json")
		if err != nil {
			fmt.Printf("Could not open `%s`: %s\n", "data.json", err)
			return
		}
		_, err = file.WriteString(data.ResumeSkeletonData)
		if err != nil {
			fmt.Printf("Could not write skeleton resumé data to `%s`: %s\n", "data.json", err)
			return
		}

		fmt.Printf("Skeleton data generated successfully. You can find it in `%s`", "data.json")
		return
	}

	var buf bytes.Buffer

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Could not open `%s`: %s\n", inputPath, err)
		return
	}

	data := new(types.ResumeData)
	err = json.NewDecoder(file).Decode(data)
	if err != nil {
		fmt.Println("Error reading JSON from file:", err)
		return
	}
	lib.GenerateResumePDF(data, &buf)

	// Save the pdf to file
	file, err = os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = buf.WriteTo(file)
	if err != nil {
		fmt.Println("Error writing PDF buffer to file:", err)
		return
	}

	fmt.Printf("Your resumé was successfully generated and saved to `%s`", outputPath)
}
