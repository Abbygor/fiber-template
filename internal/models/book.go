package models

type Book struct {
	BookID      int    `gorm:"primaryKey;autoIncrement;;column:book_id" json:"book_id"` // SERIAL PRIMARY KEY
	Title       string `gorm:"type:varchar(255);not null;column:title" json:"title"`    // VARCHAR(255) NOT NULL
	Genre       string `gorm:"type:varchar(100);column:genre" json:"genre"`             // VARCHAR(100)
	PublishDate string `gorm:"type:date;;column:publish_date" json:"publish_date"`      // DATE
	AuthorID    int    `gorm:"not null;;column:author_id" json:"author_id"`             // INT (Foreign Key)
	//Author      Author `gorm:"foreignKey:AuthorID;references:AuthorID"`                 // Relación con el autor
}

func (Book) TableName() string {
	return "books"
}
