package temp

type CustomerService struct {
	customers []Customers
	customerNum int
}

func NewCustomerService() *CustomerService  {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := NewCustomer(1,"张三","男",20,"11214563254","5346@.com")
	customerService.customers = append(customerService.customers,customer)
	return customerService
}
func (this *CustomerService) List() []Customers{
	return this.customers
}
func (this *CustomerService)Add(customer Customers) bool  {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return  true
}
func (this *CustomerService)FindById(id int) int {
	index := -1
	for i := 0;i< len(this.customers) ;i++  {
		if this.customers[i].Id == id{
			index = i
		}
	}
	return index
}
func (this *CustomerService )Delete(id int)bool  {
	index := this.FindById(id)
	if index == -1 {
		return false
	}else {
		this.customers = append(this.customers[0:index],this.customers[index+1 :]...)
	}
	return true
}

func(this *CustomerService ) AlterName(id int,name string) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}else {
		this.customers[index].Name = name
	}
	return true
}
func(this *CustomerService ) AlterGender(id int,gender string) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}else {
		this.customers[index].Gender = gender
	}
	return true
}
func(this *CustomerService ) AlterAge(id int,age int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}else {
		this.customers[index].Age = age
	}
	return true
}
func(this *CustomerService ) AlterPhone(id int,phone string) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}else {
		this.customers[index].Phone = phone
	}
	return true
}
func(this *CustomerService ) AlterEmail(id int,email string) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}else {
		this.customers[index].Email = email
	}
	return true
}