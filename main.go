package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	exp := `{"example_test":{"example_test_1":1},"exampleTest":"1234","example_test_1":1}`
	if len(os.Args) > 1 {
		exp = os.Args[1]
	}
	var v interface{}
	if err := json.Unmarshal([]byte(exp), &v); err != nil {
		fmt.Println(err)
		return
	}
	m := handleV(v)
	mv, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(mv))
}

func handleV(input interface{}) interface{} {
	if sliceV, ok := input.([]interface{}); ok {
		var ret []interface{}
		for _, val := range sliceV {
			ret = append(ret, handleV(val))
		}
		return ret
	}

	if mapV, ok := input.(map[string]interface{}); ok {
		ret := make(map[string]interface{})
		for k, v := range mapV {
			key := handK(k)
			mk, ok := ret[key]
			_, ok1 := mk.([]interface{})
			_, ok2 := mk.(map[string]interface{})
			if !ok || (!ok1 && !ok2) {
				ret[key] = handleV(v)
			}
		}
		return ret
	}
	return input
}

func handK(key string) string {
	v := strings.Split(key, "_")
	if len(v) > 1 {
		res := v[0]
		for i := 1; i < len(v); i++ {
			vs := []rune(v[i])
			if len(vs) < 1 {
				continue
			}
			if vs[0] >= 'a' && vs[0] <= 'z' {
				vs[0] = vs[0] - 'a' + 'A'
			}
			res += string(vs)
		}
		return res
	}
	return v[0]
}
