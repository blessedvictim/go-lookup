package main

import (
	"fmt"
	"log"

	"github.com/go-lookup/LOL/reflection"
)

type inputGraphQL struct {
	update struct {
		Hru string `dest:"Hru"`
		Id  string `dest:"Id"`
	} `dest:"Update"`
	lol *string `dest:"Lol"`
	lal *string `dest:"LAL"`
}

func example() {

}

func main() {
	//reflection.NonNilPaths()

	//tag := &pb.Tag{}
	//mask := &field_mask.FieldMask{}
	//err := mask.Append(tag, "hru")
	//if err != nil {
	//	log.Panic(err)
	//}
	//upd := pb.UpdateArgs{
	//	Update:    tag,
	//	FieldMask: mask,
	//}
	//log.Println(upd.FieldMask.GetPaths())

	lal := "lal"
	lol := "lol"
	input := inputGraphQL{
		lal: &lal,
		lol: &lol,
	}

	pb.UpdateArgs{
		Update:    nil,
		FieldMask: nil,
	}

	lol := "LOL"
	object := &reflection.Lol{
		Inner: &reflection.Inner{HRU: &lol},
	}
	fieldMask := reflection.NonNilPaths(object)
	fmt.Println(fieldMask)

	val, err := reflection.GetByPath(object, "Inner")
	if err != nil {
		log.Panic(err)
	}

	if realVal, ok := val.(string); ok {
		log.Println(realVal)
	}

	err = reflection.SetByPath(object, "Inner.HRU", "AAZAZAZAZA")
	if err != nil {
		log.Panic(err)
	}

	log.Println(*object.Inner.HRU)
	//request := &pb.UpdateArgs{
	//	Update: &pb.Tag{
	//		Hru: "lol",
	//	},
	//	FieldMask: &fieldmaskpb.FieldMask{Paths: []string{"hru"}},
	//}
	//mask, err := fieldmask_utils.MaskFromPaths(fieldMask, generator.CamelCase)
	//if err != nil {
	//	log.Panic(err)
	//}
	//m := map[string]interface{}{}
	//err = fieldmask_utils.StructToMap(mask, object, m)
	//if err != nil {
	//	log.Panic(err)
	//}

}
