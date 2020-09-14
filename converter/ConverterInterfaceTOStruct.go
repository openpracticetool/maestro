package converter

import "github.com/mitchellh/mapstructure"

// ConverterInterfaceTOStruct perfect cast
func ConverterInterfaceTOStruct(i interface{}, c interface{}) error {
	return mapstructure.Decode(i, &c)
}
