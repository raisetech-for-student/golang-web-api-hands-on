package model

type Book struct {
	id    int
	name  string
	price int
}

func NewBook(id int, name string, price int) Book {
	return Book{
		id:    id,
		name:  name,
		price: price,
	}
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Name() string {
	return b.name
}

func (b *Book) Price() int {
	return b.price
}
