"use strict";(self["webpackChunkscui"]=self["webpackChunkscui"]||[]).push([[2637],{73278:function(__unused_webpack_module,__webpack_exports__){__webpack_exports__["Z"]={name:"uploadRender",props:{modelValue:[String,Number,Boolean,Date,Object,Array],item:{type:Object,default:()=>{}}},data(){return{value:this.modelValue,apiObj:this.getApiObj()}},watch:{value(e){this.$emit("update:modelValue",e)}},mounted(){},methods:{getApiObj(){return eval("this."+this.item.options.apiObj)}}}},2637:function(e,t,l){l.r(t),l.d(t,{default:function(){return i}});var o=l(58222);function a(e,t,l,a,p,n){const u=(0,o.resolveComponent)("el-table-column"),i=(0,o.resolveComponent)("sc-table-select");return(0,o.openBlock)(),(0,o.createBlock)(i,{modelValue:p.value,"onUpdate:modelValue":t[0]||(t[0]=e=>p.value=e),apiObj:p.apiObj,"table-width":600,multiple:l.item.options.multiple,props:l.item.options.props,style:{width:"100%"}},{default:(0,o.withCtx)((()=>[((0,o.openBlock)(!0),(0,o.createElementBlock)(o.Fragment,null,(0,o.renderList)(l.item.options.column,((e,t)=>((0,o.openBlock)(),(0,o.createBlock)(u,{key:t,prop:e.prop,label:e.label,width:e.width},null,8,["prop","label","width"])))),128))])),_:1},8,["modelValue","apiObj","multiple","props"])}var p=l(73278),n=l(83744);const u=(0,n.Z)(p.Z,[["render",a]]);var i=u}}]);