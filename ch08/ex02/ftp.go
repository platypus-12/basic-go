package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type User struct {
	name    string
	address string
	port    string
	current string
}

func main() {
	listener, err := net.Listen("tcp", "localhost:2221")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// http://srgia.com/docs/rfc959j.html
// https://www.itbook.info/network/ftp01.html
// https://atmarkit.itmedia.co.jp/fnetwork/rensai/netpro10/ftp-responsecode.html
func handleConn(c net.Conn) {
	defer c.Close()
	_, err := io.WriteString(c, "220 Service ready for new user\n")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(c)
	user := User{}
	for scanner.Scan() {
		_text := scanner.Text()
		fmt.Println(_text)
		_req := strings.SplitN(_text, " ", 2)
		command := _req[0]
		value := ""
		if len(_req) == 2 {
			value = strings.SplitN(_text, " ", 2)[1]
		}
		switch command {
		case "USER":
			user.name = value
			io.WriteString(c, "331 User name okay, need password\n")
		case "PASS":
			io.WriteString(c, "230 "+user.name+" logged in, proceed.\n")
		case "PORT":
			_val := strings.SplitN(value, ",", 6)
			user.address = _val[0] + "." + _val[1] + "." + _val[2] + "." + _val[3]
			a, _ := strconv.Atoi(_val[4])
			b, _ := strconv.Atoi(_val[5])
			// https://support.solarwinds.com/SuccessCenter/s/article/PORT-FTP-command?language=en_US
			user.port = strconv.Itoa(a*256 + b)
			fmt.Println(user)
			io.WriteString(c, "200 Port command okay.\n")
			path, _ := filepath.Abs("./")
			user.current = path
		case "RETR":
			conn, _ := net.Dial("tcp", user.address+":"+user.port)
			io.WriteString(c, "150 File status okay; about to open data connection.\n")
			fmt.Println(value)
			targetPath := filepath.Join(user.current, value)
			content, err := ioutil.ReadFile(targetPath)
			if err != nil {
				io.WriteString(c, "553 "+err.Error()+"Closing data connection.\n")
				conn.Close()
				continue
			}
			// fmt.Println(content)
			conn.Write(content)
			conn.Close()
			io.WriteString(c, "226 Closing data connection.\n")
		case "PWD":
			if user.current == "" {
				path, _ := filepath.Abs("./")
				user.current = path
			}
			// io.WriteString(c, user.current+"\n")
			// io.WriteString(c, "250 Requested file action okay, completed pwd.\n")
			io.WriteString(c, "250 " + user.current + "\n")
		case "LIST":
			conn, _ := net.Dial("tcp", user.address+":"+user.port)
			io.WriteString(c, "150 File status okay; about to open data connection.\n")
			fmt.Println(value)
			io.WriteString(c, "226 Complete ls.\n")
			conn.Close()
		case "CWD":
			if user.current == "" {
				path, _ := filepath.Abs("./")
				user.current = path
			}
			fmt.Println(value)
			dir := filepath.Join(user.current, value, "/")
			fmt.Println(dir, "dir")
			if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
				io.WriteString(c, "553 Requested action not taken.\n")
				continue
			}
			// io.WriteString(c, "225 	Data connection open; no transfer in progress.\n")
			user.current = dir
			io.WriteString(c, "250 Requested file action okay, completed cwd.\n")
		case "QUIT":
			io.WriteString(c, "Service closing control connection.\n")
			c.Close()
		default:
			io.WriteString(c, "502 Command not implemented.\n")
		}
	}

}
