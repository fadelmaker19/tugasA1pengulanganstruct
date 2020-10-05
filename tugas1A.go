package main

import "fmt"

type sistemoperasi struct{
    jenis string;
    seri string;
    tahunrilis int;
}

func main(){
 var kumpulan []sistemoperasi

    kumpulan = []sistemoperasi{
     sistemoperasi{
         jenis: "windows",
         seri : "windows 1.0",
         tahunrilis: 1985,
     },
     sistemoperasi{
        jenis: "linux",
        seri : "ubuntu 14",
        tahunrilis: 1991,
     },
     sistemoperasi{
        jenis: "macOS",
        seri : "Lion 2011",
        tahunrilis: 2011,
     },
 }

 fmt.Println("-------Data Sistem Operasi--------")
 fmt.Println(kumpulan[0])
 fmt.Println(kumpulan[1])
 fmt.Println(kumpulan[2])

 
}

 
 
