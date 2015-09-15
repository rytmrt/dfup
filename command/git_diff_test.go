package command

import (
	"testing"
)

func TestGitDiff(t *testing.T) {
	crrentRev, nextRev := "50f0c7fb3052f170d6149ac2c9bb92b4716b2d63", "4536c2538f87a7f451b96b59b59b0f7fd9a487fd"
	modRes, addRes, delRes, err := GitDiff(crrentRev, nextRev)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	modList := []string{"README.md", "command/add.go"}
	addList := []string{"command/srv_common.go", "command/svr_common_test.go"}
	delList := []string{"dup-design.toml"}

	if len(modList) == len(modRes) {
		for i, v := range modRes {
			if modList[i] != v {
				t.Errorf("ERROR: mod response error. %v : %v", modRes[i], modList[i])
			}
		}
	} else {
		t.Errorf("ERROR: modRes length is not matched. %v : %v", len(modList), len(modRes))
	}
	if len(addList) == len(addRes) {
		for i, v := range addRes {
			if addList[i] != v {
				t.Errorf("ERROR: add response error. %v : %v", addRes[i], addList[i])
			}
		}
	} else {
		t.Errorf("ERROR: addRes length is not matched. %v : %v", len(addList), len(addRes))
	}
	if len(delList) == len(delRes) {
		for i, v := range delRes {
			if delList[i] != v {
				t.Errorf("ERROR: del response error. %v : %v", delRes[i], delList[i])
			}
		}
	} else {
		t.Errorf("ERROR: delRes length is not matched. %v : %v", len(delList), len(delRes))
	}
}
