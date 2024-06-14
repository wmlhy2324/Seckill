package GetGormFields

import (
	"reflect"
	"strings"
)

func GetGormFields(stc interface{}) []string {
	value := reflect.ValueOf(stc)
	typ := value.Type()
	if typ.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil
		} else {
			typ = typ.Elem()
			value = value.Elem()
		}
	}
	if typ.Kind() == reflect.Struct {
		columns := make([]string, 0, value.NumField())
		for i := 0; i < value.NumField(); i++ {
			fieldType := typ.Field(i)
			if fieldType.IsExported() {
				if fieldType.Tag.Get("gorm") == "-" {
					continue
				}
				name := Camel2Snake(fieldType.Name)
				if len(fieldType.Tag.Get("gorm")) > 0 {
					content := fieldType.Tag.Get("gorm")
					if strings.HasPrefix(content, "column:") {
						content = content[7:]
						pos := strings.Index(content, ";")
						if pos > 0 {
							name = content[0:pos]
						} else if pos < 0 {
							name = content
						}
					}
				}
				columns = append(columns, name)
			}
		}
		return columns
	} else {
		return nil
	}
}
func Camel2Snake(s string) string {
	if len(s) == 0 {
		return ""
	}
	t := make([]byte, 0, len(s)+4)
	if IsASCIIUpper(s[0]) {
		t = append(t, UpperLowerExchange(s[0]))
	} else {
		t = append(t, s[0])
	}
	for i := 1; i < len(s); i++ {
		c := s[i]
		if IsASCIIUpper(c) {
			t = append(t, '_', UpperLowerExchange(c))
		} else {
			t = append(t, c)
		}
	}
	return string(t)
}
func IsASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func UpperLowerExchange(c byte) byte {
	return c ^ ' '
}
