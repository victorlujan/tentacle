// package services

// import (
// 	"context"
// 	"sync"

// 	"github.com/wailsapp/wails/v2/pkg/runtime"
// )

// type JSResp struct {
// 	Success bool   `json:"success"`
// 	Msg     string `json:"msg"`
// 	Data    any    `json:"data,omitempty"`
// }

// type monitorItem struct {
// 	ch         chan string
// 	closeCh    chan struct{}
// 	enventName string
// }

// type MonitorService struct {
// 	ctx          context.Context
// 	ctxCancel    context.CancelFunc
// 	mutex        sync.Mutex
// 	monitorItems map[string]*monitorItem
// }

// var monitor *MonitorService
// var onceMonitor sync.Once

// func NewMonitorService() *MonitorService {
// 	onceMonitor.Do(func() {
// 		monitor = &MonitorService{
// 			monitorItems: make(map[string]*monitorItem),
// 		}
// 	})
// 	return monitor
// }

// func (c *MonitorService) Start(ctx context.Context) {
// 	c.ctx, c.ctxCancel = context.WithCancel(ctx)
// }

// func (m *MonitorService) AddMonitorItem(eventName string) {
// 	m.mutex.Lock()
// 	defer m.mutex.Unlock()

// 	ch := make(chan string)
// 	closeCh := make(chan struct{})
// 	item := &monitorItem{
// 		ch:         ch,
// 		closeCh:    closeCh,
// 		enventName: eventName,
// 	}
// 	m.monitorItems[eventName] = item

// }

// func (c *MonitorService) StartMonitor(server string) (resp JSResp) {
// 	c.mutex.Lock()
// 	defer c.mutex.Unlock()

// 	item, ok := c.monitorItems[server]
// 	if !ok {
// 		resp.Success = false
// 		resp.Msg = "Monitor item not found"
// 		return resp
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case <-item.closeCh:
// 				return
// 			case msg := <-item.ch:
// 				runtime.EventsEmit(c.ctx, item.enventName, msg)
// 			}
// 		}
// 	}()
// 	resp.Success = true
// 	resp.Msg = "Monitor started"
// 	return resp
// }

// func (c *MonitorService) StopMonitor(server string) (resp JSResp) {
// 	c.mutex.Lock()
// 	defer c.mutex.Unlock()

// 	item, ok := c.monitorItems[server]
// 	if !ok {
// 		resp.Success = false
// 		resp.Msg = "Monitor item not found"
// 		return resp
// 	}

// 	close(item.closeCh)
// 	resp.Success = true
// 	resp.Msg = "Monitor stopped"
// 	return resp
// }

// func (c *MonitorService) StopAll() {
// 	if c.ctxCancel != nil {
// 		c.ctxCancel()
// 	}

// 	for server := range c.monitorItems {
// 		c.StopMonitor(server)
// 	}
// }
