package rice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	RiceInfo = &Rice{}

	// 0病害,1好發月份,2水稻期數,3地區,4發生期,5栽種月份,6溫度 ,7濕度,8備註
	ColumnName = map[string]int{
		"name":       0,
		"occur":      1,
		"crop":       2,
		"area":       3,
		"happenDesc": 4,
		"plantAt":    5,
		"temp":       6,
		"RH":         7,
		"comment":    8,
	}
)

type Rice struct {
	Data []Data `json:"data"`
}

type Data struct {
	Name       string `json:"name"`       // 病害名稱
	Occur      string `json:"occur"`      // 好發月份
	Crop       string `json:"crop"`       // 水稻期數
	Area       string `json:"area"`       // 地區
	HappenDesc string `json:"happenDesc"` // 發生期
	PlantAt    string `json:"plantAt"`    // 栽種月份
	Temp       string `json:"temp"`       // 溫度
	RH         string `json:"RH"`         // 濕度
	Comment    string `json:"comment"`    // 備註
	Condition  string `json:"condition"`  // 條件
}

func Init() error {
	// Open our jsonFile
	jsonFile, err := os.Open("internal/rice/rice.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}
	fmt.Println("Successfully Opened rice.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, RiceInfo)
	if err != nil {
		return err
	}

	return nil
}
