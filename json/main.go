package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type resp1 struct {
	Page   int
	Fruits []string
}

type resp2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	fln := fmt.Println
	boolB, _ := json.Marshal(true)
	fln(string(boolB))

	intB, _ := json.Marshal(1)
	fln(string(intB))

	fltB, _ := json.Marshal(2.34)
	fln(string(fltB))

	strB, _ := json.Marshal("gopher")
	fln(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fln(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fln(string(mapB))

	res1D := &resp1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fln(string(res1B))

	res2D := &resp2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fln(string(res2B))

	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fln(dat)

	num := dat["num"].(float64)
	fln(num, reflect.TypeOf(num))

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fln(str1, reflect.TypeOf(str1))

	str := `{"page":1, "fruits":["apple", "peach"]}`
	res := resp2{}

	json.Unmarshal([]byte(str), &res)
	fln(res)
	fln(res.Fruits[0])
	fln(res.Fruits[1])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
