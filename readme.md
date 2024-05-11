## 说明
这个方法只绕过验证， 使用pro的本地功能， 不能破解在线功能，如同步资产等

> 版本低于 8.12.0
## 使用方法
### 1. 进入到Termius安装目录， 我的是 ```/opt/Termius/``` ， 安装npm,asar, 解压app.asar文件
安装 npm，asar
```bash
sudo apt -y install npm

sudo npm config set registry https://registry.npm.taobao.org --global
sudo npm install -g asar
```

解压  ```app.asar```
```bash
sudo asar extract app.asar
sudo mv app.asar app.asar.bak 
```
### 2. 修改 ```app/js/background-process.js```文件
通过脚本来修改
修改脚本中对应的安装路径，我的是这个， 修改为自己的
```yaml
/opt/Termius/resources/app/js/background-process.js
```

### 3. 运行脚本替换文件
python
```python
sudo python3 crack.py
```

go
```go
sudo go run crack.go
```
