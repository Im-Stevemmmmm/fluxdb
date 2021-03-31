package actions

func (i InsertData) Dispatch(d DispatchActionData) {
	data := map[string]interface{} {
		"key": "mykey",
		"value": "myvalue",
	}
	res := makeRequest(data, "/my_sn/my_ns_insert/_insert")
	res.prettyPrint()
}


type InsertData struct{}
