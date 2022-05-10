package comm

import (
	"log"
	"reflect"
	"strconv"
)

// get struct column list
func ListStructColumnName(aa interface{}) []string {
	var send []string
	v := reflect.ValueOf(aa)
	type_of_fields := v.Type()

	for i := 0; i < v.NumField(); i++ {
		column := type_of_fields.Field(i).Name
		send = append(send, column)
	}

	return send
}

// struct to map[key]types  types is string
func ListStructKeyType(aa interface{}) map[string]string {
	var send map[string]string // map[column]ColumnType
	v := reflect.ValueOf(aa)
	type_of_fields := v.Type()

	for i := 0; i < v.NumField(); i++ {
		column := type_of_fields.Field(i).Name
		types := type_of_fields.Field(i).Type.String()
		send[column] = types
	}
	return send
}

// struct to map[key]value  ,  value is interface,  check it type
func ListStructKeyValue(aa interface{}) map[string]interface{} {
	var send map[string]interface{} // map[column]ColumnType
	v := reflect.ValueOf(aa)
	type_of_fields := v.Type()

	for i := 0; i < v.NumField(); i++ {
		column := type_of_fields.Field(i).Name
		send[column] = v.Field(i).Interface()
	}

	return send
}

// struct to map[string]interface{} ,  this is map's interface indue type ,
func RangeStruct(aa interface{}, maps map[string]interface{}) map[string]interface{} {
	v := reflect.ValueOf(aa)
	type_of_fields := v.Type()
	for i := 0; i < v.NumField(); i++ {
		types := type_of_fields.Field(i).Type.String()
		column := type_of_fields.Field(i).Name
		val := maps[column].(string)
		// fmt.Printf(" %s [%s]: %s\n", type_of_fields.Field(i).Name, type_of_fields.Field(i).Type, val)
		if types == "string" {
			maps[types] = val
		} else if types == "int64" {
			int6464, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				log.Fatalf("%s ", err)
			}
			maps[types] = int6464
		} else if types == "time.Time" {
			// time.Time
			// maps[types] = val.(time.Time)
		}
	}

	// map value is interface but, If you run 'reflect.typeOf()' on it, it has a type.
	// and then it is not error if you use 'mapstructure.Decode(m, &t)'
	return maps
}
