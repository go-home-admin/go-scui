<template>
  <el-col :span="__SPAN__">
  <el-form-item label="__LABEL__" prop="__PROP__">
    <el-date-picker v-model="__FORM__.form.__PROP__" type="datetime" placeholder="Select date and time" __KV__ />
  </el-form-item>
  </el-col>
</template>

<script>

export default {
  data() {
    return {
      "__FORM__.form.__PROP__": ""
    }
  },
  methods: {

  }
}
</script>