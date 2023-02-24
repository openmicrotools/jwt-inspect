package main

import (
	"syscall/js"
	"time"

	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

var jsDoc js.Value

func init() {
	jsDoc = js.Global().Get("document")
}

func main() {
	js.Global().Set("inspectJwt", jwtWrapper())
	<-make(chan bool)
}

func jwtWrapper() js.Func {
	jwtFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		//get radiocheck value
		radioCheckElement := getElementByQuerySelector("input[name=radiocheck]:checked")
		radioCheckValue := radioCheckElement.Get("value").String()

		//set printEpoch bool value based on radiocheck value
		var printEpoch bool
		if radioCheckValue == "0" {
			printEpoch = false
		} else {
			printEpoch = true
		}

		inputJwt := args[0].String()
		decoded, err := jwt.DecodeJwt(inputJwt, printEpoch)

		//get decoded Header textarea
		jwtOutputHeaderTextArea := getElementById("jwtoutputheader")

		//get decoded Payload textarea
		jwtOutputPayloadTextArea := getElementById("jwtoutputpayload")

		if err != nil {
			//get alert p element and set error message in the element
			jwtAlertMessage := getElementById("jwterrormessage")
			jwtAlertMessage.Set("innerHTML", err.Error())

			//get alert div element and show
			jwtAlert := getElementById("jwtalert")
			jwtAlert.Get("style").Call("setProperty", "display", "block")

			//hide alert div element after 3 seconds
			time.AfterFunc(3*time.Second, func() {
				jwtAlert.Get("style").Call("setProperty", "display", "none")
			})
		}

		//set decoded header and payload text area if is not nil
		if decoded.Header != nil {
			jwtOutputHeaderTextArea.Set("value", jwt.ToString(decoded.Header))
		} else {
			jwtOutputHeaderTextArea.Set("value", nil)
		}
		if decoded.Payload != nil {
			jwtOutputPayloadTextArea.Set("value", jwt.ToString(decoded.Payload))
		} else {
			jwtOutputPayloadTextArea.Set("value", nil)
		}
		return nil
	})
	return jwtFunc
}

// getElementById gets the dom element by id
// if there is no such element, exit immediately
func getElementById(elementId string) js.Value {
	var element = jsDoc.Call("getElementById", elementId)
	if !element.Truthy() {
		//panic will print the message in console.log useful for debugging/testing
		panic("failed to getElementById: " + elementId)
	}
	return element
}

// getElementByQuerySelector gets the dom element using querySelector
// if there is no such element, exit immediately
func getElementByQuerySelector(query string) js.Value {
	var element = jsDoc.Call("querySelector", query)
	if !element.Truthy() {
		//panic will print the message in console.log useful for debugging/testing
		panic("failed to get element using querySelector: " + query)
	}
	return element
}
