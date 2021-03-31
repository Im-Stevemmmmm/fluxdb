package main

import "github.com/Im-Stevemmmmm/fluxdb/client/actions"

func main() {
	actions.InsertData{}.Dispatch(actions.DispatchActionData{})
}

