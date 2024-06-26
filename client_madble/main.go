package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"time"

	"layeh.com/gumble/gumble"
	_ "layeh.com/gumble/opus"
)

func main() {
	// Command line flags
    server := flag.String("server", "localhost:64738", "the server to connect to")
    username := flag.String("username", "", "the username of the client")
    password := flag.String("password", "", "the password of the server")
    insecure := flag.Bool("insecure", false, "skip server certificate verification")
    certificate := flag.String("certificate", "", "PEM encoded certificate and private key")

    flag.Parse()

  for {
    func() {
      defer func() {  
        if r := recover(); r != nil {
          fmt.Fprintf(os.Stderr, "Recovered from panic: %v\n", r)
          time.Sleep(5 * time.Second)
        }
      }()
      // Initialize
      b := Barnard{
        Config:  gumble.NewConfig(),
        Address: *server,
      }

      b.Config.Username = *username
      b.Config.Password = *password

      if *insecure {
        b.TLSConfig.InsecureSkipVerify = true
      }
      if *certificate != "" {
        cert, err := tls.LoadX509KeyPair(*certificate, *certificate)
        if err != nil {
          fmt.Fprintf(os.Stderr, "%s\n", err)
          os.Exit(1)
        }
        b.TLSConfig.Certificates = append(b.TLSConfig.Certificates, cert)
      }

      for b.start() == 1 {
        time.Sleep(5 * time.Second)
      }

      //b.Stream.StartSource()

      for b.Connected == 1{
        //b.Stream.StartSource()
        time.Sleep(5 * time.Second)
      }
    }()
  }
}

