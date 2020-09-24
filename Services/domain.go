package Services

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)

type ServiceSetup struct {
	ChaincodeID string
	Clients     map[string]*channel.Client
	Ledgers     map[string]*ledger.Client
}

var HLservice ServiceSetup

func eventRegister(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("Chaincode registering failed: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Chaincode event received: %v\n", ccEvent)
	case <-time.After(time.Second * 50):
		return fmt.Errorf("Can't receive corresponding chaincode event according to event id (%s)", eventID)
	}
	return nil
}
