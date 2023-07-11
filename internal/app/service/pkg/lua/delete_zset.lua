-- 输入参数含义：KEYS[1]--ZSet的Key ARGV[1]--ZSet中要删除的member值

-- 判断key是否存在
local isExist = 0
isExist = redis.call("EXISTS", KEYS[1])
if isExist == 1 then
    -- 在ZSet中删除对应的member值
    redis.call("ZREM", KEYS[1], ARGV[1])
end
