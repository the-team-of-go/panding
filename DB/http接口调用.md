## 模块：示例模块

### 接口：返回机器列表
* 地址：http://10.243.50.4:9090/machine/list
* 简介：无
* 返回接口格式：

```
├─ code: Number (必选) (状态码)
├─ msg: String (必选) (描述信息)
└─ data: Array (机器id列表，如：[1, 12, 3, 8, 19, ...]): Array (必选) (机器id列表，如：[1, 12, 3, 8, 19, ...])
```

样本

{"code":200,"msg":"成功","data":[4,1,3,2]}

![image-20220215153548746](C:\Users\BUBU\AppData\Roaming\Typora\typora-user-images\image-20220215153548746.png)

### 接口：获取机器基本状态信息

* 地址：http://10.243.50.4:9090/machine/info/list
* 简介：- 时间区间不传时默认返回前200条数据（根据type参数决定是否聚合）时间戳只传开始不传结束时，返回开始到最新区间的数据
- 数据需要进行聚合，将5分钟的数据聚合成一条，也就是共 24 * 12条数据
- type不传时默认返回前500条数据，不聚合(现在默认返回前3条数据，数据库样本不足)

* 请求接口格式：

```
├─ id: Number (必选) (机器id)
├─ start: Number  (开始时间戳)
├─ end: Number  (结束时间戳)
└─ type: String  (信息类型：basic（默认），group)

```

* 返回接口格式：

```
原生数据：
├─ code: Number (必选) 
├─ msg: String (必选) 
└─ data: Array 
   ├─ id: Number (必选) 
   ├─ cpu: Number (必选) (浮点型数据)
   ├─ mem: Number (必选) 
   └─ disk: Number (必选)
聚合数据：
├─ code: Number (必选) 
├─ msg: String (必选) 
└─ data: Array 
   ├─ id: Number (必选) 
   ├─ cpu: {max,min,averg}
   ├─ mem: {max,min,averg}
   └─ disk:{max,min,averg}
```

样本

type=basic

{"code":200,"msg":"成功","data":[{"nodeId":1,"cpu":21.59,"mem":49.12,"disk":21.43}{"nodeId":1,"cpu":59.32,"mem":92.32,"disk" :79.58},{"nodeId":1,"cpu":94.12,"mem":23.41,"disk":21.53}]}

![image-20220215160533382](C:\Users\BUBU\AppData\Roaming\Typora\typora-user-images\image-20220215160533382.png)

type=group

{"code":200,"msg":"成功","data":[{"nodeid":1,"cpu":{"Max":59.32,"Min":21.59,"Averg":40.455},"mem":  {"Max":92.32,"Min":49.12  ,"Averg":70.72},"disk":{"Max":79.58,"Min":21.43,"Averg":50.504999999999995}},{"nodeid":1,"cpu":{"Max":94.12,"Min":94.12, "Averg":94.12},"mem":{"Max":23.41,"Min":23.41,"Averg":23.41},  "disk":{"Max":21. 53,"Min":21.53,"Averg":21.53}},{"nodeid":1, "cpu":{"Max":98.12,"Min":20.12,"Averg":46.120000000000005},"mem":{"Max":71.32,"Min":69.32,"Averg":69.98666666666666},
"disk":{"Max":64.42,"Min":12.68,"Averg":29.926666666666666}}]}

![image-20220215160642514](C:\Users\BUBU\AppData\Roaming\Typora\typora-user-images\image-20220215160642514.png)

### 接口：返回当前机器信息

* 地址：http://10.243.50.4:9090/machine/info/now
* 简介：返回最新的机器状态
* 请求接口格式：

```
└─ id: Number (必选) (机器id)
例：http://10.243.50.4:9090/machine/info/now?id=3
```

* 返回接口格式：

```
├─ code: Number (必选) 
├─ msg: String (必选) 
└─ data: Object 
   ├─ id: Number (必选) 
   ├─ cpu: Number (必选) 
   ├─ mem: Number (必选) 
   ├─ disk: Number (必选) 
   └─ status: Number (必选) (当前机器状态，枚举类型，0：normal，1：warning，2：danger)
```

样本

{"code":200,"msg":"成功","data":{"Id":3,"CpuUsed":58.12,"MemUsed":43.23,"DiskUsed":40.12,"Status":"normal"}}

![image-20220215153440850](C:\Users\BUBU\AppData\Roaming\Typora\typora-user-images\image-20220215153440850.png)