<template>
  <el-col :span="__SPAN__">
  <el-form-item label="__LABEL__" prop="__PROP__">
    <el-checkbox-group v-model="__FORM__.form.__PROP__" __options__ clearable>
      <el-checkbox
          v-for="item in __ID__options"
          :key="item.value"
          :label="item.label">
      {{ item.value }}
      </el-checkbox>
    </el-checkbox-group>
  </el-form-item>
  </el-col>
</template>

<script>

export default {
  data() {
    return {
      "__ID__options": []
    }
  },
  methods: {
    __ID__SetData: function (data) {
      this.__ID__options = data
    },
  }
}

</script>