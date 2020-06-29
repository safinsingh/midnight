package main

func main() {
	mode, file := argparse()
	defs := jsonParse(file)
	commence(defs, mode)
}
