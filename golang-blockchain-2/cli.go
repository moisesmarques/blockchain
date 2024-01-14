package main

import (
	"flag"
	"fmt"
	"log"
	"main/shell"
	"os"
)

// CLI responsible for processing command line arguments
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  createwallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  listaddresses - Lists all addresses from the wallet file")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  reindexutxo - Rebuilds the UTXO set")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT -mine - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.")
	fmt.Println("  startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining")
	fmt.Println("  createfak -address ADDRESS - Create a FAK for wallet address")
	fmt.Println("  getwallet -address ADDRESS - Get wallet by address")
	fmt.Println("  createfaksubkey -address, -fak")
	//code added by Sameer for shell creation
	fmt.Println("  createshell -pallet PALLET - Create a shell by passing requirements pallet as json")
	//code end
	fmt.Println("  createapkeysignature -wallets, -message")
	fmt.Println("  verifyapkeysignature -wallets, -wallet, -signature, -seed, -message")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Printf("NODE_ID env. var is not set!")
		os.Exit(1)
	}

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	reindexUTXOCmd := flag.NewFlagSet("reindexutxo", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	startNodeCmd := flag.NewFlagSet("startnode", flag.ExitOnError)
	createFAKCmd := flag.NewFlagSet("createfak", flag.ExitOnError)
	getWalletCmd := flag.NewFlagSet("getWallet", flag.ExitOnError)
	createFAKSubKeyCmd := flag.NewFlagSet("createfaksubkey", flag.ExitOnError)
	createPaletteCmd := flag.NewFlagSet("createpalette", flag.ExitOnError)
	getPaletteCmd := flag.NewFlagSet("getpalette", flag.ExitOnError)
	createPaintCmd := flag.NewFlagSet("createpaint", flag.ExitOnError)
	createShellCmd := flag.NewFlagSet("createshell", flag.ExitOnError) //added by Sameer for shell creation
	createApKeySignatureCmd := flag.NewFlagSet("createapkeysignature", flag.ExitOnError)
	verifyApKeySignatureCmd := flag.NewFlagSet("verifyapkeysignature", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")
	sendMine := sendCmd.Bool("mine", false, "Mine immediately on the same node")
	startNodeMiner := startNodeCmd.String("miner", "", "Enable mining mode and send reward to ADDRESS")
	createFAKwallet := createFAKCmd.String("address", "", "The wallet address to create FAK for")
	getWalletAddress := getWalletCmd.String("address", "", "The wallet address")
	createFAKSubKeyAddress := createFAKSubKeyCmd.String("address", "", "The wallet address")
	createFAKSubKeyFak := createFAKSubKeyCmd.String("fak", "", "The Fak address")
	getPalettePaletteId := getPaletteCmd.String("paletteid", "", "The palette id")
	createPaintPaletteId := createPaintCmd.String("paletteid", "", "The palette id")
	createShellPallet := createShellCmd.String("pallet", "", "The Pallet in JSON format") //added by Sameer for shell creation

	createApKeySignatureWallets := createApKeySignatureCmd.String("wallets", "", "The list of wallets to generate the Ring")
	createApKeySignatureMessage := createApKeySignatureCmd.String("message", "", "The message to be signed")

	verifyApKeySignatureWallets := verifyApKeySignatureCmd.String("wallets", "", "The list of wallets to generate the Ring")
	verifyApKeySignatureSignature := verifyApKeySignatureCmd.String("signature", "", "The signature of the message")
	verifyApKeySignatureSeed := verifyApKeySignatureCmd.String("seed", "", "The seed value used to sign the message")
	verifyApKeySignatureMessage := verifyApKeySignatureCmd.String("message", "", "The message to be signed")

	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "reindexutxo":
		err := reindexUTXOCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "startnode":
		err := startNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createfak":
		err := createFAKCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getwallet":
		err := getWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createfaksubkey":
		err := createFAKSubKeyCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	// Code Added by Sameer for shell creation
	case "createshell":
		err := createShellCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	//Code end
	case "getpalette":
		err := getPaletteCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createpalette":
		err := createPaletteCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createpaint":
		err := createPaintCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createapkeysignature":
		err := createApKeySignatureCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "verifyapkeysignature":
		err := verifyApKeySignatureCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress, nodeID)
	}

	if createFAKCmd.Parsed() {
		if *createFAKwallet == "" {
			createFAKCmd.Usage()
			os.Exit(1)
		}
		cli.createFak(*createFAKwallet, nodeID)
	}

	if getWalletCmd.Parsed() {
		if *getWalletAddress == "" {
			getWalletCmd.Usage()
			os.Exit(1)
		}
		cli.getWallet(*getWalletAddress, nodeID)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress, nodeID)
	}

	if createWalletCmd.Parsed() {
		cli.createWallet(nodeID)
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses(nodeID)
	}

	if printChainCmd.Parsed() {
		cli.printChain(nodeID)
	}

	if reindexUTXOCmd.Parsed() {
		cli.reindexUTXO(nodeID)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount, nodeID, *sendMine)
	}

	if startNodeCmd.Parsed() {
		nodeID := os.Getenv("NODE_ID")
		if nodeID == "" {
			startNodeCmd.Usage()
			os.Exit(1)
		}
		cli.startNode(nodeID, *startNodeMiner)
	}

	if createFAKSubKeyCmd.Parsed() {
		if *createFAKSubKeyAddress == "" || *createFAKSubKeyFak == "" {
			createFAKSubKeyCmd.Usage()
			os.Exit(1)
		}

		cli.CreateFakSubKey(*createFAKSubKeyAddress, *createFAKSubKeyFak, nodeID)
	}

	if getPaletteCmd.Parsed() {

		if *getPalettePaletteId == "" {
			getPaletteCmd.Usage()
			os.Exit(1)
		}
		cli.getPalette(*getPalettePaletteId, nodeID)
	}

	if createPaletteCmd.Parsed() {
		cli.createPalette(nodeID)
	}

	if createPaintCmd.Parsed() {
		if *createPaintPaletteId == "" {
			createPaintCmd.Usage()
			os.Exit(1)
		}

		cli.createPaint(*createPaintPaletteId, nodeID)
	}

	//Code added by Sameer for shell creation
	if createShellCmd.Parsed() {
		if *createShellPallet == "" {
			createShellCmd.Usage()
			os.Exit(1)
		}

		shell.CreateShell(*createShellPallet)
	}
	//Code end

	if createApKeySignatureCmd.Parsed() {
		if *createApKeySignatureWallets == "" || *createApKeySignatureMessage == "" {
			createApKeySignatureCmd.Usage()
			os.Exit(1)
		}

		cli.createAPKeySignature(*createApKeySignatureWallets, *createApKeySignatureMessage, nodeID)
	}

	if verifyApKeySignatureCmd.Parsed() {
		if *verifyApKeySignatureWallets == "" ||
			*verifyApKeySignatureSignature == "" ||
			*verifyApKeySignatureSeed == "" ||
			*verifyApKeySignatureMessage == "" {
			verifyApKeySignatureCmd.Usage()
			os.Exit(1)
		}

		cli.verifyAPKeySignature(*verifyApKeySignatureWallets,
			*verifyApKeySignatureSignature,
			// *verifyApKeySignatureSeed,
			*verifyApKeySignatureMessage,
			nodeID)
	}
}
