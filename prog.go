// в файле wrapper.go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	files := []string{"setuppostgres.go", "createtables.go", "natssetup.go", "cacheservice.go", "httpserver.go", "webinterface.go"}
	//1.setuppostgres.go 2.createtables.go 3.natssetup.go 4.cacheservice.go 5.httpserver.go 6.webinterface.go

	for _, file := range files {
		cmd := exec.Command("go", "run", file)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error running %s: %s\n", file, err)
		}
	}

	fmt.Println("All processes have been executed.")
}
