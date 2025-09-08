package main
import ("fmt"; "math/rand")


func main(){
	var ans string

	fmt.Println("Привет, я чат-бот. Чего хотите сегодня?")
	fmt.Println("1) Рекомендации фильмов")
	fmt.Println("2) Рекомендации музыки")
	fmt.Println("3) Рассказать анекдот")
	fmt.Println("4) Каталог мерча")

	for {
		fmt.Printf("\nВаш запрос?")
		fmt.Scan(&ans)

		if ans == "Выход" {
			fmt.Printf("\nДо свидания!")
			break
		}
		processCommand(ans)
	}
}

func processCommand (input string) {
		if input == "1" || input == "Рекомендации фильмов" || input == "фильмы" {
		fmt.Println("Фильмы по жанру или по настроению?")
		fmt.Scan(&input)
		if input == "Настроение" || input == "по настроению" {
		reccomendMoviesMood()
		} else if input == "Жанр" || input == "по жанру" {
		reccomendMoviesGenre()
		}
	} else if input == "2" || input == "Рекомендации музыки" || input == "музыка" {
		reccomendMoviesMood()
	} else if input == "3" || input == "Рассказать анекдот" || input == "анекдот" {
		fmt.Println("3) Рассказать анекдот")
	} else if input == "4" || input == "Каталог мерча" || input == "мерч" {
		fmt.Println("4) Каталог мерча")
	} 
}


func reccomendMoviesGenre() {
	genres:= []string{"Ужасы", "Комедия", "Драма", "Боевик", "Приключения", "Мультики"}

	movies := map[string][]string{
		"Ужасы": {"Крик", "ФНАФ", "Кошмар на улице Вязова"},
		"Комедия": {"Семь психопатов", "Достать ножи", "Паразиты"},
		"Драма": {"Один день", "После", "Пленницы"},
		"Боевик": {"Бесславные ублюдки", "План побега", "Стрелок"},
		"Приключения": {"Эверест", "Унесенные призраками", "Назад в будущее"},
		"Мультики": {"Валли", "Тайна Коко", "Тачки"},
	}
		genre := genres[rand.Intn(len(genres))]
		movie := movies[genre][rand.Intn(len(movies[genre]))]
		fmt.Printf("Рекомендую фильм '%s' в жанре %s\n", movie, genre)
}


func reccomendMoviesMood() {
	moods:= []string{"Грустное", "Веселое", "Вдохновляющее", "Романтическое", "Страшное"}

	movies := map[string][]string{
		"Страшное": {"Крик", "Пила 9", "Кошмар на улице Вязова"},
		"Веселое": {"Елки 9", "Один дома", "1+1"},
		"Романтическое": {"Один день", "После", "Пленницы"},
		"Грустное": {"Хатико", "Мальчик в полосатой пижаме", "Зеленая миля"},
		"Вдохновляющее": {"Джуманджи", "Унесенные призраками", "Назад в будущее"},
	}
	mood := moods[rand.Intn(len(moods))]
	movie := movies[mood][rand.Intn(len(movies[mood]))]
	fmt.Printf("Рекомендую фильм '%s' под %s настроение \n", movie, mood)
}