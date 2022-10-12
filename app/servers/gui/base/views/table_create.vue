<el-dialog :title="titleMap[mode]" v-model="visible" :width="500" destroy-on-close @closed="$emit('closed')">
  <el-form :model="form" :rules="rules" :disabled="mode=='show'" ref="dialogForm" label-width="100px">
    <el-form-item label="姓名" prop="name">
      <el-input v-model="form.name" placeholder="请输入姓名" clearable></el-input>
    </el-form-item>
    <el-form-item label="性别" prop="sex">
      <el-radio-group v-model="form.sex">
        <el-radio label="男">男</el-radio>
        <el-radio label="女">女</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="邮箱" prop="email">
      <el-input v-model="form.email" placeholder="请输入邮箱" clearable></el-input>
    </el-form-item>
    <el-form-item label="评分" prop="num">
      <el-input-number v-model="form.num" :min="0" style="width: 100%;"></el-input-number>
    </el-form-item>
    <el-form-item label="状态" prop="boolean">
      <el-switch v-model="form.boolean"></el-switch>
    </el-form-item>
  </el-form>
  <template #footer>
    <el-button @click="visible=false">取 消</el-button>
    <el-button v-if="mode!='show'" type="primary" :loading="isSaveing" @click="submit()">保 存</el-button>
  </template>
</el-dialog>