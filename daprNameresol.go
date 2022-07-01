type ResolveRequest struct {
	ID        string
	Namespace string
	Port      int
	Data      map[string]string
  }
  
  type Resolver interface {
	// Init initializes name resolver.
	Init(metadata Metadata) error
	// ResolveID resolves name to address.
	ResolveID(req ResolveRequest) (string, error)
  }
  https://github.com/dapr/components-contrib/blob/c7adb917f3910136ff68f3174078d50ba0168d73/nameresolution/requests.go


  // ResolveID resolves name to address in Kubernetes.
func (k *resolver) ResolveID(req nameresolution.ResolveRequest) (string, error) {
	// Dapr requires this formatting for Kubernetes services
	return fmt.Sprintf("%s-dapr.%s.svc.%s:%d", req.ID, req.Namespace, k.clusterDomain, req.Port), nil
  }
  https://github.com/dapr/components-contrib/blob/62dbc782cb1746268f8381eb486623da855decd2/nameresolution/kubernetes/kubernetes.go


// http://github.com/zcong1993/dapr-1/blob/a8ee30180e1183e2a2e4d00c283448af6d73d0d0/pkg/grpc/grpc.go#L77-L77
func (g *Manager) GetGRPCConnection(ctx context.Context, address, id string, namespace string,
									 skipTLS, recreateIfExists, sslEnabled bool, customOpts ...grpc.DialOption) 
										(*grpc.ClientConn, error) {
	// ...
	dialPrefix := GetDialAddressPrefix(g.mode)
	// ...
	conn, err := grpc.DialContext(ctx, dialPrefix+address, opts...)
	// ...
  }
  
  // http://github.com/zcong1993/dapr-1/blob/a8ee30180e1183e2a2e4d00c283448af6d73d0d0/pkg/grpc/dial.go#L11-L11
  func GetDialAddressPrefix(mode modes.DaprMode) string {
	if runtime.GOOS == "windows" {
	  return ""
	}
  
	switch mode {
	case modes.KubernetesMode:
	  return "dns:///"
	default:
	  return ""
	}
  }



func main() {
	port := flag.String("port", ":8080", "listen port")
	maxConnectionAge := flag.Bool("maxConnectionAge", false, "If set server keepalive MaxConnectionAge.")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		panic(err)
	}

	opts := make([]grpc.ServerOption, 0)

	if *maxConnectionAge {
		opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: time.Second * 30,
		}))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterHelloServer(s, &helloService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

https://github.com/zcong1993/grpc-example/blob/d08b33fc99131e3c32610e4d561f54333c91899c/cmd/dns/server/main.go
