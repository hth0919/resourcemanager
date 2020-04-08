package main

import (
	"context"
	"fmt"
	res "github.com/hth0919/resourcecollector"
	"google.golang.org/grpc"
	"time"
)

const (
	port = ":50051"
)

var ci = &res.ClusterInfo{
	MetricValue:      []string{},
	Clustername:      "",
	KubeConfig:       "",
	AdminToken:       "",
	NodeList:         []*res.NodeInfo{},
	ClusterMetricSum: map[string]float64{},
	Host:             "",
}
var clustername *string

func D() res.SendClusterClient{
	host := "10.0.3.201" + port
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	c := res.NewSendClusterClient(conn)
	return c
}

func main() {
	clustername = new(string)
	i:=0
	for {

		fmt.Println(i)
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		aa := fillcluster()
		fmt.Println(aa.Clustername)
		fmt.Println()
		fmt.Println(ci.Host)

		r, err := D().SendCluster(ctx, aa)
		if err != nil {
			fmt.Printf("could not connect : %v", err)
		}
		*clustername = r.ClusterName
		fmt.Println(*clustername)
		i++

		time.Sleep(time.Second * time.Duration(r.Tick))
	}
}

func fillcluster() *res.ClusterInfo{
	fmt.Println("fill")
	start := time.Now()
	ci.NewClusterClient("")
	ci.NodeListInit()
	for i:=0;i<len(ci.NodeList) ;i++ {
		ci.CalculateNodeMetricSum(i)
	}
	ci.CalculateClusterMetricSum()
	elapsedTime := time.Since(start)
	ci.Clustername = *clustername
	fmt.Println("filldone ::::::", elapsedTime)
	return ci
}