import http from "@/utils/request";
import config from "@/config";

/**
 *  登陆
 * @param {{username:string,password:string}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function LoginPost(data) {
	return await http.post(config.API_URL + "/login", data);
}

/**
 *  退出登陆
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function LogoutPost(data) {
	return await http.post(config.API_URL + "/logout", data);
}

/**
 *  登陆
 * @param {{username:string,password:string}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function TokenPost(data) {
	return await http.post(config.API_URL + "/token", data);
}

/**
 *  用户创建
 * @param {{id:{},name:string,password:string,userName:string}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function AdminUserPost(data) {
	return await http.post(config.API_URL + "/admin/user", data);
}

/**
 *  用户删除
 * @param {{id:{}}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function AdminUserDelete(data) {
	return await http.delete(config.API_URL + "/admin/user", data);
}

/**
 *  登陆账户信息
 * @param {{id:{}}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function InfoGet(data) {
	return await http.get(config.API_URL + "/info", data);
}

/**
 *  菜单删除
 * @param {{ids:{}}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function SystemMenuDelDelete(data) {
	return await http.delete(config.API_URL + "/system/menu/del", data);
}

/**
 *  菜单列表
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function SystemMenuListGet(data) {
	return await http.get(config.API_URL + "/system/menu/list", data);
}

/**
 *  登陆用户的菜单
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function SystemMenuMyGet(data) {
	return await http.get(config.API_URL + "/system/menu/my", data);
}

/**
 *  菜单保存
 * @param {{id:{},parentId:{},name:string,path:string,meta:string,apiList:{},component:string}} data
 * @returns {Promise<{code:Number,data:{},message:string}>}
 * @constructor
 */
export async function SystemMenuSavePost(data) {
	return await http.post(config.API_URL + "/system/menu/save", data);
}
