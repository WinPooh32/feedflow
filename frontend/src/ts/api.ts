// tslint:disable
/// <reference path="./custom.d.ts" />
/**
 * FeedFlow
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: 0.3.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as url from "url";
import { Configuration } from "./configuration";
import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';

const BASE_PATH = "https://pooh32.tk/api".replace(/\/+$/, "");

/**
 *
 * @export
 */
export const COLLECTION_FORMATS = {
    csv: ",",
    ssv: " ",
    tsv: "\t",
    pipes: "|",
};

/**
 *  
 * @export
 * @interface RequestArgs
 */
export interface RequestArgs {
    url: string;
    options: any;
}

/**
 * 
 * @export
 * @class BaseAPI
 */
export class BaseAPI {
    protected configuration: Configuration | undefined;

    constructor(configuration?: Configuration, protected basePath: string = BASE_PATH, protected axios: AxiosInstance = globalAxios) {
        if (configuration) {
            this.configuration = configuration;
            this.basePath = configuration.basePath || this.basePath;
        }
    }
};

/**
 * 
 * @export
 * @class RequiredError
 * @extends {Error}
 */
export class RequiredError extends Error {
    name: "RequiredError" = "RequiredError";
    constructor(public field: string, msg?: string) {
        super(msg);
    }
}

/**
 * 
 * @export
 * @interface LoginRequest
 */
export interface LoginRequest {
    /**
     * 
     * @type {string}
     * @memberof LoginRequest
     */
    username: string;
    /**
     * 
     * @type {string}
     * @memberof LoginRequest
     */
    password: string;
}

/**
 * 
 * @export
 * @interface ModelError
 */
export interface ModelError {
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    code: string;
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    message: string;
}

/**
 * 
 * @export
 * @interface NewPageContent
 */
export interface NewPageContent {
    /**
     * 
     * @type {number}
     * @memberof NewPageContent
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof NewPageContent
     */
    title: string;
    /**
     * 
     * @type {string}
     * @memberof NewPageContent
     */
    content: string;
    /**
     * 
     * @type {Array<Tag>}
     * @memberof NewPageContent
     */
    tags: Array<Tag>;
}

/**
 * 
 * @export
 * @interface SigninRequest
 */
export interface SigninRequest {
    /**
     * 
     * @type {string}
     * @memberof SigninRequest
     */
    username: string;
    /**
     * 
     * @type {string}
     * @memberof SigninRequest
     */
    password: string;
    /**
     * 
     * @type {string}
     * @memberof SigninRequest
     */
    email: string;
}

/**
 * 
 * @export
 * @interface Tag
 */
export interface Tag {
    /**
     * 
     * @type {string}
     * @memberof Tag
     */
    value: string;
}

/**
 * 
 * @export
 * @interface UploadImg
 */
export interface UploadImg {
    /**
     * 
     * @type {string}
     * @memberof UploadImg
     */
    hash: string;
    /**
     * 
     * @type {any}
     * @memberof UploadImg
     */
    fileName: any;
}


/**
 * FeedApi - axios parameter creator
 * @export
 */
export const FeedApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Return next feed data chunk.
         * @param {number} since 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestChunk(since: number, options: any = {}): RequestArgs {
            // verify required parameter 'since' is not null or undefined
            if (since === null || since === undefined) {
                throw new RequiredError('since','Required parameter since was null or undefined when calling requestChunk.');
            }
            const localVarPath = `/feed/request`;
            const localVarUrlObj = url.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = Object.assign({ method: 'GET' }, baseOptions, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            if (since !== undefined) {
                localVarQueryParameter['since'] = since;
            }

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * FeedApi - functional programming interface
 * @export
 */
export const FeedApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Return next feed data chunk.
         * @param {number} since 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestChunk(since: number, options?: any): (axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<NewPageContent>> {
            const localVarAxiosArgs = FeedApiAxiosParamCreator(configuration).requestChunk(since, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = Object.assign(localVarAxiosArgs.options, {url: basePath + localVarAxiosArgs.url})
                return axios.request(axiosRequestArgs);                
            };
        },
    }
};

/**
 * FeedApi - factory interface
 * @export
 */
export const FeedApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Return next feed data chunk.
         * @param {number} since 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestChunk(since: number, options?: any) {
            return FeedApiFp(configuration).requestChunk(since, options)(axios, basePath);
        },
    };
};

/**
 * FeedApi - object-oriented interface
 * @export
 * @class FeedApi
 * @extends {BaseAPI}
 */
export class FeedApi extends BaseAPI {
    /**
     * 
     * @summary Return next feed data chunk.
     * @param {number} since 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof FeedApi
     */
    public requestChunk(since: number, options?: any) {
        return FeedApiFp(this.configuration).requestChunk(since, options)(this.axios, this.basePath);
    }

}

/**
 * PagesApi - axios parameter creator
 * @export
 */
export const PagesApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Add a new page.
         * @param {NewPageContent} [newPageContent] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        add(newPageContent?: NewPageContent, options: any = {}): RequestArgs {
            const localVarPath = `/pages/add`;
            const localVarUrlObj = url.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = Object.assign({ method: 'POST' }, baseOptions, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);
            const needsSerialization = (<any>"NewPageContent" !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.data =  needsSerialization ? JSON.stringify(newPageContent || {}) : (newPageContent || "");

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Upload a new image.
         * @param {string} hash 
         * @param {any} fileName 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        imgUpload(hash: string, fileName: any, options: any = {}): RequestArgs {
            // verify required parameter 'hash' is not null or undefined
            if (hash === null || hash === undefined) {
                throw new RequiredError('hash','Required parameter hash was null or undefined when calling imgUpload.');
            }
            // verify required parameter 'fileName' is not null or undefined
            if (fileName === null || fileName === undefined) {
                throw new RequiredError('fileName','Required parameter fileName was null or undefined when calling imgUpload.');
            }
            const localVarPath = `/pages/img/upload/`;
            const localVarUrlObj = url.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = Object.assign({ method: 'POST' }, baseOptions, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;
            const localVarFormParams = new url.URLSearchParams();

            if (hash !== undefined) {
                localVarFormParams.set('hash', hash as any);
            }

            if (fileName !== undefined) {
                localVarFormParams.set('fileName', fileName as any);
            }

            localVarHeaderParameter['Content-Type'] = 'application/x-www-form-urlencoded';

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);
            localVarRequestOptions.data = localVarFormParams.toString();

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Logs in and returns the authentication cookie.
         * @param {LoginRequest} loginRequest A JSON object containing the login and password.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login(loginRequest: LoginRequest, options: any = {}): RequestArgs {
            // verify required parameter 'loginRequest' is not null or undefined
            if (loginRequest === null || loginRequest === undefined) {
                throw new RequiredError('loginRequest','Required parameter loginRequest was null or undefined when calling login.');
            }
            const localVarPath = `/pages/login`;
            const localVarUrlObj = url.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = Object.assign({ method: 'POST' }, baseOptions, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);
            const needsSerialization = (<any>"LoginRequest" !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.data =  needsSerialization ? JSON.stringify(loginRequest || {}) : (loginRequest || "");

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Check username is taken or not.
         * @param {string} username 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        nameIsFree(username: string, options: any = {}): RequestArgs {
            // verify required parameter 'username' is not null or undefined
            if (username === null || username === undefined) {
                throw new RequiredError('username','Required parameter username was null or undefined when calling nameIsFree.');
            }
            const localVarPath = `/user/name/isfree/{username}`
                .replace(`{${"username"}}`, encodeURIComponent(String(username)));
            const localVarUrlObj = url.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = Object.assign({ method: 'HEAD' }, baseOptions, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Move page to archive.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        remove(options: any = {}): RequestArgs {
            const localVarPath = `/pages/remove`;
            const localVarUrlObj = url.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = Object.assign({ method: 'DELETE' }, baseOptions, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Singns in.
         * @param {SigninRequest} signinRequest A JSON object containing the login and password.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        signin(signinRequest: SigninRequest, options: any = {}): RequestArgs {
            // verify required parameter 'signinRequest' is not null or undefined
            if (signinRequest === null || signinRequest === undefined) {
                throw new RequiredError('signinRequest','Required parameter signinRequest was null or undefined when calling signin.');
            }
            const localVarPath = `/pages/signin`;
            const localVarUrlObj = url.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = Object.assign({ method: 'POST' }, baseOptions, options);
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = Object.assign({}, localVarUrlObj.query, localVarQueryParameter, options.query);
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            localVarRequestOptions.headers = Object.assign({}, localVarHeaderParameter, options.headers);
            const needsSerialization = (<any>"SigninRequest" !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.data =  needsSerialization ? JSON.stringify(signinRequest || {}) : (signinRequest || "");

            return {
                url: url.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * PagesApi - functional programming interface
 * @export
 */
export const PagesApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Add a new page.
         * @param {NewPageContent} [newPageContent] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        add(newPageContent?: NewPageContent, options?: any): (axios?: AxiosInstance, basePath?: string) => AxiosPromise<Response> {
            const localVarAxiosArgs = PagesApiAxiosParamCreator(configuration).add(newPageContent, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = Object.assign(localVarAxiosArgs.options, {url: basePath + localVarAxiosArgs.url})
                return axios.request(axiosRequestArgs);                
            };
        },
        /**
         * 
         * @summary Upload a new image.
         * @param {string} hash 
         * @param {any} fileName 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        imgUpload(hash: string, fileName: any, options?: any): (axios?: AxiosInstance, basePath?: string) => AxiosPromise<string> {
            const localVarAxiosArgs = PagesApiAxiosParamCreator(configuration).imgUpload(hash, fileName, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = Object.assign(localVarAxiosArgs.options, {url: basePath + localVarAxiosArgs.url})
                return axios.request(axiosRequestArgs);                
            };
        },
        /**
         * 
         * @summary Logs in and returns the authentication cookie.
         * @param {LoginRequest} loginRequest A JSON object containing the login and password.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login(loginRequest: LoginRequest, options?: any): (axios?: AxiosInstance, basePath?: string) => AxiosPromise<Response> {
            const localVarAxiosArgs = PagesApiAxiosParamCreator(configuration).login(loginRequest, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = Object.assign(localVarAxiosArgs.options, {url: basePath + localVarAxiosArgs.url})
                return axios.request(axiosRequestArgs);                
            };
        },
        /**
         * 
         * @summary Check username is taken or not.
         * @param {string} username 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        nameIsFree(username: string, options?: any): (axios?: AxiosInstance, basePath?: string) => AxiosPromise<Response> {
            const localVarAxiosArgs = PagesApiAxiosParamCreator(configuration).nameIsFree(username, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = Object.assign(localVarAxiosArgs.options, {url: basePath + localVarAxiosArgs.url})
                return axios.request(axiosRequestArgs);                
            };
        },
        /**
         * 
         * @summary Move page to archive.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        remove(options?: any): (axios?: AxiosInstance, basePath?: string) => AxiosPromise<Response> {
            const localVarAxiosArgs = PagesApiAxiosParamCreator(configuration).remove(options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = Object.assign(localVarAxiosArgs.options, {url: basePath + localVarAxiosArgs.url})
                return axios.request(axiosRequestArgs);                
            };
        },
        /**
         * 
         * @summary Singns in.
         * @param {SigninRequest} signinRequest A JSON object containing the login and password.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        signin(signinRequest: SigninRequest, options?: any): (axios?: AxiosInstance, basePath?: string) => AxiosPromise<Response> {
            const localVarAxiosArgs = PagesApiAxiosParamCreator(configuration).signin(signinRequest, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = Object.assign(localVarAxiosArgs.options, {url: basePath + localVarAxiosArgs.url})
                return axios.request(axiosRequestArgs);                
            };
        },
    }
};

/**
 * PagesApi - factory interface
 * @export
 */
export const PagesApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Add a new page.
         * @param {NewPageContent} [newPageContent] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        add(newPageContent?: NewPageContent, options?: any) {
            return PagesApiFp(configuration).add(newPageContent, options)(axios, basePath);
        },
        /**
         * 
         * @summary Upload a new image.
         * @param {string} hash 
         * @param {any} fileName 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        imgUpload(hash: string, fileName: any, options?: any) {
            return PagesApiFp(configuration).imgUpload(hash, fileName, options)(axios, basePath);
        },
        /**
         * 
         * @summary Logs in and returns the authentication cookie.
         * @param {LoginRequest} loginRequest A JSON object containing the login and password.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login(loginRequest: LoginRequest, options?: any) {
            return PagesApiFp(configuration).login(loginRequest, options)(axios, basePath);
        },
        /**
         * 
         * @summary Check username is taken or not.
         * @param {string} username 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        nameIsFree(username: string, options?: any) {
            return PagesApiFp(configuration).nameIsFree(username, options)(axios, basePath);
        },
        /**
         * 
         * @summary Move page to archive.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        remove(options?: any) {
            return PagesApiFp(configuration).remove(options)(axios, basePath);
        },
        /**
         * 
         * @summary Singns in.
         * @param {SigninRequest} signinRequest A JSON object containing the login and password.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        signin(signinRequest: SigninRequest, options?: any) {
            return PagesApiFp(configuration).signin(signinRequest, options)(axios, basePath);
        },
    };
};

/**
 * PagesApi - object-oriented interface
 * @export
 * @class PagesApi
 * @extends {BaseAPI}
 */
export class PagesApi extends BaseAPI {
    /**
     * 
     * @summary Add a new page.
     * @param {NewPageContent} [newPageContent] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PagesApi
     */
    public add(newPageContent?: NewPageContent, options?: any) {
        return PagesApiFp(this.configuration).add(newPageContent, options)(this.axios, this.basePath);
    }

    /**
     * 
     * @summary Upload a new image.
     * @param {string} hash 
     * @param {any} fileName 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PagesApi
     */
    public imgUpload(hash: string, fileName: any, options?: any) {
        return PagesApiFp(this.configuration).imgUpload(hash, fileName, options)(this.axios, this.basePath);
    }

    /**
     * 
     * @summary Logs in and returns the authentication cookie.
     * @param {LoginRequest} loginRequest A JSON object containing the login and password.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PagesApi
     */
    public login(loginRequest: LoginRequest, options?: any) {
        return PagesApiFp(this.configuration).login(loginRequest, options)(this.axios, this.basePath);
    }

    /**
     * 
     * @summary Check username is taken or not.
     * @param {string} username 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PagesApi
     */
    public nameIsFree(username: string, options?: any) {
        return PagesApiFp(this.configuration).nameIsFree(username, options)(this.axios, this.basePath);
    }

    /**
     * 
     * @summary Move page to archive.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PagesApi
     */
    public remove(options?: any) {
        return PagesApiFp(this.configuration).remove(options)(this.axios, this.basePath);
    }

    /**
     * 
     * @summary Singns in.
     * @param {SigninRequest} signinRequest A JSON object containing the login and password.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PagesApi
     */
    public signin(signinRequest: SigninRequest, options?: any) {
        return PagesApiFp(this.configuration).signin(signinRequest, options)(this.axios, this.basePath);
    }

}

