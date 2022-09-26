#!/usr/bin/lua

-- 使用redis-cell令牌桶
function redisCell()
    return redis.call('CL.THROTTLE', KEYS[1], KEYS[2], KEYS[3], KEYS[4], KEYS[5]);
end
