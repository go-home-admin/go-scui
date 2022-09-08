import http from "@/api/http"

/**
 *
 * @param url {string} 放到啊
 * @returns {Promise<{data: {template: string,data: string,methods: string,mounted: string,props: string}}>}
 */
export function loadTableView(url) {
	return http.get(url);
}
