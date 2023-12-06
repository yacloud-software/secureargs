// client create: SecureArgsServiceClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/secureargs/secureargs.proto
   gopackage : golang.conradwood.net/apis/secureargs
   importname: ai_0
   clientfunc: GetSecureArgsService
   serverfunc: NewSecureArgsService
   lookupfunc: SecureArgsServiceLookupID
   varname   : client_SecureArgsServiceClient_0
   clientname: SecureArgsServiceClient
   servername: SecureArgsServiceServer
   gsvcname  : secureargs.SecureArgsService
   lockname  : lock_SecureArgsServiceClient_0
   activename: active_SecureArgsServiceClient_0
*/

package secureargs

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_SecureArgsServiceClient_0 sync.Mutex
  client_SecureArgsServiceClient_0 SecureArgsServiceClient
)

func GetSecureArgsClient() SecureArgsServiceClient { 
    if client_SecureArgsServiceClient_0 != nil {
        return client_SecureArgsServiceClient_0
    }

    lock_SecureArgsServiceClient_0.Lock() 
    if client_SecureArgsServiceClient_0 != nil {
       lock_SecureArgsServiceClient_0.Unlock()
       return client_SecureArgsServiceClient_0
    }

    client_SecureArgsServiceClient_0 = NewSecureArgsServiceClient(client.Connect(SecureArgsServiceLookupID()))
    lock_SecureArgsServiceClient_0.Unlock()
    return client_SecureArgsServiceClient_0
}

func GetSecureArgsServiceClient() SecureArgsServiceClient { 
    if client_SecureArgsServiceClient_0 != nil {
        return client_SecureArgsServiceClient_0
    }

    lock_SecureArgsServiceClient_0.Lock() 
    if client_SecureArgsServiceClient_0 != nil {
       lock_SecureArgsServiceClient_0.Unlock()
       return client_SecureArgsServiceClient_0
    }

    client_SecureArgsServiceClient_0 = NewSecureArgsServiceClient(client.Connect(SecureArgsServiceLookupID()))
    lock_SecureArgsServiceClient_0.Unlock()
    return client_SecureArgsServiceClient_0
}

func SecureArgsServiceLookupID() string { return "secureargs.SecureArgsService" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("secureargs.SecureArgsService")
   AddService("secureargs.SecureArgsService")
}

