package main

import (
	"fmt"
    "os"
    "path/filepath"
    "io/ioutil"
    "encoding/json"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Check if package.json exists
	packageJSONPath := filepath.Join(wd, "package.json")
	_, err = os.Stat(packageJSONPath)
	if os.IsNotExist(err) {
		fmt.Println("Error: Not a Node.js project (package.json not found)")
		return
	}

	// checking for React dependencies
	// (Assuming "dependencies" and "devDependencies" fields in package.json)
	reactDependencies := []string{"react", "react-dom"}
	doesHaveDependencies, err := hasDependencies(packageJSONPath, reactDependencies)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if doesHaveDependencies {
		fmt.Println("This appears to be a React project.")
	} else {
		fmt.Println("Error: Not a React project (React dependencies not found)")
		return
	}

	// checking for src directory
	srcDir := filepath.Join(wd, "src")
	_, err = os.Stat(srcDir)
	if os.IsNotExist(err) {
		fmt.Println("Error: Not a React project (src directory not found)")
		return
	}

	fmt.Println("This appears to be a React project.")
}

func hasDependencies(packageJSONPath string, dependencies []string) (bool, error) {
	file, err := ioutil.ReadFile(packageJSONPath)
	if err != nil {
		return false, err
	}

	fmt.Println(string(file))

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