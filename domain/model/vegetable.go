package model

// Vegetable Model
type Vegetable struct {
	BaseModel
	ShopID         uint
	SellID         uint
	BuyID          uint
	Name           string
	Fee            int64
	IsChemical     bool
	ImagePath      string
	ProductionDate string
	Category       string
}

type CategorizedVegetable struct {
	Name           string
	Fee            int64
	IsChemical     bool
	ImagePath      string
	ProductionDate string
	Category       string
	Amount         int64
}

func CategorizeVegetable(v []Vegetable) ([]CategorizedVegetable, error) {
	buf := []CategorizedVegetable{}
	squashes := []Vegetable{}
	cabbages := []Vegetable{}
	sweetPotatos := []Vegetable{}
	potatos := []Vegetable{}
	radishes := []Vegetable{}
	onions := []Vegetable{}
	carrots := []Vegetable{}
	bellPeppers := []Vegetable{}
	spinaches := []Vegetable{}
	lettuces := []Vegetable{}

	for _, a := range v {
		if a.Category == "SQUASH" {
			squashes = append(squashes, a)
		} else if a.Category == "CABBAGE" {
			cabbages = append(cabbages, a)
		} else if a.Category == "SWEETPOTATO" {
			sweetPotatos = append(sweetPotatos, a)
		} else if a.Category == "POTATO" {
			potatos = append(potatos, a)
		} else if a.Category == "RADISH" {
			radishes = append(radishes, a)
		} else if a.Category == "ONION" {
			onions = append(onions, a)
		} else if a.Category == "CARROT" {
			carrots = append(carrots, a)
		} else if a.Category == "BELLPEPPER" {
			bellPeppers = append(bellPeppers, a)
		} else if a.Category == "SPINACH" {
			spinaches = append(spinaches, a)
		} else if a.Category == "LETTUCE" {
			lettuces = append(lettuces, a)
		}
	}

	if count := len(squashes); count != 0 {
		squash := CategorizedVegetable{
			Name:           squashes[0].Name,
			Fee:            squashes[0].Fee,
			IsChemical:     squashes[0].IsChemical,
			ImagePath:      squashes[0].ImagePath,
			ProductionDate: squashes[0].ProductionDate,
			Category:       squashes[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, squash)
	}
	if count := len(cabbages); count != 0 {
		cabbage := CategorizedVegetable{
			Name:           cabbages[0].Name,
			Fee:            cabbages[0].Fee,
			IsChemical:     cabbages[0].IsChemical,
			ImagePath:      cabbages[0].ImagePath,
			ProductionDate: cabbages[0].ProductionDate,
			Category:       cabbages[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, cabbage)
	}
	if count := len(sweetPotatos); count != 0 {
		sweetPotato := CategorizedVegetable{
			Name:           sweetPotatos[0].Name,
			Fee:            sweetPotatos[0].Fee,
			IsChemical:     sweetPotatos[0].IsChemical,
			ImagePath:      sweetPotatos[0].ImagePath,
			ProductionDate: sweetPotatos[0].ProductionDate,
			Category:       sweetPotatos[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, sweetPotato)
	}
	if count := len(potatos); count != 0 {
		potato := CategorizedVegetable{
			Name:           potatos[0].Name,
			Fee:            potatos[0].Fee,
			IsChemical:     potatos[0].IsChemical,
			ImagePath:      potatos[0].ImagePath,
			ProductionDate: potatos[0].ProductionDate,
			Category:       potatos[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, potato)
	}
	if count := len(radishes); count != 0 {
		radish := CategorizedVegetable{
			Name:           radishes[0].Name,
			Fee:            radishes[0].Fee,
			IsChemical:     radishes[0].IsChemical,
			ImagePath:      radishes[0].ImagePath,
			ProductionDate: radishes[0].ProductionDate,
			Category:       radishes[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, radish)
	}
	if count := len(onions); count != 0 {
		onion := CategorizedVegetable{
			Name:           onions[0].Name,
			Fee:            onions[0].Fee,
			IsChemical:     onions[0].IsChemical,
			ImagePath:      onions[0].ImagePath,
			ProductionDate: onions[0].ProductionDate,
			Category:       onions[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, onion)
	}
	if count := len(carrots); count != 0 {
		carrot := CategorizedVegetable{
			Name:           carrots[0].Name,
			Fee:            carrots[0].Fee,
			IsChemical:     carrots[0].IsChemical,
			ImagePath:      carrots[0].ImagePath,
			ProductionDate: carrots[0].ProductionDate,
			Category:       carrots[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, carrot)
	}
	if count := len(bellPeppers); count != 0 {
		bellPepper := CategorizedVegetable{
			Name:           bellPeppers[0].Name,
			Fee:            bellPeppers[0].Fee,
			IsChemical:     bellPeppers[0].IsChemical,
			ImagePath:      bellPeppers[0].ImagePath,
			ProductionDate: bellPeppers[0].ProductionDate,
			Category:       bellPeppers[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, bellPepper)
	}
	if count := len(spinaches); count != 0 {
		spinach := CategorizedVegetable{
			Name:           spinaches[0].Name,
			Fee:            spinaches[0].Fee,
			IsChemical:     spinaches[0].IsChemical,
			ImagePath:      spinaches[0].ImagePath,
			ProductionDate: spinaches[0].ProductionDate,
			Category:       spinaches[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, spinach)
	}
	if count := len(lettuces); count != 0 {
		lettuce := CategorizedVegetable{
			Name:           lettuces[0].Name,
			Fee:            lettuces[0].Fee,
			IsChemical:     lettuces[0].IsChemical,
			ImagePath:      lettuces[0].ImagePath,
			ProductionDate: lettuces[0].ProductionDate,
			Category:       lettuces[0].Category,
			Amount:         int64(count),
		}
		buf = append(buf, lettuce)
	}
	return buf, nil
}
