package main

import "fmt"
type BrowserHistory struct {
    CurrentPlace int 
    History []string
    Homepage string
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{CurrentPlace: 0, History: []string{}, Homepage: homepage}
}


func (this *BrowserHistory) Visit(url string)  {
	this.History = append(this.History, url)
	this.CurrentPlace += 1 
    return 
}


func (this *BrowserHistory) Back(steps int) (string, error) {
	// let's make sure we don't go less than 0 
	moveBack := this.CurrentPlace - steps 
	if moveBack - 1 < 0 {
		return "", fmt.Errorf("we are not able to move backwards")
	}
	// let's check if it's zero 
	if moveBack == 0 {
		this.CurrentPlace = 0 
		return this.History[0], nil 
	}
	
	// handle the rest of stuff 
	this.CurrentPlace = moveBack
    return this.History[moveBack - 1], nil 
}


func (this *BrowserHistory) Forward(steps int) (string, error) {
	// let's make sure that we don't move out of bounds 
	moveTo := this.CurrentPlace + steps 
	if moveTo > len(this.History) {
		return "", fmt.Errorf("we are not able to move forward")
	}

	this.CurrentPlace = moveTo
    return this.History[this.CurrentPlace - 1], nil 
}

func (this *BrowserHistory) PrintHistory() {
	for key, value := range this.History {
		fmt.Printf("%v = %v\n", key, value)
	}
}

func (this *BrowserHistory) PrintCurrentPlace() {
	fmt.Printf("current place = %v\n", this.CurrentPlace)
}

func main() {
	history := Constructor("test.com")
	history.Visit("facebook.com")
	history.PrintHistory()
	history.PrintCurrentPlace()
	fmt.Println("---------")
	history.Visit("test2.com")
	history.Visit("test3.com")
	history.Visit("boom.com")
	history.PrintHistory()
	history.PrintCurrentPlace()

	fmt.Println("---------")


	fmt.Printf("we are currently at %s\n", history.History[history.CurrentPlace - 1])
	backPlace, err := history.Back(1)
	if err != nil {
		fmt.Printf("an error occrurred while trying to move back: %v\n", err)
		return 
	}
	fmt.Printf("we moved back to %s\n", backPlace)

	fmt.Println("---------")

	fmt.Printf("we are currently at %s\n", history.History[history.CurrentPlace - 1])
	forwardPlace, err := history.Forward(1)
	if err != nil {
		fmt.Printf("an error occrurred while trying to move forward: %v\n", err)
		return 
	}
	fmt.Printf("we moved to forward %s\n", forwardPlace)

	fmt.Println("---------")

	fmt.Printf("we are currently at %s\n", history.History[history.CurrentPlace - 1])
	backPlace, err = history.Back(1)
	if err != nil {
		fmt.Printf("an error occrurred while trying to move back: %v\n", err)
		return 
	}
	fmt.Printf("we moved back to %s\n", backPlace)

	fmt.Println("---------")

	fmt.Println("---------")

	fmt.Printf("we are currently at %s\n", history.History[history.CurrentPlace - 1])
	backPlace, err = history.Back(1)
	if err != nil {
		fmt.Printf("an error occrurred while trying to move back: %v\n", err)
		return 
	}
	fmt.Printf("we moved back to %s\n", backPlace)

	fmt.Println("---------")

	fmt.Println("---------")

	fmt.Printf("we are currently at %s\n", history.History[history.CurrentPlace - 1])
	backPlace, err = history.Back(1)
	if err != nil {
		fmt.Printf("an error occrurred while trying to move back: %v\n", err)
		return 
	}
	fmt.Printf("we moved back to %s\n", backPlace)

	fmt.Println("---------")

	fmt.Println("---------")

	fmt.Printf("we are currently at %s\n", history.History[history.CurrentPlace - 1])
	backPlace, err = history.Back(1)
	if err != nil {
		fmt.Printf("an error occrurred while trying to move back: %v\n", err)
		return 
	}
	fmt.Printf("we moved back to %s\n", backPlace)

	fmt.Println("---------")

	fmt.Println("---------")

	fmt.Printf("we are currently at %s\n", history.History[history.CurrentPlace - 1])
	backPlace, err = history.Back(1)
	if err != nil {
		fmt.Printf("an error occrurred while trying to move back: %v\n", err)
		return 
	}
	fmt.Printf("we moved back to %s\n", backPlace)

	fmt.Println("---------")
}