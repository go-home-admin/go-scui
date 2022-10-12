<template>
  <el-dialog v-model="__ID__visible" :width="500">
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