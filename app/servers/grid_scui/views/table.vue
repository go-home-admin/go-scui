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
