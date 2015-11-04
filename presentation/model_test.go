package main

import "testing"

type User struct {
    Name string
}

func (u User) Valid() bool {
    return u.Name == "Win"
}

func TestValidUser(t *testing.T) {
    user := User{Name: "Win"}
    
    ok := user.Valid()
    if !ok {
        t.Errorf("%v should be a valid name, but was not", user.Name)
    }
}

