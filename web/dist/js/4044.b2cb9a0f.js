"use strict";(self["webpackChunkscui"]=self["webpackChunkscui"]||[]).push([[4044],{4044:function(e,t,i){i.r(t),i.d(t,{default:function(){return h}});var n=i(58222);function a(e,t,i,a,r,l){return(0,n.openBlock)(),(0,n.createElementBlock)("div",{class:"sc-code-editor",style:(0,n.normalizeStyle)({height:l._height})},[(0,n.withDirectives)((0,n.createElementVNode)("textarea",{ref:"textarea","onUpdate:modelValue":t[0]||(t[0]=e=>r.contentValue=e)},null,512),[[n.vModelText,r.contentValue]])],4)}var r=i(4631),l=i.n(r),o=(i(20017),i(96876),i(54086),{props:{modelValue:{type:String,default:""},mode:{type:String,default:"javascript"},height:{type:[String,Number],default:300},options:{type:Object,default:()=>{}},theme:{type:String,default:"idea"},readOnly:{type:Boolean,default:!1}},data(){return{contentValue:this.modelValue,coder:null,opt:{theme:this.theme,styleActiveLine:!0,lineNumbers:!0,lineWrapping:!1,tabSize:4,indentUnit:4,indentWithTabs:!0,mode:this.mode,readOnly:this.readOnly,...this.options}}},computed:{_height(){return Number(this.height)?Number(this.height)+"px":this.height}},watch:{modelValue(e){this.contentValue=e,e!==this.coder.getValue()&&this.coder.setValue(e)}},mounted(){this.init()},methods:{init(){this.coder=(0,n.markRaw)(l().fromTextArea(this.$refs.textarea,this.opt)),this.coder.on("change",(e=>{this.contentValue=e.getValue(),this.$emit("update:modelValue",this.contentValue)}))},formatStrInJson(e){return JSON.stringify(JSON.parse(e),null,4)}}}),s=i(83744);const u=(0,s.Z)(o,[["render",a],["__scopeId","data-v-a0186eba"]]);var h=u}}]);