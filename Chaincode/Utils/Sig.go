package Utils

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"
)

//绝对路径1
var keyLoc = "/Users/blu3sky/Desktop/HF1.4_project/src/FabricNetwork/crypto-config/peerOrganizations"

var KeyMap = make(map[string]*KeyPair)

type KeyPair struct {
	Skpem []byte            //`json: "Skpem"`
	Pkpem []byte            //`json: "Pkpem"`
	Sk    *ecdsa.PrivateKey `json: "sk"`
	Pk    *ecdsa.PublicKey  `json: "Pk"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randGen(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}
	return string(b)
}

func init() {

	bInitForCC := true
	var err error
	if bInitForCC {
		keyMapJsonStr := "{\"Admin@HUST.builder.com\":{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ0F1UzduSHBLUTZveC8xdUwKY0NXbU1MYmVkZmNDNUJIVk9WcEVpZ0ZCQ1graFJBTkNBQVI2Q2FqVk1xSGlkSnFPU2dqbUpOM0VoSEgvT0tqWgp3NzMvVmJGS1pZbjhYcW1JclNMQi9qRzdpOWs1QlpMYk1HenFzSVVpUmxBWjdEL0J5VjYvSGRJegotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNIRENDQWNLZ0F3SUJBZ0lSQUxJcjNBSlVHVi9PSnhOSVozdXNQTDh3Q2dZSUtvWkl6ajBFQXdJd2FURUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMkoxYVd4a1pYSXVZMjl0TVJjd0ZRWURWUVFERXc1allTNWlkV2xzClpHVnlMbU52YlRBZUZ3MHlNREF4TVRBd09EUTVNREJhRncwek1EQXhNRGN3T0RRNU1EQmFNR2N4Q3pBSkJnTlYKQkFZVEFsVlRNUk13RVFZRFZRUUlFd3BEWVd4cFptOXlibWxoTVJZd0ZBWURWUVFIRXcxVFlXNGdSbkpoYm1OcApjMk52TVE4d0RRWURWUVFMRXdaamJHbGxiblF4R2pBWUJnTlZCQU1NRVVGa2JXbHVRR0oxYVd4a1pYSXVZMjl0Ck1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRWVnbW8xVEtoNG5TYWprb0k1aVRkeElSeC96aW8KMmNPOS8xV3hTbVdKL0Y2cGlLMGl3ZjR4dTR2Wk9RV1MyekJzNnJDRklrWlFHZXcvd2NsZXZ4M1NNNk5OTUVzdwpEZ1lEVlIwUEFRSC9CQVFEQWdlQU1Bd0dBMVVkRXdFQi93UUNNQUF3S3dZRFZSMGpCQ1F3SW9BZ2dUbmpkVU9nCldMNVcxNDNjRGlCVnNQNndVZldxSVN6NG04MzlERGVxcXZBd0NnWUlLb1pJemowRUF3SURTQUF3UlFJaEFNb2QKR2pnN0c2UFZjNEMvRnQyMWtHajVWRUVPWUNIeXJacVBNN2xTdlVidkFpQVB4QU1pcWUwVjgrTjkrRjduWjI2bAppOTJNU05lUHltRlNiSENxd3IxL05BPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\"}," +
			"\"Admin@zhongjian-1-ju.constructor.com\":{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ3E0eDFZVkZrUGhjR01RL2MKLzk0M0pDNTN6RzhUWVlObThwR2hqSmc3bFIraFJBTkNBQVF5eGlvNW54akRkRFFBeGwrK3luZmw5bGMzZFdTUApzRVNNT09DWGVuMkNSQnY5SU9UVndWWWZvRDU3S01ybmV4WUMxWHpVV0NGQktzZlViZ0Q0Zk9qNQotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNKekNDQWMyZ0F3SUJBZ0lRWXRRQ1h0RlEzTGtoZHhwM0t3WldhVEFLQmdncWhrak9QUVFEQWpCeE1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVlNQllHQTFVRUNoTVBZMjl1YzNSeWRXTjBiM0l1WTI5dE1Sc3dHUVlEVlFRREV4SmpZUzVqCmIyNXpkSEoxWTNSdmNpNWpiMjB3SGhjTk1qQXdNVEV3TURnME9UQXdXaGNOTXpBd01UQTNNRGcwT1RBd1dqQnIKTVFzd0NRWURWUVFHRXdKVlV6RVRNQkVHQTFVRUNCTUtRMkZzYVdadmNtNXBZVEVXTUJRR0ExVUVCeE1OVTJGdQpJRVp5WVc1amFYTmpiekVQTUEwR0ExVUVDeE1HWTJ4cFpXNTBNUjR3SEFZRFZRUUREQlZCWkcxcGJrQmpiMjV6CmRISjFZM1J2Y2k1amIyMHdXVEFUQmdjcWhrak9QUUlCQmdncWhrak9QUU1CQndOQ0FBUXl4aW81bnhqRGREUUEKeGwrK3luZmw5bGMzZFdTUHNFU01PT0NYZW4yQ1JCdjlJT1RWd1ZZZm9ENTdLTXJuZXhZQzFYelVXQ0ZCS3NmVQpiZ0Q0Zk9qNW8wMHdTekFPQmdOVkhROEJBZjhFQkFNQ0I0QXdEQVlEVlIwVEFRSC9CQUl3QURBckJnTlZIU01FCkpEQWlnQ0I0cFRncWN5UVk0dGcxUmhWaVRIeW5ja1RNeXF3Y3FZejcvamFkeUVIT1VUQUtCZ2dxaGtqT1BRUUQKQWdOSUFEQkZBaUVBc2JnZ20zeUg2SERzUGJxOXloN1lBdFhGcWJsRXRkbURYTjZQK0VFMXh6d0NJRFZ6MFUvVwp6N09WY0ZqKzVHcjBIMkhyM0pvZUN6TnlIM3ZyM2tobGRZVjIKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\"}," +
			"\"Admin@zhongjian-2-ju.constructor.com\":{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ3E0eDFZVkZrUGhjR01RL2MKLzk0M0pDNTN6RzhUWVlObThwR2hqSmc3bFIraFJBTkNBQVF5eGlvNW54akRkRFFBeGwrK3luZmw5bGMzZFdTUApzRVNNT09DWGVuMkNSQnY5SU9UVndWWWZvRDU3S01ybmV4WUMxWHpVV0NGQktzZlViZ0Q0Zk9qNQotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNKekNDQWMyZ0F3SUJBZ0lRWXRRQ1h0RlEzTGtoZHhwM0t3WldhVEFLQmdncWhrak9QUVFEQWpCeE1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVlNQllHQTFVRUNoTVBZMjl1YzNSeWRXTjBiM0l1WTI5dE1Sc3dHUVlEVlFRREV4SmpZUzVqCmIyNXpkSEoxWTNSdmNpNWpiMjB3SGhjTk1qQXdNVEV3TURnME9UQXdXaGNOTXpBd01UQTNNRGcwT1RBd1dqQnIKTVFzd0NRWURWUVFHRXdKVlV6RVRNQkVHQTFVRUNCTUtRMkZzYVdadmNtNXBZVEVXTUJRR0ExVUVCeE1OVTJGdQpJRVp5WVc1amFYTmpiekVQTUEwR0ExVUVDeE1HWTJ4cFpXNTBNUjR3SEFZRFZRUUREQlZCWkcxcGJrQmpiMjV6CmRISjFZM1J2Y2k1amIyMHdXVEFUQmdjcWhrak9QUUlCQmdncWhrak9QUU1CQndOQ0FBUXl4aW81bnhqRGREUUEKeGwrK3luZmw5bGMzZFdTUHNFU01PT0NYZW4yQ1JCdjlJT1RWd1ZZZm9ENTdLTXJuZXhZQzFYelVXQ0ZCS3NmVQpiZ0Q0Zk9qNW8wMHdTekFPQmdOVkhROEJBZjhFQkFNQ0I0QXdEQVlEVlIwVEFRSC9CQUl3QURBckJnTlZIU01FCkpEQWlnQ0I0cFRncWN5UVk0dGcxUmhWaVRIeW5ja1RNeXF3Y3FZejcvamFkeUVIT1VUQUtCZ2dxaGtqT1BRUUQKQWdOSUFEQkZBaUVBc2JnZ20zeUg2SERzUGJxOXloN1lBdFhGcWJsRXRkbURYTjZQK0VFMXh6d0NJRFZ6MFUvVwp6N09WY0ZqKzVHcjBIMkhyM0pvZUN6TnlIM3ZyM2tobGRZVjIKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\"}," +
			"\"Admin@WH-zhijianju.supervisor.com\":{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZzBBS1ZkdnJUTWVDLzRYL0IKMWh0eDlTcXMzcGRZMUJNVjc3Q1QrTFZxV0J5aFJBTkNBQVFDdXpVYW8zMUlSNW5QZjh1a0JrcDdXNVdkTXVqbAp2Qk1HQkV4SHlEKy9yVXg5QTNjMUVYb0paT3pPUHZyaTRTUURKekxxRzMrRmtaaVVTcmFJMFNXNAotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNKRENDQWNxZ0F3SUJBZ0lRRll4VUlzTEFTT08zUHRmaGcyWVhPekFLQmdncWhrak9QUVFEQWpCdk1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVhNQlVHQTFVRUNoTU9jM1Z3WlhKMmFYTnZjaTVqYjIweEdqQVlCZ05WQkFNVEVXTmhMbk4xCmNHVnlkbWx6YjNJdVkyOXRNQjRYRFRJd01ERXhNREE0TkRrd01Gb1hEVE13TURFd056QTRORGt3TUZvd2FqRUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhEekFOQmdOVkJBc1RCbU5zYVdWdWRERWRNQnNHQTFVRUF3d1VRV1J0YVc1QWMzVndaWEoyCmFYTnZjaTVqYjIwd1dUQVRCZ2NxaGtqT1BRSUJCZ2dxaGtqT1BRTUJCd05DQUFRQ3V6VWFvMzFJUjVuUGY4dWsKQmtwN1c1V2RNdWpsdkJNR0JFeEh5RCsvclV4OUEzYzFFWG9KWk96T1B2cmk0U1FESnpMcUczK0ZrWmlVU3JhSQowU1c0bzAwd1N6QU9CZ05WSFE4QkFmOEVCQU1DQjRBd0RBWURWUjBUQVFIL0JBSXdBREFyQmdOVkhTTUVKREFpCmdDQXV0T2dNZFp6UDByZEUyTGFVNitKeVpaeE5rejZGRFJDNHhSVkIxODlRQWpBS0JnZ3Foa2pPUFFRREFnTkkKQURCRkFpRUFxSWppSGRZckF1ajJiajNTQXVoZkthVXpMdytGci95TWJFNVdCdXh0WTR3Q0lHWjZUdnU5M0Z1RApjaVBqeUpZZWVaVVc2TWtDbmp3ZkhnSk5Mb3F6U1FGMwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\"}," +
			"\"User1@builder.com\":{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ0UyUkg5aVFyd3RhUzkvL1MKS0xxSGRmZ1NxazEwQUFqMlBYdk1oL2xpMkphaFJBTkNBQVRic0p1c1F5blhlUmI3dCtPc1dDVDJySEdlQTdzdApaYnJWM3pvSkxPOFl1VDZZWWdYK1FnNkFxMWxDTWQ5ZEtYRUZFOHdMWWdHT0tRM3hwV0VmdGlxTAotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNHakNDQWNHZ0F3SUJBZ0lRYzV5cDdvdjZPeW55WldjLzF6enY4REFLQmdncWhrak9QUVFEQWpCcE1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVVNQklHQTFVRUNoTUxZblZwYkdSbGNpNWpiMjB4RnpBVkJnTlZCQU1URG1OaExtSjFhV3hrClpYSXVZMjl0TUI0WERUSXdNREV4TURBNE5Ea3dNRm9YRFRNd01ERXdOekE0TkRrd01Gb3daekVMTUFrR0ExVUUKQmhNQ1ZWTXhFekFSQmdOVkJBZ1RDa05oYkdsbWIzSnVhV0V4RmpBVUJnTlZCQWNURFZOaGJpQkdjbUZ1WTJsegpZMjh4RHpBTkJnTlZCQXNUQm1Oc2FXVnVkREVhTUJnR0ExVUVBd3dSVlhObGNqRkFZblZwYkdSbGNpNWpiMjB3CldUQVRCZ2NxaGtqT1BRSUJCZ2dxaGtqT1BRTUJCd05DQUFUYnNKdXNReW5YZVJiN3QrT3NXQ1QyckhHZUE3c3QKWmJyVjN6b0pMTzhZdVQ2WVlnWCtRZzZBcTFsQ01kOWRLWEVGRTh3TFlnR09LUTN4cFdFZnRpcUxvMDB3U3pBTwpCZ05WSFE4QkFmOEVCQU1DQjRBd0RBWURWUjBUQVFIL0JBSXdBREFyQmdOVkhTTUVKREFpZ0NDQk9lTjFRNkJZCnZsYlhqZHdPSUZXdy9yQlI5YW9oTFBpYnpmME1ONnFxOERBS0JnZ3Foa2pPUFFRREFnTkhBREJFQWlBeElYWjkKbUczVlJCeFZEY0x4QmlhU0tZSkNTWE1uaGJMQUlXT0d2MGljZEFJZ0k0bVJpMWYvbkdxQVBUNTlCWUhVQTJvNwpzVDlqd1RKWWR6cm5LRStpV21jPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\"}," +
			"\"User1@zhongjian-1-ju.constructor.com\":{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ0huMzM4WThNdmZDNldhL2oKZW9QNXhGWHdPQ2NybmVFRkJNU1gxVGtaMWNXaFJBTkNBQVJUOXRyYlVMYXVudVNwQ2llYXk5R2V6QnpPZTlKRwovYlZ1d1plVCthNHJUNDJqZFJRNVo3bFVvd0tFQlNjSVpzWjdobFJ4Qy9WcmlGS1dsL0Nkbno1QQotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNKekNDQWMyZ0F3SUJBZ0lRV3d0MExvaVQ0TUlXV002R1QrdVZuVEFLQmdncWhrak9QUVFEQWpCeE1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVlNQllHQTFVRUNoTVBZMjl1YzNSeWRXTjBiM0l1WTI5dE1Sc3dHUVlEVlFRREV4SmpZUzVqCmIyNXpkSEoxWTNSdmNpNWpiMjB3SGhjTk1qQXdNVEV3TURnME9UQXdXaGNOTXpBd01UQTNNRGcwT1RBd1dqQnIKTVFzd0NRWURWUVFHRXdKVlV6RVRNQkVHQTFVRUNCTUtRMkZzYVdadmNtNXBZVEVXTUJRR0ExVUVCeE1OVTJGdQpJRVp5WVc1amFYTmpiekVQTUEwR0ExVUVDeE1HWTJ4cFpXNTBNUjR3SEFZRFZRUUREQlZWYzJWeU1VQmpiMjV6CmRISjFZM1J2Y2k1amIyMHdXVEFUQmdjcWhrak9QUUlCQmdncWhrak9QUU1CQndOQ0FBUlQ5dHJiVUxhdW51U3AKQ2llYXk5R2V6QnpPZTlKRy9iVnV3WmVUK2E0clQ0MmpkUlE1WjdsVW93S0VCU2NJWnNaN2hsUnhDL1ZyaUZLVwpsL0Nkbno1QW8wMHdTekFPQmdOVkhROEJBZjhFQkFNQ0I0QXdEQVlEVlIwVEFRSC9CQUl3QURBckJnTlZIU01FCkpEQWlnQ0I0cFRncWN5UVk0dGcxUmhWaVRIeW5ja1RNeXF3Y3FZejcvamFkeUVIT1VUQUtCZ2dxaGtqT1BRUUQKQWdOSUFEQkZBaUVBOFJDWllaTEpOcnNaRHAwOXJ1OEttVEdlYy9hZjZoWHk5RDNwajRwSGVaa0NJR0hsQXNGYgpvRVpsdmlOK29ENDB0ZXlRLzdqOHZTbmJRMUFXR2RnWWtmSEMKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\"}," +
			"\"User1@supervisor.com\":{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ1NydnUwWWo2blZ1RktvNHcKL1dKemFhQVRiMGRUV3J6bjRMTTcxc1IyOFlhaFJBTkNBQVRxVEI4Z0pRZEFSRHI2cDl2NTg2TFJleUhzQStPdQp6UXlaQTRaTGk2aWEvK3QveHZRZG56M2E4YjFOYVBaZkcvc0oxcVJsakc1TWZ3T3FUbUlLTG1lZwotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNJekNDQWNxZ0F3SUJBZ0lRUlJmS0VNd1lhaGUrNlE3VFJzSUtEVEFLQmdncWhrak9QUVFEQWpCdk1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVhNQlVHQTFVRUNoTU9jM1Z3WlhKMmFYTnZjaTVqYjIweEdqQVlCZ05WQkFNVEVXTmhMbk4xCmNHVnlkbWx6YjNJdVkyOXRNQjRYRFRJd01ERXhNREE0TkRrd01Gb1hEVE13TURFd056QTRORGt3TUZvd2FqRUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhEekFOQmdOVkJBc1RCbU5zYVdWdWRERWRNQnNHQTFVRUF3d1VWWE5sY2pGQWMzVndaWEoyCmFYTnZjaTVqYjIwd1dUQVRCZ2NxaGtqT1BRSUJCZ2dxaGtqT1BRTUJCd05DQUFUcVRCOGdKUWRBUkRyNnA5djUKODZMUmV5SHNBK091elF5WkE0WkxpNmlhLyt0L3h2UWRuejNhOGIxTmFQWmZHL3NKMXFSbGpHNU1md09xVG1JSwpMbWVnbzAwd1N6QU9CZ05WSFE4QkFmOEVCQU1DQjRBd0RBWURWUjBUQVFIL0JBSXdBREFyQmdOVkhTTUVKREFpCmdDQXV0T2dNZFp6UDByZEUyTGFVNitKeVpaeE5rejZGRFJDNHhSVkIxODlRQWpBS0JnZ3Foa2pPUFFRREFnTkgKQURCRUFpQmNsa2Y2RzVXbFIvZmNjbktmN1VybytUK2wvQXBaemFDVnF5eHF4VmtublFJZ0FZL3AzcVREQ0VjaQpLMThhT0JGZEpRUGIrajVuY1l0YUhPbEZkaE1PWkFzPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\"}," +
			"\"User1@zhongjian-2-ju.constructor.com\":{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ0huMzM4WThNdmZDNldhL2oKZW9QNXhGWHdPQ2NybmVFRkJNU1gxVGtaMWNXaFJBTkNBQVJUOXRyYlVMYXVudVNwQ2llYXk5R2V6QnpPZTlKRwovYlZ1d1plVCthNHJUNDJqZFJRNVo3bFVvd0tFQlNjSVpzWjdobFJ4Qy9WcmlGS1dsL0Nkbno1QQotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNKekNDQWMyZ0F3SUJBZ0lRV3d0MExvaVQ0TUlXV002R1QrdVZuVEFLQmdncWhrak9QUVFEQWpCeE1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVlNQllHQTFVRUNoTVBZMjl1YzNSeWRXTjBiM0l1WTI5dE1Sc3dHUVlEVlFRREV4SmpZUzVqCmIyNXpkSEoxWTNSdmNpNWpiMjB3SGhjTk1qQXdNVEV3TURnME9UQXdXaGNOTXpBd01UQTNNRGcwT1RBd1dqQnIKTVFzd0NRWURWUVFHRXdKVlV6RVRNQkVHQTFVRUNCTUtRMkZzYVdadmNtNXBZVEVXTUJRR0ExVUVCeE1OVTJGdQpJRVp5WVc1amFYTmpiekVQTUEwR0ExVUVDeE1HWTJ4cFpXNTBNUjR3SEFZRFZRUUREQlZWYzJWeU1VQmpiMjV6CmRISjFZM1J2Y2k1amIyMHdXVEFUQmdjcWhrak9QUUlCQmdncWhrak9QUU1CQndOQ0FBUlQ5dHJiVUxhdW51U3AKQ2llYXk5R2V6QnpPZTlKRy9iVnV3WmVUK2E0clQ0MmpkUlE1WjdsVW93S0VCU2NJWnNaN2hsUnhDL1ZyaUZLVwpsL0Nkbno1QW8wMHdTekFPQmdOVkhROEJBZjhFQkFNQ0I0QXdEQVlEVlIwVEFRSC9CQUl3QURBckJnTlZIU01FCkpEQWlnQ0I0cFRncWN5UVk0dGcxUmhWaVRIeW5ja1RNeXF3Y3FZejcvamFkeUVIT1VUQUtCZ2dxaGtqT1BRUUQKQWdOSUFEQkZBaUVBOFJDWllaTEpOcnNaRHAwOXJ1OEttVEdlYy9hZjZoWHk5RDNwajRwSGVaa0NJR0hsQXNGYgpvRVpsdmlOK29ENDB0ZXlRLzdqOHZTbmJRMUFXR2RnWWtmSEMKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\"}," +
			"\"User1@supervisor.com\":{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ1NydnUwWWo2blZ1RktvNHcKL1dKemFhQVRiMGRUV3J6bjRMTTcxc1IyOFlhaFJBTkNBQVRxVEI4Z0pRZEFSRHI2cDl2NTg2TFJleUhzQStPdQp6UXlaQTRaTGk2aWEvK3QveHZRZG56M2E4YjFOYVBaZkcvc0oxcVJsakc1TWZ3T3FUbUlLTG1lZwotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNJekNDQWNxZ0F3SUJBZ0lRUlJmS0VNd1lhaGUrNlE3VFJzSUtEVEFLQmdncWhrak9QUVFEQWpCdk1Rc3cKQ1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ4TU5VMkZ1SUVaeQpZVzVqYVhOamJ6RVhNQlVHQTFVRUNoTU9jM1Z3WlhKMmFYTnZjaTVqYjIweEdqQVlCZ05WQkFNVEVXTmhMbk4xCmNHVnlkbWx6YjNJdVkyOXRNQjRYRFRJd01ERXhNREE0TkRrd01Gb1hEVE13TURFd056QTRORGt3TUZvd2FqRUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhEekFOQmdOVkJBc1RCbU5zYVdWdWRERWRNQnNHQTFVRUF3d1VWWE5sY2pGQWMzVndaWEoyCmFYTnZjaTVqYjIwd1dUQVRCZ2NxaGtqT1BRSUJCZ2dxaGtqT1BRTUJCd05DQUFUcVRCOGdKUWRBUkRyNnA5djUKODZMUmV5SHNBK091elF5WkE0WkxpNmlhLyt0L3h2UWRuejNhOGIxTmFQWmZHL3NKMXFSbGpHNU1md09xVG1JSwpMbWVnbzAwd1N6QU9CZ05WSFE4QkFmOEVCQU1DQjRBd0RBWURWUjBUQVFIL0JBSXdBREFyQmdOVkhTTUVKREFpCmdDQXV0T2dNZFp6UDByZEUyTGFVNitKeVpaeE5rejZGRFJDNHhSVkIxODlRQWpBS0JnZ3Foa2pPUFFRREFnTkgKQURCRUFpQmNsa2Y2RzVXbFIvZmNjbktmN1VybytUK2wvQXBaemFDVnF5eHF4VmtublFJZ0FZL3AzcVREQ0VjaQpLMThhT0JGZEpRUGIrajVuY1l0YUhPbEZkaE1PWkFzPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\"}}"
		err = json.Unmarshal([]byte(keyMapJsonStr), &KeyMap)
		GetKeyPairFromPem()
		bInitForCC = false
	} else {
		err = GetKeyPair()
		keyMapJson, _ := json.Marshal(KeyMap)
		if keyMapJson == nil {
			fmt.Println("wrong key map")
		}
	}
	if err != nil {
		println(err)
	}

}

func GetKeyPairFromPem() {
	for _, kp := range KeyMap {
		kp.Pk = GetPubKey(kp.Pkpem)
		kp.Sk = GetPriKey(kp.Skpem)
	}
}
func GetKeyPair() error {

	dirs, err := ioutil.ReadDir(keyLoc)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	sep := string(os.PathSeparator)

	for _, tfi := range dirs {
		sdirs, _ := ioutil.ReadDir(keyLoc + sep + tfi.Name() + sep + "users")

		for _, fi := range sdirs {
			var kp KeyPair
			var userid string
			skPath := keyLoc + sep + tfi.Name() + sep + "users" + sep + fi.Name() + sep + "msp" + sep + "keystore"
			sfi, serr := ioutil.ReadDir(skPath)
			if serr != nil {
				continue
			}
			skName := sfi[0].Name()
			skPath = skPath + sep + skName
			kp.Skpem, _ = ioutil.ReadFile(skPath)
			kp.Sk = GetPriKey(kp.Skpem)

			pkPath := keyLoc + sep + tfi.Name() + sep + "users" + sep + fi.Name() + sep + "msp" + sep + "signcerts"
			pfi, perr := ioutil.ReadDir(pkPath)
			if perr != nil {
				continue
			}
			pkName := pfi[0].Name()
			pkPath = pkPath + sep + pkName
			kp.Pkpem, _ = ioutil.ReadFile(pkPath)
			kp.Pk = GetPubKey(kp.Pkpem)
			userid = fi.Name()
			KeyMap[userid] = &kp

			for k, _ := range KeyMap {
				fmt.Println("key: " + k)
			}
		}
	}
	return nil
}

//get public key
func GetPubKey(pkpem []byte) *ecdsa.PublicKey {
	block, _ := pem.Decode(pkpem)
	if block == nil {
		panic("expected pem block")
	}
	cer, _ := x509.ParseCertificate(block.Bytes)
	pk := cer.PublicKey.(*ecdsa.PublicKey)
	return pk
}

//get private key
func parsePKCS8PrivateKeyECDSA(der []byte) (*ecdsa.PrivateKey, error) {
	key, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}
	typedKey, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("key block is not of type ECDSA")
	}
	return typedKey, nil
}
func GetPriKey(skpem []byte) *ecdsa.PrivateKey {

	block, _ := pem.Decode(skpem)
	if block == nil {
		panic("expected pem block")
	}

	sk, err := parsePKCS8PrivateKeyECDSA(block.Bytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	return sk
}

/**
  use private key to encrypt text, while text is a hash type like md5 or sha1
  with random entropy (randsign) for security
  the result is r and s together after serialization, then converted to string with hex
*/
func Sign(rawmessage []byte, userid string) (string, error) {

	sk := KeyMap[userid].Sk
	text := hashtext(rawmessage)
	//random entropy
	randSign := randGen(36) //至少36位

	r, s, err := ecdsa.Sign(strings.NewReader(randSign), sk, text)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	fmt.Println("user: " + userid + " signed")
	return hex.EncodeToString(b.Bytes()), nil
}

//fro string signature, gets r and s
func getSign(signature string) (rint, sint big.Int, err error) {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("decrypt error, " + err.Error())
		return
	}
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err != nil {
		err = errors.New("decode error," + err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	count, err := r.Read(buf)
	if err != nil {
		fmt.Println("decode = ", err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]), "+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, " + err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, " + err.Error())
		return
	}
	return

}

func Verify(text []byte, signature string, userid string) (bool, error) {

	pk := KeyMap[userid].Pk
	rint, sint, err := getSign(signature)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(pk, hashtext(text), &rint, &sint)
	if result {
		fmt.Println("user: " + userid + " verified ok")
	} else {
		fmt.Println("user " + userid + " fail to verify")
	}
	return result, nil

}

//hash encryption, using sha1

func hashtext(text []byte) []byte {
	//random salt for hash
	salt := "131ilzaw"

	sha1Inst := sha1.New()
	sha1Inst.Write(text)
	result := sha1Inst.Sum([]byte(salt))

	return result
}
