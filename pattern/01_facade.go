package pattern

import (
	"errors"
	"fmt"
	"time"
)

type Card struct {
	CardNumber string
	Balance    int
	Bank       *Bank
}

type Bank struct {
	Name  string
	Cards []Card
}

type User struct {
	Name string
	Card *Card
}

type OneProduct struct {
	ProductName string
	Price       int
}

type Shop struct {
	ShopName string
	Products []OneProduct
}

/* Я предлагаю рассмотреть структуру паттерна фасад на примере устройства магазина ,а точнее работы платежной системы.
У нас есть 3 основных структупы Банк,Покупатель и магазин у которых есть своя сложная логика взаимодействия между собой.
Паттерн Фасад помогает спрятать всю сложную логику внутри одного метода, что сильно упрощает дальнейшую разработку и
вцелом взаимодействие с данной логикой.
*/

func (c Card) CheckBalance() error {

	fmt.Printf("[Карта] Проверка баланса карты %s\n", c.CardNumber)
	time.Sleep(time.Millisecond * 500)

	return c.Bank.GetBalance(c.CardNumber)

}

func (b Bank) GetBalance(cardNumber string) error {

	fmt.Printf("[Банк] Получение остатка по карте %s\n", cardNumber)
	time.Sleep(time.Millisecond * 300)

	for _, card := range b.Cards {
		if cardNumber == card.CardNumber {
			if card.Balance <= 0 {
				return errors.New("[Банк] Баланс карты отрицателен \n")
			}
		}
	}
	fmt.Printf("[Банк] Баланс карты положителен\n")
	return nil
}

func (s Shop) SellProducts(user User, product string) error {

	fmt.Printf("[Магазин] Запрос к пользователю для получения остатка по карте\n")
	time.Sleep(time.Millisecond * 500)

	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}

	fmt.Printf("[Магазин] Достаточно ли средств для покупки у покупателя %s\n", user.Name)
	time.Sleep(time.Millisecond * 300)

	for _, prod := range s.Products {

		if prod.ProductName == product {

			if prod.Price > user.Card.Balance {
				return errors.New("[Магазин] Недостаточно средств для покупки\n")
			}

			fmt.Printf("[Магазин] Товар  %s куплен покупателем %s\n", product, user.Name)
		}
	}
	return nil
}

/*Здесь я привел простую реализацию логики работы платежной системы магазина
В данном случае фасадом выступает метод SellProducts() так как он является оберткой над логикой
работы банка и карт.Он позволяет не обращаться каждый раз к данным методам напрямую, а использовать
удобную обертку для взаимодействия
*/

var (
	bank = Bank{
		Name:  "Bank",
		Cards: []Card{},
	}

	card1 = Card{
		CardNumber: "CRD-1",
		Balance:    200,
		Bank:       &bank,
	}

	card2 = Card{
		CardNumber: "CRD-2",
		Balance:    20,
		Bank:       &bank,
	}

	user1 = User{
		Name: "user-1",
		Card: &card1,
	}

	user2 = User{
		Name: "User-2",
		Card: &card2,
	}

	product = OneProduct{
		ProductName: "Cheese",
		Price:       150,
	}

	shop = Shop{
		ShopName: "SHOP",
		Products: []OneProduct{},
	}
)

func main() {

	bank.Cards = append(bank.Cards, card1, card2)
	shop.Products = append(shop.Products, product)

	err := shop.SellProducts(user1, product.ProductName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = shop.SellProducts(user2, product.ProductName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

/*
Паттерн Фасад (Facade) применяется в разработке программного обеспечения, когда необходимо упростить сложный интерфейс или подсистему,
предоставив клиентскому коду более простой и удобный способ взаимодействия с этой системой. Вот несколько ситуаций,
в которых паттерн Фасад может быть полезным:

Сложная система с множеством подсистем:
Когда ваша система состоит из множества классов, объектов и подсистем, и клиентскому коду требуется взаимодействовать с ними, паттерн Фасад позволяет создать простой интерфейс для доступа ко всей сложной системе. Это особенно полезно в больших приложениях.

Упрощение использования библиотек и API:
При работе с внешними библиотеками, фреймворками или API может быть многословно и сложно. Фасад позволяет создать уровень абстракции, который скрывает детали внешней системы и предоставляет более простой и понятный интерфейс для клиентского кода.

Создание удобных интерфейсов к сложным операциям:
Если в вашей системе есть сложные операции, которые требуют последовательного выполнения нескольких шагов, Фасад может предоставить методы для выполнения этих операций одной командой.

Повышение безопасности и контроля:
Фасад может использоваться для контроля доступа к определенным частям системы или для обеспечения безопасности, скрывая детали реализации от клиентского кода.

Уменьшение связанности:
Паттерн Фасад помогает уменьшить связанность (coupling) между клиентским кодом и подсистемой, так как клиент взаимодействует только с фасадом, а не с каждым компонентом подсистемы.

Легкость внесения изменений:
Если в будущем необходимо внести изменения в сложную систему, это можно сделать внутри фасада, не затрагивая клиентский код. Это обеспечивает устойчивость к изменениям и упрощает обслуживание системы.
*/

/*
Плюсы паттерна Фасад:

Упрощение интерфейса: Паттерн Фасад позволяет создать упрощенный и более понятный интерфейс для взаимодействия с сложной системой или подсистемой. Это делает код клиента более чистым и понятным.

Сокрытие деталей реализации: Фасад скрывает сложные детали внутренней реализации подсистемы от клиентского кода. Клиент не знает, как именно работает подсистема, и не зависит от ее изменений.

Уменьшение связанности: Паттерн Фасад помогает уменьшить связанность между клиентским кодом и подсистемой. Клиент взаимодействует только с фасадом, а не с каждым компонентом подсистемы, что делает систему более гибкой и легкой в сопровождении.

Улучшение безопасности: Фасад может предоставлять контролируемый и безопасный интерфейс к подсистеме, что позволяет избежать нежелательных действий со стороны клиента.

Легкость внесения изменений: Если необходимо внести изменения во внутреннюю реализацию подсистемы, это можно сделать в фасаде без влияния на клиентский код. Это обеспечивает устойчивость к изменениям.

Минусы паттерна Фасад:

Увеличение числа классов: Внедрение фасада добавляет еще один уровень абстракции и может увеличить количество классов в системе, что может быть избыточным для небольших и простых приложений.

Ограничение гибкости: Фасад может скрывать некоторые возможности и функции подсистемы, что может быть нежелательным, если клиенту требуется более низкоуровневый доступ к подсистеме.

Усложнение поддержки: Если фасад становится слишком большим и выполняет множество операций, его поддержка и сопровождение могут стать сложными.

Нежелательная зависимость от фасада: Если клиенты становятся сильно зависимыми от фасада, это может усложнить переход к другой реализации подсистемы, если это станет необходимым.

Сложность тестирования: Тестирование фасада может потребовать создания специальных тестовых сценариев, что может быть сложным, если подсистема сложна.
*/

/*
Паттерн Фасад активно применяется на практике в разработке программного обеспечения. Вот несколько реальных примеров его использования:

Библиотеки для работ с базами данных:
Многие библиотеки для работы с базами данных предоставляют сложные API, требующие настройки соединения, отправки SQL-запросов и обработки результатов. Фасад может предоставить более простой и высокоуровневый интерфейс для выполнения операций с базой данных.

Библиотеки для работ с сетевыми протоколами:
При работе с HTTP, FTP, SMTP и другими сетевыми протоколами клиентский код может столкнуться с деталями обработки запросов и ответов. Фасад может скрыть эти детали и предоставить простой интерфейс для взаимодействия с сетевыми ресурсами.

Библиотеки для графического интерфейса:
Графические библиотеки, такие как Qt или JavaFX, могут иметь сложное API для создания и управления элементами интерфейса. Фасад может предоставить упрощенный способ создания окон, кнопок, текстовых полей и других элементов GUI.

Фреймворки для тестирования:
Фреймворки для автоматизированного тестирования, например, Selenium для тестирования веб-приложений, могут использовать фасад для упрощения создания и выполнения тестовых сценариев. Фасад может предоставить удобные методы для взаимодействия с элементами интерфейса и проверки результатов.

Управление жизненным циклом приложения:
В больших приложениях может быть множество компонентов, управляющих жизненным циклом (например, инициализация, конфигурация, завершение работы). Фасад может предоставить один метод для запуска и остановки всего приложения.
*/
