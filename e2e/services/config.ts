import Axios, { AxiosRequestConfig } from "axios";
import { CURRENT_TEST_ID } from "../main";

export const AXIOS_INSTANCE = Axios.create({
	baseURL: "http://localhost:8080",
});
export const customInstance = async <T>(config: AxiosRequestConfig): Promise<T> => {
	const source = Axios.CancelToken.source();
	const improvedConfig = {
		...config,
		headers: {
			testId: CURRENT_TEST_ID,
		},
	};
	const promise = AXIOS_INSTANCE({
		...improvedConfig,
		cancelToken: source.token,
	}).then(({ data }: any) => data);

	// eslint-disable-next-line
	// @ts-ignore
	promise.cancel = () => {
		source.cancel("Query was cancelled by React Query");
	};

	return promise;
};
