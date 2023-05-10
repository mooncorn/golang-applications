package main

import (
	"flag"
	"fmt"
	"github.com/manishmeganathan/peerchat/src"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const myFlag = `WELCOME TO CHAT APPLICATION - TUESDAY GROUP`

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	logrus.SetOutput(os.Stdout)
}

func main() {
	username := flag.String("user", "", "Username to use chatApp")
	chatroom := flag.String("room", "", "chatroom to join")
	loglevel := flag.String("log", "", "Level of the logs to print")
	discovery := flag.String("discover", "", "Method to use for discovery")
	flag.Parse()

	switch *loglevel {
	case "panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "info", "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	fmt.Println(myFlag)
	fmt.Println("The Chat Application is starting")
	fmt.Println("Please be patient for 30 seconds or so")
	fmt.Println()

	p2phost := src.NewP2P()
	logrus.Infoln("Completed P2P Setup")
	switch *discovery {
	case "announce":
		p2phost.AnnounceConnect()
	case "advertise":
		p2phost.AdvertiseConnect()
	default:
		p2phost.AdvertiseConnect()
	}
	logrus.Infoln("Connected to Service Peers")

	chatapp, _ := src.JoinChatRoom(p2phost, *username, *chatroom)
	logrus.Infof("Joined the %s chatroom as %s", chatapp.RoomName, chatapp.UserName)
	time.Sleep(time.Second * 5)
	ui := src.NewUI(chatapp)
	ui.Run()
}
