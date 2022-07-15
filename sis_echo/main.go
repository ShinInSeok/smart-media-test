package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", GetUsers)

	e.GET("/myid/:id", getID)

	e.GET("/show", show)

	e.POST("/save", save)

	e.POST("/save_file", save_file)

	e.POST("/connect_mysql", connectMySQL)

	e.POST("/Automigration", connectMySQLbyORM)

	e.POST("/CreateUser", CreateUser)

	e.GET("/select_user", selectUser)

	e.PUT("/update_user", updateUser)

	e.DELETE("/Delete_user", deleteUser)
	//==================================== Todo
	e.POST("/create_todo", createdTodos)

	e.GET("/select_todo", selectTodos)

	e.PUT("/update_todo", updateTodos)

	e.DELETE("/delete_todo", deleteTodos)

	e.Logger.Fatal(e.Start(":1323"))
}

func GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "user 정보")
}

func getID(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

type User struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email

	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	return c.String(http.StatusOK, "name:"+u.Name+", email:"+u.Email)
}

func save_file(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

func connectMySQL(c echo.Context) error {

	db, err := sql.Open("mysql", "root:1234@tcp(192.168.35.152:3306)/todo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from User")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var Id int
	var Name string
	var Email string

	for rows.Next() {
		err := rows.Scan(&Id, &Name, &Email)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Id : %d ,User: %s, Email: %s \n", Id, Name, Email)
	}

	return c.HTML(http.StatusOK, "<b>Thank You!</b>")

}

func connectMySQLbyORM(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	return c.HTML(http.StatusOK, "<b>Created</b>")
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
}

func CreateUser(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	user := User{
		Name:  "sis",
		Email: "is7170@naver.com",
	}

	result := db.Create(&user)

	fmt.Println(result.RowsAffected)

	return c.HTML(http.StatusOK, "<b>Created</b>")
}

func selectUser(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	user := new([]User)
	//user := map[string]interface{}{}

	result := db.Find(&user)

	fmt.Println(result.RowsAffected)

	return c.HTML(http.StatusOK, "<b>User: "+fmt.Sprintf("%v", user)+"</b>")

}

func updateUser(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	user := new(User)

	db.Where("name=?", "sis").First(&user)

	user.Email = "update@test.com"

	db.Where("name=?", "sis").Save(&user)

	return c.HTML(http.StatusOK, "<b>Updated</b>")
}

func deleteUser(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/smart_media"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	db.Where("name=?", "sis").Delete(&User{})

	return c.HTML(http.StatusOK, "<b>Delete</b>")

}

type Todos struct {
	gorm.Model
	UserID    string
	StartDate time.Time
	EndDate   time.Time
	Title     string
	Status    string
}

func createdTodos(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?charset=utf8mb4"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Todos{})

	todos := Todos{
		UserID:    "sis",
		StartDate: time.Now(),
		EndDate:   time.Now(),
		Title:     "Todolist",
		Status:    "playing",
	}

	result := db.Create(&todos)

	fmt.Println(result.RowsAffected)

	return c.HTML(http.StatusOK, "<b>CreatedTodo</b>")

}

func selectTodos(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?charset=utf8mb4"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Todos{})

	todo := new([]Todos)
	//user := map[string]interface{}{}

	result := db.Find(&todo)

	fmt.Println(result.RowsAffected)

	return c.HTML(http.StatusOK, "<b>SelectTodo</b>")

}
func updateTodos(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?charset=utf8mb4"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Todos{})

	todos := new(Todos)

	db.Where("user_id=?", "sis").First(&todos)

	todos.Status = "Update Compite!"

	db.Where("user_id=?", "sis").Save(&todos)

	return c.HTML(http.StatusOK, "<b>UpdatedTodo</b>")
}
func deleteTodos(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?charset=utf8mb4"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Todos{})

	db.Where("id=", 2).Delete(&Todos{})

	return c.HTML(http.StatusOK, "<b>DeleteTodo</b>")

}
