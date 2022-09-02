import http from "@/utils/request"

/**
 *
 * @param data {{id: number}} 放到啊
 * @returns {Promise<{data: {id: number}}>}
 */
export function token(data) {
	return http.post(this.url, data);
}

export function test_call() {
	token({id:0}).then(resp => {
		console.log(resp.data.id)
	})
}

/**
 * @returns {Promise<{data: {columns: []}}>}
 */
export async function GetAdminUser() {
	return await http.get("http://127.0.0.1:8080/admin/admin/user");
}

