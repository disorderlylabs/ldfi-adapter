// nolint:lll
// Generates the mygrpcadapter adapter's resource yaml. It contains the adapter's configuration, name, supported template
// names (metric in this case), and whether it is session or no-session based.
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -a mixer/adapter/mygrpcadapter/config/config.proto -x "-s=false -n mygrpcadapter -t authorization"

package mygrpcadapter

import (
	"context"
	"fmt"
	"github.com/ashutoshraina/myootadapter/mygrpcadapter/config"
	"net"
    "time"
    "strings"
    "strconv"
    "sync"

	"google.golang.org/grpc"

	"istio.io/api/mixer/adapter/model/v1beta1"
	policy "istio.io/api/policy/v1beta1"
	"istio.io/istio/mixer/pkg/status"
	"istio.io/istio/mixer/template/authorization"
	"istio.io/pkg/log"
)

type (
	// Server is basic server interface
	Server interface {
		Addr() string
		Close() error
		Run(shutdown chan error)
	}

	// MyGrpcAdapter supports metric template.
	MyGrpcAdapter struct {
		listener net.Listener
		server   *grpc.Server
	}
)

var _ authorization.HandleAuthorizationServiceServer = &MyGrpcAdapter{}


var (
    lineage = make(map[string][]string)
    lineageLock = sync.RWMutex{}
)

// HandleMetric records metric entries
//func (s *MyGrpcAdapter) HandleAuthorization(ctx context.Context, r *authorization.HandleAuthorizationRequest) (*v1beta1.CheckResult, error) {
func (s *MyGrpcAdapter) HandleAuthorization(ctx context.Context, r *authorization.HandleAuthorizationRequest) (*v1beta1.CheckResult, error) {

	log.Infof("received request %v\n", *r)

	cfg := &config.Params{}

	if r.AdapterConfig != nil {
		if err := cfg.Unmarshal(r.AdapterConfig.Value); err != nil {
			log.Errorf("error unmarshalling adapter config: %v", err)
			return nil, err
		}
	}

	decodeValue := func(in interface{}) interface{} {
		switch t := in.(type) {
		case *policy.Value_StringValue:
			return t.StringValue
		case *policy.Value_Int64Value:
			return t.Int64Value
		case *policy.Value_DoubleValue:
			return t.DoubleValue
		default:
			return fmt.Sprintf("%v", in)
		}
	}

	decodeValueMap := func(in map[string]*policy.Value) map[string]interface{} {
		out := make(map[string]interface{}, len(in))
		for k, v := range in {
			out[k] = decodeValue(v.GetValue())
			//fmt.Println("k:", k, "v:", v.GetValue())
		}
		return out
	}

	//log.Infof(cfg.AuthKey)

	props := decodeValueMap(r.Instance.Subject.Properties)
	//log.Infof("%v", props)

    lineageSoFar := strings.Split(props["x-req"].(string), "--")

    jaeger_id := lineageSoFar[0] 
    log.Infof("jaeger id: %s", jaeger_id)
    
    lineageLock.Lock()
    soFar, ok := lineage[jaeger_id]
    lineageLock.Unlock()

    // rethink what is going on here
    if (len(lineageSoFar) > 2) {
        //for _, i := range lineageSoFar[:1] {
        //log.Infof("len is %d, with string %s", len(lineageSoFar), props["x-req"].(string))
        //log.Infof("len is %d, with string %s", len(lineageSoFar), props["x-req"].(string))
        for _, i := range lineageSoFar[:2] {
            log.Infof("OKKKKK %s, len %d", i, len(lineageSoFar))
            // these will be separated by commas
            for _, command := range strings.Split(i, ",") {
                log.Infof("COMMAND: %s", command)

                pair := strings.Split(command, "=")
                if (pair[0] == "after") {
                    log.Infof("AFTER")
                    pairs2 := strings.Split(pair[1], "|")
                    anchor, target, fault, param := pairs2[0], pairs2[1], pairs2[2], pairs2[3]
                    //anchor, target := pairs2[0], pairs2[1]
                    if (target == props["destination_svc"] && ok) {
                        log.Infof("IN the zone")
                        for _, sf := range soFar {
                            if (sf == anchor) {
                                // dying on purpose
                                log.Infof("Dying on %s", sf)
                                //return &v1beta1.CheckResult{
                                //        Status: status.WithPermissionDenied("Unauthorized..."),
                                //}, nil
                                if (fault == "E") {
                                    return &v1beta1.CheckResult{
                                        Status: status.WithPermissionDenied("Unauthorized..."),
                                    }, nil
                                } else if (fault == "D") {
                                    if val, err := strconv.Atoi(param); err == nil {
                                        time.Sleep(time.Duration(val) * time.Second)
                                    } else {
                                        log.Infof("%s did not scan as a timeout value", param)
                                    }
                                } else {
                                    log.Infof("Unknown fault type %s", fault)
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    if (ok) {
        log.Infof("so far: %s", strings.Join(soFar, ", "))
        /*
        for _, i := range soFar {
            // cheap stopgap trick: if this is our second visit to a particular service...
            if (i == props["destination_svc"] && i != "productpage") {
                // dying on purpose
                log.Infof("Dying on %s", i)
                return &v1beta1.CheckResult{
                    Status: status.WithPermissionDenied("Unauthorized..."),
                }, nil
            }
        }
        */
    } else {
        log.Infof("EMPTY!")
        lineageLock.Lock()
        lineage[jaeger_id] = make([]string, 0)
        soFar = lineage[jaeger_id]
        lineageLock.Unlock()
    }
    
    lineageLock.Lock()
    lineage[jaeger_id] = append(soFar, props["destination_svc"].(string))
    lineageLock.Unlock()



	for k, v := range props {
		fmt.Println("k:", k, "v:", v)
	}

	return &v1beta1.CheckResult{Status: status.OK,}, nil
}

// Addr returns the listening address of the server
func (s *MyGrpcAdapter) Addr() string {
	return s.listener.Addr().String()
}

// Run starts the server run
func (s *MyGrpcAdapter) Run(shutdown chan error) {
	shutdown <- s.server.Serve(s.listener)
}

// Close gracefully shuts down the server; used for testing
func (s *MyGrpcAdapter) Close() error {
	if s.server != nil {
		s.server.GracefulStop()
	}

	if s.listener != nil {
		_ = s.listener.Close()
	}

	return nil
}

// NewMyGrpcAdapter creates a new IBP adapter that listens at provided port.
func NewMyGrpcAdapter(addr string) (Server, error) {
	if addr == "" {
		addr = "0"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
	if err != nil {
		return nil, fmt.Errorf("unable to listen on socket: %v", err)
	}
	s := &MyGrpcAdapter{
		listener: listener,
	}
	fmt.Printf("Old: listening on \"%v\"\n", s.Addr())
	s.server = grpc.NewServer()
	authorization.RegisterHandleAuthorizationServiceServer(s.server, s)
	return s, nil
}
