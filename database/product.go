package database

var ProductList []Product

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgURL      string
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Description 1",
		Price:       100.00,
		ImgURL:      "https://static.wikia.nocookie.net/colors/images/4/42/Orange_%28fruit%29.png/revision/latest/scale-to-width-down/985?cb=20250701034836",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Description 2",
		Price:       200.00,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSsT3iUXGmyhhGe8I6l5RAu1SEoz6ybubyKKQ&s",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Mango",
		Description: "Description 3",
		Price:       300.00,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Banana",
		Description: "Description 4",
		Price:       400.00,
		ImgURL:      "https://dtgxwmigmg3gc.cloudfront.net/imagery/assets/derivations/icon/512/512/true/eyJpZCI6IjIzNjFhNjAyM2YxMmY1OGM5NzcyNTgwNjM3YWZiNjAxLmpwZyIsInN0b3JhZ2UiOiJwdWJsaWNfc3RvcmUifQ?signature=7d201f5f9b7458fa9243e619222f7453c4266c06112ca287d01adaf32bdb2d1f",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Pineapple",
		Description: "Description 5",
		Price:       500.00,
		ImgURL:      "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/organic%20pineaple.png?itok=4wB5t6Tk-24IDL_Zc",
	}

	ProductList = append(ProductList, prd1, prd2, prd3, prd4, prd5)

}
