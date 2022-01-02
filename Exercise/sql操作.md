打开MySQL：net start mysql  
关闭MySQL：net stop mysql  
登陆MySQL：mysql -u root -p然后根据提示输入密码  
显示现有数据库：show databases;  
创建数据库： create database  ******;  
使用某一个数据库：Use ******;  
显示该库的表****的结构:desc ****;  
SELECT *FROM product;  

SELECT id,productName,salePrice FROM product;  

-- DISTINCT去重  

SELECT DISTINCT dir_id FROM product;  

查询所有货品的id,名称和批发价(批发价=卖价*折扣)  
SELECT id, productName ,salePrice*cutoff FROM product;  

SELECT id productName, salePrice*cutoff FROM product;  

查询所有货品的id,名称,和各进50个的成本价(成本=costPrice)  
SELECT id, productName, costPrice*50 FROM product;  

查询所有货品的id,名称,各进50个,并且每个运费1元的成本  
SELECT id, productName, 50* (costPrice+1) FROM product;  

-- 别名  
-- 使用别名用法：列as别名，用在那些地方：列，表，查询结果，视图，使用别名时as可以省略  
SELECT id, productName, 50* (costPrice+1) as 价格 FROM product;  
SELECT p.id, p.productName, 50* (p.costPrice+1) as 价格 FROM product as p;  
SELECT p.id, p.productName, 50* (p.costPrice+1) 价格 FROM product p;  


-- CONCAT拼接结果  
把多列的查询结果拼接为一列  
需求:查询商品的名字和零售价。格式: xxx商品的零售价为: xxx  

SELECT CONCAT(productName,'价格为',salePrice,'元') FROM product;  

-- 过滤查询  
查询为罗技G9XD的商品名  
SELECT *FROM product WHERE productName="罗技G9X";  

查询不为罗技G9XD的商品名  
SELECT *FROM product WHERE productName!="罗技G9X";  

查询零售价小于等于200的货品  
SELECT *FROM product WHERE salePrice<=200;  

查询id，货品名称，批发价，大于350的货品  

SELECT id,productName,salePrice*cutoff as pfj FROM product WHERE (salePrice*cutoff)>350;  

选择ID，货品名称，批发价在300-400之间货品  
SELECT id,productName,salePrice*cutoff as pfj FROM product  
WHERE salePrice*cutoff >=300 AND salePrice*cutoff<=400  
 
选择ID，货品名称，分类编号为2、4的所有货品  
SELECT *FROM product WHERE dir_id=2 OR dir_id=4  

选择分类编号不为2的所有商品  
SELECT *FROM product WHERE dir_id!=2  
SELECT *FROM product WHERE dir_id<>2  

选择id,货品名称,分类编号的货品零售价大于等于250或者是成本大于等于200  
SELECT *FROM product WHERE salePrice>=250 OR costPrice>=200  

优先级规则  
A、比较运算符 2、NOT 3、AND 4、OR  
B、小括号跨越所有运算符优先级  

------范围查询-BETWEEN AND------结果为闭区间[]  

选择id,货品名称，批发价在300-400之间的货品  
SELECT id,productName,salePrice*cutoff as pfj FROM product  
WHERE salePrice*cutoff BETWEEN 300 AND 400  

选择id,货品名称，批发价不在300-400之间的货品  
SELECT id,productName,salePrice*cutoff as pfj FROM product  
WHERE salePrice*cutoff NOT BETWEEN 300 AND 400  

--------------------集合查询--------------------
语法：  
SELECT <selectList>  
FROM table_name   
WHERE 列名 IN (值1，值2....)  

需求：  
选择id，货品名称，分类编号为2，4的所有货品  
SELECT *FROM product WHERE dir_id IN(2,4,5);  

选择id,货品名称,分类编号不为2,4的所有货品  
SELECT *FROM product WHERE dir_id NOT IN(2,4,5);  

需求: 查询商品名为NULL的所有商品信息。  
注意: 列的值是 null 和空字符串不一样,如果是空字符串应该= ' '  
SELECT *FROM product WHERE productName IS NULL;  

需求: 查询商品名不为NULL的所有商品信息。  
SELECT *FROM product WHERE productName IS NOT NULL;  

------ 模糊查询-----  
LIKE :像....  
通配符  

%匹配多个字符  
_匹配一个字符
%罗技_：罗技前面多个字符后面一个字符  
_罗技%：罗技前面一个字符后面多个字符  
_罗技_：罗技前后有一个字符  
%罗技%：前后有任意多个字符   
_罗技：前面有两个字符  
 
查询id,货品名称,货品名称包含罗技  
SELECT *FROM product WHERE productname LIKE '%罗技%'  

查询id,货品名称,分类编号,零售价大于等于200并且货品名称匹配'%罗技%'  
SELECT *FROM product WHERE salePrice>=200 AND productname LIKE '%罗技%'  


-- 结果排序--  

SELECT <selectList>  
FROM table name  
WHERE 筛选条件 ORDER BY 要排序的列1,要排序的列2排序的方式(升序 ASC,降序DESC)  
选择id,货品名称,分类编号,零售价并且按零售价降序排序  
SELECT *FROM product ORDER BY salePrice DESC;  
SELECT *FROM product ORDER BY salePrice ASC;  
 
选择id,货品名称,分类编号,零售价先按分类编号排序,再按零售价排序  
SELECT *FROM product ORDER BY dir_id DESC	   
SELECT *FROM product ORDER BY dir_id DESC,salePrice DESC;   
查询M系列并按照批发价排序(加上别名)  

查询分类为2并按照批发价排序(加上别名)  

分页：  
每页展示5条数据 5步长  
第一页: 1---5  
第二页: 6---10  
第三页: 11---15  

第n页: (当前页-1)*5+1  

如果开始的索引为0:  
第一页:0---4  
第二页: 5---9  
第三页: 10---14  

第N页 ：(当前页-1)*5  

分页: limit(开始的索引位置, 步长);  
SELECT * FROM 表 WHERE 筛选条件 ORDER BY 排序 LIMIT(开始的索引位置,每页长度)  

开始的索引位置=(当前页-1)*每页长度  
查讯产品表：每页  

第一页：SELECT *FROM product LIMIT 0,5;  
第二页：SELECT *FROM product LIMIT 5,5;  
第三页：SELECT *FROM product LIMIT 10,5;  
第四页：SELECT *FROM product LIMIT 15,5;  
第五页：SELECT *FROM product LIMIT 20,5;  
