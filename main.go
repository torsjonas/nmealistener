package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	nmea "github.com/adrianmo/go-nmea"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			if err != nil {
				log.Fatal(err)
			}
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(bufio.NewReader(conn))
	for {
		hasMore := scanner.Scan()
		if !hasMore {
			conn.Close()
			err := scanner.Err()
			if err != nil {
				log.Fatal(err)
			}

			break
		}

		printNmea(scanner.Text())
	}
}

func printNmea(line string) {
	s, err := nmea.Parse(line)
	if err != nil {
		log.Fatal(err)
	}
	m := s.(nmea.GPRMC)
	fmt.Printf("Raw sentence: %v\n", m)
	fmt.Printf("Time: %s\n", m.Time)
	fmt.Printf("Validity: %s\n", m.Validity)
	fmt.Printf("Latitude GPS: %s\n", nmea.FormatGPS(m.Latitude))
	fmt.Printf("Latitude DMS: %s\n", nmea.FormatDMS(m.Latitude))
	fmt.Printf("Longitude GPS: %s\n", nmea.FormatGPS(m.Longitude))
	fmt.Printf("Longitude DMS: %s\n", nmea.FormatDMS(m.Longitude))
	fmt.Printf("Speed: %f\n", m.Speed)
	fmt.Printf("Course: %f\n", m.Course)
	fmt.Printf("Date: %s\n", m.Date)
	fmt.Printf("Variation: %f\n", m.Variation)
}
