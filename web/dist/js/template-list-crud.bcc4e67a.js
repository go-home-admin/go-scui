"use strict";(self["webpackChunkscui"]=self["webpackChunkscui"]||[]).push([[5381],{45286:function(e,t,o){o.r(t),o.d(t,{default:function(){return B}});var l=o(58222);const a={class:"left-panel"},n=(0,l.createTextVNode)("筛选"),i={class:"right-panel"},d={class:"right-panel-search"},r=(0,l.createTextVNode)("新增"),c=(0,l.createTextVNode)(" 导出 "),s=(0,l.createTextVNode)("全部"),u=(0,l.createTextVNode)("当前页"),p=(0,l.createTextVNode)("选中的行"),h={style:{margin:"15px 0px 0px 0px","border-bottom":"1px solid var(--el-border-color-light)"}},m=(0,l.createTextVNode)("搜索"),C=(0,l.createTextVNode)("重置"),f=(0,l.createTextVNode)("查看"),V=(0,l.createTextVNode)("编辑"),w=(0,l.createTextVNode)("删除"),N=(0,l.createTextVNode)("配额"),x=(0,l.createTextVNode)("重启"),v=(0,l.createTextVNode)("停机"),b=(0,l.createTextVNode)("释放主机");function g(e,t,o,g,_,k){const y=(0,l.resolveComponent)("el-button"),T=(0,l.resolveComponent)("el-dropdown-item"),$=(0,l.resolveComponent)("el-dropdown-menu"),B=(0,l.resolveComponent)("el-dropdown"),S=(0,l.resolveComponent)("el-header"),D=(0,l.resolveComponent)("el-input"),z=(0,l.resolveComponent)("el-form-item"),j=(0,l.resolveComponent)("el-option"),E=(0,l.resolveComponent)("el-select"),A=(0,l.resolveComponent)("el-form"),O=(0,l.resolveComponent)("el-col"),F=(0,l.resolveComponent)("el-row"),I=(0,l.resolveComponent)("el-table-column"),P=(0,l.resolveComponent)("sc-status-indicator"),U=(0,l.resolveComponent)("el-popconfirm"),Z=(0,l.resolveComponent)("scTable"),K=(0,l.resolveComponent)("el-main"),G=(0,l.resolveComponent)("el-container"),R=(0,l.resolveComponent)("save-dialog"),q=(0,l.resolveComponent)("info"),H=(0,l.resolveComponent)("el-drawer");return(0,l.openBlock)(),(0,l.createElementBlock)(l.Fragment,null,[(0,l.createVNode)(G,null,{default:(0,l.withCtx)((()=>[(0,l.createVNode)(S,null,{default:(0,l.withCtx)((()=>[(0,l.createElementVNode)("div",a,[(0,l.createVNode)(y,{icon:"el-icon-Filter",onClick:k.add},{default:(0,l.withCtx)((()=>[n])),_:1},8,["onClick"])]),(0,l.createElementVNode)("div",i,[(0,l.createElementVNode)("div",d,[(0,l.createVNode)(y,{type:"primary",icon:"el-icon-plus",onClick:k.add},{default:(0,l.withCtx)((()=>[r])),_:1},8,["onClick"]),(0,l.createVNode)(B,{"split-button":"",type:"primary"},{dropdown:(0,l.withCtx)((()=>[(0,l.createVNode)($,null,{default:(0,l.withCtx)((()=>[(0,l.createVNode)(T,null,{default:(0,l.withCtx)((()=>[s])),_:1}),(0,l.createVNode)(T,null,{default:(0,l.withCtx)((()=>[u])),_:1}),(0,l.createVNode)(T,null,{default:(0,l.withCtx)((()=>[p])),_:1})])),_:1})])),default:(0,l.withCtx)((()=>[c])),_:1}),(0,l.createVNode)(y,{type:"danger",plain:"",icon:"el-icon-delete",disabled:0==_.selection.length,onClick:k.batch_del},null,8,["disabled","onClick"])])])])),_:1}),(0,l.createVNode)(K,{class:"nopadding"},{default:(0,l.withCtx)((()=>[(0,l.createElementVNode)("div",h,[(0,l.createVNode)(F,null,{default:(0,l.withCtx)((()=>[(0,l.createVNode)(O,{span:24},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(A,{model:_.form,"label-width":"120px"},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(z,{label:"Activity name"},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(D,{modelValue:_.form.name,"onUpdate:modelValue":t[0]||(t[0]=e=>_.form.name=e)},null,8,["modelValue"])])),_:1}),(0,l.createVNode)(z,{label:"Activity zone"},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(E,{modelValue:_.form.region,"onUpdate:modelValue":t[1]||(t[1]=e=>_.form.region=e),placeholder:"please select your zone"},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(j,{label:"Zone one",value:"shanghai"}),(0,l.createVNode)(j,{label:"Zone two",value:"beijing"})])),_:1},8,["modelValue"])])),_:1}),(0,l.createVNode)(z,null,{default:(0,l.withCtx)((()=>[(0,l.createVNode)(y,{type:"primary",onClick:e.onSubmit},{default:(0,l.withCtx)((()=>[m])),_:1},8,["onClick"]),(0,l.createVNode)(y,null,{default:(0,l.withCtx)((()=>[C])),_:1})])),_:1})])),_:1},8,["model"])])),_:1})])),_:1})]),(0,l.createVNode)(Z,{ref:"table",apiObj:_.list.apiObj,"row-key":"id",onSelectionChange:k.selectionChange,stripe:""},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(I,{type:"selection",width:"50"}),(0,l.createVNode)(I,{label:"姓名",prop:"name",width:"180"}),(0,l.createVNode)(I,{label:"性别",prop:"sex",width:"150"}),(0,l.createVNode)(I,{label:"邮箱",prop:"email",width:"250"}),(0,l.createVNode)(I,{label:"状态",prop:"boolean",width:"60"},{default:(0,l.withCtx)((e=>[e.row.boolean?((0,l.openBlock)(),(0,l.createBlock)(P,{key:0,type:"success"})):(0,l.createCommentVNode)("",!0),e.row.boolean?(0,l.createCommentVNode)("",!0):((0,l.openBlock)(),(0,l.createBlock)(P,{key:1,pulse:"",type:"danger"}))])),_:1}),(0,l.createVNode)(I,{label:"评分",prop:"num",width:"150"}),(0,l.createVNode)(I,{label:"操作",fixed:"right",align:"right",width:"300"},{default:(0,l.withCtx)((e=>[(0,l.createVNode)(y,{plain:"",size:"small",onClick:t=>k.table_show(e.row)},{default:(0,l.withCtx)((()=>[f])),_:2},1032,["onClick"]),(0,l.createVNode)(y,{type:"primary",plain:"",size:"small",onClick:t=>k.table_edit(e.row)},{default:(0,l.withCtx)((()=>[V])),_:2},1032,["onClick"]),(0,l.createVNode)(U,{title:"确定删除吗？",onConfirm:t=>k.table_del(e.row,e.$index)},{reference:(0,l.withCtx)((()=>[(0,l.createVNode)(y,{plain:"",type:"danger",size:"small"},{default:(0,l.withCtx)((()=>[w])),_:1})])),_:2},1032,["onConfirm"]),(0,l.createVNode)(B,null,{dropdown:(0,l.withCtx)((()=>[(0,l.createVNode)($,null,{default:(0,l.withCtx)((()=>[(0,l.createVNode)(T,null,{default:(0,l.withCtx)((()=>[N])),_:1}),(0,l.createVNode)(T,{divided:""},{default:(0,l.withCtx)((()=>[x])),_:1}),(0,l.createVNode)(T,null,{default:(0,l.withCtx)((()=>[v])),_:1}),(0,l.createVNode)(T,{divided:""},{default:(0,l.withCtx)((()=>[b])),_:1})])),_:1})])),default:(0,l.withCtx)((()=>[(0,l.createVNode)(y,{icon:"el-icon-more",size:"small"})])),_:1})])),_:1})])),_:1},8,["apiObj","onSelectionChange"])])),_:1})])),_:1}),_.dialog.save?((0,l.openBlock)(),(0,l.createBlock)(R,{key:0,ref:"saveDialog",onSuccess:k.handleSaveSuccess,onClosed:t[2]||(t[2]=e=>_.dialog.save=!1)},null,8,["onSuccess"])):(0,l.createCommentVNode)("",!0),(0,l.createVNode)(H,{modelValue:_.dialog.info,"onUpdate:modelValue":t[3]||(t[3]=e=>_.dialog.info=e),size:800,title:"详细","custom-class":"drawerBG",direction:"rtl","destroy-on-close":""},{default:(0,l.withCtx)((()=>[(0,l.createVNode)(q,{ref:"infoDialog"},null,512)])),_:1},8,["modelValue"])],64)}var _=o(98625),k=o(80927),y={name:"listCrud",components:{saveDialog:_["default"],info:k["default"]},data(){return{dialog:{save:!1,info:!1},list:{apiObj:this.$API.demo.list},selection:[],form:{name:"",region:""}}},mounted(){},methods:{onFilter(){},add(){this.dialog.save=!0,this.$nextTick((()=>{this.$refs.saveDialog.open()}))},table_edit(e){this.dialog.save=!0,this.$nextTick((()=>{this.$refs.saveDialog.open("edit").setData(e)}))},addPage(){this.$router.push({path:"/template/list/crud/detail"})},table_show(e){this.dialog.info=!0,this.$nextTick((()=>{this.$refs.infoDialog.setData(e)}))},async table_del(e,t){var o={id:e.id},l=await this.$API.demo.post.post(o);200==l.code?(this.$refs.table.removeIndex(t),this.$message.success("删除成功")):this.$alert(l.message,"提示",{type:"error"})},async batch_del(){var e=await this.$confirm(`确定删除选中的 ${this.selection.length} 项吗？`,"提示",{type:"warning",confirmButtonText:"删除",confirmButtonClass:"el-button--danger"}).catch((()=>{}));if(!e)return!1;var t=this.selection.map((e=>e.id));this.$refs.table.removeKeys(t),this.$message.success("操作成功")},selectionChange(e){this.selection=e},handleSaveSuccess(e,t){"add"==t?this.$refs.table.unshiftRow(e):"edit"==t&&this.$refs.table.updateKey(e)}}},T=o(83744);const $=(0,T.Z)(y,[["render",g]]);var B=$}}]);