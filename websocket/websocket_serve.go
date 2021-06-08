package websocket

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	//允许等待的写入时间
	writeWait = 10 * time.Second
	//time allowed to read the next pong message
	pongWait = 60 * time.Second
	//must be less than pongWait
	pingWait = (pongWait * 9) / 10
	//max message size
	maxMessageSize = 512
)

var (
	//最大连接id, +1处理
	maxConnId int64
	//ws的所有连接，用于广播
	wsConnAll map[int64]*wsConnection
	//upGrader
	upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		//check cors跨域请求, true允许所有的cors请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//客户端读写消息
type wsMessage struct {
	//websocket textMessage类型
	messageType int
	data        []byte
}

//客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn //底层websocket
	inChan   chan *wsMessage //读队列
	outChan  chan *wsMessage //写队列

	mutex     sync.Mutex //避免重复关闭管道,加锁处理
	isClosed  bool
	closeChan chan byte //关闭通知
	id        int64
}

//处理客户端request,升级为websocket
func wsHandler(resp http.ResponseWriter, req *http.Request) {
	//连接转成websocket
	wsSocket, err := upGrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Fatal("wsSocket connect failed")
		return
	}
	maxConnId++
	//todo 通过wsConnAll控制连接数
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
		id:        maxConnId,
	}
	wsConnAll[maxConnId] = wsConn
	log.Println("当前连接数量", len(wsConnAll))

	// 处理器,发送定时信息，避免意外关闭
	go wsConn.processLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}

//处理队列消息,从inChan读出消息放入到outChan
func (wsConn *wsConnection) processLoop() {
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			log.Println("获取inChan消息failed", err.Error())
			break
		}
		log.Println("从inChan接收到消息", string(msg.data))
		err = wsConn.wsWrite(msg)
		if err != nil {
			log.Println("发送消息给客户端出现错误", err.Error())
			break
		}
	}
}

//接受客户端的消息
func (wsConn *wsConnection) wsReadLoop() {
	//设置消息最大长度和deadline
	wsConn.wsSocket.SetReadLimit(maxMessageSize)
	wsConn.wsSocket.SetReadDeadline(time.Now().Add(pongWait))
	for {
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			//websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			log.Println("websocket读取消息failed", err.Error())
			wsConn.close()
			return
		}
		//处理消息放入到inChan
		msg := &wsMessage{
			messageType: msgType,
			data:        data,
		}
		select {
		case wsConn.inChan <- msg:
		case <-wsConn.closeChan:
			return
		}
	}
}

func (wsConn *wsConnection) wsWriteLoop() {
	ticker := time.NewTicker(pingWait)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case msg := <-wsConn.outChan:
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				log.Println("发送消息到客户端failed", err.Error())
				wsConn.close()
				return
			}
		case <-wsConn.closeChan:
			return
		case <-ticker.C:
			wsConn.wsSocket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := wsConn.wsSocket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		}
	}
}

//从inChan读出消息
func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
		return nil, errors.New("connect closed")
	}
}

//写入outChan
func (wsConn *wsConnection) wsWrite(msg *wsMessage) error {
	select {
	case wsConn.outChan <- msg:
	case <-wsConn.closeChan:
		return errors.New("connect had closed")
	}
	return nil
}

//关闭连接，当从客户端读取消息失败，或发送消息到客户端失败调用
func (wsConn *wsConnection) close() {
	log.Println("关闭连接")
	wsConn.wsSocket.Close()
	//加锁防止重复关闭
	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if wsConn.isClosed == false {
		wsConn.isClosed = true
		//wsConn id
		delete(wsConnAll, wsConn.id)
		close(wsConn.closeChan)
	}
}

func StartWebSocket(port string) {
	wsConnAll = make(map[int64]*wsConnection)
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(port, nil)
}
