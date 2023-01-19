import http from "@/utils/request";
import config from "@/config";

/**
 *  登陆
 * @param {{username?:string,password?:string}} data
 * @returns {Promise<{code:Number,data:{token:string,userInfo:},message:string}>}
 * @callback
 */
export async function LoginPost(data) {
	return await http.post(`${config.API_URL}/login`, data);
}

/**
 *  退出登陆
 * @returns {Promise<{code:Number,data:{tip:string},message:string}>}
 * @callback
 */
export async function LogoutPost() {
	return await http.post(`${config.API_URL}/logout`);
}

/**
 *  登陆
 * @param {{username?:string,password?:string}} data
 * @returns {Promise<{code:Number,data:{userInfo:,token:string},message:string}>}
 * @callback
 */
export async function TokenPost(data) {
	return await http.post(`${config.API_URL}/token`, data);
}

/**
 *  用户创建
 * @param {{id?:number,name?:string,password?:string,userName?:string}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @callback
 */
export async function AdminUserPost(data) {
	return await http.post(`${config.API_URL}/admin/user`, data);
}

/**
 *  用户删除
 * @param {{id?:number}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @callback
 */
export async function AdminUserDelete(data) {
	return await http.delete(`${config.API_URL}/admin/user`, data);
}

/**
 *  登陆账户信息
 * @param {{id?:number}} data
 * @returns {Promise<{code:Number,data:{roles:{}[],avatar:string,introduction:string,name:string},message:string}>}
 * @callback
 */
export async function InfoGet(data) {
	return await http.get(`${config.API_URL}/info`, data);
}

/**
 *  菜单删除
 * @param {{ids?:uint32[]}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @callback
 */
export async function SystemMenuDelDelete(data) {
	return await http.delete(`${config.API_URL}/system/menu/del`, data);
}

/**
 *  菜单列表
 * @returns {Promise<{code:Number,data:{list:{component:string,parentId:number,name:string,path:string,redirect:string,active:string,apiList:{code:string,url:string}[],children:{}[],id:number,meta:}[]},message:string}>}
 * @callback
 */
export async function SystemMenuListGet() {
	return await http.get(`${config.API_URL}/system/menu/list`);
}

/**
 *  登陆用户的菜单
 * @returns {Promise<{code:Number,data:{menu:{children:{}[],component:string,meta:,name:string,path:string}[],permissions:{}[]},message:string}>}
 * @callback
 */
export async function SystemMenuMyGet() {
	return await http.get(`${config.API_URL}/system/menu/my`);
}

/**
 *  菜单保存
 * @param {{id?:number,parentId?:number,name?:string,path?:string,meta?:,apiList?:{code?:string,url?:string}[],component?:string}} data
 * @returns {Promise<{code:Number,data:{id:number},message:string}>}
 * @callback
 */
export async function SystemMenuSavePost(data) {
	return await http.post(`${config.API_URL}/system/menu/save`, data);
}
