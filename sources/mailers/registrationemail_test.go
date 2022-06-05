package mailers

import "testing"


func TestRegistrationEmail(t *testing.T){
	got := RegistrationEmail("chilly5476@gmail.com", "Unit Testing")

	if !got{
		t.Errorf("Registration email not working")
	}
}