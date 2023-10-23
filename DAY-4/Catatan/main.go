package main

import "fmt"

// type User struct {
// 	ID    int
// 	Name  string
// 	Email string
// }

// func (usser User) changeName() string {
// 	return "hello " + usser.Name
// }

// type UserInterface interface {
// 	changeName() string
// }
// func structCatatan(){
// 	var user User = User{
// 		ID:    1,
// 		Name:  "Sayid",
// 		Email: "Sayid@gmail.com",
// 	}
// 	var ui UserInterface = &user
// 	// user.ID = 1
// 	// user.Name = "Sayid"
// 	// user.Email = "Sayid@gmail.com"
// 	// fmt.Println(user)
// 	fmt.Println(ui.changeName())
// 	// agee()
// }

func main() {

}
func Catatan2ErrorHandling() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("b")
	panic("ada panic")
	fmt.Println("c")
}

// func CatatanErrorhandling() {
// 	result, err := bagi(10, 0)
// 	fmt.Println(result, err)
// }

// func bagi(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("b tidak boleh 0")
// 	}
// 	return a / b, nil
// }

// func CatatanInterface(){
// 	tampil("alta")
// 	tampil(123)
// }
// func tampil(data interface{}) {
// 	// fmt.Println(data * 2)
// 	// fmt.Println(data.(int) * 2)
// 	switch data.(type) {
// 	case int:
// 		fmt.Println(data.(int) * 2)
// 	case string:
// 		fmt.Println("hello " + data.(string))

// 	}
// }

// func Catatanagee() {
// 	var age int = 17
// 	// var age []int = []int{3, 5, 8}
// 	fmt.Println(age)

// 	var ageAddress *int = &age
// 	changeAge(ageAddress)
// 	// changeAge(age)
// }
// func changeAge(ageAddress *int) {
// 	*ageAddress = 10
// 	fmt.Println(*ageAddress)
// }

// // func changeAge(ageAddress []int) {
// // 	ageAddress[0] = 10
// // 	fmt.Println(ageAddress)
// // 	// fmt.Println(*ageAddress)
// // }
