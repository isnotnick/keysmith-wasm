package main

import (
	"encoding/json"

	"syscall/js"

	"github.com/isnotnick/keysmith"
)

func main() {
	c := make(chan bool)
	js.Global().Set("generateCSR", js.FuncOf(jsonGenerateCSR))
	<-c

	/*subjectName := keysmith.SubjectStruct{
		CN:   "nickfrance.com",
		C:    "GB",
		SANS: "nickfrance.com,www.nickfrance.com,secure.nickfrance.com",
	}
	generateCSR := keysmith.KeyBlock{
		SubjectInfo: subjectName,
		KeyType:     "EC",
		KeySize:     384,
	}

	generateCSR.GenerateCSR()

	fmt.Println(generateCSR.CSRPEM)

	prettyPrinted, _ := json.MarshalIndent(generateCSR, "", "\t")
	fmt.Print(string(prettyPrinted))*/
}

func jsonGenerateCSR(this js.Value, inputs []js.Value) interface{} {
	jsonInput := inputs[0].String()

	csrStruct := &keysmith.KeyBlock{}
	_ = json.Unmarshal([]byte(jsonInput), csrStruct)

	csrStruct.GenerateCSR()

	csrReturn, _ := json.Marshal(csrStruct)
	return string(csrReturn)
}
