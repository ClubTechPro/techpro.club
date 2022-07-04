package common

import "testing"

func TestMongoConnnect(t *testing.T){
	status, msg, _ := Mongoconnect()

	if !status {
		t.Errorf(msg)
	}
}