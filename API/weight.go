package main
 
import (
    "fmt"
    "time"
)
 
// User struct
type Weight struct {
    ID        int64  `json:"Id"`
    ValueLeft  string `json:"valueLeft"`
    ValueRight string `json:"valueRight"`
}
 
// NewUser retrun a pointer to a User
func NewWeight() *Weight {
    return new(Weight)
}
 
 
// Save user
func (w *Weight) Save() error {
    db, err := GetDB()
    if err != nil {
        return err
    }
    defer db.Close()
    if err != nil {
        return err
    }
    dt := time.Now()
    fmt.Println("recu");
    _, err = db.Exec(
        `INSERT INTO weight (valueLeft, valueRight, dateWeight)
     VALUES (?, ?, ?)`,
        w.ValueLeft, w.ValueRight, dt.Format("2006-01-02 15:04:05"))
    if err != nil {
        return err
    }
    return nil
}
 