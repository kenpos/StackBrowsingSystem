package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Static("/", "dist/")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)
	e.GET("/getEdinetAPI", getEdinetAPI)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// GETリクエスト
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello, world")
}

func getEdinetAPI(c echo.Context) error {
	fmt.Println("getEdinetAPI")
	url := "https://disclosure.edinet-fsa.go.jp/api/v1/documents.json"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	params := request.URL.Query()
	params.Add("date", "2019-04-25")
	params.Add("type", "2")
	request.URL.RawQuery = params.Encode()

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var edinet EdinetApiStruct
	if err := json.Unmarshal(body, &edinet); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, edinet)
}

type EdinetApiStruct struct {
	Metadata struct {
		Title     string `json:"title"`
		Parameter struct {
			Date string `json:"date"`
			Type string `json:"type"`
		} `json:"parameter"`
		Resultset struct {
			Count int `json:"count"`
		} `json:"resultset"`
		ProcessDateTime string `json:"processDateTime"`
		Status          string `json:"status"`
		Message         string `json:"message"`
	} `json:"metadata"`
	Results []struct {
		SeqNumber            int         `json:"seqNumber"`
		DocID                string      `json:"docID"`
		EdinetCode           string      `json:"edinetCode"`
		SecCode              interface{} `json:"secCode"`
		JCN                  string      `json:"JCN"`
		FilerName            string      `json:"filerName"`
		FundCode             string      `json:"fundCode"`
		OrdinanceCode        string      `json:"ordinanceCode"`
		FormCode             string      `json:"formCode"`
		DocTypeCode          string      `json:"docTypeCode"`
		PeriodStart          string      `json:"periodStart"`
		PeriodEnd            string      `json:"periodEnd"`
		SubmitDateTime       string      `json:"submitDateTime"`
		DocDescription       string      `json:"docDescription"`
		IssuerEdinetCode     interface{} `json:"issuerEdinetCode"`
		SubjectEdinetCode    interface{} `json:"subjectEdinetCode"`
		SubsidiaryEdinetCode interface{} `json:"subsidiaryEdinetCode"`
		CurrentReportReason  interface{} `json:"currentReportReason"`
		ParentDocID          interface{} `json:"parentDocID"`
		OpeDateTime          interface{} `json:"opeDateTime"`
		WithdrawalStatus     string      `json:"withdrawalStatus"`
		DocInfoEditStatus    string      `json:"docInfoEditStatus"`
		DisclosureStatus     string      `json:"disclosureStatus"`
		XbrlFlag             string      `json:"xbrlFlag"`
		PdfFlag              string      `json:"pdfFlag"`
		AttachDocFlag        string      `json:"attachDocFlag"`
		EnglishDocFlag       string      `json:"englishDocFlag"`
	} `json:"results"`
}
