package main

import (
	"fmt"
)

func put_raw(data string) {
}

func put_image(img []byte) {

}

func generate_doc(params map[string]string) string {
	// Пример обработки — просто объединяем ключи и значения
	output := "Processed Params:\n"
	for k, v := range params {
		output += fmt.Sprintf("%s = %s\n", k, v)
	}
	return output
}
