package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func ProcessPDF(filePath string) (int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("erro ao abrir o arquivo PDF: %v", err)
	}
	defer f.Close()

	pdfData, err := ioutil.ReadAll(f)
	if err != nil {
		return 0, fmt.Errorf("erro ao ler o arquivo PDF: %v", err)
	}

	r := bytes.NewReader(pdfData)
	ctx, err := api.Read(r, pdfcpu.NewDefaultConfiguration())
	if err != nil {
		return 0, fmt.Errorf("erro ao processar o arquivo PDF: %v", err)
	}

	return ctx.PageCount, nil
}
