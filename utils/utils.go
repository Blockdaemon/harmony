package utils

import (
	"bytes"
	"encoding/binary"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/simple-rules/harmony-benchmark/p2p"
)

// ConvertFixedDataIntoByteArray converts an empty interface data to a byte array
func ConvertFixedDataIntoByteArray(data interface{}) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, data)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// ConvertIntoInts is to convert '1,2,3,4' into []int{1,2,3,4}.
func ConvertIntoInts(data string) []int {
	var res = []int{}
	items := strings.Split(data, ",")
	for _, value := range items {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			res = append(res, intValue)
		}
	}
	return res
}

func GetUniqueIdFromPeer(peer p2p.Peer) uint16 {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Panic("Regex Compilation Failed", "err", err)
	}
	socketId := reg.ReplaceAllString(peer.Ip+peer.Port, "") // A integer Id formed by unique IP/PORT pair
	value, _ := strconv.Atoi(socketId)
	return uint16(value)
}

// RunCmd Runs command `name` with arguments `args`
func RunCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Command running", name, args)
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Printf("Command finished with error: %v", err)
		} else {
			log.Printf("Command finished successfully")
		}
	}()
	return nil
}
