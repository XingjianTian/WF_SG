package SDKInit

import (
	"WF_SG/Services"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/status"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	contextImpl "github.com/hyperledger/fabric-sdk-go/pkg/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
	"os"
)

const ChaincodeVersion = "1.0"

var GoPath string = os.Getenv("GOPATH")
var FabricNetWorkPath string = GoPath + "/src/FabricNetwork"

type OrgInfo struct {
	OrgName               string
	OrgAdmin              string
	OrgAdminClientContext context.ClientProvider
	OrgAdminIdentity      msp.SigningIdentity
	OrgMspClient          *mspclient.Client
	OrgResMgmt            *resmgmt.Client
	OrgChannelClient      *channel.Client
	OrgLedgerClient       *ledger.Client
	OrgChannelConfig      string
}
type MultiOrgsInfo struct {
	ChannelID       string
	ChannelConfig   string
	ChaincodeID     string
	ChaincodeGoPath string
	ChaincodePath   string
	UserName        string

	OrdererOrgAdmin      string
	OrdererOrgName       string
	OrdererClientContext context.ClientProvider
	OrdererResMgmt       *resmgmt.Client
	OrgInfos             map[string]*OrgInfo
}

//全局变量
var mi MultiOrgsInfo
var ChannelIdToConfigs map[string]string

func SetupInitInfo(channelIdSelected string) error {

	return nil

	if channelIdSelected == "" {
		channelIdSelected = "hustgym"
	}
	//create sdk
	sdkConfigFile := "conf.yaml"

	sdk, err := fabsdk.New(config.FromFile(sdkConfigFile))
	if err != nil {
		return fmt.Errorf("Fabric SDK instantiation failed: %v", err)
	}
	fmt.Println("Fabric SDK instantiation succeed")

	//builder orginfo
	builderOrgInfo := OrgInfo{
		OrgName:  "HUST",
		OrgAdmin: "Admin",

		OrgAdminClientContext: nil,
		OrgAdminIdentity:      nil,
		OrgMspClient:          nil,
		OrgResMgmt:            nil,
		OrgChannelClient:      nil,

		OrgChannelConfig: FabricNetWorkPath + "/channel-artifacts/gymHUSTanchors.tx",
	}

	builder2OrgInfo := OrgInfo{
		OrgName:  "HUST",
		OrgAdmin: "Admin",

		OrgAdminClientContext: nil,
		OrgAdminIdentity:      nil,
		OrgMspClient:          nil,
		OrgResMgmt:            nil,
		OrgChannelClient:      nil,

		OrgChannelConfig: FabricNetWorkPath + "/channel-artifacts/domHUSTanchors.tx",
	}

	//constructor orginfo
	constructorOrgInfo := OrgInfo{
		OrgName:               "zhongjian-1-ju",
		OrgAdmin:              "Admin",
		OrgAdminClientContext: nil,
		OrgAdminIdentity:      nil,
		OrgMspClient:          nil,
		OrgResMgmt:            nil,
		OrgChannelClient:      nil,
		OrgChannelConfig:      FabricNetWorkPath + "/channel-artifacts/gymzhongjian-1-juanchors.tx",
	}

	constructor2OrgInfo := OrgInfo{
		OrgName:               "zhongjian-2-ju",
		OrgAdmin:              "Admin",
		OrgAdminClientContext: nil,
		OrgAdminIdentity:      nil,
		OrgMspClient:          nil,
		OrgResMgmt:            nil,
		OrgChannelClient:      nil,
		OrgChannelConfig:      FabricNetWorkPath + "/channel-artifacts/domzhongjian-2-juanchors.tx",
	}

	//supervisor orginfo
	supervisorOrgInfo := OrgInfo{
		OrgName:               "WH-zhijianju",
		OrgAdmin:              "Admin",
		OrgAdminClientContext: nil,
		OrgAdminIdentity:      nil,
		OrgMspClient:          nil,
		OrgResMgmt:            nil,
		OrgChannelClient:      nil,
		OrgChannelConfig:      FabricNetWorkPath + "/channel-artifacts/gymWH-zhijianjuanchors.tx",
	}

	supervisor2OrgInfo := OrgInfo{
		OrgName:               "WH-zhijianju",
		OrgAdmin:              "Admin",
		OrgAdminClientContext: nil,
		OrgAdminIdentity:      nil,
		OrgMspClient:          nil,
		OrgResMgmt:            nil,
		OrgChannelClient:      nil,
		OrgChannelConfig:      FabricNetWorkPath + "/channel-artifacts/domWH-zhijianjuanchors.tx",
	}

	if channelIdSelected == "hustgym" {
		mi = MultiOrgsInfo{
			ChannelID:            channelIdSelected,
			ChannelConfig:        FabricNetWorkPath + ChannelIdToConfigs[channelIdSelected],
			ChaincodeID:          "mycc-gym",
			ChaincodeGoPath:      os.Getenv("GOPATH"),
			ChaincodePath:        "HL/Chaincode/",
			UserName:             "",
			OrdererOrgAdmin:      "Admin",
			OrdererOrgName:       "orderer.gov.com",
			OrdererClientContext: nil,
			OrdererResMgmt:       nil,
			OrgInfos: map[string]*OrgInfo{
				builderOrgInfo.OrgName:     &builderOrgInfo,
				constructorOrgInfo.OrgName: &constructorOrgInfo,
				supervisorOrgInfo.OrgName:  &supervisorOrgInfo,
			},
		}
	}
	if channelIdSelected == "hustdomitory" {
		mi = MultiOrgsInfo{
			ChannelID:            channelIdSelected,
			ChannelConfig:        FabricNetWorkPath + ChannelIdToConfigs[channelIdSelected],
			ChaincodeID:          "mycc-dom",
			ChaincodeGoPath:      os.Getenv("GOPATH"),
			ChaincodePath:        "HL/Chaincode/",
			UserName:             "",
			OrdererOrgAdmin:      "Admin",
			OrdererOrgName:       "orderer.gov.com",
			OrdererClientContext: nil,
			OrdererResMgmt:       nil,
			OrgInfos: map[string]*OrgInfo{
				builder2OrgInfo.OrgName:     &builder2OrgInfo,
				constructor2OrgInfo.OrgName: &constructor2OrgInfo,
				supervisor2OrgInfo.OrgName:  &supervisor2OrgInfo,
			},
		}
	}

	//create orderer context and client
	mi.OrdererClientContext = sdk.Context(fabsdk.WithUser(mi.OrdererOrgAdmin), fabsdk.WithOrg(mi.OrdererOrgName))
	mi.OrdererResMgmt, _ = resmgmt.New(mi.OrdererClientContext)

	//create organizations context and client
	for _, org := range mi.OrgInfos {
		org.OrgAdminClientContext = sdk.Context(fabsdk.WithUser(org.OrgAdmin), fabsdk.WithOrg(org.OrgName))
		org.OrgMspClient, _ = mspclient.New(sdk.Context(), mspclient.WithOrg(org.OrgName))
		org.OrgResMgmt, _ = resmgmt.New(org.OrgAdminClientContext)
		org.OrgAdminIdentity, _ = org.OrgMspClient.GetSigningIdentity(org.OrgAdmin)
	}

	err = createAndJoinChannel()
	if err != nil {
		return fmt.Errorf("Failed to create channel or join channel: %v", err)
	}
	err = installAndInstantiateCC()
	if err != nil {
		return fmt.Errorf("Failed to install and instantiate chaincode: %v", err)
	}

	err = createChannelClients(sdk)
	if err != nil {
		return fmt.Errorf("Failed to create channel clients: %v", err)
	}
	return nil

}

func createAndJoinChannel() error {

	//Test if channel joined already
	testBuilderPeers, err := DiscoverLocalPeers(mi.OrgInfos["HUST"].OrgAdminClientContext, 1)
	bJoined, err := IsJoinedChannel(mi.ChannelID, mi.OrgInfos["HUST"].OrgResMgmt, testBuilderPeers[0])
	if bJoined {
		return nil
	}

	var lastConfigBlock uint64

	var SignIDs []msp.SigningIdentity
	for _, org := range mi.OrgInfos {
		SignIDs = append(SignIDs, org.OrgAdminIdentity)
	}
	//create a channel/orderer
	req := resmgmt.SaveChannelRequest{
		ChannelID:         mi.ChannelID,
		ChannelConfigPath: mi.ChannelConfig,
		SigningIdentities: SignIDs,
	}
	_, err = mi.OrdererResMgmt.SaveChannel(req, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(mi.OrdererOrgName))
	if err != nil {
		return fmt.Errorf("Failed to create channel: %v", err)
	}
	fmt.Println("Channel created，")

	lastConfigBlock = WaitForOrdererConfigUpdate(mi.OrdererResMgmt, mi.ChannelID, true, lastConfigBlock)

	//for anchors
	for _, org := range mi.OrgInfos {
		req = resmgmt.SaveChannelRequest{ChannelID: mi.ChannelID,
			ChannelConfigPath: org.OrgChannelConfig,
			SigningIdentities: []msp.SigningIdentity{org.OrgAdminIdentity}}
		_, err = org.OrgResMgmt.SaveChannel(req, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(mi.OrdererOrgName))
		if err != nil {
			fmt.Println("error for anchor" + err.Error())
		}

		lastConfigBlock = WaitForOrdererConfigUpdate(org.OrgResMgmt, mi.ChannelID, true, lastConfigBlock)
	}

	//join channel
	for _, org := range mi.OrgInfos {
		err = org.OrgResMgmt.JoinChannel(mi.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(mi.OrdererOrgName))
		if err != nil {
			fmt.Println(org.OrgName + " failed to join channel: " + err.Error())
		}
	}

	fmt.Println("peers succeed in joining channel.")
	return nil
}
func IsJoinedChannel(channelID string, resMgmtClient *resmgmt.Client, peer fab.Peer) (bool, error) {
	resp, err := resMgmtClient.QueryChannels(resmgmt.WithTargets(peer))
	if err != nil {
		return false, err
	}
	for _, chInfo := range resp.Channels {
		if chInfo.ChannelId == channelID {
			return true, nil
		}
	}
	return false, nil
}
func WaitForOrdererConfigUpdate(client *resmgmt.Client, channelID string, genesis bool, lastConfigBlock uint64) uint64 {

	blockNum, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
		func() (interface{}, error) {
			chConfig, err := client.QueryConfigFromOrderer(channelID, resmgmt.WithOrdererEndpoint(mi.OrdererOrgName))
			if err != nil {
				return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), err.Error(), nil)
			}

			currentBlock := chConfig.BlockNumber()
			if currentBlock <= lastConfigBlock && !genesis {
				return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("Block number was not incremented [%d, %d]", currentBlock, lastConfigBlock), nil)
			}

			block, err := client.QueryConfigBlockFromOrderer(channelID, resmgmt.WithOrdererEndpoint(mi.OrdererOrgName))
			if err != nil {
				return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), err.Error(), nil)
			}
			if block.Header.Number != currentBlock {
				return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("Invalid block number [%d, %d]", block.Header.Number, currentBlock), nil)
			}

			return &currentBlock, nil
		},
	)

	if err != nil {
		fmt.Println(err.Error())
	}
	return *blockNum.(*uint64)
}

func DiscoverLocalPeers(ctxProvider context.ClientProvider, expectedPeers int) ([]fab.Peer, error) {
	ctx, err := contextImpl.NewLocal(ctxProvider)
	if err != nil {
		return nil, errors.Wrap(err, "error creating local context")
	}

	discoveredPeers, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
		func() (interface{}, error) {
			peers, serviceErr := ctx.LocalDiscoveryService().GetPeers()
			if serviceErr != nil {
				return nil, errors.Wrapf(serviceErr, "error getting peers for MSP [%s]", ctx.Identifier().MSPID)
			}
			if len(peers) < expectedPeers {
				return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("Expecting %d peers but got %d", expectedPeers, len(peers)), nil)
			}
			return peers, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return discoveredPeers.([]fab.Peer), nil
}

func queryInstalledCC(resMgmt *resmgmt.Client, ccName, ccVersion string, peers []fab.Peer) bool {
	installed, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
		func() (interface{}, error) {
			ok := isCCInstalled(resMgmt, ccName, ccVersion, peers)
			return &ok, nil
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return *(installed).(*bool)
}
func isCCInstalled(resMgmt *resmgmt.Client, ccName, ccVersion string, peers []fab.Peer) bool {
	installedOnAllPeers := true
	for _, peer := range peers {
		resp, err := resMgmt.QueryInstalledChaincodes(resmgmt.WithTargets(peer))
		if err != nil {
			fmt.Println(err.Error())
		}
		found := false

		if resp.Chaincodes != nil {
			for _, ccInfo := range resp.Chaincodes {
				fmt.Println("Found chaincode " + ccInfo.Name + ccInfo.Version)
				if ccInfo.Name == ccName && ccInfo.Version == ccVersion {
					found = true
					break
				}
			}
		}
		if !found {
			fmt.Println("chaincode " + ccName + " is not installed on peer: " + peer.URL())
			installedOnAllPeers = false
		}
	}
	return installedOnAllPeers
}
func queryInstantiatedCC(orgID string, resMgmt *resmgmt.Client, channelID, ccName, ccVersion string, peers []fab.Peer) bool {
	var ok bool
	instantiated, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
		func() (interface{}, error) {
			ok = isCCInstantiated(resMgmt, channelID, ccName, ccVersion, peers)
			return &ok, nil
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return *(instantiated).(*bool)
}
func isCCInstantiated(resMgmt *resmgmt.Client, channelID, ccName, ccVersion string, peers []fab.Peer) bool {
	installedOnAllPeers := true
	for _, peer := range peers {
		chaincodeQueryResponse, err := resMgmt.QueryInstantiatedChaincodes(channelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithTargets(peer))
		if err != nil {
			fmt.Println(err.Error())
		}
		found := false
		if chaincodeQueryResponse.Chaincodes != nil {
			for _, chaincode := range chaincodeQueryResponse.Chaincodes {
				fmt.Println("Found instantiated chaincode Name: " + ccName + " on peer: " + peer.URL())
				if chaincode.Name == ccName && chaincode.Version == ccVersion {
					found = true
					break
				}
			}
		}
		if !found {
			fmt.Println("chaincode " + ccName + " is not instantiated on peer: " + peer.URL())
			installedOnAllPeers = false
		}
	}
	return installedOnAllPeers
}
func installAndInstantiateCC() error {

	//Test if cc installed and instantiated

	testBuilderPeers, err := DiscoverLocalPeers(mi.OrgInfos["WH-zhijianju"].OrgAdminClientContext, 1)
	binstalled := queryInstalledCC(mi.OrgInfos["WH-zhijianju"].OrgResMgmt, mi.ChaincodeID, "1.0", testBuilderPeers)
	if !binstalled {

		fmt.Println("Starting installing chaincode......")
		// creates new go lang chaincode package
		ccPkg, err := gopackager.NewCCPackage(mi.ChaincodePath, mi.ChaincodeGoPath)
		if err != nil {
			return fmt.Errorf("Failed to create chaincode package: %v", err)
		}

		// contains install chaincode request parameters
		installCCReq := resmgmt.InstallCCRequest{Name: mi.ChaincodeID, Path: mi.ChaincodePath, Version: ChaincodeVersion, Package: ccPkg}
		// allows administrators to install chaincode onto the filesystem of a peer

		for _, org := range mi.OrgInfos {
			_, err = org.OrgResMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
			if err != nil {
				return fmt.Errorf("Failed to install chaincode on "+org.OrgName+" %v", err)
			}
		}

	}
	fmt.Println("Chaincode installed")
	binstantiated := queryInstantiatedCC(mi.OrgInfos["WH-zhijianju"].OrgName,
		mi.OrgInfos["WH-zhijianju"].OrgResMgmt,
		mi.ChannelID, mi.ChaincodeID, ChaincodeVersion, testBuilderPeers)
	if !binstantiated {

		fmt.Println("Starting instantiate chaincode......")

		//  returns a policy that requires one valid
		//ccPolicy := cauthdsl.SignedByAnyAdmin([]string{"builderMSP.admin", "constructorMSP.admin", "supervisorMSP.admin"})
		ccPolicy := cauthdsl.SignedByAnyMember([]string{"HUSTMSP", "zhongjian-1-juMSP", "WH-zhijianjuMSP"})
		//ccPolicy, _ := cauthdsl.FromString("OR ('builderMSP.member','supervisorMSP.member','constructorMSP.member')")
		//ccPolicy := cauthdsl.SignedByAnyAdmin([]string{ "Admin"})

		instantiateCCReq := resmgmt.InstantiateCCRequest{Name: mi.ChaincodeID, Path: mi.ChaincodePath, Version: ChaincodeVersion, Args: [][]byte{[]byte("init")}, Policy: ccPolicy}
		// instantiates chaincode with optional custom options (specific peers, filtered peers, timeout). If peer(s) are not specified
		_, err = mi.OrgInfos["WH-zhijianju"].OrgResMgmt.InstantiateCC(mi.ChannelID, instantiateCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
		if err != nil {
			return fmt.Errorf("Failed to instantiate chaincode: %v", err)
		}
	}

	fmt.Println("Chaincode instantiation succeed")

	return nil
}

func createChannelClients(sdk *fabsdk.FabricSDK) error {

	for _, org := range mi.OrgInfos {
		clientChannelContext := sdk.ChannelContext(mi.ChannelID, fabsdk.WithUser(mi.OrdererOrgAdmin), fabsdk.WithOrg(org.OrgName))
		var err error
		org.OrgChannelClient, err = channel.New(clientChannelContext)
		org.OrgLedgerClient, err = ledger.New(clientChannelContext)

		if err != nil {
			return fmt.Errorf("Failed to create "+org.OrgName+"'s channel client: %v", err)
		}
	}

	fmt.Println("All channel client created, it can be used to query or execute transaction.")

	//instantiate service api
	Services.HLservice.ChaincodeID = mi.ChaincodeID
	Services.HLservice.Clients = make(map[string]*channel.Client)
	for _, org := range mi.OrgInfos {
		Services.HLservice.Clients[org.OrgName] = org.OrgChannelClient
	}

	Services.HLservice.Ledgers = make(map[string]*ledger.Client)
	for _, org := range mi.OrgInfos {
		Services.HLservice.Ledgers[org.OrgName] = org.OrgLedgerClient
	}

	return nil
}
