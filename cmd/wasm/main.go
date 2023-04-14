package main

import (
	"syscall/js"

	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
	"github.com/openmicrotools/jwt-inspect/pkg/text"
)

func main() {
	js.Global().Set("inspectJwt", jwtWrapper())
	<-make(chan bool)
}

func jwtWrapper() js.Func {
	jwtFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		var jsDoc = js.Global().Get("document")

		if !jsDoc.Truthy() {
			//panic will print the message in console.log useful for debugging/testing
			panic("failed to get document object")
		}

		//get radiocheck value
		radioCheckElement := getElementByQuerySelector(jsDoc, "input[name=radiocheck]:checked")
		radioCheckValue := radioCheckElement.Get("value").String()

		//set printEpoch bool value based on radiocheck value
		var printEpoch bool
		if radioCheckValue == "0" {
			printEpoch = false
		} else {
			printEpoch = true
		}

		var timezone = ""
		//get browser timezone using JS Internationalization API(Intl.DateTimeFormat().resolvedOptions().timeZone)
		if intl := js.Global().Get("Intl"); !intl.IsUndefined() {
			timezone = js.Global().Get("Intl").Call("DateTimeFormat").Call("resolvedOptions").Get("timeZone").String()
		}

		inputJwt := args[0].String()
		decoded, err := jwt.DecodeJwt(inputJwt, printEpoch, timezone)

		//get decoded Header textarea
		jwtOutputHeaderTextArea := getElementById(jsDoc, "jwtoutputheader")

		//get decoded Payload textarea
		jwtOutputPayloadTextArea := getElementById(jsDoc, "jwtoutputpayload")

		//get decoded hidden textarea
		//hidden text area is used for copying whole decoded json both header and payload to clipboard
		jwtOutputTextArea := getElementById(jsDoc, "jwtoutput")

		//get alert div element
		jwtAlert := getElementById(jsDoc, "jwtalert")

		if err != nil {
			//get alert p element and set error message in the element
			jwtAlertMessage := getElementById(jsDoc, "jwterrormessage")
			jwtAlertMessage.Set("innerHTML", text.CapitalizeFirstChar(err.Error()))

			//show alert div element
			jwtAlert.Get("style").Call("setProperty", "display", "block")
		} else {
			//if there is no error hide alert div element
			jwtAlert.Get("style").Call("setProperty", "display", "none")
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

		//set decoded json to hidden text area if it is not empty
		if (decoded != jwt.Jwt{}) {
			jwtOutputTextArea.Set("value", jwt.ToString(decoded))
		}
		return nil
	})
	return jwtFunc
}

// getElementById gets the dom element by id
// if there is no such element, exit immediately
func getElementById(jsDoc js.Value, elementId string) js.Value {
	var element = jsDoc.Call("getElementById", elementId)
	if !element.Truthy() {
		//panic will print the message in console.log useful for debugging/testing
		panic("failed to getElementById: " + elementId)
	}
	return element
}

// getElementByQuerySelector gets the dom element using querySelector
// if there is no such element, exit immediately
func getElementByQuerySelector(jsDoc js.Value, query string) js.Value {
	var element = jsDoc.Call("querySelector", query)
	if !element.Truthy() {
		//panic will print the message in console.log useful for debugging/testing
		panic("failed to get element using querySelector: " + query)
	}
	return element
}
