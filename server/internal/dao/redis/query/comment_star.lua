#!/usr/bin/lua

-- 获取并清空一半hset内容
function listHalfCommentStarsAndSetZeroLua()
    local result = {};
    local key = {}
    local value = {}
    for i, k in pairs(redis.call('keys', KEYS[1] .. '*')) do
        local nums = math.floor(redis.call('hlen', k) / 2 + 1);
        table.insert(key, k);
        local kvs = redis.call('HRANDFIELD', k, nums, 'withvalues');
        table.insert(value, kvs);
        for v = 1, table.getn(kvs), 2 do
            redis.call('hdel', k, kvs[v]);
        end
    end
    table.insert(result, key);
    table.insert(result, value);
    return result;
end

--redis.call('expire', k, 0);
-- 清空所有keys
function deleteAllCommentStarLua()
    local redisKeys = redis.call('keys', KEYS[1] .. '*');
    for i, k in pairs(redisKeys) do
        print(k)
        redis.call('expire', k, 0);
    end
end
