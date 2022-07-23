package readfs

import (
	"bytes"
	"io/ioutil"
	"strconv"
)

func Read(path string) ([]float64, error) {
	rawData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	rows := bytes.Split(rawData, []byte{'\n'})
	data := make([]float64, len(rows))
	for i, row := range rows {
		data[i], _ = strconv.ParseFloat(string(row), 64)
	}

	return data, err
}
