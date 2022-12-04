package result

type IResultManager interface {
	StoreResult() bool
	SetResult(result Result)
	SaveFile() bool
	Exists(path string) bool
	GetMessage() string
}