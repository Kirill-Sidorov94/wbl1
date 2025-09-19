package main

// Реальный пример использования
// Старый метод отдает xml, новый интерфейс ждет json

// Плюсы
// Своместимость
// Соблюдения open-closed
// Разделение отвественности

// Минусы
// Доп цепочки вызовов и новые обертки, усложение поддержки

import (
	"fmt"
)

type legacy struct {}

type legacyPingMessage struct {
	Message string
	Order   int
}

type legacyInterface interface {
	oldPing() *legacyPingMessage
}

func newLegacy() legacyInterface {
	return &legacy{}
}

func (l *legacy) oldPing() *legacyPingMessage {
	return &legacyPingMessage{Message: "old ping"}
}

type improve struct {}

func (i *improve) ping() string {
	return "new ping"
}

type adapter struct {
	legacy legacyInterface
}

func (a *adapter) ping() string {
	mgs := a.legacy.oldPing()

	return mgs.Message
}

type someInterface interface {
	ping() string 
}

func adapterDemonstaration() {
	l := newLegacy()
	i := &improve{}
	a := &adapter{legacy: l}
	someFunc(i)
	someFunc(a)
}

func someFunc(si someInterface) {
	fmt.Println(si.ping())
}