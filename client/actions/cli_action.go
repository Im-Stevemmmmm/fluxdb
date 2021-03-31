package actions

func DispatchAction(a dbAction) {

}

type dbAction interface {
	Dispatch(d DispatchActionData) error
}

type DispatchActionData struct {
	Name  *string
	Flags []*string
}

