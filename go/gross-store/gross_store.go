package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the AddItem() function")
	_, exists := units[unit]
	if !exists {
		return false
	}
	_, exists = bill[item]
	if !exists {
		bill[item] = units[unit]
	} else {
		bill[item] += units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExists := bill[item]
	if !itemExists {
		return false
	}
	_, unitExists := units[unit]
	if !unitExists {
		return false
	}
	newQuantity := bill[item] - units[unit]
	switch {
	case newQuantity == 0:
		delete(bill, item)
		return true
	case newQuantity > 0:
		bill[item] = newQuantity
		return true
	default:
		return false
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}
