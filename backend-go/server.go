package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 1291
// Hash 2556
// Hash 8577
// Hash 8086
// Hash 2456
// Hash 8536
// Hash 5429
// Hash 4466
// Hash 6007
// Hash 7868
// Hash 6881
// Hash 1650
// Hash 7795
// Hash 9375
// Hash 5980
// Hash 7215
// Hash 8115
// Hash 1380
// Hash 7356
// Hash 3462
// Hash 7322
// Hash 6942
// Hash 2603
// Hash 2419
// Hash 5528
// Hash 9162
// Hash 7994
// Hash 2972
// Hash 2086
// Hash 6219
// Hash 9079
// Hash 1530
// Hash 1522
// Hash 2993
// Hash 7740
// Hash 2625
// Hash 9106
// Hash 6797
// Hash 7358
// Hash 2602
// Hash 6517
// Hash 4726
// Hash 4120
// Hash 2997
// Hash 4010
// Hash 1995
// Hash 9670
// Hash 8877
// Hash 2687
// Hash 4589
// Hash 9484
// Hash 4708
// Hash 7068
// Hash 9136
// Hash 8932
// Hash 5833
// Hash 3763
// Hash 9604
// Hash 1808
// Hash 1342
// Hash 7104
// Hash 9715
// Hash 9355
// Hash 2485
// Hash 5735
// Hash 1237
// Hash 8719
// Hash 8185
// Hash 5869
// Hash 9672
// Hash 8587
// Hash 4071
// Hash 7926
// Hash 7328
// Hash 2147
// Hash 6077
// Hash 6937
// Hash 8544