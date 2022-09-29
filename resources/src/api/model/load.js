import http from "@/utils/request"
import config from "@/config";

/**
 *
 * @param url {string}
 * @returns {Promise<{data: {template: string,data: {},methods: string,mounted: string,props: string}}>}
 */
export function loadTableView(url) {
	return http.get(config.API_URL + url);
}
