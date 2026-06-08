package main

import (
	nats "T4/Nats"
	redisSelf "T4/Redis"
	pb "T4/proto"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	gen "T4/dataGen"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalln(err)
	}

	// 创建grpc服务
	grpcServer := grpc.NewServer()

	// 在grpc服务端中注册服务
	pb.RegisterGetDataServer(grpcServer, &server{})

	// 启动服务
	grpcServer.Serve(listen)

}

// grpc 服务端
type server struct {
	pb.UnimplementedGetDataServer
}

func (s *server) GetData(ctx context.Context, req *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	nc := nats.Connect()
	wg := sync.WaitGroup{}
	// 1. 获取数据
	wg.Add(1)
	go func() {
		defer wg.Done()
		data := gen.DataGen()
		// fmt.Println(data)
		nats.Publisher(data, nc)
		// fmt.Println("write into nats")
		time.Sleep(time.Second)
	}()

	// 3. 订阅 nats中数据 4.写入 redis
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		nats.Subscriber(nc)
		time.Sleep(time.Second)
		ch <- 1
	}()

	// 5. rpc 调用 从redis中获取数据并返回客户端
	var res pb.GetDataResponse

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ch

		dataR := redisSelf.ReadR()
		fmt.Println(dataR)

		res.Humility = dataR.Humility
		res.Tempreature = dataR.Tempreature
	}()

	wg.Wait()
	return &res, nil
}

/*
	wg := sync.WaitGroup{}
	dataChan := make(chan m.Data, 1)
	// return &pb.GetDataResponse{Tempreature: }

	wg.Add(1)
	go func(dataChan chan m.Data, wg *sync.WaitGroup) {
		defer wg.Done()
		for{
			dataChan <- gen.DataGen()
		}
		// time.Sleep(time.Second)
	}(dataChan, &wg)

	var res pb.GetDataResponse
	wg.Add(1)
	go func(dataChan chan m.Data, res *pb.GetDataResponse, wg *sync.WaitGroup) {
		defer wg.Done()
		wg.Add(2)
		defer wg.Done()
		defer wg.Done()
		// 获取生成的数据
		data := <-dataChan

		// 写入nats
		go func() {
			nat.Publisher(data)
			time.Sleep(time.Second)
		}()

		var temp m.Data
		go func(temp *m.Data) {
			nat.Subscriber()
			fmt.Println(redisSelf.ReadR())
			temp = *redisSelf.ReadR()
			ime.Sleep(time.Second)
		}(&temp)

		res.Humility = temp.Humility
		res.Tempreature = temp.Tempreature
	}(dataChan, &res, &wg)

	wg.Wait()
	return &res, nil
*/
