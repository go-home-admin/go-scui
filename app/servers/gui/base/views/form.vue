<template>
  <el-form :model="__ID__.form" :rules="__ID__.rules" ref="__ID__Ref" :loading="__ID__.loading">
    <slot id="form-item"/>
  </el-form>

  <template #footer>
    <el-button>重 置</el-button>
    <el-button type="primary" @click="__ID__onSubmit()">提 交</el-button>
  </template>
</template>