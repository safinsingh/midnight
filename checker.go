package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

////// ERROR HANDLERS //////

func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

////// FILE CONTROL FUNCTIONS //////

func keyInFile(key string, file string) (bool, string) {
	if fileExists(file) {
		f, err := ioutil.ReadFile(file)
		if err != nil {
			panic("Internal file reading error, are you in root?")
		}
		return strings.Contains(string(f), key), ""
	}

	errmsg := "File does not exist: " + file
	return false, errmsg
}

func keyPairInFile(key string, val string, file string) (bool, string) {
	if fileExists(file) {
		f, err := os.Open(file)
		if err != nil {
			panic("Internal file reading error, are you in root?")
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), key) && strings.Contains(scanner.Text(), val) {
				return true, ""
			}
		}

		if err := scanner.Err(); err != nil {
			panic("Internal file scanner error, are you in root?")
		}
	}

	errmsg := "File does not exist: " + file
	return false, errmsg
}

func fileControl(control FileCheck) (bool, string) {
	condInit, errInit := keyInFile(control.Key, control.File)
	if condInit {
		cond, err := keyPairInFile(control.Key, control.Value, control.File)
		if cond {
			return true, ""
		}
		return false, err
	}
	return false, errInit
}

////// COMMAND CHECKS //////

func cmdControl(control CommandCheck) (bool, string) {
	out, err := exec.Command("bash", "-c", control.Command).Output()
	if err != nil {
		return false, err.Error()
	}

	if strings.Contains(string(out), control.ToCheck) {
		return true, ""
	}

	return false, ""
}

////// PACKAGE CHECKS //////

func pkgControl(control PackageCheck) (bool, string) {
	out, err := exec.Command("dpkg", "-l").Output()
	if err != nil {
		return false, err.Error()
	}

	if strings.Contains(string(out), control.Package) && control.Installed {
		return true, ""
	}

	return false, ""
}

////// MAIN SWITCHER //////

func checkSwitch(control Def) (bool, string) {
	switch control.Control.Type {
	case "file":
		return fileControl(*control.Control.FileCheck())
	case "command":
		return cmdControl(*control.Control.CommandCheck())
	case "package":
		return pkgControl(*control.Control.PackageCheck())
	}

	return false, "Check not defined"
}

func commence(controls []Def, mode string) {
	var errs []string

	for i := 0; i < len(controls); i++ {
		code, err := checkSwitch(controls[i])

		if err == "" {
			if code {
				successPrint(controls[i])
			} else {
				failPrint(controls[i])
			}
		} else {
			failPrint(controls[i])
			toApp := controls[i].Description + " -> " + err
			errs = append(errs, toApp)
		}
	}

	yellow := color.New(color.FgYellow, color.Bold)
	yellow.Println("\n[-] Errors:")
	for i := 0; i < len(errs); i++ {
		yellow.Println(errs[i])
	}
}
