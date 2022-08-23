package impl

import (
	"log"
	"sync"
	"time"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Job1(wg *sync.WaitGroup) error {
	defer wg.Done()

	log.Println("job1 started")
	time.Sleep(time.Second * 10)
	log.Println("job1 finished")

	return nil
}

func (s *service) Job2(wg *sync.WaitGroup) error {
	defer wg.Done()

	log.Println("job2 started")
	time.Sleep(time.Second * 5)
	log.Println("job2 finished")

	return nil
}

func (s *service) Job3(wg *sync.WaitGroup) error {
	defer wg.Done()

	log.Println("job3 started")
	time.Sleep(time.Second * 2)
	log.Println("job3 finished")

	return nil
}
