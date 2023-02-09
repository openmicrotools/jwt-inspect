package main

import (
	"fmt"
	"syscall/js"

	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

func main() {
	js.Global().Set("inspectJwt", jwtWrapper())
	<-make(chan bool)
}

func jwtWrapper() js.Func {
	jwtFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
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
		inputJwt := args[0].String()
		decoded, err := jwt.DecodeJwt(inputJwt)

		if err != nil {
			errStr := fmt.Sprintf("unable to decode JWT. Error %s occurred\n", err)
			result := map[string]any{
				"error": errStr,
			}
			return result
		}
		jwtOutputTextArea.Set("value", decoded.ToString())
		return nil
	})
	return jwtFunc
}
