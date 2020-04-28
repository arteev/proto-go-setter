package person6

import "testing"

func TestPerson6(t *testing.T) {
	var p interface{} = &Person6{
		Id:   12,
		Name: "Foo",
		Role: nil,
		Tags: nil,
	}

	pSetRole, ok := (p).(interface {
		SetRole(role *Role)
	})
	if !ok {
		t.Errorf("not implemented SetRole(role *Role)")
	} else {
		role := &Role{Id: 7}
		pSetRole.SetRole(role)

		person := p.(*Person6)
		if person.Role != role {
			t.Errorf("expected Person.Role %v,got %v", role, person.Role)
		}
	}
}
