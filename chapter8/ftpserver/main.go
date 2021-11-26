package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintln(conn, "FTP server")
	log.SetPrefix(conn.RemoteAddr().String() + ": ")

	s := bufio.NewScanner(conn)
	for s.Scan() {
		cmd := strings.Fields(s.Text())
		if len(cmd) == 0 {
			continue
		}
		switch cmd[0] {
		case "close":
			log.Print("connection closed")
			fmt.Fprintln(conn, "connection closed")
			return
		case "ls", "pwd", "cd", "get":
			if err := execCmd(conn, cmd[0], cmd[1:]...); err != nil {
				log.Print(err)
				fmt.Fprintln(conn, err)
			}
		default:
			fmt.Fprintf(conn, "invalid command: %s\n", cmd)
		}
	}
	if err := s.Err(); err != nil {
		log.Print(err)
		return
	}
}

func execCmd(w io.Writer, cmd string, args ...string) error {
	switch cmd {
	case "cd":
		if len(args) != 1 {
			return fmt.Errorf("usage: cd [dir]")
		}
		os.Chdir(args[0])
	case "get":
		for _, file := range args {
			buf, err := os.ReadFile(file)
			if err != nil {
				return err // we could just print the errors. like we could find it but Idk =3
			}
			mustCopy(w, bytes.NewBuffer(buf))
		}
	case "ls":
		if len(args) != 0 {
			return fmt.Errorf("usage: ls")
		}
		dir, err := os.ReadDir(".")
		if err != nil {
			return err
		}
		var s []string
		for _, d := range dir {
			s = append(s, d.Name())
		}
		s = append(s, "\n")
		mustCopy(w, strings.NewReader(strings.Join(s, " ")))
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		s := []string{dir, "\n"}
		mustCopy(w, strings.NewReader(strings.Join(s, " ")))
	default:
		execCmd := exec.Command(cmd, args...)
		execCmd.Stdout = w
		if err := execCmd.Run(); err != nil {
			return err
		}
	}
	return nil
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
