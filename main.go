package main

import (
	"querycache/log"
)
type Range struct{
	startDate uint8
	endDate uint8
}

func main() {
	StartHTTP()
//	query := jsonToQuery("test")
//	println(query)
//	test := getDataFromDatabase(query)
//	write(test)
}
	//Я вижу кэширующий сервис примерно так - на вход он получает:
	//  - флаг, нужна ли почасовая разбивка.
	//	- datetime range, на котором делать выборку.
	//	- список полей, по которым нужно группировать данные (в google analytics это называется dimensions). Поля могут быть вычислемые - сервису не обязательно проверять
	// их на валидность - он полагается на проверку этих полей и их корректную генерацию клиентом, обратившимся в этот сервис
	//  - список полей, по которым нужно подсчитать агрегированные метрики (в google analytics они называются metrics). Поля также могут быть вычисляемые и их не нужно
	// проверять на валидность. Предполагается, что клиент правильно их формирует. Ограничение, которое должен соблюдать клиент - если не установлен флаг почасовой разбивки,
	// то по этим полям можно считать только sum* и count* (count* - это упрощенный вид sum*(1)). uniq, avg и прочее не подходит, т.к. потом нельзя будет смержить
	// результаты из кэша за разные интервалы. По sum* - это просто сложение. В почасовой разбивке нет ограничения на агрегирующие функции.
	//	- список условий, по которым нужно делать фильтрацию (в google analytics это custom filters). Условия - это фактически подмножество фильтров из where,
	// за исключением условий по времени.
	//на выходе отдает TSV строчки в виде:
	//toStartOfHour(datetime)  dimension values... metric values...
	//поле toStartOfHour(datetime) будет остуствовать для отчетов без почасовой разбивки.
	//
	//внутри сервис должен делать следующее:
	//1) Округлить нижнюю границу datetime range до ближайшего меньшего часа.
	//	2) Разбить datetimer range на почасовые интервалы. Последний интервал может получиться неполным.
	//	3) Составить для каждого почасового интервала запрос из входных данных.
	//	4) Проверить в кэше наличие результатов по всем запросам за полный час. В качестве ключа использовать непосредственно сам запрос. Последний интервал не имеет смысла искать в кэше, если он не покрывает целый час.
	//	5) Для всех запросов, результаты которых не найдены в кэше, отправить запрос в базу и закэшировать результат (кроме последнего интервала, если он неполный).
	//6) Смержить все результаты за почасовые интервалы и отдать в ответе. Если установлен флаг почасовой разбивки, то мерж совсем простой - просто выводить по очереди все значения по каждому интервалу начиная с минимального. Если нет почасовой разбивки, то нужно будет составить хэш-таблицу с ключом, состоящим из dimensions, и значением, состоящим из metrics для конкретного dimension'а, и пройтись по выбранным почасовым интервалам, суммируя метрики с помощью этой хэш-таблицы.
	//
	//
	//В будущем в сервис можно будет добавить поддержку сортировки с пейджинацией
	//[13:51:05] : клиентом для данного сервиса будет выступать программа, генерирующая отчеты для конечных пользователей.
	// Вся логика по валидации входных данных, формировании правильных запросов к сервису и
	// пост-обработке TSV-ответа (та же сортировка с пейджинацией) будет проходить внутри этой программы
	//[13:52:16] : большинство отчетов генерит максимум пару тысяч строчек, так что для них нет потребности в сортировке и пейджинации на нашей стороне
	// но есть пару отчетов вроде разбивки статы по доменам, где количество строк может доходить до десятков миллионов.
	// Вот для них может понадобиться сортировка с пейджинацией на нашей стороне, чтобы не грузить программу-клиента обработкой лишних данных
	//[13:57:08] : о забыл добавить в набор входных параметров имя таблицы, из которой делать выборку :)


	func cacheQuery(isHour bool, r Range, groupColumns []string, agregationColumns []string, filter []string){

    test1  := "test"
		log.Infof(test1)
	}








