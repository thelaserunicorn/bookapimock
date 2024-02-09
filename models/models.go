id title author genre price 
package models
import "gorm.io/gorm"

type Book struct {
  gorm.Model 
  Id uint `json:"id"`
  Title string `json:"title"`
  Author string `json:"author"`
  genre string `json:"genre"`
  price float32 `json:"price"`
}
