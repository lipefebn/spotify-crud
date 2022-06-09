package cli

import (
	"encoding/json"
	"fmt"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lipefebn/spotify-crud/src/adapter/relational"
)

type modelCrud[M relational.Model] struct {
	table relational.Table[M]
	nameTable string
}

func InitModel[M relational.Model](db relational.Table[M], nameTable string) modelCrud[M]{
	return modelCrud[M]{ 
		table: db,
		nameTable: nameTable,
	}
}

func (m modelCrud[M]) GetName() string { return m.nameTable }

func (m modelCrud[M]) Read(orderTable, find, table string) {
	result := widgets.NewTable()
	resultSelect := m.table.Select(orderTable, fmt.Sprintf("%s LIKE ?", table), find+"%")
	m.createTable(resultSelect, result)
	result.SetRect(53, 15, 153, 38)
	ui.Clear()
	ui.Render(result)

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<Escape>":
			ui.Clear()
			return
		}
	}
}

func (m modelCrud[M]) createTable(result []M, table *widgets.Table){
	//header
	header := []string{}
	for column := range(m.structToMap(result[0])){
		header = append(header, column)
	}
	table.Rows = [][]string{ header }
	//--

	for _, elem := range(result) {
		row := []string{}
		mapResult := m.structToMap(elem)
		for _, column := range(header){
			l := mapResult[column]
			conversion, ok := l.(float64)
			if ok {
				l = int(conversion)
			}
			row = append(row, interfaceToString(l))
		}
		table.Rows = append(table.Rows, row)
	}
}

func (m modelCrud[M]) structToMap(row M) (map[string] interface{}) {
	jsonE, _ := json.Marshal(row)
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonE), &result)
	return result
}

func interfaceToString(anyType interface{}) string {
	switch anyType.(type) {
	case int:
		return strconv.Itoa(anyType.(int))
	case bool:
		return strconv.FormatBool(anyType.(bool))
	}	
	return anyType.(string)
}