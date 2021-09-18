package cmdutil

func Header() (ansicode string) {
	ansicode = "\033[95m"

	return
}

func Blue() (ansicode string) {
	ansicode = "\033[94m"

	return
}

func Cyan() (ansicode string) {
	ansicode = "\033[96m"

	return
}

func Green() (ansicode string) {
	ansicode = "\033[92m"

	return

}

func Red() (ansicode string){
	ansicode = "\033[31m"

	return
}

func Bold() (ansicode string) {
	ansicode = "\033[1m"

	return
}

func Underline() (ansicode string) {
	ansicode = "\033[4m"

	return
}

func End() (ansicode string) {
	ansicode = "\033[0m"

	return
}
