package main

import (
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

type Server struct {
	Addr string
	Ln   net.Listener
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", s.Addr)
	s.Ln = ln
	return err
}

type RoomID int

type Room struct {
	RoomId   RoomID
	users    []User
	RoomName string
}

func (r *Room) getUsers() []User {
	return r.users
}

func (r *Room) AddUser(user User) bool {
	users := r.getUsers()
	for _, roomUser := range users {
		if roomUser.Addr.String() == user.Addr.String() {
			return false
		}
	}
	r.users = append(r.users, user)
	return true
}

func (r *Room) DistributeMsg(fromUser string, msg string) *ErrorChan {
	errChan := ErrorChan{
		ErrMap: make(map[string]error),
	}

	var wg sync.WaitGroup
	for _, usr := range r.getUsers() {
		if usr.Addr.String() != fromUser {
			wg.Add(1)
			go func(usr User) {
				defer wg.Done()
				conn := *usr.GetConnection()
				addstr := strings.Split(fromUser, ":")

				_, err := conn.Write([]byte(
					"/" + addstr[len(addstr)-1] + ": " + msg + "\n",
				))
				if err != nil {
					errChan.AddNewRoutineError(
						conn.RemoteAddr().String(),
						err,
					)
				}
			}(usr)
		}
	}

	wg.Wait()
	return &errChan
}

type User struct {
	Addr   net.Addr
	RoomId RoomID
	conn   *net.Conn
	Room   *Room
}

func (u *User) GetConnection() *net.Conn {
	return u.conn
}

func (u *User) GetARoomYouGuys() *Room {
	return u.Room
}

type ErrorChan struct {
	mu     sync.Mutex
	ErrMap map[string]error
}

func (ec *ErrorChan) AddNewRoutineError(routineAddr string, err error) {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	ec.ErrMap[routineAddr] = err
}

func main() {
	server := NewServer(":2000")
	err := server.Listen()

	var roomId RoomID = 10
	room := Room{
		RoomId:   roomId,
		RoomName: "common-room",
	}
	if err != nil {
		log.Fatalf(
			"Could not listen to port because skill issues: %s\n",
			err.Error())
	}

	for {
		conn, err := server.Ln.Accept()
		if err != nil {
			log.Printf(
				"Attemp to connect with server failed by Addr: %s\n",
				conn.RemoteAddr())
		}

		usr := User{
			Addr:   conn.RemoteAddr(),
			conn:   &conn,
			RoomId: roomId,
			Room:   &room,
		}

		room.AddUser(usr)

		go HandleIncoming(usr)
	}
}

func HandleIncoming(user User) {
	conn := *user.GetConnection()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("User disconnected: %s\n", user.Addr.String())
				return
			}
			log.Printf("Could not read message from user: %s\n", err.Error())
		}

		msg := string(buf[:n])
		room := *user.GetARoomYouGuys()

		errChan := room.DistributeMsg(conn.RemoteAddr().String(), msg)
		if len(errChan.ErrMap) != 0 {
			log.Println("Some users didn't recieve the message")
		}
	}
}
