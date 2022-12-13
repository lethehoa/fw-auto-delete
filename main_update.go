package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var vm_list []vm_information
var today = time.Now()

type vm_information struct {
	Full_path, powerState string
}

func main() {
	var path string
	_, err := os.Stat("log")
	if os.IsNotExist(err) {
		err = os.Mkdir("log", 0755)
	}
	create_off_list_fw()
	file_success, _ := os.OpenFile("log/log_success.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file_fail, _ := os.OpenFile("log/log_fail.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	for i := 0; i < len(vm_list); i++ {
		if strings.Compare(strings.TrimSpace(vm_list[i].powerState), "poweredOff") == 0 {
			fmt.Println("Deleting ", vm_list[i].Full_path)
			path = strings.TrimSpace(vm_list[i].Full_path)
			time.Sleep(4 * time.Second)
			_, err := exec.Command("govc", "vm.destroy", path).Output() //Delete fw - uncomment when needed
			if err == nil {
				write_log(file_success, path, "success")
			} else {
				write_log(file_fail, path, "fail")
			}
		}
		// fmt.Println(strings.TrimSpace(vm_list[i].Full_path), strings.TrimSpace(vm_list[i].powerState))
	}
	//create_fw()
}

func write_log(fileName *os.File, vm_path string, key string) {
	if strings.Compare(key, "success") == 0 {
		_, err := fileName.Write([]byte("Deleted " + vm_path + "\n"))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		_, err := fileName.Write([]byte("Cannot deleted " + vm_path + "\n"))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func create_off_list_fw() {
	out, _ := exec.Command("govc", "ls", "/VDC-Auto-Firewall/vm").Output()
	result := string(out)
	date_used := today.Format("02-01-2006")
	file, _ := os.Create("log/off-list-" + date_used + ".log")
	write_off_list_to_file(file, result)
}

func write_off_list_to_file(fileName *os.File, vm_list_returned string) {
	var s string
	scanner := bufio.NewScanner(strings.NewReader(vm_list_returned))
	for scanner.Scan() {
		full_path := scanner.Text()
		arr := strings.Split(full_path, "/")
		s = arr[3]
		m, _ := regexp.MatchString(`((?m)(\d{1,4}([.\-\/])\d{1,2}([.\-\/])\d{2,4}))$`, s) // get date
		if strings.Contains(s, "-off-") && m || strings.Contains(s, "-reinstall-") && m {
			vm_list = append(vm_list, return_a_struct_from_vm_info(full_path))
			_, err := fileName.Write([]byte(s + "\n"))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func return_a_struct_from_vm_info(full_path string) vm_information {
	out, _ := exec.Command("govc", "vm.info", full_path).Output()
	output := string(out)
	scn := bufio.NewScanner(strings.NewReader(output))
	var vm_test vm_information
	vm_test.Full_path = full_path
	for scn.Scan() {
		test := strings.Split(scn.Text(), ":")
		if strings.Compare(strings.TrimSpace(test[0]), "Power state") == 0 {
			vm_test.powerState = test[1]
		}
	}
	return vm_test
}
func create_fw() {
	_, err := exec.Command("/bin/sh", "list_fw.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
}
