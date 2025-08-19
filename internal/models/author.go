package models

type Author struct {
	AuthorID    int    `gorm:"primaryKey;autoIncrement;column:author_id" json:"author_id"`     // SERIAL PRIMARY KEY
	FirstName   string `gorm:"type:varchar(100);not null;column:first_name" json:"first_name"` // VARCHAR(100) NOT NULL
	LastName    string `gorm:"type:varchar(100);not null;column:last_name" json:"last_name"`   // VARCHAR(100) NOT NULL
	BirthDate   string `gorm:"type:date;column:birth_date" json:"birth_date"`                  // DATE
	Nationality string `gorm:"type:varchar(100);column:nationality" json:"nationality"`        // VARCHAR(100)
	//Books       []Book    `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relación con libros
}

func (Author) TableName() string {
	return "authors"
}
