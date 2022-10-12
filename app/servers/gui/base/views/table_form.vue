<template>
  <sc-form ref="formref" :config="filter.config" v-model="filter.form" :loading="filter.loading" style="width: 100%">
    <el-button type="primary" @click="onSubmit()">提 交</el-button>
    <el-button>重 置</el-button>
  </sc-form>
</template>