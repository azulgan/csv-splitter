package main

type MyArgs struct {
	InputFileName string
	OutputFileName string
	Comma rune
	comma string
}

func (a *MyArgs) parseArgs(args []string) {
	a.InputFileName = "input.csv"
	a.OutputFileName = "output.csv"
	a.comma = ";"
	for {
		if len(args) == 0 {
			break
		}
		args, param := a.nextParam(args)
		switch param {
		case "-i": args, a.InputFileName = a.nextParam(args)
			break
		case "-o": args, a.OutputFileName = a.nextParam(args)
			break
		case "-s": args, a.comma = a.nextParam(args)
			break
		}
	}
	a.Comma = rune(a.comma[0])
}

func (a MyArgs) nextParam(args []string) ([]string, string) {
	var ret string
	ret = ""
	if len(args) > 0 {
		ret = args[0]
		args = args[1:]
	}
	return args, ret
}