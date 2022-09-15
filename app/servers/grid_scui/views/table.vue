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
  <el-main class="nopadding">
    <scTable ref="table" tableName="listCustomColumn" :apiObj="{get: get}" :column="columns" row-key="id" stripe>
      <el-table-column type="selection" width="50"></el-table-column>
      <template #progress="scope">
        <el-progress :percentage="scope.row.progress"/>
      </template>
    </scTable>
  </el-main>
</el-container>
