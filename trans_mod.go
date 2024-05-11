package trans_mod

import (
	"encoding/json"
	"os"
	"sort"
)

func Do(src string, dest string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	type patient struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	// читаем данные
	dec := json.NewDecoder(f)
	srcRows := make([]patient, 0, 6)
	for dec.More() {
		var p patient
		err := dec.Decode(&p)
		if err != nil {
			return err
		}
		srcRows = append(srcRows, p)
	}

	//  сортировка данных
	sort.Slice(srcRows, func(i, j int) bool { return srcRows[i].Age < srcRows[j].Age })

	// собираем и пишем данные
	nf, err := os.Create(dest)
	if err != nil {
		return err
	}
	errwrt := json.NewEncoder(nf).Encode(srcRows)
	if errwrt != nil {
		return errwrt
	}
	errclose := nf.Close()
	if errclose != nil {
		return errclose
	}

	return nil
}
