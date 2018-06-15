package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main(){

	cmd := exec.Command("ls","-l")


	//Just for funning some command
	/*log.Printf("Running command and waiting for it to finish...")
	if err := cmd.Run();err!=nil{
		log.Printf("Command finished with error: %+v", err)
		return
	}
	log.Println("Command finished with no error")*/


	//For getting some resault
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

}
