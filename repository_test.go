package agollo

import (
	"testing"
	"github.com/zouyx/agollo/test"
	"time"
)

//init param
func init()  {
	//wait 1s for another go routine update apollo config
	time.Sleep(1*time.Second)

	configs:=make(map[string]interface{},0)
	//string
	configs["string"]="value"
	//int
	configs["int"]="1"
	//float
	configs["float"]="190.3"
	//bool
	configs["bool"]="true"

	currentApolloConfig.Configurations=configs
}

func TestGetIntValue(t *testing.T) {
	defaultValue:=100000

	//test default
	v:=GetIntValue("joe",defaultValue)

	test.Equal(t,defaultValue,v)

	//normal value
	v=GetIntValue("int",defaultValue)

	test.Equal(t,1,v)

	//error type
	v=GetIntValue("float",defaultValue)

	test.Equal(t,defaultValue,v)
}

func TestGetFloatValue(t *testing.T) {
	defaultValue:=100000.1

	//test default
	v:=GetFloatValue("joe",defaultValue)

	test.Equal(t,defaultValue,v)

	//normal value
	v=GetFloatValue("float",defaultValue)

	test.Equal(t,190.3,v)

	//error type
	v=GetFloatValue("int",defaultValue)

	test.Equal(t,float64(1),v)
}

func TestGetBoolValue(t *testing.T) {
	defaultValue:=false

	//test default
	v:=GetBoolValue("joe",defaultValue)

	test.Equal(t,defaultValue,v)

	//normal value
	v=GetBoolValue("bool",defaultValue)

	test.Equal(t,true,v)

	//error type
	v=GetBoolValue("float",defaultValue)

	test.Equal(t,defaultValue,v)
}

func TestGetStringValue(t *testing.T) {
	defaultValue:="j"

	//test default
	v:=GetStringValue("joe",defaultValue)

	test.Equal(t,defaultValue,v)

	//normal value
	v=GetStringValue("string",defaultValue)

	test.Equal(t,"value",v)
}