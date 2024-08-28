package api

import (
	"os"
)

func Login() (user []byte, pass []byte) {
	greeting()
	user = userName()
	pass = password()
	return
}

func greeting() {
	o := os.Stdout
	mess := []byte("Welcome!" + "\n" + "Let's get started! \n")
	_, err := o.Write(mess)
	if err != nil {
		panic(err)
	}
}

func userName() []byte {
	o := os.Stdout
	mess := []byte("Please enter your username: ")
	_, err := o.Write(mess)
	if err != nil {
		panic(err)
	}
	i := os.Stdin
	name := make([]byte, 16)
	_, err = i.Read(name)
	if err != nil {
		return nil
	}
	return name
}

func password() []byte {
	o := os.Stdout
	mess := []byte("Please enter your password: ")
	_, err := o.Write(mess)
	if err != nil {
		panic(err)
	}
	i := os.Stdin
	pass := make([]byte, 16)
	_, err = i.Read(pass)
	if err != nil {
		return nil
	}
	return pass
}
