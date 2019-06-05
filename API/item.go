package main
 
import (
)
 
// User struct
type Item struct {
    ID        int64  `json:"Id"`
    Name  string `json:"name"`
    IdTag string `json:"idTag"`
    IsInBag  int64 `json:"isInBag"`
}
 
// NewUser retrun a pointer to a User
func NewItem() *Item {
    return new(Item)
}
 
 
// Update existing user
func (i *Item) Update(newItem Item) error {
    db, err := GetDB()
    if err != nil {
        return err
    }
    _, err = db.Exec(
        `UPDATE item SET isInBag = 1 WHERE idTag = ?`,
        newItem.IdTag)
    if err != nil {
        return err
    }
    return nil
}