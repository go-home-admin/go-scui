"use strict";(self["webpackChunkscui"]=self["webpackChunkscui"]||[]).push([[976],{16284:function(e,t,a){a.r(t),a.d(t,{default:function(){return C}});var l=a(58222);const o={class:"left-panel"},i=(0,l.createTextVNode)("分配角色"),r=(0,l.createTextVNode)("密码重置"),s={class:"right-panel"},n={class:"right-panel-search"},d=(0,l.createTextVNode)("查看"),c=(0,l.createTextVNode)("编辑"),h=(0,l.createTextVNode)("删除");function u(e,t,a,u,p,m){const g=(0,l.resolveComponent)("el-input"),f=(0,l.resolveComponent)("el-header"),C=(0,l.resolveComponent)("el-tree"),b=(0,l.resolveComponent)("el-main"),w=(0,l.resolveComponent)("el-container"),N=(0,l.resolveComponent)("el-aside"),v=(0,l.resolveComponent)("el-button"),V=(0,l.resolveComponent)("el-table-column"),x=(0,l.resolveComponent)("el-avatar"),k=(0,l.resolveComponent)("el-popconfirm"),y=(0,l.resolveComponent)("el-button-group"),$=(0,l.resolveComponent)("scTable"),_=(0,l.resolveComponent)("save-dialog"),D=(0,l.resolveDirective)("loading");return(0,l.openBlock)(),(0,l.createElementBlock)(l.Fragment,null,[(0,l.createVNode)(w,null,{default:(0,l.withCtx)((()=>[(0,l.withDirectives)(((0,l.openBlock)(),(0,l.createBlock)(N,{width:"200px"},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(w,null,{default:(0,l.withCtx)((()=>[(0,l.createVNode)(f,null,{default:(0,l.withCtx)((()=>[(0,l.createVNode)(g,{placeholder:"输入关键字进行过滤",modelValue:p.groupFilterText,"onUpdate:modelValue":t[0]||(t[0]=e=>p.groupFilterText=e),clearable:""},null,8,["modelValue"])])),_:1}),(0,l.createVNode)(b,{class:"nopadding"},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(C,{ref:"group",class:"menu","node-key":"id",data:p.group,"current-node-key":"","highlight-current":!0,"expand-on-click-node":!1,"filter-node-method":m.groupFilterNode,onNodeClick:m.groupClick},null,8,["data","filter-node-method","onNodeClick"])])),_:1})])),_:1})])),_:1})),[[D,p.showGrouploading]]),(0,l.createVNode)(w,null,{default:(0,l.withCtx)((()=>[(0,l.createVNode)(f,null,{default:(0,l.withCtx)((()=>[(0,l.createElementVNode)("div",o,[(0,l.createVNode)(v,{type:"primary",icon:"el-icon-plus",onClick:m.add},null,8,["onClick"]),(0,l.createVNode)(v,{type:"danger",plain:"",icon:"el-icon-delete",disabled:0==p.selection.length,onClick:m.batch_del},null,8,["disabled","onClick"]),(0,l.createVNode)(v,{type:"primary",plain:"",disabled:0==p.selection.length},{default:(0,l.withCtx)((()=>[i])),_:1},8,["disabled"]),(0,l.createVNode)(v,{type:"primary",plain:"",disabled:0==p.selection.length},{default:(0,l.withCtx)((()=>[r])),_:1},8,["disabled"])]),(0,l.createElementVNode)("div",s,[(0,l.createElementVNode)("div",n,[(0,l.createVNode)(g,{modelValue:p.search.name,"onUpdate:modelValue":t[1]||(t[1]=e=>p.search.name=e),placeholder:"登录账号 / 姓名",clearable:""},null,8,["modelValue"]),(0,l.createVNode)(v,{type:"primary",icon:"el-icon-search",onClick:m.upsearch},null,8,["onClick"])])])])),_:1}),(0,l.createVNode)(b,{class:"nopadding"},{default:(0,l.withCtx)((()=>[(0,l.createVNode)($,{ref:"table",apiObj:p.apiObj,onSelectionChange:m.selectionChange,stripe:"",remoteSort:"",remoteFilter:""},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(V,{type:"selection",width:"50"}),(0,l.createVNode)(V,{label:"ID",prop:"id",width:"80",sortable:"custom"}),(0,l.createVNode)(V,{label:"头像",width:"80","column-key":"filterAvatar",filters:[{text:"已上传",value:"1"},{text:"未上传",value:"0"}]},{default:(0,l.withCtx)((e=>[(0,l.createVNode)(x,{src:e.row.avatar,size:"small"},null,8,["src"])])),_:1}),(0,l.createVNode)(V,{label:"登录账号",prop:"userName",width:"150",sortable:"custom","column-key":"filterUserName",filters:[{text:"系统账号",value:"1"},{text:"普通账号",value:"0"}]}),(0,l.createVNode)(V,{label:"姓名",prop:"name",width:"150",sortable:"custom"}),(0,l.createVNode)(V,{label:"所属角色",prop:"groupName",width:"200",sortable:"custom"}),(0,l.createVNode)(V,{label:"加入时间",prop:"date",width:"170",sortable:"custom"}),(0,l.createVNode)(V,{label:"操作",fixed:"right",align:"right",width:"160"},{default:(0,l.withCtx)((e=>[(0,l.createVNode)(y,null,{default:(0,l.withCtx)((()=>[(0,l.createVNode)(v,{text:"",type:"primary",size:"small",onClick:t=>m.table_show(e.row,e.$index)},{default:(0,l.withCtx)((()=>[d])),_:2},1032,["onClick"]),(0,l.createVNode)(v,{text:"",type:"primary",size:"small",onClick:t=>m.table_edit(e.row,e.$index)},{default:(0,l.withCtx)((()=>[c])),_:2},1032,["onClick"]),(0,l.createVNode)(k,{title:"确定删除吗？",onConfirm:t=>m.table_del(e.row,e.$index)},{reference:(0,l.withCtx)((()=>[(0,l.createVNode)(v,{text:"",type:"primary",size:"small"},{default:(0,l.withCtx)((()=>[h])),_:1})])),_:2},1032,["onConfirm"])])),_:2},1024)])),_:1})])),_:1},8,["apiObj","onSelectionChange"])])),_:1})])),_:1})])),_:1}),p.dialog.save?((0,l.openBlock)(),(0,l.createBlock)(_,{key:0,ref:"saveDialog",onSuccess:m.handleSuccess,onClosed:t[2]||(t[2]=e=>p.dialog.save=!1)},null,8,["onSuccess"])):(0,l.createCommentVNode)("",!0)],64)}var p=a(11058),m={name:"user",components:{saveDialog:p["default"]},data(){return{dialog:{save:!1},showGrouploading:!1,groupFilterText:"",group:[],apiObj:this.$API.system.user.list,selection:[],search:{name:null}}},watch:{groupFilterText(e){this.$refs.group.filter(e)}},mounted(){this.getGroup()},methods:{add(){this.dialog.save=!0,this.$nextTick((()=>{this.$refs.saveDialog.open()}))},table_edit(e){this.dialog.save=!0,this.$nextTick((()=>{this.$refs.saveDialog.open("edit").setData(e)}))},table_show(e){this.dialog.save=!0,this.$nextTick((()=>{this.$refs.saveDialog.open("show").setData(e)}))},async table_del(e,t){var a={id:e.id},l=await this.$API.demo.post.post(a);200==l.code?(this.$refs.table.tableData.splice(t,1),this.$message.success("删除成功")):this.$alert(l.message,"提示",{type:"error"})},async batch_del(){this.$confirm(`确定删除选中的 ${this.selection.length} 项吗？`,"提示",{type:"warning"}).then((()=>{const e=this.$loading();this.selection.forEach((e=>{this.$refs.table.tableData.forEach(((t,a)=>{e.id===t.id&&this.$refs.table.tableData.splice(a,1)}))})),e.close(),this.$message.success("操作成功")})).catch((()=>{}))},selectionChange(e){this.selection=e},async getGroup(){this.showGrouploading=!0;var e=await this.$API.system.dept.list.get();this.showGrouploading=!1;var t={id:"",label:"所有"};e.data.unshift(t),this.group=e.data},groupFilterNode(e,t){return!e||-1!==t.label.indexOf(e)},groupClick(e){var t={groupId:e.id};this.$refs.table.reload(t)},upsearch(){this.$refs.table.upData(this.search)},handleSuccess(e,t){"add"==t?(e.id=(new Date).getTime(),this.$refs.table.tableData.unshift(e)):"edit"==t&&this.$refs.table.tableData.filter((t=>t.id===e.id)).forEach((t=>{Object.assign(t,e)}))}}},g=a(83744);const f=(0,g.Z)(m,[["render",u]]);var C=f}}]);