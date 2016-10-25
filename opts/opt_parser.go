package opts

type Options struct {
	Command string
	Query   string
}

const DefaultCommand string = "search"

func ParseArgs(rawArgs []string) *Options {
	command := DefaultCommand
	query := ""

	args := rawArgs[1:]
	firstArg := args[0]

	if helpWanted(firstArg) {
		command = "help"
	} else {
		query = firstArg
	}

	return &Options{command, query}
}

func helpWanted(arg string) bool {
	helpTerms := [4]string{"help", "--help", "-h"}

	for _, helpTerm := range helpTerms {
		if arg == helpTerm {
			return true
		}
	}

	return false
}
