<el-container>
  <el-header>
    <div class="left-panel">
      <el-button type="primary" icon="el-icon-plus"></el-button>
      <el-button type="danger" plain icon="el-icon-delete"></el-button>
    </div>
    <div class="right-panel">
      <el-alert title="请在右下角第二个按钮体验自定义列" type="warning" :closable="false"/>
    </div>
  </el-header>
  <div v-if="filter.config.length !==0" style="border-bottom:0.8px solid #dcdfe6;background-color: #ffffff;padding: 13px 15px">
    <el-container>
      <sc-form ref="formref" :config="filter.config" v-model="filter.form" :loading="filter.loading" style="width: 100%">
        <el-button type="primary" >搜 索</el-button>
        <el-button>重 置</el-button>
      </sc-form>
    </el-container>
  </div>

  <el-main class="nopadding">
    <scTable ref="table" tableName="listCustomColumn" :apiObj="{get: get}" :column="columns" row-key="id" stripe>
      <el-table-column type="selection" width="50"></el-table-column>
      <template #progress="scope">
        <el-progress :percentage="scope.row.progress"/>
      </template>
    </scTable>
  </el-main>
</el-container>
