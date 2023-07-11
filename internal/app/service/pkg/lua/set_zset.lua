-- 输入参数含义：KEYS[1]--ZSet的Key ARGV[1]--ZSet的阈值长度 ARGV[2]--ZSet中新增元素的score ARGV[3]--ZSet中新增元素的member
-- ARGV[4]--为ZSet重新设置过期时间

-- 判断key是否存在
local isExist = 0
isExist = redis.call("EXISTS", KEYS[1])
if isExist == 1 then
    -- 读取ZSet的元素个数，记录到length变量中
    local zSetLength = 0
    zSetLength = redis.call("ZCARD", KEYS[1])
    -- 如果长度等于阈值就删除分数最低的那个元素
    if zSetLength == ARGV[1] then
        redis.call("ZREMRANGEBYRANK", KEYS[1], 0, 0)
    end
    -- 向ZSet中添加元素
    redis.call("ZADD", KEYS[1], ARGV[2], ARGV[3])
    redis.call("EXPIRE", KEYS[1], ARGV[4])
end
