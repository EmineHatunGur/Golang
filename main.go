package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Kullanıcıdan string değer alma

	reader := bufio.NewReader(os.Stdin)// bufio giriş/ çıkış işlemleri için kullanılan pakettir.NewReader ile yeni bir okuyucu oluşturmasını istedik.
	//os işletim sistemi içersindeki Stdinput ile dışarıdan gelecek olan veriyi yakalayarak uygulama içerisinde kullanmayı amaçlar.
	fmt.Print("enter stringValues : ")//Kullanıcıdan girdi alma
	str, _ := reader.ReadString('\n') //Girilen değişkenin okunup hata mesajı ile beraber dönmesi,burda hata mesajını boş bıraktık.
	fmt.Println(str)//Girdiyi ekrana yazma
	

	/* Kullanıcının girdiği değer her zaman string tir.Alınan değeri dönüştürme işlemini burda gerçekleştiririz*/

	values := bufio.NewReader(os.Stdin)
	fmt.Print("enter numbersValues : ")
	val, _ := values.ReadString('\n')
	aNumber, err := strconv.ParseFloat(strings.TrimSpace(val), 64) // Alınan giridyi strconv paketi ile floata dönüştürme,ParseFloat içinde iki
	// parametre olur bir girilen değerdir tabi ki mantıklı olan başındaki ve sonundaki boşlukları yok etmektir.Go boşlukları bir değer olarak alır ve sağlıklı bir dönüştürme olmaz.
	// diğer ifade 64 yada 32 bit olacağına karar vermek içindir.
	if err != nil { //nil null ile aynı anlama gelir.error nesnesi boş değil ise yani bir hata varsa bunu dönderir.
		fmt.Println(err)
	} else {// yoksa sonuçu ekrana yazar.
		fmt.Println("Number values: ", aNumber)
	}

}
