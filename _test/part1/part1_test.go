package part1_test

import (
	"reflect"
	"testing"

	"github.com/kara9renai/todoapp-go/controller/dto"
	"github.com/kara9renai/todoapp-go/model/entity"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	entityTestcases := map[string]struct {
		Target    interface{}
		FieldNmae string
		WantKinds []reflect.Kind
		WantType  reflect.Type
	}{
		"TodoEntity has ID field": {
			Target:    entity.TodoEntity{},
			FieldNmae: "ID",
			WantKinds: []reflect.Kind{reflect.Int},
		},
		"TodoEntity has Title field": {
			Target:    entity.TodoEntity{},
			FieldNmae: "Title",
			WantKinds: []reflect.Kind{reflect.String},
		},
		"TodoEntity has Content field": {
			Target:    entity.TodoEntity{},
			FieldNmae: "Content",
			WantKinds: []reflect.Kind{reflect.String},
		},
	}

	for name, tc := range entityTestcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			tp := reflect.TypeOf(tc.Target)
			f, ok := tp.FieldByName(tc.FieldNmae)
			if !ok {
				t.Error(tc.FieldNmae + "field not found")
				return
			}

			notFound := true
			for _, k := range tc.WantKinds {
				if f.Type.Kind() == k {
					notFound = false
					break
				}
			}
			if notFound {
				t.Errorf(tc.FieldNmae+"is not kind which is expected: got %s, want %s", f.Type.Kind(), tc.WantKinds)
				return
			}

			if tc.WantType == nil {
				return
			}

			if f.Type != tc.WantType {
				t.Errorf(tc.FieldNmae+"is not Type which is expected: got %s, want %s", f.Type, tc.WantType)
			}
		})
	}

	JSONTestCases := map[string]struct {
		Target       interface{}
		FieldName    string
		WantKinds    []reflect.Kind
		WantTypes    []reflect.Type
		JSONTagValue string
	}{
		"TodoResponse has ID field": {
			Target:       dto.TodoResponse{},
			FieldName:    "ID",
			WantKinds:    []reflect.Kind{reflect.Int},
			JSONTagValue: "id",
		},
		"TodoResponse has Title field": {
			Target:       dto.TodoResponse{},
			FieldName:    "Title",
			WantKinds:    []reflect.Kind{reflect.String},
			JSONTagValue: "title",
		},
		"TodoResponse has Content field": {
			Target:       dto.TodoResponse{},
			FieldName:    "Content",
			WantKinds:    []reflect.Kind{reflect.String},
			JSONTagValue: "content",
		},
		"TodoRequest has Title field": {
			Target:       dto.TodoRequest{},
			FieldName:    "Title",
			WantKinds:    []reflect.Kind{reflect.String},
			JSONTagValue: "title",
		},
		"TodoRequest has Content field": {
			Target:       dto.TodoRequest{},
			FieldName:    "Content",
			WantKinds:    []reflect.Kind{reflect.String},
			JSONTagValue: "content",
		},
		"TodosResponse has Todos field": {
			Target:       dto.TodosResponse{},
			FieldName:    "Todos",
			WantKinds:    []reflect.Kind{reflect.Slice},
			WantTypes:    []reflect.Type{reflect.TypeOf([]dto.TodoResponse{})},
			JSONTagValue: "todos",
		},
	}

	for name, tc := range JSONTestCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			tp := reflect.TypeOf(tc.Target)
			f, ok := tp.FieldByName(tc.FieldName)
			if !ok {
				t.Error(tc.FieldName + " field not found")
				return
			}

			notFound := true
			for _, k := range tc.WantKinds {
				if f.Type.Kind() == k {
					notFound = false
					break
				}
			}
			if notFound {
				t.Errorf(tc.FieldName+"is not want kind, got = %s, want = %s", f.Type.Kind(), tc.WantKinds)
				return
			}
			if tc.WantTypes != nil {
				notFound = true
				for _, et := range tc.WantTypes {
					if f.Type == et {
						notFound = false
						break
					}
				}
				if notFound {
					t.Errorf(tc.FieldName+" is not want type, got = %s, want = %s", f.Type, tc.WantTypes)
					return
				}
			}

			v, ok := f.Tag.Lookup("json")
			if !ok {
				t.Error("json tag not found")
				return
			}

			if v != tc.JSONTagValue {
				t.Errorf("json tag is not expected, got = %s, want = %s", v, tc.JSONTagValue)
				return
			}

		})
	}

}
