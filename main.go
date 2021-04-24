package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
	"vezdecode-password-generator/generator"
)

func main() {
	rand.Seed(time.Now().Unix())

	lengthFlag := flag.Int("l", 6, "required length")
	requiredDigitsFlag := flag.Bool("d", false, "required digits")
	requiredUppercaseFlag := flag.Bool("u", false, "required uppercase")
	requiredSpecialCharacterFlag := flag.Bool("s", false, "required special characters")

	inputFileFlag := flag.String("file", "", "file with passwords for bulk checking")
	reportFlag := flag.Bool("report", false, "print report")

	flag.Parse()

	action := flag.Arg(0)

	if action == "" {
		fmt.Println("You should specify the action that you want: [GENERATE], [CHECK]")
		os.Exit(1)
	}

	gen := generator.NewGenerator(
		int64(*lengthFlag),
		*requiredUppercaseFlag,
		*requiredSpecialCharacterFlag,
		*requiredDigitsFlag,
	)

	switch action {
	case "GENERATE":
		pass := gen.Generate()
		fmt.Println(pass)
		os.Exit(0)
	case "CHECK":

		if *inputFileFlag != "" {
			goodPasswords, badPasswords, checks := IterateFile(*inputFileFlag, gen)
			if *reportFlag {
				fmt.Println("REPORT:")
				fmt.Printf("GOOD PASSWORDS (%d):\n", len(goodPasswords))
				for _, p := range goodPasswords {
					fmt.Println(p)
				}
				fmt.Printf("BAD PASSWORDS (%d):\n", len(badPasswords))
				for _, p := range badPasswords {
					fmt.Println(p)
				}
				fmt.Printf("Total checks: %d", checks)
			}
			os.Exit(0)
		}

		passToCheck := flag.Arg(1)
		err := gen.CheckPassword(passToCheck)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		} else {
			fmt.Println("Password meets the requirements")
		}
		break
	default:
		fmt.Println("Wrong action. Available actions: [GENERATE] or [CHECK]")
		os.Exit(1)
	}
}

func IterateFile(filename string, gen *generator.Generator) ([]string, []string, int64) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("cannot able to read the file: %s", err)
		os.Exit(1)
	}
	defer f.Close()

	checks := int64(0)
	goodPasswords := make([]string, 0)
	badPasswords := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pass := scanner.Text()
		checkErr := gen.CheckPassword(pass)
		if checkErr != nil {
			fmt.Printf("`%s` %s\n", pass, checkErr)
			badPasswords = append(badPasswords, pass)
		} else {
			fmt.Printf("`%s` meets the requirements\n", pass)
			goodPasswords = append(goodPasswords, pass)
		}
		checks = checks + 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return goodPasswords, badPasswords, checks
}
