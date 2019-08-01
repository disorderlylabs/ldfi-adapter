Once you are done with the development of the adapter. The procedure to test the adapter locally is below.

Starting the server :

mixs server

```
⋊> araina1 at LM-SJC-11016345 in ~/go/src/github.com/ashutoshraina/myootadapter/mygrpcadapter (header-attempt|●11✚2) since Fri Jul 12 13:51:47
$ $GOPATH/out/darwin_amd64/release/mixs server --configStoreURL=fs:///Users/araina1/go/src/github.com/ashutoshraina/myootadapter/mygrpcadapter/testdata
Mixer started with
MaxMessageSize:  1048576
MaxConcurrentStreams:  1024
APIWorkerPoolSize:  1024
AdapterWorkerPoolSize:  1024
APIPort:  9091
APIAddress:
MonitoringPort:  15014
EnableProfiling:  true
SingleThreaded:  false
NumCheckCacheEntries:  1500000
ConfigStoreURL:  fs:///Users/araina1/go/src/github.com/ashutoshraina/myootadapter/mygrpcadapter/testdata
CertificateFile:  /etc/certs/cert-chain.pem
KeyFile:  /etc/certs/key.pem
CACertificateFile:  /etc/certs/root-cert.pem
ConfigDefaultNamespace:  istio-system
ConfigWaitTimeout:  2m0s
LoggingOptions: log.Options{OutputPaths:[]string{"stdout"}, ErrorOutputPaths:[]string{"stderr"}, RotateOutputPath:"", RotationMaxSize:104857600, RotationMaxAge:30, RotationMaxBackups:1000, JSONEncoding:false, LogGrpc:true, outputLevels:"default:info", logCallers:"", stackTraceLevels:"default:none"}
TracingOptions: tracing.Options{ZipkinURL:"", JaegerURL:"", LogTraceSpans:false, SamplingRate:0}
IntrospectionOptions: ctrlz.Options{Port:0x2694, Address:"localhost"}
UseTemplateCRDs: true
LoadSheddingOptions: loadshedding.Options{Mode:0, AverageLatencyThreshold:0, SamplesPerSecond:1.7976931348623157e+308, SampleHalfLife:1000000000, MaxRequestsPerSecond:0, BurstSize:0}
UseAdapterCRDs: true
```

Case 1 :

**Correct header**

Response on the client

```
⋊> araina1 at LM-SJC-11016345 in ~/go/src/github.com/ashutoshraina/myootadapter/mygrpcadapter (header-attempt|●11✚4) since Fri Jul 12 18:08:53
$ $GOPATH/out/darwin_amd64/release/mixc check -s destination.service="svc.cluster.local" --stringmap_attributes "request.headers=x-ebay-ldfi:abc"
2019-07-13T01:09:34.973858Z	info	parsed scheme: ""
2019-07-13T01:09:34.973905Z	info	scheme "" not registered, fallback to default scheme
2019-07-13T01:09:34.973948Z	info	ccResolverWrapper: sending update to cc: {[{localhost:9091 0  <nil>}] }
2019-07-13T01:09:34.973963Z	info	ClientConn switching balancer to "pick_first"
2019-07-13T01:09:34.974089Z	info	pickfirstBalancer: HandleSubConnStateChange: 0xc000294010, CONNECTING
2019-07-13T01:09:34.975046Z	info	pickfirstBalancer: HandleSubConnStateChange: 0xc000294010, READY
Check RPC completed successfully. Check status was OK
  Valid use count: 10000, valid duration: 1m0s
  Referenced Attributes
    context.reporter.kind ABSENCE
    destination.namespace ABSENCE
    request.headers::x-ebay-ldfi EXACT
```

Response on the mixs server
```
> araina1 at LM-SJC-11016345 in ~/go/src/github.com/ashutoshraina/myootadapter/mygrpcadapter (header-attempt|●11✚4) since Fri Jul 12 18:09:14
$ go run cmd/main.go
listening on "[::]:50376"
2019-07-13T01:09:34.976124Z	info	received request {&InstanceMsg{Subject:&SubjectMsg{User:,Groups:,Properties:map[string]*v1beta1.Value{custom_token_header: &Value{Value:&Value_StringValue{StringValue:abc,},},},},Action:nil,Name:icheck.instance.istio-system,} &Any{TypeUrl:type.googleapis.com/adapter.mygrpcadapter.config.Params,Value:[10 3 97 98 99],XXX_unrecognized:[],} 9056492982683158448}

2019-07-13T01:09:34.976162Z	info	abc
k: custom_token_header v: &Value_StringValue{StringValue:abc,}
2019-07-13T01:09:34.976183Z	info	map[custom_token_header:abc]
k: custom_token_header v: abc
2019-07-13T01:09:34.976189Z	info	success!!
```


Case 2 :

**Bad header**

Response on the client ( mixc)

```
⋊> araina1 at LM-SJC-11016345 in ~/go/src/github.com/ashutoshraina/myootadapter/mygrpcadapter (header-attempt|●11✚4) since Fri Jul 12 18:09:35
$ $GOPATH/out/darwin_amd64/release/mixc check -s destination.service="svc.cluster.local" --stringmap_attributes "request.headers=x-ebay-noldfi:abc"
2019-07-13T01:17:31.495993Z	info	parsed scheme: ""
2019-07-13T01:17:31.496042Z	info	scheme "" not registered, fallback to default scheme
2019-07-13T01:17:31.496079Z	info	ccResolverWrapper: sending update to cc: {[{localhost:9091 0  <nil>}] }
2019-07-13T01:17:31.496089Z	info	ClientConn switching balancer to "pick_first"
2019-07-13T01:17:31.496205Z	info	pickfirstBalancer: HandleSubConnStateChange: 0xc0002000b0, CONNECTING
2019-07-13T01:17:31.496243Z	info	blockingPicker: the picked transport is not ready, loop back to repick
2019-07-13T01:17:31.497213Z	info	pickfirstBalancer: HandleSubConnStateChange: 0xc0002000b0, READY
Check RPC completed successfully. Check status was PERMISSION_DENIED (h1.handler.istio-system:Unauthorized...)
  Valid use count: 0, valid duration: 0s
  Referenced Attributes
    context.reporter.kind ABSENCE
    destination.namespace ABSENCE
    request.headers::x-ebay-ldfi ABSENCE
```

In the adapter :

```
2019-07-13T01:17:31.497875Z	info	received request {&InstanceMsg{Subject:&SubjectMsg{User:,Groups:,Properties:map[string]*v1beta1.Value{custom_token_header: &Value{Value:&Value_StringValue{StringValue:,},},},},Action:nil,Name:icheck.instance.istio-system,} &Any{TypeUrl:type.googleapis.com/adapter.mygrpcadapter.config.Params,Value:[10 3 97 98 99],XXX_unrecognized:[],} 9056492982683158449}

2019-07-13T01:17:31.497896Z	info	abc
k: custom_token_header v: &Value_StringValue{StringValue:,}
2019-07-13T01:17:31.497916Z	info	map[custom_token_header:]
k: custom_token_header v:
2019-07-13T01:17:31.497932Z	info	failure; header not provided
```