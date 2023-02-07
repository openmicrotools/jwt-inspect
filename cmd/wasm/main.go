package main

import (
	"fmt"
	"syscall/js"

	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

func main() {
	js.Global().Set("inspectJwt", jsonWrapper())
	<-make(chan bool)
}

func jsonWrapper() js.Func {
	jwtfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			result := map[string]any{
				"error": "Invalid no of arguments passed",
			}
			return result
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "Unable to get document object",
			}
			return result
		}
		jwtOutputTextArea := jsDoc.Call("getElementById", "jwtoutput")
		if !jwtOutputTextArea.Truthy() {
			result := map[string]any{
				"error": "Unable to get output text area",
			}
			return result
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		content, err := jwt.DecodeJwt(inputJSON)
		if err != nil {
			errStr := fmt.Sprintf("unable to parse JSON. Error %s occurred\n", err)
			result := map[string]any{
				"error": errStr,
			}
			return result
		}
		jwtOutputTextArea.Set("value", content)
		return nil
	})
	return jwtfunc
}
