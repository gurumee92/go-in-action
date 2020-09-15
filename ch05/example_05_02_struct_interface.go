package main

import "fmt"

type Notifier interface {
	Notify()
}

type Person struct {
	Name string
	Mail string
}

func (p Person) GetMail() string {
	return p.Mail
}

func (p *Person) SetMail(mail string) {
	p.Mail = mail
}

func (p *Person) Notify() {
	fmt.Println(p.Name, p.Mail, "로 전송합니다.")
}

type Admin struct {
	Person
	Level string
}

func (a *Admin) Notify() {
	fmt.Println(a.Name, a.Mail, a.Level, "로 전송합니다.[관리자]")
}

func main() {
	p := Person{
		Name: "Gurumee",
		Mail: "gurumee@example.com",
	}
	p.Notify()
	SendNotify(&p)

	a := Admin{
		Person: Person{
			Name: "Gurumee",
			Mail: "gurumee@example.com",
		},
		Level: "SUPER",
	}
	SendNotify(&a)
}

func SendNotify(n Notifier) {
	n.Notify()
}
