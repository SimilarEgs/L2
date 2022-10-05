package main

import (
	"log"
)

// Это поведенчекский паттерн проектирования, который позволяет нам передавать запросы по цепочке обработчиков.
// Получив запрос, каждый обработчик решает: либо обработать запрос, либо передать его по цепочке далее, следующему обработчику
//
// Достоинвства:
// - Изоляция между каждым шагом
// - Гибкость. Мы можем менять/убирать/добовлять последовательность выполнения цепочки обработчиков, в случае если это потребовалось
// - Контроль порядка обработки запросов
// - Соотвествует принципам: единой ответственности | открытый/закрытый принцип
//
// Недостатки:
// - некоторые запросы могут остаться не обработанными
//
// Пример реалезации:
// Поисковый сервис, который принимает поискрвые реквесты и отвечает респонсоми
// Разложим приложение на составляющие:
// local casche -> redis -> elastic search
// Приминив паттерн цепочки отвественности мы сможем в любой момент заменить сервис кеширования, при этом сохранив бизнеслогику, не повлияв на работу программы.
// Либое же, мы можем убрать целиком этап локал-кеша, и направлять запросы напрямую в редис.
//
// Устройсво паттерна:
// · Общий интерфейс для все обработчиков - в основном содержит один метод для обработки запросов и еще один метод для передачи следующему обработчику по цепочке
// · Обработчик - содержит логику обработки запроса. Получив запрос, каждый обработчик решает, обрабатывать его или передавать по цепочке

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
	// [Info] downloading necessary shaders...
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
