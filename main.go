package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var args []string
var Version string = "undefined"

func main() {
	args = os.Args[1:]
	namespaceFound := false
	namespace := "undefined"
	parseInp := ""
	buf := ""

	if len(args) > 0 { // check if there is any args after "csmarkdoc"
		for _, str := range args {
			dat, err := os.ReadFile(expandPath(str))
			if err != nil {
				fmt.Println("FAILED reading " + expandPath(str))
			}
			// remove indentation & replace CRLF with LF
			buf = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(string(dat), "\r\n", "\n"), "\t", ""), "    ", "")
			lines := strings.Split(buf, "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "namespace") && !namespaceFound {
					// if it's a namespace declaration (namespace.*), just set the namespace for the project and don't print anything
					namespace = strings.Replace(strings.Replace(strings.Split(line, " ")[1], ";", "", 1), "{", "", 1)
					namespaceFound = true
				} else if strings.HasPrefix(line, "using") {
				} else if strings.HasPrefix(line, "public") { // brainfuck code be like:
					if strings.Contains(line, "class") {
						// If it's a class,
						// then print like this:
						//
						// ```cs
						// public class Class
						// ```
						// ## Class `Class`
						// ---

						parseInp += "```cs\n" + strings.ReplaceAll(line, "{", "") + "\n```\n## Class `" + strings.Split(line, " ")[2] + "`\n---\n"
					} else if strings.Contains(line, "struct") {
						// If it's a struct,
						// then print like this:
						//
						// ```cs
						// public struct Struct
						// ```
						// ## Struct `Struct`
						// ---

						parseInp += "```cs\n" + strings.ReplaceAll(line, "{", "") + "\n```\n## Struct `" + strings.Split(line, " ")[2] + "`\n---\n"
					} else {
						// Otherwise,
						// then print like this:
						//
						// ```cs
						// public static void Method(string arg)
						// ```
						// ---
						parseInp += "```cs\n" + strings.ReplaceAll(strings.ReplaceAll(line, ";", ""), "{", "") + "\n```\n---\n"
					}
				} else if strings.HasPrefix(line, "///") && !strings.HasPrefix(line, "/// <") {
					// if it's a valid XML comment (///.*), and is not a operator (/// <.*),
					// Just print the XML comment without the ///. we'll handle <see> tags in a future version.
					parseInp += strings.ReplaceAll(strings.ReplaceAll(line, "/// ", ""), "///", "") + "\n"
				}
			}
		}

		fmt.Println("# " + namespace)
		fmt.Println(strings.Replace(buf, buf, "", 1))
		fmt.Println(parseInp)
		fmt.Println("(generated by [csmarkdoc](https://github.com/malinoOS/csmarkdoc))")
	} else {
		fmt.Println("csmarkdoc: no files")
		printHelp()
	}
}

func printHelp() {
	fmt.Println("csmarkdoc v" + Version)
	fmt.Println("Usage: csmarkdoc [files]")
	os.Exit(2)
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "/") {
		return path
	} else {
		curDir, err := os.Getwd()
		if err != nil {
			return ""
		}
		return filepath.Join(curDir, path)
	}
}
