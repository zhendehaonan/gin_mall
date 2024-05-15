package model

// 商品模型
import (
	"gin_mall/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint   `gorm:"not null"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

// 商品点击数
func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}
func (product *Product) AddView() {
	//增加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))                   //商品点击数+1
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID))) //排行榜点击数+1
}
