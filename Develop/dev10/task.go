package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	host    string
	port    int
	timeout time.Duration
)

var RootCmd = &cobra.Command{
	Use:   "telnet",
	Short: "Simple telnet usage example",
}

func init() {
	RootCmd.AddCommand(TelnetServerCmd)
	RootCmd.AddCommand(TelnetClientCmd)
	TelnetClientCmd.Flags().DurationVar(&timeout, "timeout", 10*time.Second, "connection timeout")
}

var TelnetClientCmd = &cobra.Command{
	Use:   "telnet_client",
	Short: "run telnet client",
	Run: func(cmd *cobra.Command, args []string) {
		err := parseArgs(args)
		if err != nil {
			log.Fatal("wrong args: ", err)
		}

		addr := fmt.Sprintf("%s:%v", host, port)
		fmt.Printf("trying %s, timeout: %s...\n", addr, timeout)
		dialer := &net.Dialer{Timeout: timeout}
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		conn, err := dialer.DialContext(ctx, "tcp", addr)
		if err != nil {
			log.Fatalf("cannot connect: %v", err)
		}
		fmt.Printf("connected to %s\n", addr)

		go readRoutine(ctx, conn, cancel)
		go writeRoutine(ctx, conn, cancel)

		select {
		case <-ctx.Done():
			fmt.Println("shutdown signal received")
		}
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("connection closed")
	},
}

func parseArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("should be host and port")
	}

	host = args[0]
	if host == "" {
		return errors.New("host should be not empty")
	}

	port, _ = strconv.Atoi(args[1])
	if port == 0 {
		return errors.New("port can't be zero")
	}

	return nil
}

func readRoutine(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	defer cancel()
	scanner := bufio.NewScanner(conn)

OUTER:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("break read")
			break OUTER
		default:
			if scanner.Err() != nil {
				fmt.Println(scanner.Err())
			}
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN from conn")
				break OUTER
			}
			text := scanner.Text()
			log.Printf("from server: %s", text)
		}
	}

	log.Printf("finished read")
}

func writeRoutine(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	defer cancel()
	scanner := bufio.NewScanner(os.Stdin)

OUTER:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("break write")
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			_, err := conn.Write([]byte(fmt.Sprintf("%s\n", str)))
			if err != nil {
				log.Printf("error on write to server: %v\n", err)
				break OUTER
			}
		}
	}

	log.Printf("finished write")
}

var TelnetServerCmd = &cobra.Command{
	Use:   "telnet_server",
	Short: "run telnet server",
	Run: func(cmd *cobra.Command, args []string) {
		l, err := net.Listen("tcp", "0.0.0.0:3302")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer l.Close()

		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatalf("cannot accept: %v", err)
			}

			go handleConnection(conn)
		}
	},
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte(fmt.Sprintf("welcome to %s, friend from %s\n", conn.LocalAddr(), conn.RemoteAddr())))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("received: %s", text)
		if text == "quit" || text == "exit" {
			break
		}

		conn.Write([]byte(fmt.Sprintf("i have received '%s'\n", text)))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error happened on connection with %s: %v", conn.RemoteAddr(), err)
	}

	log.Printf("closing connection with %s", conn.RemoteAddr())
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
