// Package gojsonq provides a simple, elegant and fast ODM like API to access/query JSON document.
//
// JSON document can be read from file, string or io.Reader.
// Accessing the value of json property or querying document is simple as the examples below:
//
//	 package main
//
//	 import "github.com/thedevsaddam/gojsonq"
//
//	 const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
//
//	 func main() {
//		 name := gojsonq.New().JSONString(json).Find("name.first")
//		 println(name.(string)) // Tom
//	 }
//
// For more details, see the documentation and examples.
package gojsonq
