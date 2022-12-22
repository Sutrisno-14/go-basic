package main

import (
	"fmt"
	"strings"
)

func main() {

	// Struct adalah sebuah tipe data berupa kumpulan dari berbagai macam property/field dan juga method
	type User struct {
		Name 	string
		Address string
		Age 	int
	}

	// memberikan nilai ke struct
	var user User

	user.Name = "Ridho"
	user.Address = "Sleman, Yogyakarta"
	user.Age = 22
	fmt.Println("Nama :", user.Name)
	fmt.Println("Alamat :",user.Address)
	fmt.Println("Umur :", user.Age)
	fmt.Printf("Saya %s tinggal di %s dan umur saya %d\n", user.Name, user.Address, user.Age)

	fmt.Println("================inisialisasi struct===============")
	//menginialisasi struct
	user2 := User{Name:"Dimas",Address: "Bantul, Yogya", Age: 23}
	fmt.Println(user2)
	fmt.Println(user)
	
	fmt.Println(strings.Repeat("#", 40))
	//pointer pada sebuah struct
	var user3 *User = &user2

	fmt.Println("user2 name :", user2.Name)
	fmt.Println("user3 name :", user3.Name)

	fmt.Println(strings.Repeat("#", 40))
	user3.Name = "Sinta"
	fmt.Println("user2 name :", user2.Name)
	fmt.Println("user3 name :", user3.Name)

	fmt.Println(strings.Repeat("#", 40))
	//Embedded struct -> struct mengandung tipe data struct lainnya

	type Person struct {
		Name string
		Age  int
	}
	type Employee struct {
		Division string
		Person   Person
	}

	var employee1 = Employee{}

	employee1.Person.Name = "Prasetian"
	employee1.Person.Age = 23
	employee1.Division = "Backend Developer"

	fmt.Printf("%+v \n", employee1)

	fmt.Println(strings.Repeat("#",40))
	//anonymous struct -> sebuah struct yang tidak di deklarasikan diawal sebagai sebuah tipe data
						// struct baru, melainkan langsung dideklarasikan bersamaan dengan pembuatan variabel

	type Karyawan struct {
		Nama 	string
		Alamat  string
	}

	var karyawan1 = struct {
		Karyawan Karyawan
		Umur     int
	}{}

	karyawan1.Karyawan = Karyawan{Nama: "Dimas", Alamat: "Yogya"}
	karyawan1.Umur = 23

	karyawan2 := struct {
		Karyawan Karyawan
		Umur     int
	}{
		Karyawan: Karyawan{Nama: "Reinaldo", Alamat: "Palembang"},
		Umur: 23,
	}

	fmt.Printf("karyawan1 : %+v \n", karyawan1)
	fmt.Printf("karyawan2 : %+v \n", karyawan2)

	fmt.Println(strings.Repeat("#", 40))
	//slice of struct

	type Student struct {
		Name string
		Age  int
	}

	student := []Student {
		{Name: "Ridho", Age: 23},
		{Name: "Sutrisno", Age: 23},
		{Name: "Reinaldo", Age: 23},
		{Name: "Dimas", Age: 23},
	}

	for i :=0; i < len(student); i++ {      // for biasa
		fmt.Printf("%+v \n", student[i])
	}
	fmt.Println("===================")		
	for _, v := range student {				// for range
		fmt.Printf("%+v\n", v)
	}

	fmt.Println(strings.Repeat("#", 40))
	//slice of anonymous struct
	
	student2 := []struct {       // ambil dari struct Student
		Student   Student
		Kelas     string
	}{
		{Student: Student{Name: "Maulana", Age: 23}, Kelas: "IPA 1"},
		{Student: Student{Name: "Edo", Age: 23}, Kelas: "IPA1"},
		{Student: Student{Name: "Dimas", Age: 23}, Kelas: "IPS2"},
	}

	for _, v := range student2 {
		fmt.Printf("%+v\n", v)
	}
}