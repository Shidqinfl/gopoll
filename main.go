package main

import (
	"flag"
	"fmt"
	Polling "gopoll/polling"
)

type Domain struct {
	Port    string
	Baud    uint
	Id      uint8
	Address uint16
	Count   uint16
}

func main() {
	id := flag.Int64("i", 1, "insert id")
	port := flag.String("p", "dev/ttyUSB0", "insert port")
	baud := flag.Uint("b", 9600, "insert baudrate")
	reg := flag.Uint("r", 1, "insert start register")
	count := flag.Uint("c", 1, "insert count register")
	flag.Parse()

	domain := Domain{}
	domain.Port = *port //windows
	// domain.Port = "/dev/ttyUSB0" // linux
	domain.Baud = *baud
	domain.Id = uint8(*id) // 44// BMS id
	domain.Address = uint16(*reg)
	domain.Count = uint16(*count)
	fmt.Println("Holding Registers 0x03")


		// fmt.Println("--------------------", i, "--------------------")
		for {
			// fmt.Println("####################", domain.Id, "###################")
			

			data, err := Polling.Poll(domain.Port, domain.Baud, domain.Id, domain.Address, domain.Count)
			if err != nil {
				// fmt.Println(err)
				fmt.Println("try again !")
			} else {
				fmt.Println("Data[", domain.Id, "] : ", data)
				// D.Id = data[1]
				// D.liquid = data[2]
				// // D.Temperature.T1 = float32(data[2] + (data[3] / 100))
				// D.Temperature.T1 = float32(data[3]) + (float32(data[4]) / 100)
				// D.Temperature.Tstate = data[5]
				// // D.Voltages.V1 = float32(data[5] + (data[6] / 100))
				// D.Voltages.V1 = float32(data[6]) + (float32(data[7]) / 100)
				// D.Voltages.Vstate = data[8]
				// fmt.Println("Device ", D.Id, ": ", data)
				// fmt.Println("Liquid 			:", D.liquid)
				// fmt.Println("Temperature 		:", D.Temperature.T1)
				// fmt.Println("Temperature State 	:", D.Temperature.Tstate)
				// fmt.Println("Voltages 		:", D.Voltages.V1)
				// fmt.Println("VOltage State 		:", D.Voltages.Vstate)
			}
		}
}
