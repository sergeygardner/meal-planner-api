package handler

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"reflect"
)

func printTable(caption string, entities any, kind any) {
	tableStack := make([]*table.Writer, 0)

	setTable(&tableStack, caption, entities, kind)

	for _, tableStackItem := range tableStack {
		fmt.Println((*tableStackItem).Render())
	}
}

func setTable(tableStack *[]*table.Writer, caption string, entities any, kind any) {
	var tableRow table.Row

	tableWriter := table.NewWriter()
	tableWriter.SetCaption(caption)

	*tableStack = append(*tableStack, &tableWriter)

	reflectEntity := reflect.ValueOf(kind)
	tableRow = table.Row{"#"}

	if reflectEntity.Kind() == reflect.Pointer {
		reflectEntity = reflectEntity.Elem()
	}

	for i := 0; i < reflectEntity.NumField(); i++ {
		tableRow = append(tableRow, reflectEntity.Type().Field(i).Name)
	}

	tableWriter.AppendHeader(tableRow)

	reflectEntities := reflect.ValueOf(entities)

	switch reflectEntities.Kind() {
	case reflect.Slice:
		for row := 0; row < reflectEntities.Len(); row++ {
			reflectEntity = reflectEntities.Index(row)

			setTableRow(tableStack, row, tableWriter, reflectEntity)
		}
	case reflect.Struct:
		setTableRow(tableStack, 0, tableWriter, reflectEntities)
	case reflect.Pointer:
		setTableRow(tableStack, 0, tableWriter, reflectEntities)
	default:
		return
	}
}

func setTableRow(tableStack *[]*table.Writer, row int, tableWriter table.Writer, reflectEntity reflect.Value) {
	var tableRow table.Row

	tableRow = table.Row{row}

	if reflectEntity.Kind() == reflect.Pointer {
		reflectEntity = reflectEntity.Elem()
	}

	for i := 0; i < reflectEntity.NumField(); i++ {
		field := reflectEntity.Field(i)

		switch field.Kind() {
		case reflect.Slice:
			if field.Len() > 0 {
				setTable(tableStack, field.Type().String(), field.Interface(), field.Index(0).Interface())
				tableRow = append(tableRow, field.Type().String())
			}
		case reflect.Struct:
			if field.Type().Name() != "Time" {
				setTable(tableStack, field.Type().String(), field.Interface(), field.Interface())
				tableRow = append(tableRow, field.Type().String())
			} else {
				tableRow = append(tableRow, field.Interface())
			}
		case reflect.Pointer:
			setTable(tableStack, field.Type().String(), field.Interface(), field.Interface())
			tableRow = append(tableRow, field.Type().String())
		default:
			if field.CanSet() {
				tableRow = append(tableRow, field.Interface())
			} else {
				tableRow = append(tableRow, "-")
			}
		}
	}

	tableWriter.AppendRow(tableRow)
}
