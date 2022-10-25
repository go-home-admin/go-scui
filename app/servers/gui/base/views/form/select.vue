<template>
  <el-form-item label="__LABEL__" prop="__PROP__">
    <el-select v-model="__FORM__.form.__PROP__" __options__ clearable>
      <el-option
          v-for="item in __ID__options"
          :key="item.value"
          :label="item.label"
          :value="item.value">
      </el-option>
    </el-select>
  </el-form-item>
</template>

<script>

export default {
  data() {
    return {
      "__ID__options": [],
      "__ID__URL": ""
    }
  },
  mounted() {
    this.init__ID__Options()
  },
  methods: {
    init__ID__Options() {
      if (this.__ID__options.length === 0 && this.__ID__URL !== "") {
        this.$HTTP.get(this.__ID__URL).then(res=>{
          this.__ID__options = res.data.options
        })
      }
    }
  }
}
</script>