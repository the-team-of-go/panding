# panding



接入服务参考client客户端，将传入的信息写在req中，上传req即可

![image-20220211132910689](C:\Users\BUBU\AppData\Roaming\Typora\typora-user-images\image-20220211132910689.png)



管理服务更新系统配置参考clientSql客户端，将要更新的信息写在req中，上传req即可

![image-20220211133012193](C:\Users\BUBU\AppData\Roaming\Typora\typora-user-images\image-20220211133012193.png)

管理服务更新系统联系人参考clientAdmin客户端，每次请求都需要将数据库所有联系人发送过来，将管理员姓名与邮箱依次加入name和email数组，然后调用即可

![image-20220211133228490](C:\Users\BUBU\AppData\Roaming\Typora\typora-user-images\image-20220211133228490.png)