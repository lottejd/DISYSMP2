package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lottejd/DISYSMP2/ChittyChat"
	"google.golang.org/grpc"
)

const (
	address       = "localhost:8080"
	clientLogFile = "clientLogId_"
)

var (
	clientId                    int
	latestClientVectorTimeStamp []int32
)

func main() {
	// init
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// create client
	chat := ChittyChat.NewChittyChatServiceClient(conn)
	ctx := context.Background()

	JoinChat(ctx, chat)

	// constantly trying to get latest broadcast
	go GetBroadcast(ctx, chat)

	for {
		// to ensure "enter" has been hit before publishing - skud ud til Mie
		// remove newline windows format "\r\n"
		reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		input := strings.TrimSuffix(reader, "\r\n")
		//inputMac := strings.TrimSuffix(reader, "\n");

		if err != nil {
			Logger("bad bufio input", FormatLogFile(clientId))
		}
		if strings.EqualFold(input, "leave") {
			fmt.Println("Do you want to leave? \"y/n\"")
			_reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
			_input := strings.TrimSuffix(_reader, "\r\n")
			checkErr(err)
			if strings.EqualFold(_input, "y") {

				leaveRequest := ChittyChat.LeaveChatRequest{ClientId: int32(clientId)}
				response, err := chat.LeaveChat(ctx, &leaveRequest)
				checkErr(err)
				LoggerVectorClock(response.Msg, latestClientVectorTimeStamp, clientLogFile)

				time.Sleep(500 * time.Millisecond)
				os.Exit(0)
			}

		}
		if len(input) > 0 {
			PublishFromClient(input, ctx, chat)
		}
	}
}

func GetBroadcast(ctx context.Context, chat ChittyChat.ChittyChatServiceClient) {
	var latestError error
	for {
		//overvej at sende alle de seneste broadcasts, sorter lokalt ved clienten via lamport/vector clock
		time.Sleep(time.Millisecond * 30)

		response, err := chat.GetBroadcast(ctx, &ChittyChat.GetBroadcastRequest{})
		if err != nil {
			if err.Error() == latestError.Error() {
				// to avoid spamming the log with the same error
				continue
			} else {
				latestError = err
				Logger(err.Error(), FormatLogFile(clientId))
				continue
			}

		}

		// intent check vector clock to adjust latest broadcast
		broadCastIsNewer := false
		vectorClockFromServer := response.GetClientsConnected()
		if len(vectorClockFromServer) > len(latestClientVectorTimeStamp) {
			broadCastIsNewer = true
		}

	latest:
		for i := 0; i < len(vectorClockFromServer); i++ {
			if broadCastIsNewer {
				break latest
			}
			if vectorClockFromServer[i] > latestClientVectorTimeStamp[i] {
				broadCastIsNewer = true
			}
		}
		if broadCastIsNewer {
			latestClientVectorTimeStamp = vectorClockFromServer
			msg := response.Msg + ", by " + strconv.Itoa(int(response.GetClientId()))
			LoggerVectorClock(msg, latestClientVectorTimeStamp, FormatLogFile(clientId))
		}
	}
}

func PublishFromClient(input string, ctx context.Context, chittyServer ChittyChat.ChittyChatServiceClient) {
	inputFromClient := &ChittyChat.PublishRequest{Request: input, ClientId: int32(clientId)}
	response, err := chittyServer.Publish(ctx, inputFromClient)
	checkErr(err)

	LoggerVectorClock(response.GetMsg(), latestClientVectorTimeStamp, FormatLogFile(clientId))
}

func JoinChat(ctx context.Context, chittyServer ChittyChat.ChittyChatServiceClient) {
	response, err := chittyServer.JoinChat(ctx, &ChittyChat.JoinChatRequest{})
	checkErr(err)

	clientId = int(response.GetClientId())

	msg := fmt.Sprintf("connected with id: %v", clientId)
	Logger(msg, FormatLogFile(clientId))
}

func LeaveChat(ctx context.Context, chittyServer ChittyChat.ChittyChatServiceClient) {

	response, err := chittyServer.LeaveChat(ctx, &ChittyChat.LeaveChatRequest{ClientId: int32(clientId)})
	checkErr(err)

	msg := fmt.Sprintf("%s, %v", response.GetMsg(), clientId)
	Logger(msg, FormatLogFile(clientId))
}

//help methods
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
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

func FormatLogFile(clientId int) string {
	return fmt.Sprintf("%s%v", clientLogFile, clientId)
}

func Logger(message string, logFileName string) {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
}

func LoggerVectorClock(message string, vectorClock []int32, logFileName string) {
	Logger(message+", VectorClock: "+FormatVectorClock(vectorClock), logFileName)
}
