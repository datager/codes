-- sum = 0
-- num = 1
-- while num <= 100 do
--     sum = sum + num
--     num = num + 1
-- end
-- print("sum = ", sum)

-- if age == 40 and sex =="Male" then
--     print("男人四十一枝花")
-- elseif age > 60 and sex ~="Female" then
--     print("old man without country!")
-- elseif age < 20 then
--     io.write("too young, too naive!\n")
-- else
--     local age = io.read()
--     print("Your age is "..age)
-- end

-- sum = 2
-- repeat
--    sum = sum ^ 2 --幂操作
--    print(sum)
-- until sum >1000

-- function fib(n)
--   if n < 2 then return 1 end
--   return fib(n - 2) + fib(n - 1)
-- end

-- function newCounter()
--     local i = 0
--     return function()     -- anonymous function
--        i = i + 1
--         return i
--     end
-- end

-- c1 = newCounter()
-- print(c1())  --> 1
-- print(c1())  --> 2

-- function myPower(x)
--     return function(y) return y^x end
-- end
-- power2 = myPower(2)
-- power3 = myPower(3)
-- print(power2(4)) --4的2次方
-- print(power3(5)) --5的3次方

-- function getUserInfo(id)
--     print(id)
--     return "haoel", 37, "haoel@hotmail.com", "https://coolshell.cn"
-- end
-- name, age, email, website, bGay = getUserInfo()
-- print(name)
-- print(age)
-- print(email)
-- print(website)
-- print(bGay)

-- haoel = {name="ChenHao", age=37, handsome=True}

-- haoel.website="https://coolshell.cn/"
-- local age = haoel.age
-- haoel.handsome = false
-- haoel.name=nil

-- t = {[20]=100, ['name']="ChenHao", [3.14]="PI"} 

-- arr = {10,20,30,40,50}

arr = {"string", 100, "haoel", function() print("coolshell.cn") end}

-- arr[4]()

for i=1, #arr do
    print(arr[i])
end

-- for k, v in pairs(arr) do
--     print(k, v)
-- end



