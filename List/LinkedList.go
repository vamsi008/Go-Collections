package List



const MAX_LENGHT = 10
type Node struct{

	Next *Node
	Value interface{}
}


type ArrayList struct{
	Value interface{}
}

func NewArrayList() ArrayList{
	var a [MAX_LENGHT]string
	list:=ArrayList{a}
	return list
}

func (list *ArrayList ) put(){

	

}


