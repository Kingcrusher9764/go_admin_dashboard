package dataset

type StatusType int 
const (
    Active StatusType = iota
    Paid
    Shipped
    Delivered
)
func (s StatusType) String() string{
    return [...]string{"Active", "Paid", "Shipped", "Delivered"}[s]
}
func (s StatusType) EnumIndex() int{
    return int(s)
}

type Customer struct{
    OrderId int
    Name, 
    Email,
    Product string
    Status StatusType
    Quantity,
    Price int
    Origin, 
    CreatedAt string
}
var Customers = []Customer {
    {
        OrderId: 1,
        Name: "Alan",
        Email: "alan@gmail.com",
        Product: "Product 1",
        Status: Active,
        Quantity: 90,
        Price: 280,
        Origin: "Dash",
        CreatedAt: "date and time",
    },
    {
        OrderId: 2,
        Name: "Alan2",
        Email: "alan@gmail.com",
        Product: "Product 1",
        Status: Active,
        Quantity: 90,
        Price: 280,
        Origin: "Dash",
        CreatedAt: "date and time",
    },
    {
        OrderId: 3,
        Name: "Alan3",
        Email: "alan@gmail.com",
        Product: "Product 1",
        Status: Delivered,
        Quantity: 90,
        Price: 280,
        Origin: "Dash",
        CreatedAt: "date and time",
    },
    {
        OrderId: 4,
        Name: "Alan4",
        Email: "alan@gmail.com",
        Product: "Product 1",
        Status: Shipped,
        Quantity: 90,
        Price: 280,
        Origin: "Dash",
        CreatedAt: "date and time",
    },
    {
        OrderId: 5,
        Name: "Alan5",
        Email: "alan@gmail.com",
        Product: "Product 1",
        Status: Active,
        Quantity: 90,
        Price: 280,
        Origin: "Dash",
        CreatedAt: "date and time",
    },
    {
        OrderId: 6,
        Name: "Alan6",
        Email: "alan@gmail.com",
        Product: "Product 1",
        Status: Paid,
        Quantity: 90,
        Price: 280,
        Origin: "Web",
        CreatedAt: "date and time",
    },
}
