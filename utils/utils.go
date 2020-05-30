package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateRandomStr(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXY"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func JsonDumps(input interface{}) string {
	rst, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	}
	return string(rst)
}

func JsonLoads(input string) map[string]interface{} {
	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(input), &jsonObj)
	if err != nil {
		fmt.Println(err)
	}
	return jsonObj
}

func Int2str(num int) string {
	return strconv.Itoa(num)
}
