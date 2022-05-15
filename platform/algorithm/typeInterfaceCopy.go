package algorithm

type CopyInterface interface {
	GetProcessed() (list *[]Point)
	GetOriginal() (list *[]Point)
}
