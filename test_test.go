package main

import (
	"GolangDesignPatterns/internal/AdapterPattern"
	"GolangDesignPatterns/internal/Bridge"
	"GolangDesignPatterns/internal/BuilderPattern"
	"GolangDesignPatterns/internal/ChainofResponsibilityPattern"
	"GolangDesignPatterns/internal/CommandPattern"
	"GolangDesignPatterns/internal/CompositePattern"
	"GolangDesignPatterns/internal/DecoratorPattern"
	"GolangDesignPatterns/internal/FacadePattern"
	"GolangDesignPatterns/internal/FactoryPattern"
	"GolangDesignPatterns/internal/FilterPattern"
	"GolangDesignPatterns/internal/FlyweightPattern"
	"GolangDesignPatterns/internal/InterpreterPattern"
	"GolangDesignPatterns/internal/IteratorPattern"
	"GolangDesignPatterns/internal/MediatorPattern"
	"GolangDesignPatterns/internal/MementoPattern"
	"GolangDesignPatterns/internal/NullObjectPattern"
	"GolangDesignPatterns/internal/ObserverPattern"
	"GolangDesignPatterns/internal/ProxyPattern"
	"GolangDesignPatterns/internal/SingletonPattern"
	"GolangDesignPatterns/internal/StatePattern"
	"GolangDesignPatterns/internal/StrategyPattern"
	"GolangDesignPatterns/internal/TemplatePattern"
	"GolangDesignPatterns/internal/VisitorPattern"
	"fmt"
	"testing"
)

//简单工厂
func TestSimplenessFactory(t *testing.T) {
	f := new(FactoryPattern.SimplenessFactory)
	var s FactoryPattern.Shape
	s, ok := f.GetShape("Rectangle")
	if ok {
		s.Draw()
	}
}

//抽象工厂
func TestAbstractFactory(t *testing.T) {
	f := new(FactoryPattern.AbsFactory)
	color := f.GetColor("Red")
	color.Fill()
	shape := f.GetShape("Rectangle")
	shape.Draw()
}

//单例
func TestSingleton(t *testing.T) {
	s1 := SingletonPattern.GetInstance1()
	s1.Count = 5
	fmt.Println(s1)
	s2 := SingletonPattern.GetInstance1()
	fmt.Println(s2)
}

//建造者
func TestBuilderPattern(t *testing.T) {
	builder := new(BuilderPattern.ComputerBuilder)
	director := BuilderPattern.Director{Builder: builder}
	computer := director.Create("I7", "32G", "4T")
	fmt.Println(*computer)
}

//适配器
func TestAdapterPattern(t *testing.T) {
	audioPlayer := AdapterPattern.AudioPlayer{}

	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}

//桥接
func TestBridge(t *testing.T) {
	redCircle := Bridge.Circle{}
	redCircle.Circle(100, 100, 10, &Bridge.RedCircle{})
	greenCircle := Bridge.Circle{}
	greenCircle.Circle(100, 100, 10, &Bridge.GreenCircle{})
	redCircle.Draw()
	greenCircle.Draw()
}

//过滤器
func TestFilterPattern(t *testing.T) {
	var persons []FilterPattern.Person
	persons = append(persons, FilterPattern.GetPerson("Robert", "Male", "Single"))
	persons = append(persons, FilterPattern.GetPerson("John", "Male", "Married"))
	persons = append(persons, FilterPattern.GetPerson("Laura", "Female", "Married"))
	persons = append(persons, FilterPattern.GetPerson("Diana", "Female", "Single"))
	persons = append(persons, FilterPattern.GetPerson("Mike", "Male", "Single"))
	persons = append(persons, FilterPattern.GetPerson("Bobby", "Male", "Single"))
	male := new(FilterPattern.CriteriaMale)
	fmt.Println(male.MeetCriteria(persons))
	female := new(FilterPattern.CriteriaFemale)
	fmt.Println(female.MeetCriteria(persons))
	single := new(FilterPattern.CriteriaSingle)
	fmt.Println(single.MeetCriteria(persons))
	singleMale := new(FilterPattern.AndCriteria)
	singleMale.AndCriteria(single, male)
	fmt.Println(singleMale.MeetCriteria(persons))
	singleFemale := new(FilterPattern.AndCriteria)
	singleFemale.AndCriteria(single, female)
	fmt.Println(singleFemale.MeetCriteria(persons))
}

//组合模式
func TestCompositePattern(t *testing.T) {
	ceo := CompositePattern.GetEmployee("John", "CEO", 30000)

	headSales := CompositePattern.GetEmployee("Robert", "Head Sales", 20000)
	headMarketing := CompositePattern.GetEmployee("Michel", "Head Marketing", 20000)

	ceo.Add(headSales)
	ceo.Add(headMarketing)

	salesExecutive1 := CompositePattern.GetEmployee("Richard", "Sales", 10000)
	salesExecutive2 := CompositePattern.GetEmployee("Rob", "Sales", 10000)

	headSales.Add(salesExecutive1)
	headSales.Add(salesExecutive2)

	clerk1 := CompositePattern.GetEmployee("Laura", "Marketing", 10000)
	clerk2 := CompositePattern.GetEmployee("Bob", "Marketing", 10000)

	headMarketing.Add(clerk1)
	headMarketing.Add(clerk2)

	ceo.PrintSubordinates()

	ceo.Remove(headSales)
	ceo.Remove(headMarketing)

	ceo.PrintSubordinates()
}

//装饰器
func TestDecoratorPattern(t *testing.T) {
	circle := DecoratorPattern.Circle{}
	redCircle := DecoratorPattern.RedShapeDecorator{}
	redCircle.RedShapeDecorator(new(DecoratorPattern.Circle))

	redRectangle := DecoratorPattern.RedShapeDecorator{}
	redRectangle.RedShapeDecorator(new(DecoratorPattern.Rectangle))

	fmt.Println("Circle with normal border")
	circle.Draw1()
	fmt.Println("Circle of red border")
	redCircle.Draw1()
	fmt.Println("Rectangle of red border")
	redRectangle.Draw1()
}

//外观
func TestFacadePattern(t *testing.T) {
	shapeMaker := new(FacadePattern.ShapeMaker)
	shapeMaker.ShapeMaker()

	shapeMaker.DrawCircle()
	shapeMaker.DrawRectangle()
	shapeMaker.DrawSquare()
}

//享元设计模式
func TestFlyweightPattern(t *testing.T) {
	flyweightPatternDemo := FlyweightPattern.FlyweightPatternDemo{}
	shapeFactory := new(FlyweightPattern.ShapeFactory)
	for i := 0; i < 20; i++ {
		circle := shapeFactory.GetCircle(flyweightPatternDemo.GetRandomColor())
		circle.SetX(flyweightPatternDemo.GetRandomXAndY())
		circle.SetY(flyweightPatternDemo.GetRandomXAndY())
		circle.SetRadius(100)
		circle.Draw3()
	}
}

//代理设计模式
func TestProxyPattern(t *testing.T) {
	proxyImage := ProxyPattern.NewProxyImage("1.jpg")
	proxyImage.Display()
	fmt.Println("...")
	proxyImage.Display()
}

//责任链模式
func TestChainofResponsibilityPattern(t *testing.T) {
	logger := ChainofResponsibilityPattern.TestChainofResponsibilityPattern()

	logger.LogMessage(ChainofResponsibilityPattern.INFO, "This is an information.", nil)

	logger.LogMessage(ChainofResponsibilityPattern.DEBUG, "This is a debug level information.", nil)

	logger.LogMessage(ChainofResponsibilityPattern.ERROR, "This is an error information.", nil)
}

//命令模式
func TestCommandPattern(t *testing.T) {
	abcStock := CommandPattern.NewStock()

	buyStockOrder := new(CommandPattern.BuyStock)
	buyStockOrder.BuyStock(abcStock)

	sellStockOrder := new(CommandPattern.SellStock)
	sellStockOrder.SellStock(abcStock)

	broker := new(CommandPattern.Broker)
	broker.TakeOrder(buyStockOrder)
	broker.TakeOrder(sellStockOrder)

	broker.PlaceOrders()
}

//解释器模式
func TestInterpreterPattern(t *testing.T) {
	isMale := InterpreterPattern.GetMaleExpression()
	isMarriedWoman := InterpreterPattern.GetMarriedWomanExpression()
	fmt.Println("John is male? ", isMale.Interpret("John"))
	fmt.Println("Julie is a married women? ", isMarriedWoman.Interpret("Married Julie"))
}

//迭代模式
func TestIteratorPattern(t *testing.T) {
	namesRepository := new(IteratorPattern.NameRepository)
	for iter := namesRepository.GetIterator(); iter.HasNext(); {
		fmt.Println(iter.Next())
	}
}

//中介者模式
func TestMediatorPattern(t *testing.T) {
	robert := new(MediatorPattern.User)
	robert.User("Robert")

	john := new(MediatorPattern.User)
	john.User("John")

	robert.SendMessage("Hi John")
	john.SendMessage("Hi Robert")
}

//备忘录模式
func TestMementoPattern(t *testing.T) {
	originator := new(MementoPattern.Originator)
	careTaker := new(MementoPattern.CareTaker)
	originator.SetState("State #1")
	originator.SetState("State #2")
	careTaker.Add(originator.SaveStateToMemento())
	originator.SetState("State #3")
	careTaker.Add(originator.SaveStateToMemento())
	originator.SetState("State #4")

	fmt.Println("Current State: " + originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(0))
	fmt.Println("First saved State: " + originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Println("Second saved State: " + originator.GetState())
}

//观察者模式
func TestObserverPattern(t *testing.T) {
	subject := new(ObserverPattern.Subject)
	new(ObserverPattern.HexaObserver).HexaObserver(subject)
	new(ObserverPattern.OctalObserver).OctalObserver(subject)
	new(ObserverPattern.BinaryObserver).BinaryObserver(subject)
	fmt.Println("First state change:15")
	subject.SetState(15)
	fmt.Println("Second state change:10")
	subject.SetState(10)
}

//状态模式
func TestStatePattern(t *testing.T) {
	context := new(StatePattern.Context)

	startState := new(StatePattern.StartState)
	startState.DoAction(context)
	fmt.Println(context.GetState().ToString())

	stopState := new(StatePattern.StopState)
	stopState.DoAction(context)

	fmt.Println(context.GetState().ToString())
}

//空对象模式
func TestNullObjectPattern(t *testing.T) {
	customerFactory := new(NullObjectPattern.CustomerFactory)
	customerFactory.Names = []string{"Rob", "Joe", "Julie"}

	customer1 := customerFactory.GetCustomer("Rob")
	customer2 := customerFactory.GetCustomer("Bob")
	customer3 := customerFactory.GetCustomer("Julie")
	customer4 := customerFactory.GetCustomer("Laura")

	fmt.Println(customer1.GetName())
	fmt.Println(customer2.GetName())
	fmt.Println(customer3.GetName())
	fmt.Println(customer4.GetName())
}

//策略模式
func TestStrategyPattern(t *testing.T) {
	context := new(StrategyPattern.Context).Context(new(StrategyPattern.OperationAdd))
	fmt.Println("10+5:", context.ExecuteStrategy(10, 5))

	context.Context(new(StrategyPattern.OperationSubstract))
	fmt.Println("10-5:", context.ExecuteStrategy(10, 5))

	context.Context(new(StrategyPattern.OperationMultiply))
	fmt.Println("10*5:", context.ExecuteStrategy(10, 5))
}

//模版模式
func TestTemplatePattern(t *testing.T) {
	cricket := new(TemplatePattern.Cricket)
	cricket.Play()

	football := new(TemplatePattern.Football)
	football.Play()
}

//访问者模式
func TestVisitorPattern(t *testing.T) {
	computer := new(VisitorPattern.Computer).Computer()
	computer.Accept(new(VisitorPattern.ComputerPartDisplayVisitor))
}
