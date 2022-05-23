package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "flag"
    "bufio"
    "strings"
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/ethclient"
)

func createKs(pass string) {
    ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
    password := pass
    account, err := ks.NewAccount(password)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Wallet is created here is address =>")
    fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs(filePath string, pass string) {
    file := filePath
    ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
    jsonBytes, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }

    password := pass
    account, err := ks.Import(jsonBytes, password, password)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Account imported here is address =>")
    fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

    if err := os.Remove(file); err != nil {
        log.Fatal(err)
    }
}

func main() {
    url := flag.String("url", "https://main-rpc.linkpool.io/", "Test net url to connect.")
    importKey := flag.String("import", "", "Import key file. This must be file location")
    createKeys := flag.Bool("createKeys",false, "Create new key pair.")
    flag.Parse()
    reader := bufio.NewReader(os.Stdin)
    if len(os.Args)< 2 {
	flag.PrintDefaults()
	os.Exit(1)
    }
    if *importKey=="" {
	if *createKeys == false {
		flag.PrintDefaults()
	}else{
	  fmt.Println("Enter password for your new wallet")
	  input, err := reader.ReadString('\n')
 	  if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	  }
          input = strings.TrimSuffix(input, "\n")
	  createKs(input)
 	}
    }else{
	if *createKeys {
	 fmt.Println("You cannot use createKeys and import same time")
	 os.Exit(1)
	}else{
	
	  fmt.Println("Enter password for imported wallet")
	  input, err := reader.ReadString('\n')
 	  if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	  }
          input = strings.TrimSuffix(input, "\n")
	  importKs(*importKey,input)
	}
    }
    fmt.Println("Making connection to ", *url, " testnet")
    client, err := ethclient.Dial(*url)
    if err != nil {
                log.Fatalf("Oops! There was a problem", err)
    }else {
                fmt.Println("Success! you are connected to the ", *url, " Network")
    }
    _ = client
}
