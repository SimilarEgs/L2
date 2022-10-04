package main

import (
	"log"
)

// Это поведенчекский паттерн проектирования, который позволяет нам передавать запросы по цепочке обработчиков.
// Получив запрос, каждый обработчик решает: либо обработать запрос, либо передать его по цепочке далее, следующему обработчику

type Data struct {
	packageDone bool
	cacheDone   bool
	shadersDone bool
}

type Responsability interface {
	Download(*Data)
	setNext(Responsability)
}

func main() {

	data := Data{}
	packages := &PackagesResponsability{}
	cache := &CacheResponsability{}
	shaders := &ShadersResponsability{}
	final := &FinalStageResponsability{}
	packages.setNext(cache)
	cache.setNext(shaders)
	shaders.setNext(final)

	packages.Download(&data)
	// [Info] downloading necessary packages...
	// [Info] downloading necessary cache...
	// [Info] downloading necessary cache...
	// [Info] all data was successfully download

}

type PackagesResponsability struct {
	next Responsability
}

func (p *PackagesResponsability) Download(data *Data) {

	if data.packageDone {
		log.Println("[Info] packages loaded successfully")
		p.next.Download(data)
		return
	}

	log.Println("[Info] downloading necessary packages...")
	data.packageDone = true
	p.next.Download(data)

}

func (p *PackagesResponsability) setNext(next Responsability) {
	p.next = next
}

type CacheResponsability struct {
	next Responsability
}

func (c *CacheResponsability) Download(data *Data) {

	if data.cacheDone {
		log.Println("[Info] cache loaded successfully")
		c.next.Download(data)
		return
	}

	log.Println("[Info] downloading necessary cache...")
	data.cacheDone = true
	c.next.Download(data)
}

func (c *CacheResponsability) setNext(next Responsability) {
	c.next = next
}

type ShadersResponsability struct {
	next Responsability
}

func (s *ShadersResponsability) Download(data *Data) {

	if data.shadersDone {
		log.Println("[Info] shaders loaded successfully")
		s.next.Download(data)
		return
	}

	log.Println("[Info] downloading necessary shaders...")
	data.shadersDone = true
	s.next.Download(data)
}

func (s *ShadersResponsability) setNext(next Responsability) {
	s.next = next
}

type FinalStageResponsability struct {
	next Responsability
}

func (f *FinalStageResponsability) Download(data *Data) {
	log.Println("[Info] all data was successfully download")
}

func (f *FinalStageResponsability) setNext(next Responsability) {
	f.next = next
}
