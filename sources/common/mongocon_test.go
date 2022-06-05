package common

import "testing"

func TestMongoConnnect(t *testing.T){
	_, got := Mongoconnect()

	if !got {
		t.Errorf("Connection to mongodb failed")
	}
}