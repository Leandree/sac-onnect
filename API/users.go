package main
 
import (
)
 
// User struct
type User struct {
    ID        int64  `json:"Id"`
    Name  string `json:"name"`
    Weight string `json:"weight"`
    Size  string `json:"size"`
}
 
// NewUser retrun a pointer to a User
func NewUser() *User {
    return new(User)
}
 
 
// Save user
func (u *User) Save() error {
    db, err := GetDB()
    if err != nil {
        return err
    }
    defer db.Close()
    if err != nil {
        return err
    }
    _, err = db.Exec(
        `INSERT INTO users (name, weight, size)
     VALUES (?, ?, ?)`,
        u.Name, u.Weight, u.Size)
    if err != nil {
        return err
    }
    return nil
}
 