package export

import (
	"encoding/json"
	"io/ioutil"
)

//func CsvFile(data []byte) error {
//	outputFile, err := os.Create("output.csv")
//	if err != nil {
//		return err
//	}
//	defer outputFile.Close()
//
//	writer := csv.NewWriter(outputFile)
//	defer writer.Flush()
//
//	header := []string{"accountid", "accountname", "groups"}
//	if err := writer.Write(header); err != nil {
//		return err
//	}
//
//	for _, r := range data {
//		fmt.Println(fmt.Sprint(r))
//	}
//	return nil
//}

func JsonFile(data []byte) {
	file, _ := json.MarshalIndent(data, "", "	")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
