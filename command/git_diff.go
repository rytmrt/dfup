package command

import (
	"os/exec"
	"strings"
)

func GitDiff(crrentRev string, nextRev string) (modList []string, addList []string, delList []string, err error) {
	out, e := exec.Command("git", "diff", "--name-status", crrentRev, nextRev).Output()
	if e != nil {
		err = e
	}
	tmp := strings.Replace(string(out), "\n", "\t", -1)
	str := strings.Split(tmp, "\t")

	modList = make([]string, 0, len(str)/2)
	addList = make([]string, 0, len(str)/2)
	delList = make([]string, 0, len(str)/2)
	for i := 0; i < len(str); i += 2 {
		switch {
		case str[i] == "M":
			modList = append(modList, str[i+1])
		case str[i] == "A":
			addList = append(addList, str[i+1])
		case str[i] == "D":
			delList = append(delList, str[i+1])
		}
	}
	return
}
