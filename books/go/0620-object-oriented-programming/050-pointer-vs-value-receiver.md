Title: Pointer Vs Value receiver
Id: 27428
Score: 0
Body:
the receiver of a method is usually a pointer for performance reason because we wouldn't make a copy of the instance, as it would be the case in value receiver, this is especially true if the receiver type is a struct. anoter reason to make the receiver type a pointer would be so we could modify the data the receiver points to.

a value receiver is used to avoid modification of the data the receiver contains, a vaule receiver may cause a performance hit if the receiver is a large struct.

    package main

    type User struct {
        ID uint64
        FullName, Email string
    }

    // We do no require any special syntax to access field because receiver is a pointer
    func (user *User) SendEmail(email string) {
        fmt.Printf("Sent email to: %s\n", user.Email)
    }    

    // ChangeMail will modify the users email because the receiver type is a ponter
    func (user *User) ChangeEmail(email string) {
        user.Email = email;
    }

    func main() {
        user := User{
            1,
            "Zelalem Mekonen",
            "zola.mk.27@gmail.com",
        }

        user.SendEmail("Hello") // Sent email to: zola.mk.27@gmail.com

        user.ChangeEmail("zola@gmail.com")

        fmt.Println(user.Email) // zola@gmail.com
    }
|======|
