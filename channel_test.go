package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
*
dalam membuat channel harus dipastikan
ada yang mengirim data ke channel dan menerina data dari channel
jika salah satunya tidak ada maka akan menyebabkan error
*/
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Dani Yudistira Maulana" // jika data dari channel ini belum diambil, maka dia akan blocking, kode dibawah tidak akan dijalankan
		fmt.Println("Selesai mengirim data ke channel")
	}()

	// data := <-channel // ini akan deadlock jika channel kosong/tidak pernah mengirim data (kondisi channel kosong)
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

	// channel <- "Dani" // cara mengirim data ke channel

	// data := <-channel // cara menerima data dari channel
	// // or
	// // var data string
	// // data = <- channel // cata menerima data dari channel
	// fmt.Println(<-channel) // menampilkan data dari channel
}

// Channel Sebagai Parameter

/*
*
prosedur (channel secara default pass by reference)
1. buat variabel dengan tipe data channel
2. kirim channel ke func GiveMeResponse yang diterima oleh paramater channel
3. channel dari parameter akan mengirim data "Dani Yudistira Maulana"
4. variable channel dari func TestChannelAsParameter berisi data "Dani Yudistira Maulana"
5. variable data dari func TestChannelAsParameter menerima data dari variable channel yang telah dikirim dari func GiveMeResponse
*/
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dani Yudistira Maulana"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Channel In dan Out
// menandai func, untuk memberitahu func tsb untuk mengirim data ke channel dan menerima data dari channel

/** func yang hanya digunakan untuk mengirim data ke channel */
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Dani Yudistira Maulana"
}

/** func yang hanya digunakan untuk menerima data dari channel */
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel

/*
*
Buffered Channel tidak bersifat blocking
jika sebuah channel menggunakan buffer baik itu capacity hanya 1 buffer
dan mengirim data ke channel, data tsb seolah-olah berada dalam buffernya(tempat penyimpanan)
sehingga jika data channel tsb tidak ada yang mengambil tidak masalah dan tidak akan menyebabkan error/deadlock.
tapi jika mengirim data melebihi capacity buffer nya, dan tidak ada yang mengambilnya makan tetap akan deadlock
*/
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Dani"
		channel <- "Veronica"
		channel <- "Dian"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")

	// channel <- "Dani"
	// channel <- "Veronica"
	// channel <- "Dian"

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// fmt.Println("Selesai")
}

// Range Channel

/*
*

	prosedur
	1. variabel channel dibuat
	2. goroutine dijalankan dan loopnya dijalankan
	3. ketika loop dijalankan perulangan pertama data akan dikirim ke channel
	4. setelah dikirim looping range channel akan menerima dan menapilkan data pertama dari channel
	5. looping terus berulang sampai ke 10
	6. dan fungsi close(channel) baru dijalankan
	7. menampilkan fmt.Println("Selesai")

	jika tidak diclose range channel akan menunggu data baru dikirim dari channel, sehingga menyebabkan deadlock
*/
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data ", data)
	}

	fmt.Println("Selesai")
}

// Select Channel

/** mengambil langsung dari beberapa channel */
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// Default Select Channel
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}
