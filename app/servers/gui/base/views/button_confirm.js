function ActionsConfirm (row) {
    this.$confirm('这是一个敏感的操作，是否继续?', '敏感操作提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
        center: true,
    }).then(() => {
        this.$HTTP.get(__url__)
    })
}