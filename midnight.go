package main

func main() {
	mode, file := argparse()
	defs := jsonParse(file)
	handle(defs, mode)
}
