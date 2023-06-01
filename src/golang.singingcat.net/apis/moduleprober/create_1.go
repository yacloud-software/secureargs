// client create: ModuleProberClient
/*
  Created by /srv/home/cnw/devel/go/go-tools/src/golang.conradwood.net/gotools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.singingcat.net/apis/moduleprober/moduleprober.proto
   gopackage : golang.singingcat.net/apis/moduleprober
   importname: ai_0
   clientfunc: GetModuleProber
   serverfunc: NewModuleProber
   lookupfunc: ModuleProberLookupID
   varname   : client_ModuleProberClient_0
   clientname: ModuleProberClient
   servername: ModuleProberServer
   gscvname  : moduleprober.ModuleProber
   lockname  : lock_ModuleProberClient_0
   activename: active_ModuleProberClient_0
*/

package moduleprober

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_ModuleProberClient_0 sync.Mutex
  client_ModuleProberClient_0 ModuleProberClient
)

func GetModuleProberClient() ModuleProberClient { 
    if client_ModuleProberClient_0 != nil {
        return client_ModuleProberClient_0
    }

    lock_ModuleProberClient_0.Lock() 
    if client_ModuleProberClient_0 != nil {
       lock_ModuleProberClient_0.Unlock()
       return client_ModuleProberClient_0
    }

    client_ModuleProberClient_0 = NewModuleProberClient(client.Connect(ModuleProberLookupID()))
    lock_ModuleProberClient_0.Unlock()
    return client_ModuleProberClient_0
}

func ModuleProberLookupID() string { return "moduleprober.ModuleProber" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.
