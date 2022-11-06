<template>
  <el-form :model="__ID__.form" :rules="__ID__.rules" ref="__ID__Ref" __FORM_OPT__>
    <el-row :gutter="15">
      <slot id="form-item"/>
    </el-row>
  </el-form>

  <template #footer>
    <el-button>重 置</el-button>
    <el-button type="primary" @click="__ID__onSubmit()">提 交</el-button>
  </template>
</template>

<script>

export default {
  data() {
    return {
      "__ID__": {
        "form": {},
        "rules": {},
        "url": ""
      }
    }
  },
  methods: {
    __ID__SetData: function (form) {
      this.__ID__.form = form
    },
    __ID__onSubmit: function () {
      this.$HTTP.post(this.__ID__.url, this.__ID__.form).then(()=>{
        this.$refs.table.refresh()
      })
    },
  }
}
</script>