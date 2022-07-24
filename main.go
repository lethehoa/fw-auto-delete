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
	create_off_list()
	file_success, _ := os.OpenFile("log/log_success.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file_fail, _ := os.OpenFile("log/log_fail.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	for i := 0; i < len(vm_list); i++ {
		if strings.Compare(strings.TrimSpace(vm_list[i].powerState), "poweredOff") == 0 {
			fmt.Println("Deleting ", vm_list[i].Full_path)
			path = strings.TrimSpace(vm_list[i].Full_path)
			time.Sleep(6 * time.Second)
			_, err := exec.Command("govc", "vm.destroy", path).Output()
			if err == nil {
				write_log(file_success, path, "success")
			} else {
				write_log(file_fail, path, "fail")
			}
		}
		fmt.Println(strings.TrimSpace(vm_list[i].Full_path), strings.TrimSpace(vm_list[i].powerState))
	}
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

func return_list_fw() string {
	out, _ := exec.Command("govc", "ls", "/VDC-Auto-Firewall/vm").Output()
	return string(out)
}

func create_off_list() {
	var f *os.File
	f, _ = os.Open("log/fw_list.log")
	write_off_list_to_file(f)
}

func write_off_list_to_file(fileName *os.File) {
	var s, vm_date_time string
	date_used := today.Format("02-01-2006")
	file, _ := os.Create("log/off-list-" + date_used + ".log")
	scanner := bufio.NewScanner(strings.NewReader(return_list_fw()))

	for scanner.Scan() {
		full_path := scanner.Text()
		arr := strings.Split(full_path, "/")
		s = arr[3]
		m, _ := regexp.MatchString(`(?m)(\d{1,4}([.\-\/])\d{1,2}([.\-\/])\d{1,4})`, s) // get date
		if strings.Contains(s, "-off-") && m {
			vm_date_time = return_date_time(s) //check if vm_path end with date
			if vm_date_time != "not_valid" {
				if compare_date_time_with_current(vm_date_time) { //delete if day off > 2 days
					vm_list = append(vm_list, return_a_struct_from_vm_info(full_path))
					_, err := file.Write([]byte(s + "\n"))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}
}

func return_date_time(vm_path string) string {
	var re = regexp.MustCompile(`(?m)(\d{1,4}([.\-\/])\d{1,2}([.\-\/])\d{1,4})`) //return date-time
	var str_final = re.FindString(vm_path)
	str_final_array := strings.Split(str_final, "-")
	if !strings.HasSuffix(vm_path, str_final) {
		return "not_valid"
	}
	if strings.Contains(str_final, "-22") || strings.Contains(str_final, "-2") { //fix pattern
		str_final_array[2] = "2022"
		str_final = strings.Join(str_final_array[:], "-")
		return str_final
	}
	return str_final

}

func compare_date_time_with_current(date string) bool {
	today := time.Now()
	t, err := time.Parse("2-1-2006", date)
	if err != nil {
		fmt.Println(err)
	}
	if today.YearDay()-t.YearDay() > 2 {
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
