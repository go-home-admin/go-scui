"use strict";(self["webpackChunkscui"]=self["webpackChunkscui"]||[]).push([[936],{45316:function(e,t,l){l.r(t),l.d(t,{default:function(){return M}});var o=l(58222);const a=(0,o.createTextVNode)("自定义插槽"),i=(0,o.createTextVNode)("切换multiple"),s=(0,o.createTextVNode)("切换hideUpload");function r(e,t,l,r,n,c){const d=(0,o.resolveComponent)("el-alert"),u=(0,o.resolveComponent)("el-button"),p=(0,o.resolveComponent)("sc-file-select"),m=(0,o.resolveComponent)("el-card"),f=(0,o.resolveComponent)("el-col"),h=(0,o.resolveComponent)("el-row"),g=(0,o.resolveComponent)("el-main");return(0,o.openBlock)(),(0,o.createBlock)(g,null,{default:(0,o.withCtx)((()=>[(0,o.createVNode)(h,{gutter:15},{default:(0,o.withCtx)((()=>[(0,o.createVNode)(f,{lg:18},{default:(0,o.withCtx)((()=>[(0,o.createVNode)(m,{shadow:"never"},{default:(0,o.withCtx)((()=>[(0,o.createVNode)(d,{title:"资源库选择器即将弃用,将不会维护更新,与1.7版本之后将移除此组件",type:"error",style:{"margin-bottom":"20px"}}),(0,o.createVNode)(p,{modelValue:n.file,"onUpdate:modelValue":t[0]||(t[0]=e=>n.file=e),multiple:n.multiple,hideUpload:n.hideUpload,max:99,onSubmit:c.submit},{do:(0,o.withCtx)((()=>[(0,o.createVNode)(u,null,{default:(0,o.withCtx)((()=>[a])),_:1})])),_:1},8,["modelValue","multiple","hideUpload","onSubmit"])])),_:1})])),_:1}),(0,o.createVNode)(f,{lg:6},{default:(0,o.withCtx)((()=>[(0,o.createVNode)(m,{shadow:"never",header:"参数和方法"},{default:(0,o.withCtx)((()=>[(0,o.createVNode)(u,{type:"primary",onClick:t[1]||(t[1]=e=>n.multiple=!n.multiple)},{default:(0,o.withCtx)((()=>[i])),_:1}),(0,o.createVNode)(u,{type:"primary",onClick:t[2]||(t[2]=e=>n.hideUpload=!n.hideUpload)},{default:(0,o.withCtx)((()=>[s])),_:1})])),_:1})])),_:1})])),_:1})])),_:1})}const n=e=>((0,o.pushScopeId)("data-v-48ded1ad"),e=e(),(0,o.popScopeId)(),e),c={class:"sc-file-select"},d={class:"sc-file-select__side"},u={class:"sc-file-select__side-menu"},p={class:"el-tree-node__label"},m={key:0,class:"sc-file-select__side-msg"},f=(0,o.createTextVNode)(" 已选择 "),h=(0,o.createTextVNode)(" / "),g=(0,o.createTextVNode)(" 项 "),k={class:"sc-file-select__files"},v={class:"sc-file-select__top"},_={key:0,class:"upload"},y=(0,o.createTextVNode)("本地上传"),N={class:"tips"},C={class:"keyword"},V={class:"sc-file-select__list"},x={class:"sc-file-select__item__file"},w={class:"sc-file-select__item__upload"},b=["onClick"],B={class:"sc-file-select__item__file"},P={key:0,class:"sc-file-select__item__checkbox"},E={key:1,class:"sc-file-select__item__select"},S=n((()=>(0,o.createElementVNode)("div",{class:"sc-file-select__item__box"},null,-1))),z={key:3,class:"item-file item-file-doc"},D={key:1,class:"sc-icon-file-list-fill",style:{color:"#999"}},I=["title"],L={class:"sc-file-select__pagination"},U={class:"sc-file-select__do"},O=(0,o.createTextVNode)("确 定");function T(e,t,l,a,i,s){const r=(0,o.resolveComponent)("el-icon-folder"),n=(0,o.resolveComponent)("el-icon"),T=(0,o.resolveComponent)("el-tree"),j=(0,o.resolveComponent)("el-button"),q=(0,o.resolveComponent)("el-upload"),$=(0,o.resolveComponent)("el-icon-warning"),F=(0,o.resolveComponent)("el-input"),A=(0,o.resolveComponent)("el-empty"),K=(0,o.resolveComponent)("el-progress"),Z=(0,o.resolveComponent)("el-image"),J=(0,o.resolveComponent)("el-icon-check"),M=(0,o.resolveComponent)("el-scrollbar"),R=(0,o.resolveComponent)("el-pagination"),G=(0,o.resolveDirective)("loading");return(0,o.openBlock)(),(0,o.createElementBlock)("div",c,[(0,o.withDirectives)(((0,o.openBlock)(),(0,o.createElementBlock)("div",d,[(0,o.createElementVNode)("div",u,[(0,o.createVNode)(T,{ref:"group",class:"menu",data:i.menu,"node-key":i.treeProps.key,props:i.treeProps,"current-node-key":i.menu.length>0?i.menu[0][i.treeProps.key]:"","highlight-current":"",onNodeClick:s.groupClick},{default:(0,o.withCtx)((({node:e})=>[(0,o.createElementVNode)("span",p,[(0,o.createVNode)(n,{class:"icon"},{default:(0,o.withCtx)((()=>[(0,o.createVNode)(r)])),_:1}),(0,o.createTextVNode)((0,o.toDisplayString)(e.label),1)])])),_:1},8,["data","node-key","props","current-node-key","onNodeClick"])]),l.multiple?((0,o.openBlock)(),(0,o.createElementBlock)("div",m,[f,(0,o.createElementVNode)("b",null,(0,o.toDisplayString)(i.value.length),1),h,(0,o.createElementVNode)("b",null,(0,o.toDisplayString)(l.max),1),g])):(0,o.createCommentVNode)("",!0)])),[[G,i.menuLoading]]),(0,o.withDirectives)(((0,o.openBlock)(),(0,o.createElementBlock)("div",k,[(0,o.createElementVNode)("div",v,[l.hideUpload?(0,o.createCommentVNode)("",!0):((0,o.openBlock)(),(0,o.createElementBlock)("div",_,[(0,o.createVNode)(q,{class:"sc-file-select__upload",action:"",multiple:"","show-file-list":!1,accept:i.accept,"on-change":s.uploadChange,"before-upload":s.uploadBefore,"on-progress":s.uploadProcess,"on-success":s.uploadSuccess,"on-error":s.uploadError,"http-request":s.uploadRequest},{default:(0,o.withCtx)((()=>[(0,o.createVNode)(j,{type:"primary",icon:"el-icon-upload"},{default:(0,o.withCtx)((()=>[y])),_:1})])),_:1},8,["accept","on-change","before-upload","on-progress","on-success","on-error","http-request"]),(0,o.createElementVNode)("span",N,[(0,o.createVNode)(n,null,{default:(0,o.withCtx)((()=>[(0,o.createVNode)($)])),_:1}),(0,o.createTextVNode)("大小不超过"+(0,o.toDisplayString)(l.maxSize)+"MB",1)])])),(0,o.createElementVNode)("div",C,[(0,o.createVNode)(F,{modelValue:i.keyword,"onUpdate:modelValue":t[0]||(t[0]=e=>i.keyword=e),"prefix-icon":"el-icon-search",placeholder:"文件名搜索",clearable:"",onKeyup:(0,o.withKeys)(s.search,["enter"]),onClear:s.search},null,8,["modelValue","onKeyup","onClear"])])]),(0,o.createElementVNode)("div",V,[(0,o.createVNode)(M,{ref:"scrollbar"},{default:(0,o.withCtx)((()=>[0==i.fileList.length&&0==i.data.length?((0,o.openBlock)(),(0,o.createBlock)(A,{key:0,description:"无数据","image-size":80})):(0,o.createCommentVNode)("",!0),((0,o.openBlock)(!0),(0,o.createElementBlock)(o.Fragment,null,(0,o.renderList)(i.fileList,((e,t)=>((0,o.openBlock)(),(0,o.createElementBlock)("div",{key:t,class:"sc-file-select__item"},[(0,o.createElementVNode)("div",x,[(0,o.createElementVNode)("div",w,[(0,o.createVNode)(K,{type:"circle",percentage:e.progress,width:70},null,8,["percentage"])]),(0,o.createVNode)(Z,{src:e.tempImg,fit:"contain"},null,8,["src"])]),(0,o.createElementVNode)("p",null,(0,o.toDisplayString)(e.name),1)])))),128)),((0,o.openBlock)(!0),(0,o.createElementBlock)(o.Fragment,null,(0,o.renderList)(i.data,(e=>((0,o.openBlock)(),(0,o.createElementBlock)("div",{key:e[i.fileProps.key],class:(0,o.normalizeClass)(["sc-file-select__item",{active:i.value.includes(e[i.fileProps.url])}]),onClick:t=>s.select(e)},[(0,o.createElementVNode)("div",B,[l.multiple?((0,o.openBlock)(),(0,o.createElementBlock)("div",P,[(0,o.createVNode)(n,null,{default:(0,o.withCtx)((()=>[(0,o.createVNode)(J)])),_:1})])):((0,o.openBlock)(),(0,o.createElementBlock)("div",E,[(0,o.createVNode)(n,null,{default:(0,o.withCtx)((()=>[(0,o.createVNode)(J)])),_:1})])),S,s._isImg(e[i.fileProps.url])?((0,o.openBlock)(),(0,o.createBlock)(Z,{key:2,src:e[i.fileProps.url],fit:"contain",lazy:""},null,8,["src"])):((0,o.openBlock)(),(0,o.createElementBlock)("div",z,[i.files[s._getExt(e[i.fileProps.url])]?((0,o.openBlock)(),(0,o.createElementBlock)("i",{key:0,class:(0,o.normalizeClass)(i.files[s._getExt(e[i.fileProps.url])].icon),style:(0,o.normalizeStyle)({color:i.files[s._getExt(e[i.fileProps.url])].color})},null,6)):((0,o.openBlock)(),(0,o.createElementBlock)("i",D))]))]),(0,o.createElementVNode)("p",{title:e[i.fileProps.fileName]},(0,o.toDisplayString)(e[i.fileProps.fileName]),9,I)],10,b)))),128))])),_:1},512)]),(0,o.createElementVNode)("div",L,[(0,o.createVNode)(R,{small:"",background:"",layout:"prev, pager, next",total:i.total,"page-size":i.pageSize,currentPage:i.currentPage,"onUpdate:currentPage":t[1]||(t[1]=e=>i.currentPage=e),onCurrentChange:s.reload},null,8,["total","page-size","currentPage","onCurrentChange"])]),(0,o.createElementVNode)("div",U,[(0,o.renderSlot)(e.$slots,"do",{},void 0,!0),(0,o.createVNode)(j,{type:"primary",disabled:i.value.length<=0,onClick:s.submit},{default:(0,o.withCtx)((()=>[O])),_:1},8,["disabled","onClick"])])])),[[G,i.listLoading]])])}var j=l(7877),q={apiObj:j.Z.common.upload,menuApiObj:j.Z.common.file.menu,listApiObj:j.Z.common.file.list,successCode:200,maxSize:30,max:99,uploadParseData:function(e){return{id:e.data.id,fileName:e.data.fileName,url:e.data.src}},listParseData:function(e){return{rows:e.data.rows,total:e.data.total,msg:e.message,code:e.code}},request:{page:"page",pageSize:"pageSize",keyword:"keyword",menuKey:"groupId"},menuProps:{key:"id",label:"label",children:"children"},fileProps:{key:"id",fileName:"fileName",url:"url"},files:{doc:{icon:"sc-icon-file-word-2-fill",color:"#409eff"},docx:{icon:"sc-icon-file-word-2-fill",color:"#409eff"},xls:{icon:"sc-icon-file-excel-2-fill",color:"#67C23A"},xlsx:{icon:"sc-icon-file-excel-2-fill",color:"#67C23A"},ppt:{icon:"sc-icon-file-ppt-2-fill",color:"#F56C6C"},pptx:{icon:"sc-icon-file-ppt-2-fill",color:"#F56C6C"}}},$={props:{modelValue:null,hideUpload:{type:Boolean,default:!1},multiple:{type:Boolean,default:!1},max:{type:Number,default:q.max},onlyImage:{type:Boolean,default:!1},maxSize:{type:Number,default:q.maxSize}},data(){return{keyword:null,pageSize:20,total:0,currentPage:1,data:[],menu:[],menuId:"",value:this.multiple?[]:"",fileList:[],accept:this.onlyImage?"image/gif, image/jpeg, image/png":"",listLoading:!1,menuLoading:!1,treeProps:q.menuProps,fileProps:q.fileProps,files:q.files}},watch:{multiple(){this.value=this.multiple?[]:"",this.$emit("update:modelValue",JSON.parse(JSON.stringify(this.value)))}},mounted(){this.getMenu(),this.getData()},methods:{async getMenu(){this.menuLoading=!0;var e=await q.menuApiObj.get();this.menu=e.data,this.menuLoading=!1},async getData(){this.listLoading=!0;var e={[q.request.menuKey]:this.menuId,[q.request.page]:this.currentPage,[q.request.pageSize]:this.pageSize,[q.request.keyword]:this.keyword};this.onlyImage&&(e.type="image");var t=await q.listApiObj.get(e),l=q.listParseData(t);this.data=l.rows,this.total=l.total,this.listLoading=!1,this.$refs.scrollbar.setScrollTop(0)},groupClick(e){this.menuId=e.id,this.currentPage=1,this.keyword=null,this.getData()},reload(){this.getData()},search(){this.currentPage=1,this.getData()},select(e){const t=e[this.fileProps.url];this.multiple?this.value.includes(t)?this.value.splice(this.value.findIndex((e=>e==t)),1):this.value.push(t):this.value.includes(t)?this.value="":this.value=t},submit(){const e=JSON.parse(JSON.stringify(this.value));this.$emit("update:modelValue",e),this.$emit("submit",e)},uploadChange(e,t){e.tempImg=URL.createObjectURL(e.raw),this.fileList=t},uploadBefore(e){const t=e.size/1024/1024<this.maxSize;if(!t)return this.$message.warning(`上传文件大小不能超过 ${this.maxSize}MB!`),!1},uploadRequest(e){var t=q.apiObj;const l=new FormData;l.append("file",e.file),l.append([q.request.menuKey],this.menuId),t.post(l,{onUploadProgress:t=>{e.onProgress(t)}}).then((t=>{e.onSuccess(t)})).catch((t=>{e.onError(t)}))},uploadProcess(e,t){t.progress=Number((e.loaded/e.total*100).toFixed(2))},uploadSuccess(e,t){this.fileList.splice(this.fileList.findIndex((e=>e.uid==t.uid)),1);var l=q.uploadParseData(e);this.data.unshift({[this.fileProps.key]:l.id,[this.fileProps.fileName]:l.fileName,[this.fileProps.url]:l.url}),this.multiple||(this.value=l.url)},uploadError(e){this.$notify.error({title:"上传文件错误",message:e})},_isImg(e){const t=[".jpg",".jpeg",".png",".gif",".bmp"],l=e.substring(e.lastIndexOf("."));return-1!=t.indexOf(l)},_getExt(e){return e.substring(e.lastIndexOf(".")+1)}}},F=l(83744);const A=(0,F.Z)($,[["render",T],["__scopeId","data-v-48ded1ad"]]);var K=A,Z={name:"fileselect",components:{scFileSelect:K},data(){return{file:"",multiple:!1,hideUpload:!1,upload:"",upload2:""}},mounted(){},methods:{submit(e){console.log(e),this.$message("返回值请查看F12控制台console.log()")}}};const J=(0,F.Z)(Z,[["render",r]]);var M=J}}]);