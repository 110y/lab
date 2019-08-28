local redis = require 'redis'
local client = redis.connect('127.0.0.1', 6379)

local key = 'foo'
client:set(key, 123)
local val = client:get(key)
print(val)
