package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phpdave11/gofpdf"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/report/{filename}", GenerateReport).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GenerateReport(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	filename := vars["filename"]
	report, err := PDFReport()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, _ = response.Write([]byte(err.Error()))
		return
	}

	response.Header().Add("Content-Type", "application/pdf")
	response.Header().Add("Content-Disposition", `attachment; filename="`+filename+`"`)

	_, _ = response.Write(report)
}

func PDFReport() (data []byte, err error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Sample PDF")

	var buff bytes.Buffer
	pdfData := io.Writer(&buff)
	err = pdf.Output(pdfData)
	if err == nil {
		data = buff.Bytes()
	}
	return
}
