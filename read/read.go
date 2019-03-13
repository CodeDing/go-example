package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type AntiItem struct {
	Tk       string `json:"tk"`
	RuleCode string `json:"rule_code"`
	RuleName string `json:"rule_name"`
	Ch       string `json:"ch"`
	Source   string `json:"source"`
	Device   string `json:"device"`
}

//ACG6rSiskObpRB_9utnX-sEsqVBmS2_A225odXpob25n^@111^@运行按键精灵^@^@android^@865401030394286

func json2File(dst string, v interface{}) error {
	if v == nil {
		return nil
	}
	file, err := os.Create(dst)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(file)
	err = enc.Encode(&v)
	if err != nil {
		return err
	}
	return nil
}

func Convert2Items(src, dst string, mode int8) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	items := make([]*AntiItem, 0, 100000)
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				//TODO
				ret := make([]byte, 0, 100000)
				if mode == 1 {
					for _, item := range items {
						bs, _ := json.Marshal(item)
						ret = append(ret, bs...)
						ret = append(ret, '\n')
					}
					ioutil.WriteFile(dst, ret, 0644)
				} else {
					json2File(dst, &items)
				}
				return nil
			}
			return err
		}
		parts := strings.Split(strings.TrimSpace(line), "\000")

		if len(parts) != 6 {
			fmt.Printf("not regular string, line: %s", line)
			continue
		}
		items = append(items, &AntiItem{Tk: parts[0], RuleCode: parts[1], RuleName: parts[2], Ch: parts[3], Source: parts[4], Device: parts[5]})
	}
	return nil

}

func main() {
	err := Convert2Items("data/anti-data", "anti_items", 0)
	if err != nil {
		panic(err)
	}
	err = Convert2Items("data/anti-data", "anti_items_array", 1)
	if err != nil {
		panic(err)
	}
}
