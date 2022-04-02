package lib

import (
	"github.com/google/uuid"
	"log"
)

func Strptr(s string) *string {
	return &s
}

func Intptr(i int) *int {
	return &i
}

func Int64ptr(i int64) *int64 {
	return &i
}

func Float64ptr(f float64) *float64 {
	return &f
}

func UUIDPtr(u uuid.UUID) *uuid.UUID {
	return &u
}

func BoolPtr(b bool) *bool {
	return &b
}

func PrintPtr(obj interface{}) interface{} {
	//switch data := obj.(type) {
	//case string:
	//	log.Println(data, "string")
	//default:
	//	log.Println(data)
	//}
	var resultArrayInterface []interface{}
	var resultInterface = new(interface{})
	var resultMap = make(map[string]interface{})
	err := Merge(obj, &resultMap)
	if err != nil {
		log.Println(err)
		err = Merge(obj, &resultInterface)
		if err != nil {
			log.Println(err)
			err = Merge(obj, &resultArrayInterface)
			if err != nil {
				log.Println(err)
				return nil
			}
			return resultArrayInterface
		}
		return resultInterface
	}
	return resultMap
}
