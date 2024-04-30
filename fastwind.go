package main

import (
	"strings"
	"encoding/json"
	// "bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	if (checkingForViteReactProject()) {
		fmt.Println("\nThis appears to be a React project.")
	} else {
		printInRed("\nThis is not a React (Vite) project.")
		return;
	}

	// npm install -D tailwindcss postcss autoprefixer
	cmd := "npm"
	args := []string{"install", "-D", "tailwindcss", "postcss", "autoprefixer"}

	runCommand(cmd, args)
	
	// npx tailwindcss init -p
	cmd = "npx"
	args = []string{"tailwindcss", "init", "-p"}

	runCommand(cmd, args)

	modifyContentOfTailwindConfigForViteReact()

	modifyIndexCSS()

	printInMagenta("\nTailwind Successfully Installed!\n")


}

func modifyContentOfTailwindConfigForViteReact() {

	printInCyan("\nModifying content of tailwind.config.js...")

	content, err := ioutil.ReadFile("tailwind.config.js")
	if err != nil {
		panic(err)
	}

	contentStr := string(content)

	index := strings.Index(contentStr, "content: []")
	if index == -1 {
		return;
	}

	newContentStr := contentStr[:index] + `content: [
		"./index.html",
		"./src/**/*.{js,ts,jsx,tsx}",
	]` + contentStr[index+len("content: []"):]

	err = ioutil.WriteFile("tailwind.config.js", []byte(newContentStr), 0644)
	if err != nil {
		panic(err)
	}

	printInGreen("\ntailwind.config.js modified successfully.")
	return;
}

func modifyIndexCSS() {

	printInCyan("\nModifying content of index.css ...")

	newContentStr := 
	`@tailwind base;
@tailwind components;
@tailwind utilities;`

	err := ioutil.WriteFile("./src/index.css", []byte(newContentStr), 0644)
	if err != nil {
		panic(err)
	}

	printInGreen("\nindex.css modified successfully.")
	return;
}

func runCommand(command string, args []string) string {
	// Create command object
	cmd := exec.Command(command, args...)

	printInCyan("\nRunning command:", command+" "+filepath.Join(args...))

	output, err := cmd.CombinedOutput()
	if err != nil {
		printInCyan("Error executing command:", err)
		return ""
	}
	fmt.Println("\nCommand output:", string(output))
	printInGreen("Command executed successfully.")
	return string(output)
}

func checkingForViteReactProject() (bool) {
	printInYellow("\nChecking for React project...")

	wd, err := os.Getwd()
	if err != nil {
		printInRed("Error:", err)
		return false;
	}

	// Check if package.json exists
	packageJSONPath := filepath.Join(wd, "package.json")
	_, err = os.Stat(packageJSONPath)
	if os.IsNotExist(err) {
		printInRed("\nError: Not a Node.js project (package.json not found)")
		return false;
	}

	// Check if vite.config.js exists
	viteConfigPath := filepath.Join(wd, "vite.config.js")
	_, err = os.Stat(viteConfigPath)
	if os.IsNotExist(err) {
		printInRed("\nError: Not a Vite project (vite.config.js not found)")
		return false
	}

	// checking for React dependencies
	// (Assuming "dependencies" and "devDependencies" fields in package.json)
	reactDependencies := []string{"react", "react-dom"}
	doesHaveDependencies, err := hasDependencies(packageJSONPath, reactDependencies)
	if err != nil {
		printInRed("\nError:", err)
		return false;
	}

	if !doesHaveDependencies {
		printInRed("\nError: Not a React project (React dependencies not found)")
		return false
	}

	// checking for src directory
	srcDir := filepath.Join(wd, "src")
	_, err = os.Stat(srcDir)
	if os.IsNotExist(err) {
		printInRed("\nError: Not a React project (src directory not found)")
		return false;
	}

	return true
}

func doesTailwindConfigExist() (bool) {
	wd, err := os.Getwd()
	if err != nil {
		printInRed("Error:", err)
		return false;
	}

	// Check if package.json exists
	packageJSONPath := filepath.Join(wd, "tailwind.config.js")
	_, err = os.Stat(packageJSONPath)
	if os.IsNotExist(err) {
		return false;
	}

	return true
}

func hasDependencies(packageJSONPath string, dependencies []string) (bool, error) {
	file, err := ioutil.ReadFile(packageJSONPath)
	if err != nil {
		return false, err
	}

	var packageJSON struct {
		Dependencies   map[string]string `json:"dependencies"`
		DevDependencies map[string]string `json:"devDependencies"`
	}

	// Unmarshal JSON content into the struct
	if err := json.Unmarshal(file, &packageJSON); err != nil {
		return false, err
	}

	for _, dep := range dependencies {
		if _, ok := packageJSON.Dependencies[dep]; ok {
			return true, nil
		}
		if _, ok := packageJSON.DevDependencies[dep]; ok {
			return true, nil
		}
	}

	return false, nil
}

const (
    reset   = "\033[0m"
    red     = "\033[31m"
    green   = "\033[32m"
    yellow  = "\033[33m"
    blue    = "\033[34m"
    magenta = "\033[35m"
    cyan    = "\033[36m"
)

func printInYellow(args ...interface{}) {
    var concatenatedString string
    for _, arg := range args {
        concatenatedString += fmt.Sprintf("%v ", arg)
    }
    fmt.Println(yellow + concatenatedString + reset)
}

func printInRed(args ...interface{}) {
    var concatenatedString string
    for _, arg := range args {
        concatenatedString += fmt.Sprintf("%v ", arg)
    }
    fmt.Println(red + concatenatedString + reset)
}

func printInGreen(args ...interface{}) {
    var concatenatedString string
    for _, arg := range args {
        concatenatedString += fmt.Sprintf("%v ", arg)
    }
    fmt.Println(green + concatenatedString + reset)
}

func printInBlue(args ...interface{}) {
    var concatenatedString string
    for _, arg := range args {
        concatenatedString += fmt.Sprintf("%v ", arg)
    }
    fmt.Println(blue + concatenatedString + reset)
}

func printInMagenta(args ...interface{}) {
    var concatenatedString string
    for _, arg := range args {
        concatenatedString += fmt.Sprintf("%v ", arg)
    }
    fmt.Println(magenta + concatenatedString + reset)
}

func printInCyan(args ...interface{}) {
    var concatenatedString string
    for _, arg := range args {
        concatenatedString += fmt.Sprintf("%v ", arg)
    }
    fmt.Println(cyan + concatenatedString + reset)
}