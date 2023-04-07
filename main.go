package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    for {
        var pemain, dadu int
        fmt.Print("Masukkan jumlah pemain: ")
        fmt.Scan(&pemain)
        fmt.Print("Masukkan jumlah dadu: ")
        fmt.Scan(&dadu)

        var pemenang int
        var skorpemenang int
        daduPemain := make([][]int, pemain)
        skor := make([]int, pemain)
        rand.Seed(time.Now().UnixNano())
        for i := 0; i < pemain; i++ {
            daduPemain[i] = make([]int, dadu)
            for j := 0; j < dadu; j++ {
                daduPemain[i][j] = rand.Intn(6) + 1 
              
            }
        }
        
        for{
            fmt.Println(daduPemain)
            sisaPemain := len(daduPemain)
            indexGiven := 0
            for i := 0; i < sisaPemain; i++ {
                if len(daduPemain[i]) == 0 {
                    daduPemain = append(daduPemain[:i], daduPemain[i+1:]...)
                    sisaPemain--
                    i--
                }
                totalnumberSix := 0
                for j := 0; j < len(daduPemain[i]); j++ {

                        if daduPemain[i][j] == 6 {
                        daduPemain[i] = append(daduPemain[i][:j],daduPemain[i][j+1:]...)
                        skor[i]++
                        totalnumberSix++
                        j--
                    } else if daduPemain[i][j] == 1 {
                        if j == indexGiven && j != 0{
                            continue
                        }
                        pemainSelanjutnya := (i+1) % sisaPemain
                    
                        daduPemain[i] = append(daduPemain[i][:j],daduPemain[i][j+1:]...)
                        daduPemain[pemainSelanjutnya] = append(daduPemain[pemainSelanjutnya], 1)
                        j--
                        indexGiven = len(daduPemain[pemainSelanjutnya]) - totalnumberSix 
                    }
                }

            }
            // sisaDaduPemain := make([][]int, len(daduPemain))


            for i := range daduPemain {
                for j := range daduPemain[i] {
                    daduPemain[i][j] = rand.Intn(6)+1
                }
            }

            // rand.Seed(time.Now().UnixNano())
            // for i := 0; i < sisaPemain; i++ {
            //     daduPemain[i] = make([]int, dadu)
            //     for j := 0; j < dadu; j++ {
            //         daduPemain[i][j] = rand.Intn(6) + 1 // lempar dadu acak
                
            //     }
            // }
            if sisaPemain == 1 {
                fmt.Println(skor)
                for i := range skor {
                    skorpemenang = skor[0]
                    pemenang = 1
                    fmt.Printf("Skor pemain %d : %d\n", i+1, skor[i])

                    if skor[i] > skorpemenang {
                        pemenang = i + 1
                        skorpemenang = skor[i]
                    } 
                }
                break
            }
        }
        
        fmt.Printf("Pemenang adalah pemain %d dengan skor %d", pemenang, skorpemenang)


        var lanjutBermain string
        fmt.Print("\n\nApakah kamu mau lanjut bermain? (Y/N) \n\n")
        fmt.Scan(&lanjutBermain)

        if lanjutBermain != "Y" && lanjutBermain != "y" {
            fmt.Println("\n\nTerima kasih telah bermain game dadu")
            break
        }
    }

}