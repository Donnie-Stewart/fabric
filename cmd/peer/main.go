/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/internal/peer/chaincode"
	"github.com/hyperledger/fabric/internal/peer/channel"
	"github.com/hyperledger/fabric/internal/peer/common"
	"github.com/hyperledger/fabric/internal/peer/lifecycle"
	"github.com/hyperledger/fabric/internal/peer/node"
	"github.com/hyperledger/fabric/internal/peer/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
//added
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"html"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	//
// )
// type Trainer struct {
// 	Name string
// 	Age  int
// 	City string
// }
// The main command describes the service and
// defaults to printing the help message.
var mainCmd = &cobra.Command{Use: "peer"}

func main() {
	// For environment variables.
	viper.SetEnvPrefix(common.CmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Define command-line flags that are valid for all peer commands and
	// subcommands.
	mainFlags := mainCmd.PersistentFlags()

	mainFlags.String("logging-level", "", "Legacy logging level flag")
	viper.BindPFlag("logging_level", mainFlags.Lookup("logging-level"))
	mainFlags.MarkHidden("logging-level")

	cryptoProvider := factory.GetDefault()

	mainCmd.AddCommand(version.Cmd())
	mainCmd.AddCommand(node.Cmd())
	mainCmd.AddCommand(chaincode.Cmd(nil, cryptoProvider))
	mainCmd.AddCommand(channel.Cmd(nil))
	mainCmd.AddCommand(lifecycle.Cmd(cryptoProvider))

	///////////////////////////////////////////addressAutoDetect
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	//
	// http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
	// 		fmt.Fprintf(w, "Hi")
	// 			// Set client options
	// 			clientOptions := options.Client().ApplyURI("mongodb://admin:adminpw@mongodb0:27017")
	//
	// 			// Connect to MongoDB
	// 			client, err := mongo.Connect(context.TODO(), clientOptions)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	//
	// 			// Check the connection
	// 			err = client.Ping(context.TODO(), nil)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	//
	// 			fmt.Println("Connected to MongoDB!")
	//
	// 			// Get a handle for your collection
	// 			collection := client.Database("test3").Collection("trainers")
	//
	// 			// Some dummy data to add to the Database
	// 			ash := Trainer{"Ash", 10, "Pallet Town"}
	// 			misty := Trainer{"Misty", 10, "Cerulean City"}
	// 			brock := Trainer{"Brock", 15, "Pewter City"}
	//
	// 			trainers := []interface{}{ash, misty, brock}
	//
	// 			insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 			fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	// })
	//
	// log.Fatal(http.ListenAndServe(":7051", nil))

/////////////////////////////////////////

	// On failure Cobra prints the usage message and error string, so we only
	// need to exit with a non-0 status
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
