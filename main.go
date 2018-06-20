package main

import (
	"fmt"
	"os"
)


const(
	FileName = "main.go"
)


func main(){
	f, err := os.OpenFile("dat", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}

	defer f.Close()

	//if _, err = f.WriteString("Third row"); err != nil {
	//	fmt.Printf("Error: %+v", err)
	//}
	if len(os.Args) == 4{

		switch os.Args[1]{

			case "-a", "--add":

				str := fmt.Sprintf("\n%s %s", os.Args[2], os.Args[3])

				n, err := f.WriteString(str)
				if err != nil {
					fmt.Printf("Error: %+v", err)
				}
				fmt.Println(n)

			case "-o","--open":

			case "-h","--help":
			fmt.Printf("Usage: gorun %+v <parameter> <value>\n For example: gorun %+v -o goland /opt/GoLand-2018.1.4/bin/goland.sh", FileName, FileName)

			default:
				fmt.Printf("Command '%+v' not found.See gorun %+v --help",os.Args[1], FileName)
		}
	}else if len(os.Args) == 1{
		fmt.Println("You need to enter count and then program names.")
	}


	//args := []string{os.Args[1]}
	//cmd := exec.Command("ps", args...)
	//fmt.Println(cmd)

	//Just for funning some command
	/*log.Printf("Running command and waiting for it to finish...")
	if err := cmd.Run();err!=nil{
		log.Printf("Command finished with error: %+v", err)
		return
	}
	log.Println("Command finished with no error")*/


	//For getting some resault
	//out, err := cmd.Output()
	//
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//
	//print(string(out))

}
