package messenger

type Writer interface {
	//New(lineNumber int, lineName string, text interface{}) msgs
	Error(lineNumber int, lineName string, text any)
	//Fatal(lineNumber int, lineName string, text any)
}
