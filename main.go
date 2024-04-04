package main

import (
	"PnoT/auth"
	db_file "PnoT/db/file"
	"PnoT/util"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func BanchMark() {
	// e := auth.CreateUser("test", "test")
	// fmt.Println(e)
	// t, e := auth.LoginUser("test", "test")
	// fmt.Println(t, e)
	// t, e = auth.ValidateJWT(t)
	// fmt.Println(t, e)

	var wg sync.WaitGroup

	n := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e := db_file.AddFile("aa", fmt.Sprintf("test%d", i), "## test", false)
			if e != nil {
				fmt.Println("Error:", e)
			}
		}(i)
	}

	wg.Wait()
	et := time.Now()
	fmt.Println("100 Item Add", et.Sub(n))

	n = time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, e := db_file.GetFile("aa", fmt.Sprintf("test%d", 1))
			if e != nil {
				fmt.Println("Error:", e)
			}
		}(i)
	}

	wg.Wait()
	et = time.Now()
	fmt.Println("100 Item Get", et.Sub(n))

	n = time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				_, e := db_file.GetFile("aa", fmt.Sprintf("test%d", 1))
				if e != nil {
					fmt.Println("Error:", e)
				}
			} else {
				e := db_file.AddFile("aa", fmt.Sprintf("test%d", i), "## test", false)
				if e != nil {
					fmt.Println("Error:", e)
				}
			}
		}(i)
	}

	wg.Wait()
	et = time.Now()
	fmt.Println("100 Item Get and Add", et.Sub(n))
}

func main() {
	r := gin.Default()

	a := r.Group("api/")
	{
		at := a.Group("auth/")
		{
			at.POST("login", func(c *gin.Context) {
				type req struct {
					Username string `json:"username" binding:"required"`
					Password string `json:"password" binding:"required"`
				}
				var data req
				if util.BindJSON(c, &data) != nil {
					util.Error(c, 400, "Invalid request")
					return
				}
				fmt.Println(data)

				token, err := auth.LoginUser(data.Username, data.Password)
				if err != nil {
					util.Error(c, 400, "Invalid username or password")
					return
				}
				c.SetCookie("token", token, 3600, "/", "localhost", false, true)
				c.JSON(200, gin.H{
					"token": token,
				})
			})

			at.POST("register", func(c *gin.Context) {
				type req struct {
					Username string `json:"username" binding:"required"`
					Password string `json:"password" binding:"required"`
				}
				var data req
				if util.BindJSON(c, &data) != nil {
					util.Error(c, 400, "Invalid request")
					return
				}

				err := auth.CreateUser(data.Username, data.Password)
				if err != nil {
					util.Error(c, 400, "Username already exists")
					return
				}
				c.JSON(200, gin.H{
					"message": "User created",
				})
			})
		}
		a.Use(func(c *gin.Context) {
			token, err := c.Cookie("token")
			if err != nil {
				// util.Error(c, 400, "Invalid token")
				return
			}
			username, err := auth.ValidateJWT(token)
			if err != nil {
				// util.Error(c, 400, "Invalid token")
				return
			}
			c.Set("username", username)
			c.Next()
		})

		a.GET("/file", func(c *gin.Context) {
			user := c.Query("username")
			path := c.Query("path")
			if path == "" {
				util.Error(c, 400, "Invalid path")
				return
			}
			fmt.Println(user, path)
			file, err := db_file.GetFile(user, path)
			if err != nil || (!file.Public && c.GetString("username") != file.Autor) {
				fmt.Println(err)
				util.Error(c, 404, "File not found")
				return
			}

			c.JSON(200, file)
		})
		a.POST("/file", func(c *gin.Context) {
			type req struct {
				Path    string `json:"path" binding:"required"`
				Content string `json:"content" binding:"required"`
				Public  bool   `json:"public"`
			}
			var data req
			e := util.BindJSON(c, &data)
			if e != nil {
				fmt.Println(e)
				util.Error(c, 400, "Invalid request")
				return
			} else if data.Path == "" {
				util.Error(c, 400, "Invalid path")
				return
			} else if data.Content == "" {
				util.Error(c, 400, "Invalid content")
				return
			} else if c.GetString("username") == "" {
				util.Error(c, 400, "Invalid user")
				return
			}

			if db_file.FileExists(c.GetString("username"), data.Path) {
				err := db_file.UpdateFile(c.GetString("username"), data.Path, data.Content, data.Public)
				if err != nil {
					util.Error(c, 400, "Wow")
					return
				}
				c.JSON(200, gin.H{
					"message": "File updated",
				})
			} else {
				err := db_file.AddFile(c.GetString("username"), data.Path, data.Content, data.Public)
				if err != nil {
					util.Error(c, 400, "Super Wow")
					return
				}
				c.JSON(200, gin.H{
					"message": "File created",
				})
			}
		})
	}

	r.Run(":8080")
}
