// Code generated by 'goexports log'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlib

import (
	"log"
	"reflect"
)

func init() {
	Symbols["log"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Fatal":         reflect.ValueOf(log.Fatal),
		"Fatalf":        reflect.ValueOf(log.Fatalf),
		"Fatalln":       reflect.ValueOf(log.Fatalln),
		"Flags":         reflect.ValueOf(log.Flags),
		"LUTC":          reflect.ValueOf(log.LUTC),
		"Ldate":         reflect.ValueOf(log.Ldate),
		"Llongfile":     reflect.ValueOf(log.Llongfile),
		"Lmicroseconds": reflect.ValueOf(log.Lmicroseconds),
		"Lshortfile":    reflect.ValueOf(log.Lshortfile),
		"LstdFlags":     reflect.ValueOf(log.LstdFlags),
		"Ltime":         reflect.ValueOf(log.Ltime),
		"New":           reflect.ValueOf(log.New),
		"Output":        reflect.ValueOf(log.Output),
		"Panic":         reflect.ValueOf(log.Panic),
		"Panicf":        reflect.ValueOf(log.Panicf),
		"Panicln":       reflect.ValueOf(log.Panicln),
		"Prefix":        reflect.ValueOf(log.Prefix),
		"Print":         reflect.ValueOf(log.Print),
		"Printf":        reflect.ValueOf(log.Printf),
		"Println":       reflect.ValueOf(log.Println),
		"SetFlags":      reflect.ValueOf(log.SetFlags),
		"SetOutput":     reflect.ValueOf(log.SetOutput),
		"SetPrefix":     reflect.ValueOf(log.SetPrefix),

		// type definitions
		"Logger": reflect.ValueOf((*log.Logger)(nil)),
	}
}
