本目录不是真实的vue模版，不会支持全部语法, 固定支持以下模版有的结构，没有的不支持
使用

    Render := LoadVue(file)

````vue

<template>
  <el-container>
    data 只能json格式, 必须是对象
    函数只支持 onFilter: function () { },
  </el-container>
</template>

<script>

export default {
  data() {
    return {
      "key": ""
    }
  },
  mounted() {

  },
  methods: {
    //点击筛选按钮
    onFilter: function () {

    },
  }
}
</script>
````