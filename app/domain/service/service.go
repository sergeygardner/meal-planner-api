package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

var updateError = fmt.Errorf("an error occurred while updating the target with the source")

func Update(target any, sources ...any) (*reflect.Value, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic(r)
		}
	}()
	reflectValueTarget := reflect.ValueOf(target)
	sourcesCombined := sources

	for _, source := range sourcesCombined {
		reflectTypeSource := reflect.TypeOf(source)
		reflectValueSource := reflect.ValueOf(source)
		reflectValueSourceElements := reflectValueSource.Elem()
		reflectTypeSourceElements := reflectTypeSource.Elem()
		max := reflectValueSourceElements.NumField()

		for ii := 0; ii < max; ii++ {
			reflectValueSourceElementField := reflectValueSourceElements.Field(ii)

			switch reflectValueSourceElementField.Kind() {
			case reflect.Bool:
				reflectValueTargetField := reflectValueTarget.Elem().FieldByName(reflectTypeSourceElements.Field(ii).Name)

				if reflectValueTargetField.IsValid() {
					reflectValueTargetField.Set(reflectValueSourceElements.Field(ii))
				} else {
					return nil, updateError
				}
			default:
				if !reflectValueSourceElementField.IsZero() {
					reflectValueTargetField := reflectValueTarget.Elem().FieldByName(reflectTypeSourceElements.Field(ii).Name)

					if reflectValueTargetField.IsValid() {
						reflectValueTargetField.Set(reflectValueSourceElements.Field(ii))
					} else {
						return nil, updateError
					}
				}
			}

		}
	}

	return &reflectValueTarget, nil
}
