package main

type Animal struct {
	Name string
	mean bool
}

type AnimalSounder interface {
	MakeNoise()
}

//匿名组合: 继承Animal的方法
type Dog struct {
	Animal
	BarkStrength int
}

type Catt struct {
	Basics Animal
	MeowStrength int
}
//添加Animal的方法
func (animal *Animal) PerformNoise(strength int, sound string) {
	if animal.mean == true {
		strength = strength * 2
	}
	for voice := 0; voice < strength; voice++ {
		println(sound)
	}
	println("")
}
//实现接口继承、多肽
func (dog *Dog) MakeNoise()  {
	dog.PerformNoise(dog.BarkStrength, "BARK")
}

func (cat *Catt) MakeNoise() {
	cat.Basics.PerformNoise(cat.MeowStrength, "MEOW")
}

//声明函数，参数为接口，根据对象的不同调用对象自己的方法
func MakeSomeNoise(animalsounder AnimalSounder) {
	animalsounder.MakeNoise()
}

func main() {
	//匿名组合的赋值
	dog := &Dog{Animal{Name: "saw", mean: false}, 2}
	cat := &Catt{Basics: Animal{Name: "kitty", mean: true}, MeowStrength: 2}
	MakeSomeNoise(dog)
	MakeSomeNoise(cat)
}
