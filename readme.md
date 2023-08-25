# IMA 1C XML CUSTOM PROPERTY PARSER



## How does it work?  

1. It finds all xml in folder
2. Iterates over all found filenames
3. Finds all orders 
4. Stores information in redis (using order number as key)
5. deletes successfully parsed xml
6. in case xml has bad or malformed structure it will not be automatically deleted


## Schemas
XML structure
```go
...
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
...
```
Resulting JSON structure
```json

 {
    "orderNumber": "477294",
    "properties": [
        {
            "name": "Отменен",
            "value": "false"
        },
        {
            "name": "Проведен",
            "value": "true"
        },
        {
            "name": "Адрес доставки",
            "value": "000000, Хабаровский край, Хабаровск"
        },
        {
            "name": "Количество мест",
            "value": "1"
        },
        {
            "name": "ОК_СтоимостьДоставки",
            "value": "0"
        },
        {
            "name": "Стоимость доставки для клиента",
            "value": "332"
        },
        {
            "name": "Стоимость доставки ТК",
            "value": "431"
        }
]}
```
This json will then be written to be available by `477294` key


# Command line arguments

example:
>go run main.go -path C:/Users/legen/GolandProjects/xmls/ -redisAddress 127.0.0.1 -redisPort 6379 -redisDbIndex 7

Parameters:

>-help
Show help

>-path string
Path to XML Documents folder (default "./")

>-redisAddress string
Redis IP flag (default "localhost")

>-redisDbIndex int
Redis DB index (default 10)

>-redisPassword string
Redis Password (leave empty for NoAuth)

>-redisPort string
Redis Port Flag (default "6379")
