-- 输入参数含义：KEYS[1]--ZSet的Key ARGV[1]--max ARGV[2]--min ARGV[3]--offset ARGV[4]--count ARGV[5]--isLimit
-- 返回参数含义：array=[0]--key不存在；array=其他--key存在且array中的值为从redis中读出的值

-- 判断key是否存在
local isExist = 0
local array = {0}
isExist = redis.call("EXISTS", KEYS[1])
if isExist == 0 then
    return array
end
if isExist == 1 then
    -- 在ZSet中删除对应的member值
    if ARGV[5] == 1 then
        array = redis.call("ZREVRANGEBYSCORE", KEYS[1], ARGV[1], ARGV[2], "LIMIT", ARGV[3], ARGV[4])
    else
        array = redis.call("ZREVRANGEBYSCORE", KEYS[1], ARGV[1], ARGV[2])
    end 
    return array
end
