package main

import (
  "testing"
  "gotest.tools/assert"
)

func addFunc(a int, b int) (int){
  return a+b
}

// Make sure to add 'Test' to the beginning of your test method!
func TestAddFunc(t *testing.T){
  assert.Equal(t, 3, addFunc(0,3))
}

func compareMap(t *testing.T, mapA map[string]interface{}, mapB map[string]interface{}){
  for key, value := range (mapA){
    valInOther, existsInOther := mapB[key]
    assert.Assert(t, existsInOther)
    assert.Equal(t, value, valInOther)
  }
  for key, value := range (mapB){
    valInOther, existsInOther := mapA[key]
    assert.Assert(t, existsInOther)
    assert.Equal(t, value, valInOther)
  }
}

func TestTesterFunc(t *testing.T){
  // Leverage json_response_struct struct type and testerFunc from hello.go
  var jsonbody json_response_struct
  testerFunc(&jsonbody)
  testbody := json_response_struct{
    Struct: map[string]interface{}{"test":"blah2"},
    List: []map[string]interface{}{},
  }
  compareMap(t, jsonbody.Struct, testbody.Struct)
}