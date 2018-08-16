package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type Worker struct {
	tasks               chan (Task)          // Канал для отправки задач на выполнение
	tickerStoppersToAdd chan (chan struct{}) // Канал для добавления каналов-стопперов для таймеров
	tickerStoppers      []chan struct{}      // Слайс для хранения каналов-стопперов для таймеров
	smoothFinish        sync.WaitGroup       // Для того, чтобы при Stop() дождаться завершения всех основных и вспомогательных задач
}

func main() {
	worker := NewWorker()
	worker.Start()
	worker.AddTicker(time.Second*3, Task1)
	worker.AddTicker(time.Second*15, Task2)
	time.Sleep(time.Second * 30)
	worker.Stop()
}

func Task1() {
	fmt.Println("Started task 1")
	time.Sleep(time.Second)
	fmt.Println("Ended task 1")
}

func Task2() {
	fmt.Println("Started task 2")
	time.Sleep(time.Second * 5)
	fmt.Println("Ended task 2")
}

func (w *Worker) Start() {
	go w.tasksProcessor()
	go w.tickersProcessor()
}

func (w *Worker) tasksProcessor() {
	// Читаем подряд все сообщения из канала tasks, пока он не закроется
	for task := range w.tasks {
		w.smoothFinish.Add(1)
		task()
		w.smoothFinish.Done()
	}
	// Здесь мы окажемся, когда канал tasks закроется
}

func (w *Worker) tickersProcessor() {
	// Обработка события добавления нового "тикера"
	for tickerStopper := range w.tickerStoppersToAdd {
		w.tickerStoppers = append(w.tickerStoppers, tickerStopper)
	}
	// Здесь мы оказываемся, когда канал w.tickerStoppersToAdd закрылся
	// Сообщаем всем тикерам, что лавочка закрывается (для этого мы их и добавляли в слайс)
	for _, tickerStopper := range w.tickerStoppers {
		tickerStopper <- struct{}{}
	}
	// Сообщаем, что закончили закрывать все тикеры
	w.smoothFinish.Done()
}

func (w *Worker) AddTicker(interval time.Duration, task Task) {
	tickerStopper := make(chan struct{})
	w.tickerStoppersToAdd <- tickerStopper
	go func() {
		// Главный цикл тикера
		for {
			// Пытаемся выполнить задачу
			select {
			case w.tasks <- task:
			case <-tickerStopper:
				return
			}
			// Ждем интервал между выполнениями
			select {
			case <-time.After(interval):
			case <-tickerStopper:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	// Инкрементим waitGroup, чтобы смочь дождаться закрытия всех тикеров
	w.smoothFinish.Add(1)
	// Закрываем канал добавления тикеров
	close(w.tickerStoppersToAdd)
	// Ждем, пока все тикеры закроются и все задачи завершатся
	// Очень важно дождаться тут закрытия всех тикеров
	// потому что если тикер попытается записать в закрытый канал tasks, будет паника
	w.smoothFinish.Wait()
	// Закрываем канал в задачами, в результате чего цикл воркера завершится
	close(w.tasks)
}

func NewWorker() *Worker {
	return &Worker{
		tasks:               make(chan Task),
		tickerStoppersToAdd: make(chan (chan struct{})),
	}
}
