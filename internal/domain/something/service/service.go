package service

import "sync"

type SomethingService interface {
	Job1(wg *sync.WaitGroup) error
	Job2(wg *sync.WaitGroup) error
	Job3(wg *sync.WaitGroup) error
}
