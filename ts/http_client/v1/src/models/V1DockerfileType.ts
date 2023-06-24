/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc22
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 *
 * @export
 * @interface V1DockerfileType
 */
export interface V1DockerfileType {
    /**
     *
     * @type {string}
     * @memberof V1DockerfileType
     */
    image?: string;
    /**
     *
     * @type {{ [key: string]: string; }}
     * @memberof V1DockerfileType
     */
    env?: { [key: string]: string; };
    /**
     *
     * @type {Array<string>}
     * @memberof V1DockerfileType
     */
    path?: Array<string>;
    /**
     *
     * @type {Array<object>}
     * @memberof V1DockerfileType
     */
    copy?: Array<object>;
    /**
     *
     * @type {Array<object>}
     * @memberof V1DockerfileType
     */
    post_run_copy?: Array<object>;
    /**
     *
     * @type {Array<string>}
     * @memberof V1DockerfileType
     */
    run?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof V1DockerfileType
     */
    langEnv?: string;
    /**
     *
     * @type {number}
     * @memberof V1DockerfileType
     */
    uid?: number;
    /**
     *
     * @type {number}
     * @memberof V1DockerfileType
     */
    gid?: number;
    /**
     *
     * @type {number}
     * @memberof V1DockerfileType
     */
    username?: number;
    /**
     *
     * @type {string}
     * @memberof V1DockerfileType
     */
    filename?: string;
    /**
     *
     * @type {string}
     * @memberof V1DockerfileType
     */
    workdir?: string;
    /**
     *
     * @type {string}
     * @memberof V1DockerfileType
     */
    workdirPath?: string;
    /**
     *
     * @type {string}
     * @memberof V1DockerfileType
     */
    shell?: string;
}

/**
 * Check if a given object implements the V1DockerfileType interface.
 */
export function instanceOfV1DockerfileType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1DockerfileTypeFromJSON(json: any): V1DockerfileType {
    return V1DockerfileTypeFromJSONTyped(json, false);
}

export function V1DockerfileTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1DockerfileType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'image': !exists(json, 'image') ? undefined : json['image'],
        'env': !exists(json, 'env') ? undefined : json['env'],
        'path': !exists(json, 'path') ? undefined : json['path'],
        'copy': !exists(json, 'copy') ? undefined : json['copy'],
        'post_run_copy': !exists(json, 'post_run_copy') ? undefined : json['post_run_copy'],
        'run': !exists(json, 'run') ? undefined : json['run'],
        'langEnv': !exists(json, 'langEnv') ? undefined : json['langEnv'],
        'uid': !exists(json, 'uid') ? undefined : json['uid'],
        'gid': !exists(json, 'gid') ? undefined : json['gid'],
        'username': !exists(json, 'username') ? undefined : json['username'],
        'filename': !exists(json, 'filename') ? undefined : json['filename'],
        'workdir': !exists(json, 'workdir') ? undefined : json['workdir'],
        'workdirPath': !exists(json, 'workdirPath') ? undefined : json['workdirPath'],
        'shell': !exists(json, 'shell') ? undefined : json['shell'],
    };
}

export function V1DockerfileTypeToJSON(value?: V1DockerfileType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'image': value.image,
        'env': value.env,
        'path': value.path,
        'copy': value.copy,
        'post_run_copy': value.post_run_copy,
        'run': value.run,
        'langEnv': value.langEnv,
        'uid': value.uid,
        'gid': value.gid,
        'username': value.username,
        'filename': value.filename,
        'workdir': value.workdir,
        'workdirPath': value.workdirPath,
        'shell': value.shell,
    };
}

