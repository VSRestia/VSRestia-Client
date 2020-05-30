package core

import (
	"errors"
	"fmt"
	"github.com/VSRestia/VSRestia-Client/utils"
	"net"
)

func Hello() {
	fmt.Println("Hello Restia!")
}

func Run(config map[string]interface{}) {
	//Listen for TCP requests on port 25565(default)
	LocalProxyPort := utils.Int2str(int(config["LocalProxyPort"].(float64)))
	LocalProxyAddress := "127.0.0.1:" + LocalProxyPort
	listener, listenerErr := net.Listen("tcp", LocalProxyAddress)
	if listenerErr != nil {
		fmt.Println(listenerErr)
		return
	}
	fmt.Println("Local proxy listen on:", LocalProxyAddress)
	//Endless loop for handle requests
	for {
		client, acceptErr := listener.Accept()
		if acceptErr != nil {
			fmt.Println(acceptErr)
		}
		//Start handle each request
		go func() {
			handlerErr := handleRequest(client)
			if handlerErr != nil {
				fmt.Println(handlerErr)
			}
		}()
	}

}
func handleRequest(client net.Conn) error {
	//Declare the err which will be return
	var err error

	//Check if the client is nil
	if client == nil {
		err = errors.New("client is nil")
		return err
	}

	//Show Remote address
	fmt.Println("Receive from:", client.RemoteAddr())

	//

	//Close the request
	closeErr := client.Close()
	if closeErr != nil {
		err = closeErr
	}

	//Return error
	return err
}
