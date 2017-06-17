package beegofreecache

import "testing"
import (
    "github.com/astaxie/beego/cache"
    "time"
    . "beegofreecache"
    "strconv"
)

var bm, err = cache.NewCache("beegofreecache", `{"size":64}`)
func TestBeeFreeCacheGP(t *testing.T)  {

    if err!=nil {
        println(err.Error())
    }
    key:="key"
    bm.Put(key,1200,10*time.Second)
    val:=bm.Get(key)
    var ii int
    GobDecode(val.([]byte),&ii)
    println(val,ii)

}

func TestBeeFreeCacheDel(t *testing.T)  {
    //var bm, err = cache.NewCache("beegofreecache", `{"size":64}`)
    if err!=nil {
        println(err.Error())
    }
    key:="key"
    bm.Put(key,1200,10*time.Second)
    val:=bm.Get(key)
    var ii int
    GobDecode(val.([]byte),&ii)
    println(val,ii)
    bm.Delete(key)
    val=bm.Get(key)
    println(val)
}

func TestBeeFreeCacheGetMu(t *testing.T)  {
    //var bm, err = cache.NewCache("beegofreecache", `{"size":64}`)
    if err!=nil {
        println(err.Error())
    }
    key:="key"
    for i:=0;i<10 ;i++  {
        bm.Put(key+strconv.Itoa(i),i*23,10*time.Second)
    }

    vals:=bm.GetMulti([]string{"key0","key1","key2","key3"})
    for _,val:=range vals{
        var ii int
        GobDecode(val.([]byte),&ii)
        println(val,ii)
    }
}

func TestBeeFreeCacheCr(t *testing.T)  {
   // var bm, err = cache.NewCache("beegofreecache", `{"size":64}`)
    if err!=nil {
        println(err.Error())
    }
    key:="key"
    bm.Put(key,int64(0),10*time.Second)

    val:=bm.Get(key)
    var ii int64
    GobDecode(val.([]byte),&ii)
    println("get",val,ii)

    err=bm.Incr(key)
    val=bm.Get(key)
    GobDecode(val.([]byte),&ii)
    println("Incr",err,val,ii)
    bm.Decr(key)
    val=bm.Get(key)
    GobDecode(val.([]byte),&ii)
    println(val,ii)
}

func TestBeeFreeCacheEx(t *testing.T)  {
    key:="exist"
    bm.Put(key,"TestBeeFreeCacheEx",10*time.Second)
    b:=bm.IsExist(key)
    if b{
        val:=bm.Get(key)
        var ii string
        GobDecode(val.([]byte),&ii)
        println("get",val,ii)
    }
    b=bm.IsExist("noexist")
    println("noexist",b)
}

func TestBeeFreeCacheClear(t *testing.T)  {

    key:="key"
    for i:=0;i<10 ;i++  {
        bm.Put(key+strconv.Itoa(i),i*23,10*time.Second)
    }
    for i:=0;i<10 ;i++  {
        println(key+strconv.Itoa(i),bm.IsExist(key+strconv.Itoa(i)))
    }
    bm.ClearAll()

    for i:=0;i<10 ;i++  {
        println(key+strconv.Itoa(i),bm.IsExist(key+strconv.Itoa(i)))
    }
}


func TestBeeFreeCacheGetValue(t *testing.T)  {
    free:= NewFree(16)

    key:="key"
    free.Put(key,1200,10*time.Second)
    var ii int
    val:=free.GetValue(key,&ii)
    println("GetValue",val,ii)
}