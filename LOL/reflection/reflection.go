package reflection

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

type Lol struct {
	Name  *string `dest:"NamePB"`
	Inner *Inner  `dest:"InnerPB"`
}

type Inner struct {
	HRU         *string `dest:"HRUPB"`
	nonexported string
}

type pathValue struct {
	path  string
	value interface{}
}

func NonNilPaths(src interface{}) map[string]string {
	return genPathsForNonNil(src, nil, "", "")
}

const pathDelimiter = "."

func genPathsForNonNil(src interface{}, paths map[string]string, rootPath, rootPathDst string) map[string]string {
	if paths == nil {
		paths = make(map[string]string, 0)
	}

	var value reflect.Value
	if val, ok := src.(reflect.Value); ok {
		value = val
	} else {
		value = reflect.ValueOf(src)
	}
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return paths
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return paths
	}

	for i := 0; i < value.NumField(); i++ {
		curField := value.Field(i)
		switch curField.Kind() {
		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			switch {
			case !curField.IsNil() && curField.Elem().Kind() == reflect.Struct:
				var newRootPath = value.Type().Field(i).Name
				if rootPath != "" {
					newRootPath = strings.Join([]string{rootPath, newRootPath}, pathDelimiter)
				}
				var newRootPathDst = value.Type().Field(i).Tag.Get("dest")
				if rootPathDst != "" {
					newRootPathDst = strings.Join([]string{rootPathDst, newRootPathDst}, pathDelimiter)
				}
				paths = genPathsForNonNil(curField.Elem(), paths, newRootPath, newRootPathDst)
			case !curField.IsNil():
				if rootPath != "" {
					paths[strings.Join([]string{rootPath, value.Type().Field(i).Name}, pathDelimiter)] = strings.Join([]string{rootPathDst, value.Type().Field(i).Tag.Get("dest")}, pathDelimiter)
					//paths = append(paths, strings.Join([]string{rootPath, value.Type().Field(i).Name}, pathDelimiter))
				} else {
					//paths = append(paths, curField.Type().Field(0).Name)
					paths[curField.Type().Field(0).Name] = value.Type().Field(i).Tag.Get("dest")

				}
			}
		case reflect.Struct:
			var newRootPath = value.Type().Field(i).Name
			if rootPath != "" {
				newRootPath = strings.Join([]string{rootPath, newRootPath}, pathDelimiter)
			}
			var newRootPathDst = value.Type().Field(i).Tag.Get("dest")
			if rootPathDst != "" {
				newRootPathDst = strings.Join([]string{rootPathDst, newRootPathDst}, pathDelimiter)
			}
			paths = genPathsForNonNil(curField, paths, newRootPath, newRootPathDst)
		default:
			//fieldName := value.Type().Field(i).Name
			if rootPath != "" {
				//paths = append(paths, strings.Join([]string{rootPath, fieldName}, pathDelimiter))
				paths[strings.Join([]string{rootPath, value.Type().Field(i).Name}, pathDelimiter)] = strings.Join([]string{rootPathDst, value.Type().Field(i).Tag.Get("dest")}, pathDelimiter)
			} else {
				//paths = append(paths, fieldName)
				paths[curField.Type().Field(0).Name] = value.Type().Field(i).Tag.Get("dest")
			}
		}
	}

	return paths
}

// GetByPath return value by path
func GetByPath(obj interface{}, path string) (interface{}, error) {
	value, err := LookupString(obj, path)
	if err != nil {
		return nil, err
	}

	return value.Interface(), nil
}

// SetByPath set value by path. obj must be pointer
func SetByPath(obj interface{}, path string, val interface{}) error {
	value, err := LookupString(obj, path)
	if err != nil {
		return err
	}
	if !value.Type().AssignableTo(reflect.ValueOf(val).Type()) {
		return errors.New("not assignable")
	}
	value.Set(reflect.ValueOf(val))

	return nil
}

// MapValuesByMapping - осуществляет маппинг полей из структуры src в структуру dst. src и dst must be non-nil.
func FillValuesByMapping(src interface{}, dst interface{}) (interface{}, error) {
	if src == nil || dst == nil {
		return nil, errors.New("nil")
	}

	mapping := NonNilPaths(src)
	for srcPath, dstPath := range mapping {
		v, err := GetByPath(src, srcPath)
		if err != nil {
			return nil, errors.Wrap(err, "get by path")
		}
		err = SetByPath(dst, dstPath, v)
		if err != nil {
			return nil, errors.Wrap(err, "set by path")
		}
	}

	return nil, nil
}

//func mapByMapping(src, dst interface{}) error {
//	if src == nil || dst == nil {
//		return errors.New("nil")
//	}
//
//	srcValue := reflect.ValueOf(src)
//	dstValue := reflect.ValueOf(dst)
//	if {
//
//	}
//}

//func printTypes(item interface{}) {
//	t := reflect.ValueOf(item)
//	for i := 0; i < t.NumField(); i++ {
//		ft := t.Field(i).Type
//		if ft.Kind() == reflect.Ptr {
//			ft = ft.Elem()
//		}
//		fmt.Println(ft.Kind())
//	}
//}
