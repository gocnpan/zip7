<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-zip7.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/zip7.v1"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev"></a>
  <a href="https://kaos.sh/w/zip7/ci"><img src="https://kaos.sh/w/zip7/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/zip7"><img src="https://kaos.sh/r/zip7.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/zip7"><img src="https://codebeat.co/badges/11fb655d-8da8-4694-a32b-b95ff9eed602" alt="Codebeat badge" /></a>
  <a href="https://kaos.sh/w/zip7/codeql"><img src="https://kaos.sh/w/zip7/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#compatibility-and-os-support">Compatibility and OS support</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`zip7` package provides methods for working with 7z archives (`p7zip` wrapper).

# 新增功能
重命名: rn (Rename) command
- 参考: https://documentation.help/7-Zip-18.0/rename.htm
- `rn <archive_name> <src_file_1> <dest_file_1> [ <src_file_2> <dest_file_2> ... ]`
- `7z rn a.7z old.txt new.txt 2.txt folder\2new.txt`

可以在追加文件后, 通过重命名的方式将文件追加到指定位置

# 追加对安卓的支持
需要将`p7zip`编译后将当前Android系统架构的7z可执行文件释放到当前程序同级目录下

名称为`7za`

## p7zip编译

这里主要是Android的编译

首先是下载`p7zip`源码: https://jaist.dl.sourceforge.net/project/p7zip/p7zip/16.02/p7zip_16.02_src_all.tar.bz2

## 配置`android ndk`
然后是配置`android ndk`, 这里需要用到`16.1.4479499`的版本，版本越高越可能出错
- `ndk历史版本`:https://github.com/android/ndk/wiki/Unsupported-Downloads
- `ndk下载地址`:https://dl.google.com/android/repository/android-ndk-r16b-linux-x86_64.zip

要在Ubuntu上安装Android NDK 16.1.4479499，按照以下步骤操作：

### 1. 下载NDK
首先，下载NDK 16.1.4479499版本。

```bash
wget https://dl.google.com/android/repository/android-ndk-r16b-linux-x86_64.zip
```

### 2. 解压NDK
下载完成后，解压文件：

```bash
unzip android-ndk-r16b-linux-x86_64.zip -d ~/
```

### 3. 设置环境变量
将NDK路径添加到环境变量中，方便使用。

打开`~/.bashrc`文件：

```bash
nano ~/.bashrc
```

在文件末尾添加：

```bash
export ANDROID_NDK_HOME=~/android-ndk-r16b
export PATH=$PATH:$ANDROID_NDK_HOME
```

保存并退出，然后运行以下命令使更改生效：

```bash
source ~/.bashrc
```

### 4. 验证安装
检查NDK是否安装成功：

```bash
ndk-build --version
```

如果显示版本信息，说明安装成功。

### 5. 使用NDK
现在你可以使用NDK编译Android项目的本地代码了。

### 注意事项
- 如果使用其他shell（如zsh），请修改对应的配置文件（如`~/.zshrc`）。
- 确保系统已安装`unzip`工具，若未安装，可通过以下命令安装：

```bash
sudo apt-get install unzip
```

完成以上步骤后，你已成功在Ubuntu上安装了Android NDK 16.1.4479499。

Android ndk编译时报错：[armeabi-v7a] Compile++ thumb: 7zr <= 7zCompressionMode.cpp
/home/ap/android-ndk/android-ndk-r16b/toolchains/llvm/prebuilt/linux-x86_64/bin/clang++: error while loading shared libraries: libncurses.so.5: cannot open shared object file: No such file or directory
make: *** [/home/ap/data/workspace/eos-store/app/plugin/p7zip_16.02/CPP/ANDROID/7zr/obj/local/armeabi-v7a/objs/7zr/__/__/__/__/CPP/7zip/Archive/7z/7zCompressionMode.o] Error 127

解决方法：
```shell
sudo apt-get update
sudo apt-get install libncurses5
```

## 开始编译
```shell
cd plugin/p7zip_16.02/CPP/ANDROID/7z/jni

ndk-build
```

不同版本的`7z`：
- 7z：功能最全面，支持多种格式和插件。
- 7za：轻量级版本，支持有限格式（如 7z 和 zip），独立运行。
- 7zr：最精简的版本，仅支持 7z 格式，适合资源受限的场景。

# Installation

Make sure you have a working Go 1.17+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get github.com/gocnpan/zip7
```

If you want to update `zip7` to latest stable release, do:

```
go get -u github.com/gocnpan/zip7
```

# Compatibility and OS support

|      Version |     1.x |
|--------------|---------|
|  `p7zip 9.x` | Partial |
| `p7zip 15.x` |    Full |
| `p7zip 16.x` |    Full |

| OS       | Support            |
|----------|--------------------|
| Linux    | :heavy_check_mark: |
| Mac OS X | :heavy_check_mark: |
| FreeBSD  | :heavy_check_mark: |
| Windows  | :heavy_check_mark: |

Running this project on Windows pre-requisite 7zip and 7zip standalone console version from [7zip](https://www.7-zip.org/download.html) 

Add Windows `%PATH%` to `C:\Program Files\7-Zip` or where the program is.

# Build Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://kaos.sh/w/zip7/ci.svg?branch=master)](https://kaos.sh/w/zip7/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/zip7/ci.svg?branch=develop)](https://kaos.sh/w/zip7/ci?query=branch:develop) |

# Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

# License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
