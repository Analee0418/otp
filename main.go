package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	// TODO 请填充你的 secret 配置
	aliasMapping := map[string]string{
		"wallet": "xxx",
	}

	var s string
	ok := false
	alias := os.Args[1]
	if s, ok = aliasMapping[alias]; !ok {
		fmt.Println("not found ", alias)
		return
	}

	vk, err := base32.StdEncoding.DecodeString(strings.ToUpper(s))
	if err != nil {
		print(err)
		return
	}

	barray := make([]byte, 8)
	binary.BigEndian.PutUint64(barray, uint64(time.Now().Unix()/30))

	hash := hmac.New(sha1.New, vk)
	hash.Write(barray)
	h := hash.Sum(nil)

	o := (h[19] & 15)
	var header uint32
	r := bytes.NewReader(h[o : o+4])
	err = binary.Read(r, binary.BigEndian, &header)
	if err != nil {
		fmt.Println(err)
		return
	}

	h12 := (int(header) & 0x7fffffff) % 1000000
	totp := strconv.Itoa(int(h12))
	fmt.Println(totp)

	cmd := exec.Command("pbcopy")
	str, _ := cmd.StdinPipe()
	go func() {
		defer str.Close()
		io.WriteString(str, totp)
	}()
	cmd.CombinedOutput()

}
