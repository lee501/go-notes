#### go依赖包的版本升级和降级

##### 升级
    * 使用go list -m -versions 查看依赖包的相关版本
        
        ```shell
            go list -m -versions github.com/gin-gonic/gin
        ```
    * 查看当前使用版本

        ```shell
            go list -m github.com/gin-gonic/gin
        ```

    * 使用go get 指定版本
      
        ```shell
            go get github.com/gin-gonic/gin@v1.7.3
        ```

##### 降级

    * 首先使用go mod edit来修改版本
    ```shell
        go mod edit -require=github.com/gin-gonic/gin@v1.7.3
    ```

    *然后使用go mod tidy, 降级成功使用go mod -m查看版本信息
