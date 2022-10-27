<template>
  <el-form-item label="__LABEL__" prop="__PROP__">
    <el-upload
        ref="upload"
        action=""
        :on-preview="handlePreview"
        :on-remove="handleRemove"
        :file-list="__ID__options"
        list-type="picture"
        :auto-upload="false">
      <el-button size="small" type="primary">选取文件</el-button>
      <div class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
    </el-upload>
  </el-form-item>
</template>

<script>

export default {
  data() {
    return {
      "__ID__options": []
    }
  },
  methods: {
    handlePreview: function () {

    },
    handleRemove: function () {

    }
  }
}

</script>