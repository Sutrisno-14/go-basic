package main

import (
	"fmt"
)


func main() {

	//Data Types in Golang
	//string -> default is string kosong
	//bool (boolea) -> default is false
	//int is same with int32 or int64
	//int, int8, int16, int32, int64
	//uint is same with uint32 or uint64
	//uint, uint8, uint16, uint32, uint64
	//byte is same with uint8
	//rune is same with int32

	var pertama uint8 = 89;
	var kedua int8 = -12;
	var float float32 = 3.65;
	var boolean bool = false

	var str string
	var in int32
	var bo bool
	var uin uint32
	var fl float32
	fmt.Println(pertama)
	fmt.Println(kedua)
	fmt.Printf("ini float: %f\n", float)
	fmt.Printf("ini flaot kedua %.4f\n", float)

	fmt.Printf("is it permited ? %t\n", boolean)

	fmt.Println(str)
	fmt.Println(in)
	fmt.Println(bo)
	fmt.Println(uin)
	fmt.Println(fl)
	fmt.Println("===============")
	fmt.Println(len("ridho"))
	fmt.Println("ridho"[0])
	fmt.Println("ridho maula"[1])

	var (
		name = "dina"
		last = "putri"
	)

	fmt.Println(name)
	fmt.Println(last)

	var (
		right = true
		left = false
	)

	var salah = right && left
	var coba = !right
	fmt.Println(salah)
	fmt.Println(coba)

	//konversi tipe of data 

	var (
		nilai32 int32 = 23768
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)
	)
	fmt.Printf("ini nilai int32: %d\n", nilai32)
	fmt.Printf("ini nilai int64 : %d\n", nilai64)
	fmt.Printf("ini nilai int16 : %d\n", nilai16)
	//konversi tipe of data string
	fmt.Println("================")
	var (
		nama = "Ridho Maulana"
		r    = nama[0]
		eString = string(r)
	)

	fmt.Println(nama)
	fmt.Println(r)
	fmt.Println(eString)

	//type data declarations
	type noKtp string
	var ktpRidho noKtp = "12837162"

	fmt.Println("ini ktp ridho : ", ktpRidho)
	fmt.Println(noKtp("9826253713"))

	//UNARY OPERATOR

	var j = 1
	j++
	j++

	fmt.Println(j)
	//
	fmt.Println("=========")
	//program operasi boolean
	var (
		nilaiAkhir = 90
		absensi = 80
		lulusNilaiAkhir bool = nilaiAkhir > 80
		lulusAbsensi bool = absensi > 80
		lulus bool = lulusNilaiAkhir && lulusAbsensi
	)
	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)
	fmt.Printf("INI TIPE DATA : %t",lulus)

	var (
		nilaia = 78
		nilaib = "B"
		nilaic = "A"
	)

	_,_,_ = nilaia, nilaib, nilaic

}