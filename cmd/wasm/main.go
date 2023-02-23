package main

import (
	"syscall/js"
	"time"

	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

func main() {
	js.Global().Set("inspectJwt", jwtWrapper())
	<-make(chan bool)
}

func jwtWrapper() js.Func {
	jwtFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		jsDoc := js.Global().Get("document")

		//get radiocheck value
		radioCheck := jsDoc.Call("querySelector", "input[name=radiocheck]:checked").Get("value").String()

		//set printEpoch bool value based on radiocheck value
		var printEpoch bool
		if radioCheck == "0" {
			printEpoch = false
		} else {
			printEpoch = true
		}

		inputJwt := args[0].String()
		decoded, err := jwt.DecodeJwt(inputJwt, printEpoch)

		//get decoded Header textarea
		jwtOutputHeaderTextArea := jsDoc.Call("getElementById", "jwtoutputheader")

		//get decoded Payload textarea
		jwtOutputPayloadTextArea := jsDoc.Call("getElementById", "jwtoutputpayload")

		if err != nil {
			//get alert p element and set error message in the element
			jwtAlertMessage := jsDoc.Call("getElementById", "jwterrormessage")
			jwtAlertMessage.Set("innerHTML", err.Error())

			//get alert div element and show
			jwtAlert := jsDoc.Call("getElementById", "jwtalert")
			jwtAlert.Get("style").Call("setProperty", "display", "block")

			//hide alert div element after 3 seconds
			time.AfterFunc(3*time.Second, func() {
				jwtAlert.Get("style").Call("setProperty", "display", "none")
			})
		}

		//set decoded header and payload text area
		if decoded.Header != nil {
			jwtOutputHeaderTextArea.Set("value", decoded.Header.ToString())
		} else {
			jwtOutputHeaderTextArea.Set("value", nil)
		}
		if decoded.Payload != nil {
			jwtOutputPayloadTextArea.Set("value", decoded.Payload.ToString())
		} else {
			jwtOutputPayloadTextArea.Set("value", nil)
		}
		return nil
	})
	return jwtFunc
}
