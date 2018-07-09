package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
	"log"
)

func check(err error){

	if err != nil {
		fmt.Printf("Error: %+v", err)
	}

}

const(
	FileName = "main.go"
)

func DecodeData(fileHandle *os.File) (map[string]string, error){


	list := make(map[string]string)
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {

		al,pa := "",""
		str := fileScanner.Text()
		flag := false

		for i := 0; i < len(str); i++{

			if flag == false && str[i] != ' '{
				al = fmt.Sprintf("%v%v",al,string(str[i]))
			}else if flag == false && str[i] == ' '{
				flag = true
			}else if flag == true{
				pa = fmt.Sprintf("%v%v",pa,string(str[i]))
			}

		}

		list[al] = pa

	}


	return list, nil
}

func main(){
	fileHandle, _ := os.Open("dat")
	defer fileHandle.Close()
	list , err := DecodeData(fileHandle)
	check(err)




	//if _, err = f.WriteString("Third row"); err != nil {
	//	fmt.Printf("Error: %+v", err)
	//}


		switch os.Args[1]{

			case "-a", "--add":
				if (len(os.Args) == 4) {
					fileHandle, _ := os.OpenFile("dat", os.O_APPEND|os.O_WRONLY, 0666)
					defer fileHandle.Close()
					writer := bufio.NewWriter(fileHandle)
					n, err := fmt.Fprintln(writer, fmt.Sprintf("%s %s", os.Args[2], os.Args[3]))
					check(err)
					err = writer.Flush()
					check(err)
					fmt.Println(n)
				}else{
					fmt.Printf("Usage: gorun %+v <alias> <path>\n For example: gorun %+v -a goland /opt/GoLand-2018.1.4/bin/goland.sh\n", FileName, FileName)
				}

			case "-o","--open":
				if len(os.Args) == 3{
					if val, ok := list[os.Args[2]]; ok {
						cmd := exec.Command(val,"")
						out, err := cmd.CombinedOutput()
						if err != nil {
									log.Fatalf("cmd.Run() failed with %s\n", err)
						}
						fmt.Printf("combined out:\n%s\n", string(out))
					}
				}

			case "-l","--list":
				for k, v := range list {
					fmt.Printf("%s   %s\n", k, v)
				}


			case "-h","--help":
			fmt.Printf("usage: gorun %+v <parameter> <value>\n" +
				"Options:\n" +
				" -o, --open   open file by alias\n" +
				" -a, --add    add new alias and path\n" +
				" -l, --list   show all aliases and relaited\n" +
				" -h, --help   display help message\n " +
				" For example: gorun %+v -a goland /opt/GoLand-2018.1.4/bin/goland.sh\n", FileName, FileName)

			default:
				fmt.Printf("Command '%+v' not found.See gorun %+v --help",os.Args[1], FileName)
		}
	if len(os.Args) == 1{
		fmt.Println("You need to enter count and then program names.")
	}

	//cmd := exec.Command("ps","aux")
	//out, err := cmd.CombinedOutput()
	//if err != nil {
	//			log.Fatalf("cmd.Run() failed with %s\n", err)
	//}
	//fmt.Printf("combined out:\n%s\n", string(out))
}
