package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	//连接数据库
	dsn := "root:Wsy1817396846@tcp(127.0.0.1:3306)/go-crud-demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic("数据库连接失败！")
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	// 结构体
	type List struct {
		gorm.Model
		Name    string `gorm:"type:varchar(20);not null ;column:name" json:"name" binding:"required"`
		Status  string `gorm:"type:varchar(20);not null" json:"status" binding:"required"`
		Phone   string `gorm:"type:varchar(20);not null" json:"phone" binding:"required"`
		Email   string `gorm:"type:varchar(40);not null" json:"email" binding:"required"`
		Address string `gorm:"type:varchar(200);not null" json:"address" binding:"required"`
	}

	//迁移
	db.AutoMigrate(&List{})
	//接口
	r := gin.Default()

	//增加
	r.POST("/user/add", func(c *gin.Context) {
		var list List
		if err := c.ShouldBindJSON(&list); err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "参数错误",
				"data": gin.H{},
			})
			return
		}
		db.Create(&list)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "添加成功",
			"data": list,
		})
	})
	//删除
	r.DELETE("/user/delete/:id", func(c *gin.Context) {
		var list []List
		id := c.Param("id")

		db.Where("id = ?", id).Find(&list)
		if len(list) == 0 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "用户不存在,删除失败",
			})
			return
		} else {
			//操作数据库删除
			db.Where("id = ?", id).Delete(&list)
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "删除成功",
			})
		}

	})
	//修改
	r.PUT("/user/update/:id", func(c *gin.Context) {
		// 1.找到对应的id所对应的条目
		// 2.判断id是否存在
		// 3.修改对应条目
		// 4.返回id，没有找到
		var list List
		id := c.Param("id")
		db.Where("id = ?", id).Find(&list)
		if list.ID == 0 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "用户id没有找到",
			})
		} else {
			err := c.ShouldBindJSON(&list)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  "修改失败",
				})
			} else {
				db.Where("id = ?", id).Updates(&list)
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "修改成功",
					"data": list,
				})

			}
		}
	})
	//查询所有
	r.GET("/user/list", func(c *gin.Context) {
		var lists []List

		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		pageNum, _ := strconv.Atoi(c.Query("pageNum"))

		//判断是否需要分页
		if pageSize == 0 {
			pageSize = -1
		}
		if pageNum == 0 {
			pageNum = -1
		}

		offsetVal := (pageNum - 1) * pageSize
		if pageNum == -1 && pageSize == -1 {
			offsetVal = -1
		}

		//返回一个总数
		var total int64

		//查询数据库
		db.Model(lists).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&lists)
		if len(lists) == 0 {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "没有查询到数据",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "查询成功",
				"data": gin.H{
					"total":    total,
					"list":     lists,
					"pageSize": pageSize,
					"pageNum":  pageNum,
				},
			})
		}

	})
	//根据name查询
	r.GET("/user/list/:name", func(c *gin.Context) {
		name := c.Param("name")

		var lists []List
		db.Where("name = ?", name).Find(&lists)

		if len(lists) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "用户不存在",
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "查询成功",
				"data": lists,
			})
		}

	})

	r.Run(":3300")
}
