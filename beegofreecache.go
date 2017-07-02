package beegofreecache

import (
    "github.com/coocood/freecache"
    "encoding/gob"
    "bytes"
    "time"
    "fmt"
    "github.com/astaxie/beego/cache"
    "encoding/json"
    "errors"
)

//缺点:
//     1.当需要缓存的数据占用超过提前分配缓存的 1/1024 则不能缓存
//     2.当分配内存过多则会导致内存占用高 最好不要超过100MB的内存分配
//     3.key的长度需要小于65535字节

var (
    DefaultSize = 64 //64M
    free *freecache.Cache = nil
)
type Cache struct {
    Free *freecache.Cache
}

// 如果是需要集成到beego，则init函数必须打开，反之可以注释掉
func init() {
    gob.Register(map[string]interface{}{})
    gob.Register(map[string][]int{})
    gob.Register(map[string][]int64{})
    cache.Register("beegofreecache", NewFreeRegister)
}

// 该函数是为了集成到beego框架中
func NewFreeRegister() cache.Cache {
    return &Cache{}
}

// 该函数是可以在任意地方使用，初始化Free进程缓存
func NewFree(m int) *Cache {
    if free == nil {
        cacheSize := m * 1024 * 1024
        free = freecache.NewCache(cacheSize)
    }
    //beeFree := Cache{}
    beeFree :=new(Cache)
    beeFree.Free = free
    return beeFree
}

//根据key获取对应的value,如果不是在beego框架中,建议使用GetValue
func (free *Cache) Get(key string) interface{} {
    cache, err := free.Free.Get([]byte(key))
    if err != nil || len(cache) <= 0 {
        return nil
    }

    return cache
}

// 推荐使用,但是beego框架不支持
func (free *Cache) GetValue(key string, value interface{}) error {
    cache, err := free.Free.Get([]byte(key))
    if len(cache) > 0 && err == nil {
        err = GobDecode(cache, value)
        return nil
    } else {
        return err
    }
}

//批量获取keys
func (free *Cache) GetMulti(keys []string) []interface{} {
    retList := make([]interface{}, 0)
    for i := 0; i < len(keys); i++ {
        retList = append(retList, free.Get(keys[i]))
    }
    return retList
}

//设置缓存
func (free *Cache) Put(key string, val interface{}, timeout time.Duration) error {
    var cache []byte = make([]byte, 0)
    var err error
    // 如果value是[]byte就不需要转化了
    switch val.(type) {
    case []byte:
        cache = val.([]byte)
    default:
        cache, err = GobEncode(val)
    }
    if err == nil {
        err = free.Free.Set([]byte(key), cache, int(timeout.Seconds()))
    }
    return err
}

// 删除key
func (free *Cache) Delete(key string) error {
    b := free.Free.Del([]byte(key))
    if b {
        return nil
    } else {
        return errors.New("del" + key + " error!!!")
    }
}

//对key值为int64的加1
func (free *Cache) Incr(key string) error {
    /*buf:=free.Get(key)
    if buf ==nil {
        return errors.New("key is empty")
    }

    var value int64
    err:=GobDecode(buf.([]byte),&value)*/
    var value int64
    err := free.GetValue(key, &value)
    if err != nil {
        return errors.New("value is not an integer or out of range")
    }
    t, err := free.Free.TTL([]byte(key))
    free.Put(key, value+1, time.Duration(t)*time.Second)
    return err
}

//对key值为int64的减1
func (free *Cache) Decr(key string) error {
    var value int64
    err := free.GetValue(key, &value)
    if err != nil {
        return errors.New("value is not an integer or out of range")
    }
    t, err := free.Free.TTL([]byte(key))
    free.Put(key, value-1, time.Duration(t)*time.Second)
    return err
}

//判断指定key是否存在
func (free *Cache) IsExist(key string) bool {
    buf := free.Get(key)
    if buf == nil {
        return false
    }
    return true
}

//清出所有缓存
func (free *Cache) ClearAll() error {
    free.Free.Clear()
    return nil
}

//在beego框架中，注册时会自动执行该函数初始化
func (cache *Cache) StartAndGC(config string) error {
    var cf map[string]int
    json.Unmarshal([]byte(config), &cf)
    if _, ok := cf["size"]; !ok {
        cf = make(map[string]int)
        cf["size"] = DefaultSize
    }
    //NewFree(int(cf["size"]))
    if free == nil {
        cacheSize := int(cf["size"]) * 1024 * 1024
        free = freecache.NewCache(cacheSize)
    }

    cache.Free = free
    cache.Free.ResetStatistics()
   // free = nil
    return nil
}

//输出cache状态
func (free *Cache) String() string {
    info := fmt.Sprintf("EntryCount is %d,ExpiredCount is %d,HitCount is %d,HitRate is %f,EvacuateCount Is %d,AverageAccessTime is %d,LookupCount is %d .", free.Free.EntryCount(), free.Free.ExpiredCount(),
        free.Free.HitCount(), free.Free.HitRate(), free.Free.EvacuateCount(), free.Free.AverageAccessTime(), free.Free.LookupCount())
    return info
}

//获取cache状态
func (free *Cache) CacheStatus() map[string]interface{} {

    infoMap := map[string]interface{}{
        "HitCount":          free.Free.HitCount(),
        "HitRate":           free.Free.HitRate(),
        "EvacuateCount":     free.Free.EvacuateCount(),
        "AverageAccessTime": free.Free.AverageAccessTime(),
        "LookupCount":       free.Free.LookupCount(),
        "EntryCount":        free.Free.EntryCount(),
        "ExpiredCount":      free.Free.ExpiredCount(),

    }
    return infoMap
}

// Gob序列化
func GobEncode(data interface{}) ([]byte, error) {

    buf := bytes.NewBuffer(nil)
    enc := gob.NewEncoder(buf)
    err := enc.Encode(data)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// Gob反序列化
func GobDecode(data []byte, to interface{}) error {

    buf := bytes.NewBuffer(data)
    dec := gob.NewDecoder(buf)
    return dec.Decode(to)
}
