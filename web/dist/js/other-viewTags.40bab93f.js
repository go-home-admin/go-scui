"use strict";(self["webpackChunkscui"]=self["webpackChunkscui"]||[]).push([[6488],{93653:function(e,t,o){o.r(t),o.d(t,{default:function(){return N}});var r=o(58222);const a=(0,r.createTextVNode)("打开个人信息"),l=(0,r.createTextVNode)("打开后执行"),i=(0,r.createTextVNode)("刷新当前"),n=(0,r.createTextVNode)("关闭当前"),c=(0,r.createTextVNode)("关闭其他"),s=(0,r.createTextVNode)("关闭后执行"),u=(0,r.createTextVNode)("设置标题"),p=(0,r.createTextVNode)("fullpage");function m(e,t,o,m,d,h){const f=(0,r.resolveComponent)("el-button"),C=(0,r.resolveComponent)("el-alert"),v=(0,r.resolveComponent)("el-card"),w=(0,r.resolveComponent)("el-input"),x=(0,r.resolveComponent)("el-form-item"),V=(0,r.resolveComponent)("el-form"),N=(0,r.resolveComponent)("el-main");return(0,r.openBlock)(),(0,r.createBlock)(N,null,{default:(0,r.withCtx)((()=>[(0,r.createVNode)(v,{shadow:"never",header:"打开"},{default:(0,r.withCtx)((()=>[(0,r.createVNode)(f,{type:"primary",plain:"",onClick:h.open1},{default:(0,r.withCtx)((()=>[a])),_:1},8,["onClick"]),(0,r.createVNode)(f,{type:"primary",plain:"",onClick:h.open2},{default:(0,r.withCtx)((()=>[l])),_:1},8,["onClick"]),(0,r.createVNode)(C,{title:"打开后执行原理: 路由push时,在当前路由对象中插入一个特殊标识, 在目标视图中beforeRouteEnter获取判断是否需要执行特殊方法",style:{"margin-top":"20px"}})])),_:1}),(0,r.createVNode)(v,{shadow:"never",header:"刷新",style:{"margin-top":"15px"}},{default:(0,r.withCtx)((()=>[(0,r.createVNode)(f,{type:"primary",plain:"",onClick:h.refresh1},{default:(0,r.withCtx)((()=>[i])),_:1},8,["onClick"])])),_:1}),(0,r.createVNode)(v,{shadow:"never",header:"关闭",style:{"margin-top":"15px"}},{default:(0,r.withCtx)((()=>[(0,r.createVNode)(f,{type:"primary",plain:"",onClick:h.close1},{default:(0,r.withCtx)((()=>[n])),_:1},8,["onClick"]),(0,r.createVNode)(f,{type:"primary",plain:"",onClick:h.close2},{default:(0,r.withCtx)((()=>[c])),_:1},8,["onClick"]),(0,r.createVNode)(f,{type:"primary",plain:"",onClick:h.close3},{default:(0,r.withCtx)((()=>[s])),_:1},8,["onClick"])])),_:1}),(0,r.createVNode)(v,{shadow:"never",header:"设置",style:{"margin-top":"15px"}},{default:(0,r.withCtx)((()=>[(0,r.createVNode)(V,{inline:!0},{default:(0,r.withCtx)((()=>[(0,r.createVNode)(x,null,{default:(0,r.withCtx)((()=>[(0,r.createVNode)(w,{modelValue:d.input,"onUpdate:modelValue":t[0]||(t[0]=e=>d.input=e),placeholder:"请输入内容"},null,8,["modelValue"])])),_:1}),(0,r.createVNode)(x,null,{default:(0,r.withCtx)((()=>[(0,r.createVNode)(f,{type:"primary",plain:"",onClick:h.set1},{default:(0,r.withCtx)((()=>[u])),_:1},8,["onClick"])])),_:1})])),_:1})])),_:1}),(0,r.createVNode)(v,{shadow:"never",header:"整页路由",style:{"margin-top":"15px"}},{default:(0,r.withCtx)((()=>[(0,r.createVNode)(f,{type:"primary",plain:"",onClick:h.fullpage},{default:(0,r.withCtx)((()=>[p])),_:1},8,["onClick"]),(0,r.createVNode)(C,{title:"变更路由的层级关系,向上推至顶级达到在layout视图中显示. 只需要在路由中设置 meta.fullpage 即可",style:{"margin-top":"20px"}})])),_:1})])),_:1})}var d=o(74865),h=o.n(d),f=o(2314),C=o(24239),v={refresh(){h().start();const e=f.Z.currentRoute.value;C.Z.commit("removeKeepLive",e.name),C.Z.commit("setRouteShow",!1),(0,r.nextTick)((()=>{C.Z.commit("pushKeepLive",e.name),C.Z.commit("setRouteShow",!0),h().done()}))},close(e){const t=e||f.Z.currentRoute.value;C.Z.commit("removeViewTags",t),C.Z.commit("removeIframeList",t),C.Z.commit("removeKeepLive",t.name);const o=C.Z.state.viewTags.viewTags,r=o.slice(-1)[0];r?f.Z.push(r):f.Z.push("/")},closeNext(e){const t=f.Z.currentRoute.value;if(C.Z.commit("removeViewTags",t),C.Z.commit("removeIframeList",t),C.Z.commit("removeKeepLive",t.name),e){const t=C.Z.state.viewTags.viewTags;e(t)}},closeOther(){const e=f.Z.currentRoute.value,t=[...C.Z.state.viewTags.viewTags];t.forEach((t=>{if(t.meta&&t.meta.affix||e.fullPath==t.fullPath)return!0;this.close(t)}))},setTitle(e){C.Z.commit("updateViewTagsTitle",e)}},w={name:"viewTags",data(){return{input:"newTabName"}},mounted(){},methods:{open1(){this.$router.push("/usercenter")},open2(){this.$router.push("/usercenter"),this.$route.is=!0},refresh1(){v.refresh()},close1(){v.close()},close2(){v.closeOther()},close3(){v.closeNext((e=>{console.log(e),this.$router.push("/usercenter"),this.$route.is=!0}))},set1(){v.setTitle(this.input)},fullpage(){this.$router.push("/other/fullpage")}}},x=o(83744);const V=(0,x.Z)(w,[["render",m]]);var N=V}}]);