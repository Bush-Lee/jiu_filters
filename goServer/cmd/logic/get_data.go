package logic

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

const CSV_NAME string = "document.csv"
const ALL_FILES_NAME string = "all_files.json"
const COL_DESCRIPTION string = "col_description.json"
const COL_OPTIONS string = "col_options.json"

type Document struct {
	Obiectiv_politica     string
	Titlu                 string
	Prioritate_nr         string
	Prioritate_denumire   string
	Obiectivul_specific   string
	Activitati_finante    interface{}
	Beneficiari_eligibili interface{}
	Beneficiari_finali    interface{}
	Indicatori            interface{}
	Beneficii             interface{}
	Provocari             interface{}
	Link                  string
}

type DocumentList struct {
	Obiectiv_politica     []string
	Titlu                 []string
	Prioritate_nr         []string
	Prioritate_denumire   []string
	Obiectivul_specific   []string
	Activitati_finante    []string
	Beneficiari_eligibili []string
	Beneficiari_finali    []string
	Indicatori            []string
	Beneficii             []string
	Provocari             []string
}

type GeneralDocument interface {
	[]Document | DocumentList | map[string][]string
}

func getCsv(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error geting the csv \n%v", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Response not OK/200 \n%v\n", err)
	}

	byte_data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response \n%v\n", err)
	}

	return byte_data
}

func writeCSV(csv_data []byte, path string) {
	file, err := os.Create(path + CSV_NAME)
	if err != nil {
		log.Fatalf("Can't open/create file \n%s\n", err)
	}
	defer file.Close()

	_, err = file.Write(csv_data)
	if err != nil {
		log.Fatalf("Can't write to csv file \n%v\n", err)
	}
}

func writeJSON(path string, filename string, data interface{}) {
	file, err := os.Create(path + filename)
	if err != nil {
		log.Fatalf("Can't write to file \n%s\n", err)
	}
	defer file.Close()

	if reflect.TypeOf(data).Kind() == reflect.Map || (reflect.TypeOf(data).Kind() == reflect.Slice && reflect.TypeOf(data).Elem().Kind() == reflect.Struct) {
		dataMarshaled, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			log.Fatalf("Can't marshal data \n%s\n", err)
		}
		file.Write(dataMarshaled)
	} else {
		log.Fatal("Invalid data type. It should be either map[string][]string or []Document.")
	}
}

func convertToCorrectType(posibleList string) interface{} {
	posibleList = strings.Replace(posibleList, "\n", " ", -1)
	posibleList = strings.Replace(posibleList, "\r", " ", -1)
	if strings.Contains(posibleList, ";") {
		return strings.Split(posibleList, ";")
	}
	return posibleList
}

func genDocuments(csvData [][]string) []Document {
	documente := make([]Document, len(csvData[0]))
	for index_col := 0; index_col < len(csvData[0]); index_col++ {
		for index, row := range csvData {
			switch index {
			case 0:
				row[index_col] = strings.Replace(row[index_col], "\n", " ", -1)
				row[index_col] = strings.Replace(row[index_col], "\r", " ", -1)

				documente[index_col].Obiectiv_politica = row[index_col]
			case 1:
				row[index_col] = strings.Replace(row[index_col], "\n", " ", -1)
				row[index_col] = strings.Replace(row[index_col], "\r", " ", -1)

				documente[index_col].Titlu = row[index_col]
			case 2:
				row[index_col] = strings.Replace(row[index_col], "\n", " ", -1)
				row[index_col] = strings.Replace(row[index_col], "\r", " ", -1)

				documente[index_col].Prioritate_nr = row[index_col]
			case 3:
				row[index_col] = strings.Replace(row[index_col], "\n", " ", -1)
				row[index_col] = strings.Replace(row[index_col], "\r", " ", -1)

				documente[index_col].Prioritate_denumire = row[index_col]
			case 4:
				row[index_col] = strings.Replace(row[index_col], "\n", " ", -1)
				row[index_col] = strings.Replace(row[index_col], "\r", " ", -1)

				documente[index_col].Obiectivul_specific = row[index_col]
			case 5:
				documente[index_col].Activitati_finante = convertToCorrectType(row[index_col])
			case 6:
				documente[index_col].Beneficiari_eligibili = convertToCorrectType(row[index_col])
			case 7:
				documente[index_col].Beneficiari_finali = convertToCorrectType(row[index_col])
			case 8:
				documente[index_col].Indicatori = convertToCorrectType(row[index_col])
			case 9:
				documente[index_col].Beneficii = convertToCorrectType(row[index_col])
			case 10:
				documente[index_col].Provocari = convertToCorrectType(row[index_col])
			case 11:
				row[index_col] = strings.Replace(row[index_col], "\n", "", -1)
				row[index_col] = strings.Replace(row[index_col], "\r", "", -1)

				documente[index_col].Link = row[index_col]
			}
		}
	}
	return documente
}

func appendIfUniqueVal(element string, mapList map[string][]string, key string) {
	if element == "" {
		return
	}

	for _, el := range mapList[key] {
		if el == element {
			return
		}
	}
	mapList[key] = append(mapList[key], element)
}

func genColOptions(full_data []Document) map[string][]string {
	unique_values := make(map[string][]string)
	for _, document := range full_data {
		appendIfUniqueVal(document.Obiectiv_politica, unique_values, "Obiectiv_politica")
		appendIfUniqueVal(document.Titlu, unique_values, "Titlu")
		appendIfUniqueVal(document.Prioritate_nr, unique_values, "Prioritate_nr")
		appendIfUniqueVal(document.Prioritate_denumire, unique_values, "Prioritate_denumire")
		appendIfUniqueVal(document.Obiectivul_specific, unique_values, "Obiectivul_specific")

		// TODO: Make it work for interface types
		switch v := document.Activitati_finante.(type) {
		case string:
			appendIfUniqueVal(v, unique_values, "Activitati_finante")
		case []string:
			for _, value := range v {
				appendIfUniqueVal(value, unique_values, "Activitati_finante")
			}
		}

		switch v := document.Beneficiari_eligibili.(type) {
		case string:
			appendIfUniqueVal(v, unique_values, "Beneficiari_eligibili")
		case []string:
			for _, value := range v {
				appendIfUniqueVal(value, unique_values, "Beneficiari_eligibili")
			}
		}

		switch v := document.Beneficiari_finali.(type) {
		case string:
			appendIfUniqueVal(v, unique_values, "Beneficiari_finali")
		case []string:
			for _, value := range v {
				appendIfUniqueVal(value, unique_values, "Beneficiari_finali")
			}
		}

		switch v := document.Indicatori.(type) {
		case string:
			appendIfUniqueVal(v, unique_values, "Indicatori")
		case []string:
			for _, value := range v {
				appendIfUniqueVal(value, unique_values, "Indicatori")
			}
		}

		switch v := document.Beneficii.(type) {
		case string:
			appendIfUniqueVal(v, unique_values, "Beneficii")
		case []string:
			for _, value := range v {
				appendIfUniqueVal(value, unique_values, "Beneficii")
			}
		}

		switch v := document.Provocari.(type) {
		case string:
			appendIfUniqueVal(v, unique_values, "Provocari")
		case []string:
			for _, value := range v {
				appendIfUniqueVal(value, unique_values, "Provocari")
			}
		}

	}
	return unique_values
}

func GenStaticFolder(url string, path string) {
	fmt.Println("Getting data ...")

	csv_data := getCsv(url)
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		fmt.Println("Folder allready exists, writing into folder ...")
	}

	writeCSV(csv_data, path)
	file, err := os.Open(path + CSV_NAME)
	if err != nil {
		log.Fatalf("Can't open csv file\n %s", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Can't read csv file\n %s", err)
	}

	documente := genDocuments(csvData)
	writeJSON(path, COL_DESCRIPTION, documente[:1])
	writeJSON(path, ALL_FILES_NAME, documente[2:])

	colOptions := genColOptions(documente[2:])
	writeJSON(path, COL_OPTIONS, colOptions)
}
