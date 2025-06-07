package tags

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Bold   = "\033[1m"
	Err    = Red + Bold + "[ERR]\t" + Reset
	Log    = Yellow + Bold + "[LOG]\t" + Reset
	Info   = Green + Bold + "[INFO]\t" + Reset
)
