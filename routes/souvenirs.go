package routes

type Souvenir struct {
	Name      string
	Date      string
	PhotoPath string
}

func NewSouvenir(name, date, photoPath string) *Souvenir {
	return &Souvenir{
		Name:      name,
		Date:      date,
		PhotoPath: photoPath,
	}
}

func generateSouvenirs() []Souvenir {
	return []Souvenir{
		*NewSouvenir("Paris1", "09/01/2021", "IMG_1206"),
		*NewSouvenir("Randos1", "12/06/2021", "IMG_2012"),
		*NewSouvenir("LeLot1", "19/06/2021", "IMG_2090"),
		*NewSouvenir("Famille", "24/07/2021", "IMG_2275"),
		*NewSouvenir("LaReunion", "05/09/2021", "IMG_2957"),
		*NewSouvenir("Noel1", "24/12/2021", "IMG_3846"),
		*NewSouvenir("Taiga2", "16/06/2022", "IMG_4533"),
		*NewSouvenir("Suisse", "15/07/2022", "IMG_4738"),
		*NewSouvenir("Malte", "31/08/2022", "IMG_5018"),
		*NewSouvenir("Beauval", "24/09/2022", "IMG_5225"),
		*NewSouvenir("Noel2", "24/12/2022", "IMG_5429"),
		*NewSouvenir("Correze", "24/05/2023", "IMG_5876"),
		*NewSouvenir("Mariages", "15/07/2023", "IMG_4800"),
		*NewSouvenir("Croatia", "02/09/2023", "IMG_6373"),
		*NewSouvenir("Halloween", "28/10/2023", "IMG_6640"),
		*NewSouvenir("BONUS", "18/12/2020-INFINI", "IMG_"),
	}
}
