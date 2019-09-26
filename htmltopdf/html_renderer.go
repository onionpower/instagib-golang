package htmltopdf

import (
	"fmt"
	p "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
)

func GenPdf(htmlUri string) []byte {
	pdfg, err := p.NewPDFGenerator()
	//Create new PDF generator
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(30)
	pdfg.Orientation.Set(p.OrientationLandscape)
	//pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := p.NewPage(htmlUri)

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	pdfg.AddPage(page)

	// Add to document

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	buf := pdfg.Buffer()
	b := make([]byte, buf.Len())
	_, err = buf.Read(b)
	if err != nil {
		log.Println(err)
	}

	//err = pdfg.WriteFile("./simplesample.pdf")
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println("Done")
	// Output: Done
	return b
}
