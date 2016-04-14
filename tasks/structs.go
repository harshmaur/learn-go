package tasks

import (
	"encoding/json"
	"fmt"
)

type NewPerson struct {
	First       string
	Last        string
	Age         int
	notExported int    //this will not show up in Marshal
	Middle      string `json:"-"`            // explicitly excluding from json
	Mobile      string `json:"wisdom score"` // Mobile will be replaced with wisdom score
}

// JSONMarshal is used to convert struct type to json
func JSONMarshal() {
	p1 := NewPerson{"Harsh", "Maur", 21, 007, "None", "Hey"}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}

// JSONUnmarshal is used to convert json to struct
func JSONUnmarshal() {
	var p1 NewPerson
	bs := []byte(`{"First":"Harsh","Last":"Maur","Age":21,"Middle":"None","wisdom score":"9899901974"}`)
	json.Unmarshal(bs, &p1)
	fmt.Println(p1)
}
