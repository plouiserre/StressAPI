package result

type ResultManagerMock struct {
	IsStoreResultCalled bool
}

func (rm *ResultManagerMock) StoreResult() bool {
	rm.IsStoreResultCalled = true
	return true
}

func (rm ResultManagerMock) SaveFile() bool {
	return true
}

func (rm ResultManagerMock) Exists(path string) bool {
	return true
}

func (rm ResultManagerMock) GetMessage() string {
	return ""
}

func (rm *ResultManagerMock) SetResult(result Result) {

}
