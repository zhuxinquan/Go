##源码文件的分类和含义

1. 命令源码文件
2. 库源码文件
3. 测试源码文件

功能测试函数

func TestFind(t *testing.T){

}

基准(性能)测试函数

func BenchmarkFind(b *testing.B){

}

##代码包相关

 1. 声明与导入路径的区别
    1. 代码包**声明语句**中的包名称只包含代码包导入路径的最右子路径
 2. 代码包导入
    1. 代码包导入语句中使用的包名称应该与其导入路径一致
    2. 带别名的导入: import str "strings"     str.HasPrefix("abc", "a")
    3. 本地化导入: import . "strings"        HasPrefix("abc", "a")，调用代码包中的程序实体如同调用当前代码包的程序实体一样，可以不写任何前缀
    4. 仅仅初始化： import _ "strings"  导入了代码包，但并不调用其中的任何程序实体，仅仅执行初始化代码包的动作
 3. 代码包初始化函数：无参数声明和结果声明的init函数，init函数可以被声明在任何文件中，且可以有多个
 4. init的执行时机
    1. 在单一代码包内，当导入某个代码包时，首先对代码包中的全局变量进行求值，求值完成之后，执行所有的init函数，在同一个代码包中init的执行顺序是不定的
    2. 不同代码包之间(当前代码包和导入的代码包都包含init函数)，被导入的代码包中的init函数先被执行，然后执行导入它的那个代码包的init函数<font color="red">不应在同一个代码包中被导入的多个代码包的init函数的执行顺序作出假设</font>
    3. 所有涉及到的代码包，init函数先执行，main函数才会被执行，并且每个init函数只会被执行一次
    
##命令基础

 1. go run
    1. 用于运行命令源码文件
    2. 只能接受一个命令源码文件以及若干个库源码文件作为文件参数
    3. 内部操作：先编译源码文件再运行，编译结果包括命令源码文件，被编译后生成的可执行文件以及相关库源码文件编译后所生成的归档文件
    4. go run常用标记的使用：
        1. -a:强制编译相关代码，不论他们的编译结果是否已是最新的
        2. -n:打印编译过程中所需运行的命令，但*不是真正执行它们*，最终没有程序的执行结果
        3. -p n:并行编译，其中n为并行的数量（建议为CPU逻辑个数（2CPU，2核=》逻辑CPU
        为4个）-p 4）
        4. -v:列出被编译的代码包的名称
        5. -a -v:列出所有被编译的代码包的名称
        6. -work:显示编译时创建的临时工作目录的路径，并且在执行完成后不删除（默认删除）
        7. -x:打印编译过程中所需运行的命令，_并执行它们_
 2. go build
    1. 用于编译源码文件或代码包
    2. 编译非命令源码文件不会产生任何结果文件（检查库源码文件的有效性）
    3. 编译命令源码文件时会在该命令的执行目录中生成一个可执行文件
    4. 执行该命令时不追加任何参数时，它会试图把当前目录作为代码包并编译
    5. 执行该命令且以代码包的导入路径作为参数时，该代码包及其依赖会被编译
    6. 执行该命令且以若干源码文件作为参数时，只有这些文件会被编译
 3. go install
    1. 用于编译并安装代码包或源码文件
    2. 安装代码包会在当前工作区的pkg/<平台相关目录>下生成归档文件
    3. 安装命令源码文件会在当前工作区的bin目录或$GOBIN目录下生成可执行文件
    4. 执行该命令时不追加任何参数时，它会试图把当前目录作为代码包并安装
    5. 执行该命令且以代码包的导入路径作为参数时，该代码包及其依赖会被安装
    6. 执行该命令且以若干源码文件作为参数时，只有这些文件会被编译并安装
 4. go get
    1. -d:只执行下载动作，而不执行安装动作
    2. -fix:在下载代码包后先执行修正动作，而后再进行编译和安装（比如修正版本之间的差异）
    3. -u:利用网络来更新已有的代码包及其依赖包（upadte）
    