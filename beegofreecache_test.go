package beegofreecache

import "testing"
import (
    "github.com/astaxie/beego/cache"
    "time"
    . "beegofreecache"
)


func TestBeeFreeCache(t *testing.T)  {
    var bm, err = cache.NewCache("beegofreecache", `{"size":64}`)
    if err!=nil {
        println(err.Error())
    }

    bm.Put("key",1200,10*time.Second)
    val:=bm.Get("key")
    var ii int
    GobDecode(val.([]byte),&ii)
    println(val,ii)

}
