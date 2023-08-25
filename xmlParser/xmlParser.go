package xmlParser

import (
	"IMAXMLParser/redis"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Path string

// Struct where I dump xml contents
type XMLDump struct {
	XMLName    xml.Name `xml:"КоммерческаяИнформация"`
	Containers []struct {
		Documents []struct {
			Operation string `xml:"ХозОперация"`
			Number    string `xml:"Номер"`
			Requisits []struct {
				RequisitValue []struct {
					Name  string `xml:"Наименование"`
					Value string `xml:"Значение"`
				} `xml:"ЗначениеРеквизита"`
			} `xml:"ЗначенияРеквизитов"`
		} `xml:"Документ"`
	} `xml:"Контейнер"`
}

// Struct that I will use for storage in Redis
type Order struct {
	OrderNumber string            `json:"orderNumber"`
	Properties  []OrderProperties `json:"properties"`
}

type OrderProperties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CommercialInfo struct {
	Containers []Container `xml:"КоммерческаяИнформация"`
}

type Container struct {
	Document []Document `xml:"Документ"`
}

type Document struct {
	Operation string `xml:"ХозОперация"`
}

func ParseXML(filename string) {
	xmlFile, err := os.ReadFile(Path + filename)
	if err != nil {
		return
	}

	dump := new(XMLDump)
	err = xml.Unmarshal(xmlFile, dump)

	if err != nil {

		//fmt.Printf("error: %v", err, "BAD XML \n")
		return
	}

	//fmt.Printf("%+v\n", dump)

	//filtering loop

	//orderProperties := new(OrderProperties)

	for _, container := range dump.Containers {
		for _, document := range container.Documents {
			if document.Operation != "Заказ товара" {
				continue
			}

			//debug print
			//if document.Number == "477294" {
			//	fmt.Println("----------DEBUG PRINT STARTS----------")
			//	fmt.Println(document)
			//}
			order := new(Order)
			order.Properties = make([]OrderProperties, 0)

			order.OrderNumber = document.Number
			for _, requisit := range document.Requisits {
				for _, value := range requisit.RequisitValue {

					props := OrderProperties{Name: value.Name, Value: value.Value}
					order.Properties = append(order.Properties, props)
					//fmt.Println(order.OrderNumber, " <-------OrderNum")
					//fmt.Println(value.Name, " <-------NAME")
					//fmt.Println(value.Value, " <-------VALUE \n")
				}

			}
			jsonData, err := json.MarshalIndent(order, "", "    ")

			if err != nil {
				return
				//log.Fatal(err)
			}

			//fmt.Println(string(jsonData))

			err = redis.Redis.Set(redis.Ctx, order.OrderNumber, jsonData, 0).Err()
			if err != nil {
				return
				//log.Fatal(err)
			}

			//Write to redis

		}
	}
	err = os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
	//print all redis keys

	//keys, err := redis.Redis.Keys(redis.Ctx, "*").Result()
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//for _, key := range keys {
	//	value, err := redis.Redis.Get(redis.Ctx, key).Result()
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		continue
	//	}
	//
	//	fmt.Println("\n Key:", key, "\n Value: \n", value)
	//}

}

func GetAllXMLFilePaths(directoryPath string) ([]string, error) {
	var filePaths []string

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".xml") {
			filePaths = append(filePaths, filepath.Join(directoryPath, file.Name()))
		}
	}

	return filePaths, nil
}
