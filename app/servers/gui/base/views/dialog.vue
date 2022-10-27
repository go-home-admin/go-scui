<template>
  <el-dialog v-model="__ID__visible" __OPT__>
    <slot id="context"/>
  </el-dialog>
</template>

<script>

export default {
  data() {
    return {
      "__ID__visible": false
    }
  }
}
</script>