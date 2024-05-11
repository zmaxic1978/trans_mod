package trans_mod

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"sort"
)

func Do(src string, dest string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	type Patient struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	// читаем данные
	dec := json.NewDecoder(f)
	srcRows := make([]Patient, 0, 6)
	for dec.More() {
		var p Patient
		err := dec.Decode(&p)
		if err != nil {
			return err
		}
		srcRows = append(srcRows, p)
	}

	//  сортировка данных
	sort.Slice(srcRows, func(i, j int) bool { return srcRows[i].Age < srcRows[j].Age })

	type patients struct{ Patient []Patient }
	pats := patients{Patient: srcRows}

	// собираем и пишем данные
	nf, err := os.Create(dest)
	if err != nil {
		return err
	}
	nf.WriteString(xml.Header)
	enc := xml.NewEncoder(nf)
	enc.Indent("", "    ")
	errwrt := enc.Encode(pats)
	if errwrt != nil {
		return errwrt
	}
	errclose := nf.Close()
	if errclose != nil {
		return errclose
	}

	return nil
}
