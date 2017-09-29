package main

import (
	"fmt"
)

type Gifter interface {
	SendGifts() error
}

type Notifier interface {
	Notify() error
}

// Admin has embedded type User defined in it
type Admin struct {
	User
	Level string
}

type User struct {
	name  string
	email string
}

func (u *User) Notify() error {
	fmt.Println("Sending email to regular user " + u.email)
	return nil
}

func (a *Admin) Notify() error {
	fmt.Println("Sending email to admin" + a.User.email)
	return nil
}

//Note: Only user has implemented the gifter interface, and admin embeds this type
//so calling Notify on admin type objects will call this function only
//uncomment the below SendGifts for admin, and it will be overridden
func (u *User) SendGifts() error {
	fmt.Println("Gifts sent by user " + u.email)
	return nil
}

// func (a *Admin) SendGifts() error {
// 	fmt.Println("Gifts sent by admin " + a.email)
// 	return nil
// }

//NewUser is the factory-method to create objects for the User type
func NewUser(name string, email string) Notifier {
	return &User{
		name:  name,
		email: email,
	}
}

//NewAdmin is the factory-method to create Admin's objects without having to
//know about Admin representation, only data is known to the consumer
func NewAdmin(name string, email string) Notifier {
	return &Admin{
		User: User{
			name:  name,
			email: email,
		},
		Level: "super",
	}
}

//NewUser is the factory-method to create objects for the User type
func NewGifterUser(name string, email string) Gifter {
	return &User{
		name:  name,
		email: email,
	}
}

//NewGifterAdmin is the factory-method to create Admin's objects without having to
//know about Admin representation, only data is known to the consumer
func NewGifterAdmin(name string, email string) Gifter {
	return &Admin{
		User: User{
			name:  name,
			email: email,
		},
		Level: "super",
	}
}

const (
	TypeAdmin = "Admin"
	TypeUser  = "User"
)

//NewNotifier is the top-level factory-method which calls other factory-methods to create objects of
//the required type
func NewNotifier(name string, email string, t string) Notifier {
	switch t {
	case TypeAdmin:
		return NewAdmin(name, email)
	case TypeUser:
		return NewUser(name, email)
	default:
		return NewUser(name, email)
	}
}

//NewGifter is the top-level factory-method which calls other factory-methods to create objects of
//the required type
func NewGifter(name string, email string, t string) Gifter {
	switch t {
	case TypeAdmin:
		return NewGifterAdmin(name, email)
	case TypeUser:
		return NewGifterUser(name, email)
	default:
		return NewGifterUser(name, email)
	}
}

func main() {

	user := NewNotifier("Alok", "al2k.vzse@gmail.com", TypeAdmin)
	user.Notify()
	admin := NewNotifier("Vicky", "mail.alokewsq@gmail.com", TypeUser)
	admin.Notify()

	giftUser := NewGifter("GAlok", "gal2k.vzse@gmail.com", TypeAdmin)
	giftUser.SendGifts()
	giftAdmin := NewGifter("GVicky", "gmail.alokewsq@gmail.com", TypeUser)
	giftAdmin.SendGifts()

}
