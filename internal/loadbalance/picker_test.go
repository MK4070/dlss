package loadbalance_test

import (
	"testing"

	"github.com/MK4070/dlss/internal/loadbalance"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
)

func TestPickerNoSubConnAvailable(t *testing.T) {
	picker := &loadbalance.Picker{}
	for _, method := range []string{
		"/log.vX.Log/Produce",
		"/log.vX.Log/Consume",
	} {
		info := balancer.PickInfo{
			FullMethodName: method,
		}
		result, err := picker.Pick(info)
		require.Equal(t, balancer.ErrNoSubConnAvailable, err)
		require.Nil(t, result.SubConn)
	}
}

func TestPickerProducesToLeader(t *testing.T) {
	picker, subConns := setupTest()
	info := balancer.PickInfo{
		FullMethodName: "log.vX.Log/Produce",
	}
	for range 5 {
		gotPick, err := picker.Pick(info)
		require.NoError(t, err)
		require.Equal(t, subConns[0], gotPick.SubConn)
	}
}

func TestPickerConsumesFromFollowers(t *testing.T) {
	picker, subConns := setupTest()
	info := balancer.PickInfo{
		FullMethodName: "log.vX.Log/Consume",
	}
	for i := range 5 {
		gotPick, err := picker.Pick(info)
		require.NoError(t, err)
		require.Equal(t, subConns[i%2+1], gotPick.SubConn)
	}
}

func setupTest() (*loadbalance.Picker, []balancer.SubConn) {
	var subConns []balancer.SubConn
	buildInfo := base.PickerBuildInfo{
		ReadySCs: make(map[balancer.SubConn]base.SubConnInfo),
	}
	cc := &testClientConn{}
	cc.newSubConn = func(a []resolver.Address, nsco balancer.NewSubConnOptions) (balancer.SubConn, error) {
		return &subConn{}, nil
	}
	for i := range 3 {
		addr := resolver.Address{
			Attributes: attributes.New("is_leader", i == 0),
		}
		sc, _ := cc.NewSubConn([]resolver.Address{addr}, balancer.NewSubConnOptions{})
		buildInfo.ReadySCs[sc] = base.SubConnInfo{Address: addr}
		subConns = append(subConns, sc)
	}
	picker := &loadbalance.Picker{}
	picker.Build(buildInfo)
	return picker, subConns
}

type testClientConn struct {
	balancer.ClientConn
	newSubConn func([]resolver.Address, balancer.NewSubConnOptions) (balancer.SubConn, error)
}

func (c *testClientConn) NewSubConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (balancer.SubConn, error) {
	return c.newSubConn(addrs, opts)
}

type subConn struct {
	balancer.SubConn
}
