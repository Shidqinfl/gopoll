package Polling

import (
	"errors"
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
)

func Poll(port string, baud uint, id uint8, addr uint16, count uint16) (datas []uint16, err error) {
	url := "rtu://" + port
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      url,
		Speed:    baud,
		DataBits: 8,
		Parity:   modbus.PARITY_NONE,
		StopBits: 1,
		Timeout:  1000 * time.Millisecond,
	})
	if err != nil {
		fmt.Println("[not connect]")
	}
	err = client.Open()
	if err != nil {
		fmt.Println("[can't open port]")
	}
	defer client.Close()
	client.SetUnitId(id)

	datas, err = client.ReadRegisters(addr, count, modbus.HOLDING_REGISTER)
	datas, err = client.ReadRegisters(addr, count, modbus.HOLDING_REGISTER)
	if err != nil {
		fmt.Println("[Timeout]")
		err = errors.New("[Failed to get data via modbus]")
	}
	// fmt.Println("=====================================")
	// fmt.Println("Address:", addr)
	// for i, data := range datas {
	// 	fmt.Printf("[%v] %v\n", i, data)
	// }

	// fmt.Printf("data:%v", datas)
	return datas, err
}
