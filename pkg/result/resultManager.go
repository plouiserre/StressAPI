package result

import (
	"os"
	"strconv"
	"time"
)

//TODO correct bug for save during the first call
type ResultManager struct {
	StoreFolder string
	Result Result
}

func (rm ResultManager) StoreResult() bool{
	isOk := true 
	isExist := rm.Exists(rm.StoreFolder)
	if !isExist {
		err := os.Mkdir(rm.StoreFolder, 0222)
		if err != nil {
			isOk = false
		} 
	}
	if isOk {
		isOk = rm.SaveFile()		
	}
	return isOk
}

func (rm ResultManager) SaveFile() bool{
	isOk := true
	currentDate := time.Now()
	msg := rm.GetMessage()
	fileName := rm.StoreFolder + "\\" + currentDate.Format("20060102")+".txt"
	fileExists := rm.Exists(fileName)
	if !fileExists {
		datas := []byte(msg)
		err := os.WriteFile(fileName, datas, 0222)
 		if err != nil {
			isOk = false
		}
	} else {
		f, err := os.OpenFile(fileName, os.O_APPEND, 0222)
		if err != nil {
			isOk = false
		}
		defer f.Close()
		if isOk {
			_, err2 := f.WriteString(msg)
			if err2 != nil {
				isOk = false
			}
		}
	}
	return isOk
}

func (rm ResultManager) Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {     
		return true
	} else {
		return false
	}
}

func (rm ResultManager) GetMessage() string{
	dateNow := time.Now()
	dateMsg := dateNow.Format("2006-01-02 15:04:05.000000")	
	httpCode := strconv.Itoa(rm.Result.HttpCode)
	msg := dateMsg+": httpCode :"+httpCode+" timing request \""+ rm.Result.RequestDuration+ "\" req :\""+rm.Result.UriCalled+"\" body :\""+rm.Result.Body+"\"  response :"+rm.Result.Response
	return msg
}

func (rm *ResultManager)SetResult(result Result){
	rm.Result = result
}
