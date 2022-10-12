<template>
  <el-container>
    <el-header>
      <div class="left-panel">
        <el-button type="primary" icon="el-icon-filter" @click="onFilter()"></el-button>
      </div>
      <div class="right-panel">
        <el-button type="danger" plain icon="el-icon-delete"></el-button>
        <el-button type="primary" icon="el-icon-plus"></el-button>
      </div>
    </el-header>
    <div v-if="filter.config.length !==0" v-show="filterShow"
         style="border-bottom:0.8px solid #dcdfe6;background-color: #ffffff;padding: 13px 15px">
      <el-container>
        <slot id="search"/>
      </el-container>
    </div>

    <el-main class="nopadding">
      <scTable ref="table" tableName="listCustomColumn" :apiObj="{get: getData}" :column="columns" row-key="id" stripe>
        <el-table-column type="selection" width="50"></el-table-column>

        <el-table-column label="操作" fixed="right" align="right" width="300">
          <template #default="{row}">
            <slot id="actions"/>
          </template>
        </el-table-column>
      </scTable>
    </el-main>
  </el-container>
</template>

<script>

export default {
  data() {
    return {
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
  }
}
</script>