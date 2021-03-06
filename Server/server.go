package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/lottejd/DISYSMP2/ChittyChat"
	"google.golang.org/grpc"
)

const (
	port          = ":8080"
	serverLogFile = "serverLog"
)

var (
	vectorClock     []int32
	broadCastBuffer chan (bufferedMessage)
	clientCount     int
	lock            sync.Mutex
	latestBroadCast bufferedMessage
)

type Server struct {
	ChittyChat.UnimplementedChittyChatServiceServer
}

type bufferedMessage struct {
	message         string
	vectorTimeStamp []int32
	clientId        int32
}

func main() {

	//init
	vectorClock = make([]int32, 0, 1)
	broadCastBuffer = make(chan bufferedMessage, 10)
	lock = sync.Mutex{}
	grpcServer := grpc.NewServer()

	//setup listen on port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// constantly updating the latestBroadCast
	go EvalLatestBroadCast(broadCastBuffer)

	Logger("server is running", vectorClock, serverLogFile)

	// start the service / server on the specific port
	ChittyChat.RegisterChittyChatServiceServer(grpcServer, &Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s  %v", port, err)
	}

}

func (s *Server) GetBroadcast(ctx context.Context, _ *ChittyChat.GetBroadcastRequest) (*ChittyChat.Response, error) {
	if len(latestBroadCast.vectorTimeStamp) < 1 {
		return nil, errors.New("no broadcasts")
	}

	return &ChittyChat.Response{Msg: latestBroadCast.message, ClientId: latestBroadCast.clientId, ClientsConnected: latestBroadCast.vectorTimeStamp}, nil
}

func (s *Server) Publish(ctx context.Context, message *ChittyChat.PublishRequest) (*ChittyChat.Response, error) {
	validateMessage, err := ValidateMessage(message.GetRequest())
	if validateMessage {
		//logging
		lock.Lock()
		vectorClock[int(message.GetClientId())]++
		lock.Unlock()

		msg := "Request by: " + strconv.Itoa(int(message.GetClientId())) + " was accepted"
		Logger(msg, vectorClock, serverLogFile)

		Broadcast(message.GetRequest(), int(message.GetClientId()))
		return &ChittyChat.Response{Msg: "Request was accepted"}, nil
	} else {
		//logging
		msg := message.GetRequest() + ", error msg " + err.Error()
		Logger(msg, vectorClock, "ServerErrorLog")

		return &ChittyChat.Response{Msg: err.Error()}, err
	}
}

func (s *Server) JoinChat(ctx context.Context, _ *ChittyChat.JoinChatRequest) (*ChittyChat.JoinResponse, error) {

	// add a client
	vectorClock = append(vectorClock, 0)
	clientId := clientCount
	lock.Lock()
	vectorClock[clientId]++
	clientCount++
	lock.Unlock()

	//logging
	msg := "client: " + strconv.Itoa(clientId) + ", succesfully joined the chat"
	Logger(msg, vectorClock, serverLogFile)

	Broadcast(msg, clientId)
	return &ChittyChat.JoinResponse{ClientId: int32(clientId)}, nil
}

func (s *Server) LeaveChat(ctx context.Context, request *ChittyChat.LeaveChatRequest) (*ChittyChat.LeaveResponse, error) {
	clientId := request.GetClientId()

	lock.Lock()
	vectorClock[int(clientId)]++
	lock.Unlock()

	//logging
	msg := "client: " + strconv.Itoa(int(clientId)) + ", succesfully left the chat"
	Logger(msg, vectorClock, serverLogFile)

	Broadcast(msg, int(clientId))
	return &ChittyChat.LeaveResponse{Msg: msg}, nil
}

func Broadcast(msg string, clientId int) {
	//locking because using global variables is scary, we probably should think of something different?
	lock.Lock()

	// increment clock and add latest broadcast to the buffer
	vectorClock[clientId]++
	vectorClock := vectorClock
	broadCastBuffer <- bufferedMessage{message: msg, vectorTimeStamp: vectorClock, clientId: int32(clientId)}

	//logging
	Logger(msg+", by: "+strconv.Itoa(clientId), vectorClock, serverLogFile)

	lock.Unlock()
}

// help method
func EvalLatestBroadCast(broadCastBuffer chan (bufferedMessage)) {
	for {
		select {
		case temp := <-broadCastBuffer:
			latestBroadCast = temp
			fmt.Println("new broadcast by, " + strconv.Itoa(int(temp.clientId)) + ": " + temp.message)
			time.Sleep(50 * time.Millisecond)
		default:
		}
	}
}

//??
func ValidateMessage(message string) (bool, error) {
	valid := utf8.Valid([]byte(message))
	if !valid {
		fmt.Println(message)
		return false, errors.New("not UTF-8")
	}
	if len(message) > 128 {
		return false, errors.New("too long")
	}
	return true, nil
}

func Logger(message string, vectorClock []int32, logFileName string) {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message + ", VectorClock: " + FormatVectorClock(vectorClock))
}

func FormatVectorClock(clock []int32) string {
	var sb = strings.Builder{}
	sb.WriteString("<")
	for i := 0; i < len(clock); i++ {
		sb.WriteString(" ")
		sb.WriteString(strconv.Itoa(int(clock[i])))
		sb.WriteString(",")
	}
	sb.WriteString(" >")
	return sb.String()
}
