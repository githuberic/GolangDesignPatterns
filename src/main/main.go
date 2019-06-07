package main

import (
	"AdapterPattern"
	"Bridge"
	"BuilderPattern"
	"ChainofResponsibilityPattern"
	"CommandPattern"
	"CompositePattern"
	"DecoratorPattern"
	"FacadePattern"
	"FactoryPattern"
	"FilterPattern"
	"FlyweightPattern"
	"InterpreterPattern"
	"IteratorPattern"
	"MediatorPattern"
	"MementoPattern"
	"ObserverPattern"
	"ProxyPattern"
	"SingletonPattern"
	"StatePattern"
	"fmt"
)

func main() {
	testStatePattern()
}

//简单工厂
func testSimplenessFactory() {
	f := new(FactoryPattern.SimplenessFactory)
	var s FactoryPattern.Shape
	s, ok := f.GetShape("Rectangle")
	if ok {
		s.Draw()
	}
}

//抽象工厂
func testAbstractFactory() {
	f := new(FactoryPattern.AbsFactory)
	color := f.GetColor("Red")
	color.Fill()
	shape := f.GetShape("Rectangle")
	shape.Draw()
}

//单例
func testSingleton() {
	s1 := SingletonPattern.GetInstance1()
	s1.Count = 5
	fmt.Println(s1)
	s2 := SingletonPattern.GetInstance1()
	fmt.Println(s2)
}

//建造者
func testBuilderPattern() {
	builder := new(BuilderPattern.ComputerBuilder)
	director := BuilderPattern.Director{Builder: builder}
	computer := director.Create("I7", "32G", "4T")
	fmt.Println(*computer)
}

//适配器
func testAdapterPattern() {
	audioPlayer := AdapterPattern.AudioPlayer{}

	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}

//桥接
func testBridge() {
	redCircle := Bridge.Circle{}
	redCircle.Circle(100, 100, 10, &Bridge.RedCircle{})
	greenCircle := Bridge.Circle{}
	greenCircle.Circle(100, 100, 10, &Bridge.GreenCircle{})
	redCircle.Draw()
	greenCircle.Draw()
}

//过滤器
func testFilterPattern() {
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
func testCompositePattern() {
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
func testDecoratorPattern() {
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
func testFacadePattern() {
	shapeMaker := new(FacadePattern.ShapeMaker)
	shapeMaker.ShapeMaker()

	shapeMaker.DrawCircle()
	shapeMaker.DrawRectangle()
	shapeMaker.DrawSquare()
}

//享元设计模式
func testFlyweightPattern() {
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
func testProxyPattern() {
	proxyImage := ProxyPattern.NewProxyImage("1.jpg")
	proxyImage.Display()
	fmt.Println("...")
	proxyImage.Display()
}

//责任链模式
func testChainofResponsibilityPattern() {
	logger := ChainofResponsibilityPattern.TestChainofResponsibilityPattern()

	logger.LogMessage(ChainofResponsibilityPattern.INFO, "This is an information.", nil)

	logger.LogMessage(ChainofResponsibilityPattern.DEBUG, "This is a debug level information.", nil)

	logger.LogMessage(ChainofResponsibilityPattern.ERROR, "This is an error information.", nil)
}

//命令模式
func testCommandPattern() {
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
func testInterpreterPattern() {
	isMale := InterpreterPattern.GetMaleExpression()
	isMarriedWoman := InterpreterPattern.GetMarriedWomanExpression()
	fmt.Println("John is male? ", isMale.Interpret("John"))
	fmt.Println("Julie is a married women? ", isMarriedWoman.Interpret("Married Julie"))
}

//迭代模式
func testIteratorPattern() {
	namesRepository := new(IteratorPattern.NameRepository)
	for iter := namesRepository.GetIterator(); iter.HasNext(); {
		fmt.Println(iter.Next())
	}
}

//中介者模式
func testMediatorPattern() {
	robert := new(MediatorPattern.User)
	robert.User("Robert")

	john := new(MediatorPattern.User)
	john.User("John")

	robert.SendMessage("Hi John")
	john.SendMessage("Hi Robert")
}

//备忘录模式
func testMementoPattern() {
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
func testObserverPattern() {
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
func testStatePattern()  {
	context:=new(StatePattern.Context)

	startState:=new(StatePattern.StartState)
	startState.DoAction(context)
	fmt.Println(context.GetState().ToString())

	stopState:=new(StatePattern.StopState)
	stopState.DoAction(context)

	fmt.Println(context.GetState().ToString())
}
