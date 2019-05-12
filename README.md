# sbrun

用sb的方式运行sb(spring boot)程序

## 简介

每次运行 `nohup mvn spring-boot:run &` 都需要 `exit`，每次重新运行都需要找进程杀了才行，真的好烦噢，于是就写了这个命令。可是就算写完了这个程序，我还是连 `netstat` 的用法都记不住哎，这下有了这个命令，就更记不住了23333

至于这个命令具体做了什么。emm...打比方有个spring boot项目要在11111这个端口运行，那么它做这样的事情

```
netstat -lnp|grep 11111
# 假设上一步找到的pid是2415，就执行这一步，否则跳过
kill -s 9 2415
nohup mvn spring-boot:run &
```

这个命令是通过正则优先去读 `src/main/resources/application.properties` 下的端口配置，如果找不到，就尝试去读 `src/main/resources/application.yaml`，也就是说需要在项目根目录下执行

## 食用方法

```
cd /usr/bin

wget https://github.com/ChenViVi/sbrun/raw/master/sbrun

chmod +x sbrun

cd /your_spring_boot_application_path

sbrun
```