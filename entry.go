package logger

type Entry struct {
	Message string
	Level   Level
	Context *map[string]interface{}
}
