package main

import (
	"log"
)

// Chain Of Responsibility:
//
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
//
// Цели использования:
// - Инкапсулировать и обернуть какое-либо действие/информацию в небольшой пакет, чтобы воспользоваться им где бы то ни было


type Data struct {
	packageDone bool
	cacheDone   bool
	shadersDone bool
}

type Responsibility interface {
	Download(*Data)
	setNext(Responsibility)
}

func main() {

	data := Data{}
	packages := &PackagesResponsibility{}
	cache := &CacheResponsibility{}
	shaders := &ShadersResponsibility{}
	final := &FinalStageResponsibility{}
	packages.setNext(cache)
	cache.setNext(shaders)
	shaders.setNext(final)

	packages.Download(&data)
	// [Info] downloading necessary packages...
	// [Info] downloading necessary cache...
	// [Info] downloading necessary shaders...
	// [Info] all data was successfully download

}

type PackagesResponsibility struct {
	next Responsibility
}

func (p *PackagesResponsibility) Download(data *Data) {

	if data.packageDone {
		log.Println("[Info] packages loaded successfully")
		p.next.Download(data)
		return
	}

	log.Println("[Info] downloading necessary packages...")
	data.packageDone = true
	p.next.Download(data)

}

func (p *PackagesResponsibility) setNext(next Responsibility) {
	p.next = next
}

type CacheResponsibility struct {
	next Responsibility
}

func (c *CacheResponsibility) Download(data *Data) {

	if data.cacheDone {
		log.Println("[Info] cache loaded successfully")
		c.next.Download(data)
		return
	}

	log.Println("[Info] downloading necessary cache...")
	data.cacheDone = true
	c.next.Download(data)
}

func (c *CacheResponsibility) setNext(next Responsibility) {
	c.next = next
}

type ShadersResponsibility struct {
	next Responsibility
}

func (s *ShadersResponsibility) Download(data *Data) {

	if data.shadersDone {
		log.Println("[Info] shaders loaded successfully")
		s.next.Download(data)
		return
	}

	log.Println("[Info] downloading necessary shaders...")
	data.shadersDone = true
	s.next.Download(data)
}

func (s *ShadersResponsibility) setNext(next Responsibility) {
	s.next = next
}

type FinalStageResponsibility struct {
	next Responsibility
}

func (f *FinalStageResponsibility) Download(data *Data) {
	log.Println("[Info] all data was successfully download")
}

func (f *FinalStageResponsibility) setNext(next Responsibility) {
	f.next = next
}
