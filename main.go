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

type vm_information struct {
	Full_path, powerState string
}

// import "govc"

func main() {
	var path string
	write_to_file()
	create_off_list()
	file_success, _ := os.OpenFile("log/log_success.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file_fail, _ := os.OpenFile("log/log_fail.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	for i := 0; i < len(vm_list); i++ {
		if strings.Compare(strings.TrimSpace(vm_list[i].powerState), "poweredOff") == 0 {
			// fmt.Println("Deleting ", vm_list[i].Full_path)
			path = strings.TrimSpace(vm_list[i].Full_path)
			time.Sleep(10 * time.Second)
			_, err := exec.Command("govc", "vm.info", path).Output()
			if err != nil {
				fmt.Println(err)
				write_fail_log(file_fail, path)
			} else {
				write_sucess_log(file_success, path)
			}
		}
	}
}

func write_sucess_log(fileName *os.File, success_vm_path string) {
	_, err := fileName.Write([]byte("Deleted " + success_vm_path + "\n"))
	if err != nil {
		fmt.Println(err)
	}
}

func write_fail_log(fileName *os.File, fault_vm_path string) {
	_, err := fileName.Write([]byte("Cannot deleted " + fault_vm_path + "\n"))
	if err != nil {
		fmt.Println(err)
	}
}

func write_to_file() {
	// fmt.Scanln(&input_from_user)
	var f *os.File
	out, _ := exec.Command("govc", "ls", "/VDC-Auto-Firewall/vm").Output()
	output := string(out)
	f, _ = os.Create("log/fw_list.txt")
	_, err := f.WriteString(output)
	if err != nil {
		fmt.Println("error")
	}
	// for i := 0; i < len(datacenters); i++ {
	// 	out, _ := exec.Command("govc", "ls", datacenters[i]).Output()
	// 	output := string(out)
	// 	switch i {
	// 	case 0:
	// 		f, _ = os.Create("v2_list.txt")
	// 	case 1:
	// 		f, _ = os.Create("v3_list.txt")
	// 	case 2:
	// 		f, _ = os.Create("v4_list.txt")
	// 	case 3:
	// 		f, _ = os.Create("fw_list.txt")
	// 	default:
	// 		fmt.Println("Wrong input")
	// 		return
	// 	}
	// 	_, err := f.WriteString(output)
	// 	if err != nil {
	// 		fmt.Println("error")
	// 	}
	// }
}

func create_off_list() {
	var f *os.File
	f, _ = os.Open("log/fw_list.txt")
	write_off_list_to_file(f)
	// for i := 0; i < 4; i++ {
	// 	switch i {
	// 	case 0:
	// 		f, _ = os.Open("v2_list.txt")
	// 		write_off_list_to_file(f)
	// 	case 1:
	// 		f, _ = os.Open("v3_list.txt")
	// 		write_off_list_to_file(f)
	// 	case 2:
	// 		f, _ = os.Open("v4_list.txt")
	// 		write_off_list_to_file(f)
	// 	case 3:
	// 		f, _ = os.Open("fw_list.txt")
	// 		write_off_list_to_file(f)
	// 	default:
	// 		fmt.Println("Wrong")
	// 		return
	// 	}
	// }

}

// func read_from_file
func write_off_list_to_file(fileName *os.File) {
	var s, vm_date_time string
	today := time.Now().Format("02-01-2006")
	file, _ := os.Create("log/off-list-" + today + ".txt")
	// file, _ := os.OpenFile("off-list-"+today+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	scanner := bufio.NewScanner(fileName)

	for scanner.Scan() {
		full_path := scanner.Text()
		arr := strings.Split(full_path, "/")
		s = arr[3]
		m, _ := regexp.MatchString(`(?m)(\d{1,4}([.\-\/])\d{1,2}([.\-\/])\d{1,4})`, s) // check for date - time
		if strings.Contains(s, "-off-") && m {
			vm_date_time = return_date_time(s) //check if vm_path end with date
			if vm_date_time != "not_valid" {
				// if compare_date_time_with_current(vm_date_time) {
				vm_list = append(vm_list, return_a_struct_from_vm_info(full_path))
				// fmt.Println(full_path)
				_, err := file.Write([]byte(s + "\n"))
				if err != nil {
					log.Fatal(err)
				}
				// }
			}
		}
	}
}

func return_date_time(vm_path string) string {
	var re = regexp.MustCompile(`(?m)(\d{1,4}([.\-\/])\d{1,2}([.\-\/])\d{1,4})`) //return date-time
	var str_final = re.FindString(vm_path)
	if !strings.HasSuffix(vm_path, str_final) {
		return "not_valid"
	}
	if strings.Contains(str_final, "-22") {
		str_final_1 := strings.Replace(str_final, "22", "2022", -1)
		return str_final_1
	}
	return str_final

}

func compare_date_time_with_current(date string) bool {
	today := time.Now()
	t, err := time.Parse("2-1-2006", date)
	if err != nil {
		fmt.Println(err)
	}
	if today.YearDay()-t.YearDay() > 7 {
		return true // xoa
	}
	return false // khong xoa

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
