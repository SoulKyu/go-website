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
		*NewSouvenir("Paris2", "15/02/2022", "IMG_4025"),
		*NewSouvenir("Taiga1", "17/04/2022", "IMG_4429"),
		*NewSouvenir("Taiga2", "16/06/2022", "IMG_"),
		*NewSouvenir("Suisse", "15/07/2022", "IMG_"),
		*NewSouvenir("Malte", "31/08/2022", "IMG_"),
		*NewSouvenir("Beauval", "24/09/2022", "IMG_"),
		*NewSouvenir("Noel2", "24/12/2022", "IMG_"),
		*NewSouvenir("Randos2", "15/04/2023", "IMG_"),
		*NewSouvenir("Correze", "24/05/2023", "IMG_"),
		*NewSouvenir("Mariage1", "15/07/2023", "IMG_"),
		*NewSouvenir("Mariage2", "20/08/2023", "IMG_"),
		*NewSouvenir("Croatia", "02/09/2023", "IMG_"),
		*NewSouvenir("Halloween", "28/10/2023", "IMG_"),
		*NewSouvenir("BONUS", "18/12/2020-INFINI", "IMG_"),

		// Ajoutez d'autres souvenirs ici
	}
}
