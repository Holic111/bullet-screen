package public

type ResponseMap struct {
	Code int
	Data map[string]interface{}
	Msg string
}

type ResponseData struct {
	Code int
	Data interface{}
	Msg string
}