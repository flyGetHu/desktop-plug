package test

import (
	"bailidaming.com/anjun/destop/utils"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestFor(t *testing.T) {
	var count = 100

	for {
		if count < 0 {
			break
		}
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			t.Error("随机数为0")
			break
		}
		count--
	}
	if count == 0 {
		t.Error("退出")
	} else {
		t.Error("错误退出")
	}
}

type person struct {
	Name string `json:"name"`
}

func TestJson(t *testing.T) {
	str := `
		{
			"name":"huan",
			"age":[22,33]
		}
	`
	m := map[string]string{}
	json.Unmarshal([]byte(str), &m)
	log.Println(m)
}

func TestStrSerial(t *testing.T) {
	var split = "\u000200014\u0003\u000200026\u0003\u000200154\u0003\u000200226\u0003"
	var tempStr = ""
	for _, num := range split {
		if len(tempStr) == 5 && utils.StrIsNum(tempStr) {
			break
		}
		if len(tempStr) != 0 && !utils.StrIsNum(string(num)) {
			return
		}
		if utils.StrIsNum(string(num)) {
			tempStr += string(num)
		}
	}
	var num = ""
	if utils.StrIsNum(tempStr) && len(tempStr) == 5 {
		for i, item := range tempStr {
			if i == 3 {
				num += "."
			}
			num += string(item)
		}
	}
	if weight, err := strconv.ParseFloat(num, 64); err == nil {
		log.Println(weight)
	}
	log.Println(tempStr)
}
