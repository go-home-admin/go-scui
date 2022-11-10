<template>
  <el-container>
    <el-header>
      <div class="left-panel">
        <el-button type="primary" icon="el-icon-filter" @click="onFilter()"></el-button>
      </div>
      <div class="right-panel">
        <slot id="header"/>
      </div>
    </el-header>
    <div v-if="filterShow" style="border-bottom:0.8px solid #dcdfe6;background-color: #ffffff;padding: 13px 15px">
      <el-container>
        <slot id="search"/>
      </el-container>
    </div>

    <el-main class="nopadding">
      <scTable ref="table" tableName="listCustomColumn" :apiObj="getData" :column="columns" row-key="id" stripe>
        <el-table-column type="selection" width="50"></el-table-column>
        <slot id="el-table-column"/>
      </scTable>
    </el-main>
  </el-container>
</template>

<script>

export default {
  data() {
    return {
      "primary": "id",
      "filterShow": false
    }
  },
  methods: {
    //点击筛选按钮
    onFilter: function () {
      if (this.filterShow) {
        this.filterShow = false
      } else {
        this.filterShow = true
      }
    },
    getData: async function (params) {
      return await this.$HTTP.get(this.url, params);
    },
  }
}
</script>